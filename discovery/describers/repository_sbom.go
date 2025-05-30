package describers

import (
	"context"

	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func GetAllRepositoriesSBOMs(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient

	var repositoryName string
	if value := ctx.Value(paramKeyRepoName); value != nil {
		repositoryName = value.(string)
	}

	organization, err := GetOrganizationAdditionalData(ctx, githubClient.RestClient, organizationName)
	if err != nil {
		return nil, err
	}

	if repositoryName != "" {
		repo, err := getRepositoryDetails(ctx, githubClient.RestClient, organizationName, repositoryName)
		if err != nil {
			return nil, err
		}
		repoValue, err := GetRepositorySBOMs(ctx, githubClient, organizationName, repositoryName, repo.GetID(), organization.GetID())
		if err != nil {
			return nil, err
		}
		return []models.Resource{*repoValue}, nil
	}

	repositories, err := getRepositories(ctx, client, organizationName)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValue, err := GetRepositorySBOMs(ctx, githubClient, organizationName, repo.GetName(), repo.GetID(), organization.GetID())
		if err != nil {
			return nil, err
		}
		if stream != nil {
			if err := (*stream)(*repoValue); err != nil {
				return nil, err
			}
		} else {
			values = append(values, *repoValue)
		}
	}
	return values, nil
}

func GetRepositorySBOMs(ctx context.Context, githubClient model.GitHubClient, owner, repo string, repoId int64, orgId int64) (*models.Resource, error) {
	client := githubClient.RestClient
	SBOM, _, err := client.DependencyGraph.GetSBOM(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	repoFullName := formRepositoryFullName(owner, repo)
	value := models.Resource{
		ID:   SBOM.GetSBOM().GetSPDXID(),
		Name: SBOM.GetSBOM().GetName(),
		Description: model.RepoSBOMDescription{
			RepositoryID:       repoId,
			RepositoryFullName: repoFullName,
			SPDXID:             SBOM.GetSBOM().GetSPDXID(),
			SPDXVersion:        SBOM.GetSBOM().GetSPDXVersion(),
			CreationInfo:       SBOM.GetSBOM().GetCreationInfo(),
			Name:               SBOM.GetSBOM().GetName(),
			DataLicense:        SBOM.GetSBOM().GetDataLicense(),
			DocumentDescribes:  SBOM.GetSBOM().DocumentDescribes,
			DocumentNamespace:  SBOM.GetSBOM().GetDocumentNamespace(),
			Packages:           SBOM.GetSBOM().Packages,
			Organization:       owner,
			OrganizationID:     orgId,
			RepositoryName:     repo,
		},
	}
	return &value, nil
}
