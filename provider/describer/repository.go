package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
)

// MAX_REPO as requested
const MAX_REPO = 250

// GetRepositoryList returns a list of all active (non-archived, non-disabled) repos in the organization.
// By default, no excludes are applied, so this returns only active repositories.
func GetRepositoryList(
	ctx context.Context,
	githubClient GitHubClient,
	organizationName string,
	stream *models.StreamSender,
) ([]models.Resource, error) {
	// Call the helper with default options (no excludes)
	return GetRepositoryListWithOptions(ctx, githubClient, organizationName, stream, false, false)
}

// GetRepositoryListWithOptions returns a list of all active repos in the organization with options to exclude archived or disabled.
// It paginates through the results up to MAX_REPO and does minimal processing: it only extracts 'id' and 'name' fields.
// Instead of marshaling into JSON bytes, we store the map[string]interface{} directly in Resource.Description.Value,
// so the streaming code won't break when it re-encodes that data.
func GetRepositoryListWithOptions(
	ctx context.Context,
	githubClient GitHubClient,
	organizationName string,
	stream *models.StreamSender,
	excludeArchived bool,
	excludeDisabled bool,
) ([]models.Resource, error) {

	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	var allResources []models.Resource
	perPage := 100
	page := 1

	for len(allResources) < MAX_REPO {
		endpoint := fmt.Sprintf("/orgs/%s/repos?per_page=%d&page=%d", organizationName, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}

		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("error fetching repos: %w", err)
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
		}

		// Decode into a slice of generic maps to extract 'id', 'name', etc.
		var repos []map[string]interface{}
		if err := json.Unmarshal(resp.Data, &repos); err != nil {
			return nil, fmt.Errorf("error decoding repos list: %w", err)
		}
		if len(repos) == 0 {
			break
		}

		for _, r := range repos {
			// Apply filters
			if excludeArchived {
				if archived, ok := r["archived"].(bool); ok && archived {
					continue
				}
			}
			if excludeDisabled {
				if disabled, ok := r["disabled"].(bool); ok && disabled {
					continue
				}
			}

			var idStr string
			if idVal, ok := r["id"]; ok {
				idStr = fmt.Sprintf("%v", idVal)
			}

			var nameStr string
			if nameVal, ok := r["name"].(string); ok {
				nameStr = nameVal
			}

			// Instead of marshalling to []byte, we store the map object directly in Value
			// This prevents "cannot unmarshal string into Go value of type map[string]interface{}"
			resource := models.Resource{
				ID:   idStr,
				Name: nameStr,
				Description: JSONAllFieldsMarshaller{
					Value: r, // store the entire map
				},
			}

			// Stream the resource if possible
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, fmt.Errorf("streaming resource failed: %w", err)
				}
			}

			allResources = append(allResources, resource)
			if len(allResources) >= MAX_REPO {
				break
			}
		}

		if len(repos) < perPage {
			break
		}
		page++
	}

	return allResources, nil
}

// GetRepository returns details for a given repo: fetches from GitHub, transforms it,
// fetches languages, enriches metrics, and returns a single Resource. If a stream is provided,
// it also streams the resource. This function is fully independent of GetRepositoryListWithOptions.
func GetRepository(
	ctx context.Context,
	githubClient GitHubClient,
	organizationName string,
	repositoryName string,
	resourceID string,
	stream *models.StreamSender,
) (*models.Resource, error) {

	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	// Fetch the single repo's details
	repoDetail, err := util_fetchRepoDetails(sdk, organizationName, repositoryName)
	if err != nil {
		return nil, fmt.Errorf("error fetching repository details for %s/%s: %w",
			organizationName, repositoryName, err)
	}

	// Transform into final structure
	finalDetail := util_transformToFinalRepoDetail(repoDetail)

	// Fetch repository languages
	langs, err := util_fetchLanguages(sdk, organizationName, repositoryName)
	if err == nil && len(langs) > 0 {
		finalDetail.Languages = langs
	}

	// Enrich with metrics (commits, issues, branches, PRs, releases, tags)
	if err := util_enrichRepoMetrics(sdk, organizationName, repositoryName, finalDetail); err != nil {
		log.Printf("Error enriching repo metrics for %s/%s: %v",
			organizationName, repositoryName, err)
	}

	// Build final Resource
	value := models.Resource{
		ID:   strconv.Itoa(finalDetail.GitHubRepoID),
		Name: finalDetail.Name,
		Description: JSONAllFieldsMarshaller{
			// Storing the struct directly, not as raw JSON bytes
			Value: finalDetail,
		},
	}

	// (Optional) Print to terminal
	fmt.Println(value)

	// Stream if provided
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, fmt.Errorf("streaming resource failed: %w", err)
		}
	}

	return &value, nil
}

// ----------------------------------------------------------------------------
// All utility/helper functions now have the prefix "util_"
// ----------------------------------------------------------------------------

// util_fetchRepoDetails fetches a single repository's details from GitHub.
func util_fetchRepoDetails(
	sdk *resilientbridge.ResilientBridge,
	owner, repo string,
) (*model.RepoDetail, error) {

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s", owner, repo),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching repo details: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var detail model.RepoDetail
	if err := json.Unmarshal(resp.Data, &detail); err != nil {
		return nil, fmt.Errorf("error decoding repo details: %w", err)
	}
	return &detail, nil
}

// util_transformToFinalRepoDetail transforms a raw model.RepoDetail into model.RepositoryDescription.
func util_transformToFinalRepoDetail(detail *model.RepoDetail) *model.RepositoryDescription {
	var parent *model.RepositoryDescription
	if detail.Parent != nil {
		parent = util_transformToFinalRepoDetail(detail.Parent)
	}
	var source *model.RepositoryDescription
	if detail.Source != nil {
		source = util_transformToFinalRepoDetail(detail.Source)
	}

	var finalOwner *model.Owner
	if detail.Owner != nil {
		finalOwner = &model.Owner{
			Login:   detail.Owner.Login,
			ID:      detail.Owner.ID,
			NodeID:  detail.Owner.NodeID,
			HTMLURL: detail.Owner.HTMLURL,
			Type:    detail.Owner.Type,
		}
	}

	var finalOrg *model.Organization
	if detail.Organization != nil {
		finalOrg = &model.Organization{
			Login:        detail.Organization.Login,
			ID:           detail.Organization.ID,
			NodeID:       detail.Organization.NodeID,
			HTMLURL:      detail.Organization.HTMLURL,
			Type:         detail.Organization.Type,
			UserViewType: detail.Organization.UserViewType,
			SiteAdmin:    detail.Organization.SiteAdmin,
		}
	}

	dbObj := map[string]string{"name": detail.DefaultBranch}
	dbBytes, _ := json.Marshal(dbObj)

	isActive := !(detail.Archived || detail.Disabled)
	isEmpty := (detail.Size == 0)

	var licenseJSON json.RawMessage
	if detail.License != nil {
		if data, err := json.Marshal(detail.License); err == nil {
			licenseJSON = data
		}
	}

	finalDetail := &model.RepositoryDescription{
		GitHubRepoID:            detail.ID,
		NodeID:                  detail.NodeID,
		Name:                    detail.Name,
		NameWithOwner:           detail.FullName,
		Description:             detail.Description,
		CreatedAt:               detail.CreatedAt,
		UpdatedAt:               detail.UpdatedAt,
		PushedAt:                detail.PushedAt,
		IsActive:                isActive,
		IsEmpty:                 isEmpty,
		IsFork:                  detail.Fork,
		IsSecurityPolicyEnabled: false, // default
		Owner:                   finalOwner,
		HomepageURL:             detail.Homepage,
		LicenseInfo:             licenseJSON,
		Topics:                  detail.Topics,
		Visibility:              detail.Visibility,
		DefaultBranchRef:        dbBytes,
		Permissions:             detail.Permissions,
		Organization:            finalOrg,
		Parent:                  parent,
		Source:                  source,
		Languages:               nil, // set by util_fetchLanguages
		RepositorySettings: model.RepositorySettings{
			HasDiscussionsEnabled:     detail.HasDiscussions,
			HasIssuesEnabled:          detail.HasIssues,
			HasProjectsEnabled:        detail.HasProjects,
			HasWikiEnabled:            detail.HasWiki,
			MergeCommitAllowed:        detail.AllowMergeCommit,
			MergeCommitMessage:        detail.MergeCommitMessage,
			MergeCommitTitle:          detail.MergeCommitTitle,
			SquashMergeAllowed:        detail.AllowSquashMerge,
			SquashMergeCommitMessage:  detail.SquashMergeCommitMessage,
			SquashMergeCommitTitle:    detail.SquashMergeCommitTitle,
			HasDownloads:              detail.HasDownloads,
			HasPages:                  detail.HasPages,
			WebCommitSignoffRequired:  detail.WebCommitSignoffRequired,
			MirrorURL:                 detail.MirrorURL,
			AllowAutoMerge:            detail.AllowAutoMerge,
			DeleteBranchOnMerge:       detail.DeleteBranchOnMerge,
			AllowUpdateBranch:         detail.AllowUpdateBranch,
			UseSquashPRTitleAsDefault: detail.UseSquashPRTitleAsDefault,
			CustomProperties:          detail.CustomProperties,
			ForkingAllowed:            detail.AllowForking,
			IsTemplate:                detail.IsTemplate,
			AllowRebaseMerge:          detail.AllowRebaseMerge,
			Archived:                  detail.Archived,
			Disabled:                  detail.Disabled,
			Locked:                    detail.Locked,
		},
		SecuritySettings: model.SecuritySettings{
			VulnerabilityAlertsEnabled:               false,
			SecretScanningEnabled:                    false,
			SecretScanningPushProtectionEnabled:      false,
			DependabotSecurityUpdatesEnabled:         false,
			SecretScanningNonProviderPatternsEnabled: false,
			SecretScanningValidityChecksEnabled:      false,
		},
		RepoURLs: model.RepoURLs{
			GitURL:   detail.GitURL,
			SSHURL:   detail.SSHURL,
			CloneURL: detail.CloneURL,
			SVNURL:   detail.SVNURL,
			HTMLURL:  detail.HTMLURL,
		},
		Metrics: model.Metrics{
			Stargazers:  detail.StargazersCount,
			Forks:       detail.ForksCount,
			Subscribers: detail.SubscribersCount,
			Size:        detail.Size,
			OpenIssues:  detail.OpenIssuesCount,
		},
	}

	return finalDetail
}

// util_fetchLanguages fetches repository languages and returns map[string]int.
func util_fetchLanguages(
	sdk *resilientbridge.ResilientBridge,
	owner, repo string,
) (map[string]int, error) {

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/repos/%s/%s/languages", owner, repo),
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching languages: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	var langs map[string]int
	if err := json.Unmarshal(resp.Data, &langs); err != nil {
		return nil, fmt.Errorf("error decoding languages: %w", err)
	}
	return langs, nil
}

// util_enrichRepoMetrics populates metrics (commits, issues, branches, etc.) in finalDetail.
func util_enrichRepoMetrics(
	sdk *resilientbridge.ResilientBridge,
	owner, repoName string,
	finalDetail *model.RepositoryDescription,
) error {

	var dbObj map[string]string
	if finalDetail.DefaultBranchRef != nil {
		if err := json.Unmarshal(finalDetail.DefaultBranchRef, &dbObj); err != nil {
			return err
		}
	}
	defaultBranch := dbObj["name"]
	if defaultBranch == "" {
		defaultBranch = "main"
	}

	commitsCount, err := util_countCommits(sdk, owner, repoName, defaultBranch)
	if err != nil {
		return fmt.Errorf("counting commits: %w", err)
	}
	finalDetail.Metrics.Commits = commitsCount

	issuesCount, err := util_countIssues(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting issues: %w", err)
	}
	finalDetail.Metrics.Issues = issuesCount

	branchesCount, err := util_countBranches(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting branches: %w", err)
	}
	finalDetail.Metrics.Branches = branchesCount

	prCount, err := util_countPullRequests(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting PRs: %w", err)
	}
	finalDetail.Metrics.PullRequests = prCount

	releasesCount, err := util_countReleases(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting releases: %w", err)
	}
	finalDetail.Metrics.Releases = releasesCount

	tagsCount, err := util_countTags(sdk, owner, repoName)
	if err != nil {
		return fmt.Errorf("counting tags: %w", err)
	}
	finalDetail.Metrics.Tags = tagsCount

	return nil
}

// util_countTags counts how many tags the repo has.
func util_countTags(
	sdk *resilientbridge.ResilientBridge,
	owner, repoName string,
) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/tags?per_page=1", owner, repoName)
	return util_countItemsFromEndpoint(sdk, endpoint)
}

// util_countCommits counts how many commits are in the default branch.
func util_countCommits(
	sdk *resilientbridge.ResilientBridge,
	owner, repoName, defaultBranch string,
) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/commits?sha=%s&per_page=1", owner, repoName, defaultBranch)
	return util_countItemsFromEndpoint(sdk, endpoint)
}

// util_countIssues counts how many issues (open, closed, etc.) are in the repo.
func util_countIssues(
	sdk *resilientbridge.ResilientBridge,
	owner, repoName string,
) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/issues?state=all&per_page=1", owner, repoName)
	return util_countItemsFromEndpoint(sdk, endpoint)
}

// util_countBranches counts how many branches the repo has.
func util_countBranches(
	sdk *resilientbridge.ResilientBridge,
	owner, repoName string,
) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/branches?per_page=1", owner, repoName)
	return util_countItemsFromEndpoint(sdk, endpoint)
}

// util_countPullRequests counts how many PRs (open, closed, merged, etc.) are in the repo.
func util_countPullRequests(
	sdk *resilientbridge.ResilientBridge,
	owner, repoName string,
) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/pulls?state=all&per_page=1", owner, repoName)
	return util_countItemsFromEndpoint(sdk, endpoint)
}

// util_countReleases counts how many releases the repo has.
func util_countReleases(
	sdk *resilientbridge.ResilientBridge,
	owner, repoName string,
) (int, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/releases?per_page=1", owner, repoName)
	return util_countItemsFromEndpoint(sdk, endpoint)
}

// util_countItemsFromEndpoint tries to parse the 'Link' header for a "last page" or, if none, uses the length of the returned array.
func util_countItemsFromEndpoint(
	sdk *resilientbridge.ResilientBridge,
	endpoint string,
) (int, error) {

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %w", err)
	}
	// Some repos return 409 for certain endpoints (e.g. empty repos with no branches).
	if resp.StatusCode == 409 {
		return 0, nil
	}
	if resp.StatusCode >= 400 {
		return 0, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	// Attempt to parse 'Link' header for last page
	var linkHeader string
	for k, v := range resp.Headers {
		if strings.ToLower(k) == "link" {
			linkHeader = v
			break
		}
	}

	if linkHeader == "" {
		// If there's no Link header, see if the response is a JSON array
		if len(resp.Data) > 2 {
			var items []interface{}
			if err := json.Unmarshal(resp.Data, &items); err != nil {
				// If we can't parse it as array, assume at least 1 item
				return 1, nil
			}
			return len(items), nil
		}
		return 0, nil
	}

	lastPage, err := util_parseLastPage(linkHeader)
	if err != nil {
		return 0, fmt.Errorf("could not parse last page: %w", err)
	}
	return lastPage, nil
}

// util_parseLastPage reads the "last" page link from the Link header (e.g. `rel="last"&page=5`).
func util_parseLastPage(linkHeader string) (int, error) {
	re := regexp.MustCompile(`page=(\d+)>; rel="last"`)
	matches := re.FindStringSubmatch(linkHeader)
	if len(matches) < 2 {
		// If no "last" link, assume only 1 page
		return 1, nil
	}
	var lastPage int
	if _, err := fmt.Sscanf(matches[1], "%d", &lastPage); err != nil {
		return 0, err
	}
	return lastPage, nil
}
