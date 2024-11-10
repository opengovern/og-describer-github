package describer

import (
	"context"
	"fmt"
	"github.com/google/go-github/v55/github"
	"github.com/shurcooL/githubv4"
	"slices"
)

const (
	maxPagesCount = 100
	pageSize      = 100
	repoPageSize  = 50
)

func appendRepoColumnIncludes(m *map[string]interface{}, cols []string) {
	optionals := map[string]string{
		"allow_update_branch":              "includeAllowUpdateBranch",
		"archived_at":                      "includeArchivedAt",
		"auto_merge_allowed":               "includeAutoMergeAllowed",
		"can_administer":                   "includeCanAdminister",
		"can_create_projects":              "includeCanCreateProjects",
		"can_subscribe":                    "includeCanSubscribe",
		"can_update_topics":                "includeCanUpdateTopics",
		"code_of_conduct":                  "includeCodeOfConduct",
		"contact_links":                    "includeContactLinks",
		"created_at":                       "includeCreatedAt",
		"default_branch_ref":               "includeDefaultBranchRef",
		"delete_branch_on_merge":           "includeDeleteBranchOnMerge",
		"description":                      "includeDescription",
		"disk_usage":                       "includeDiskUsage",
		"fork_count":                       "includeForkCount",
		"forking_allowed":                  "includeForkingAllowed",
		"funding_links":                    "includeFundingLinks",
		"has_discussions_enabled":          "includeHasDiscussionsEnabled",
		"has_issues_enabled":               "includeHasIssuesEnabled",
		"has_projects_enabled":             "includeHasProjectsEnabled",
		"has_starred":                      "includeHasStarred",
		"has_vulnerability_alerts_enabled": "includeHasVulnerabilityAlertsEnabled",
		"has_wiki_enabled":                 "includeHasWikiEnabled",
		"homepage_url":                     "includeHomepageUrl",
		"interaction_ability":              "includeInteractionAbility",
		"is_archived":                      "includeIsArchived",
		"is_blank_issues_enabled":          "includeIsBlankIssuesEnabled",
		"is_disabled":                      "includeIsDisabled",
		"is_empty":                         "includeIsEmpty",
		"is_fork":                          "includeIsFork",
		"is_in_organization":               "includeIsInOrganization",
		"is_locked":                        "includeIsLocked",
		"is_mirror":                        "includeIsMirror",
		"is_private":                       "includeIsPrivate",
		"is_security_policy_enabled":       "includeIsSecurityPolicyEnabled",
		"is_template":                      "includeIsTemplate",
		"is_user_configuration_repository": "includeIsUserConfigurationRepository",
		"issue_templates":                  "includeIssueTemplates",
		"license_info":                     "includeLicenseInfo",
		"lock_reason":                      "includeLockReason",
		"merge_commit_allowed":             "includeMergeCommitAllowed",
		"merge_commit_message":             "includeMergeCommitMessage",
		"merge_commit_title":               "includeMergeCommitTitle",
		"mirror_url":                       "includeMirrorUrl",
		"open_graph_image_url":             "includeOpenGraphImageUrl",
		"open_issues_total_count":          "includeOpenIssues",
		"possible_commit_emails":           "includePossibleCommitEmails",
		"primary_language":                 "includePrimaryLanguage",
		"projects_url":                     "includeProjectsUrl",
		"pull_request_templates":           "includePullRequestTemplates",
		"pushed_at":                        "includePushedAt",
		"rebase_merge_allowed":             "includeRebaseMergeAllowed",
		"repository_topics_total_count":    "includeRepositoryTopics",
		"security_policy_url":              "includeSecurityPolicyUrl",
		"squash_merge_allowed":             "includeSquashMergeAllowed",
		"squash_merge_commit_message":      "includeSquashMergeCommitMessage",
		"squash_merge_commit_title":        "includeSquashMergeCommitTitle",
		"ssh_url":                          "includeSshUrl",
		"stargazer_count":                  "includeStargazerCount",
		"subscription":                     "includeSubscription",
		"updated_at":                       "includeUpdatedAt",
		"url":                              "includeUrl",
		"uses_custom_open_graph_image":     "includeUsesCustomOpenGraphImage",
		"visibility":                       "includeVisibility",
		"watchers_total_count":             "includeWatchers",
		"web_commit_signoff_required":      "includeWebCommitSignoffRequired",
		"your_permission":                  "includeYourPermission",
	}
	for key, value := range optionals {
		(*m)[value] = githubv4.Boolean(slices.Contains(cols, key))
	}
}

func appendBranchColumnIncludes(m *map[string]interface{}, cols []string) {
	protectionIncluded := githubv4.Boolean(slices.Contains(cols, "protected") || slices.Contains(cols, "branch_protection_rule"))
	(*m)["includeBranchProtectionRule"] = protectionIncluded
	(*m)["includeAllowsDeletions"] = protectionIncluded
	(*m)["includeAllowsForcePushes"] = protectionIncluded
	(*m)["includeBlocksCreations"] = protectionIncluded
	(*m)["includeCreator"] = protectionIncluded
	(*m)["includeBranchProtectionRuleId"] = protectionIncluded
	(*m)["includeDismissesStaleReviews"] = protectionIncluded
	(*m)["includeIsAdminEnforced"] = protectionIncluded
	(*m)["includeLockAllowsFetchAndMerge"] = protectionIncluded
	(*m)["includeLockBranch"] = protectionIncluded
	(*m)["includePattern"] = protectionIncluded
	(*m)["includeRequireLastPushApproval"] = protectionIncluded
	(*m)["includeRequiredApprovingReviewCount"] = protectionIncluded
	(*m)["includeRequiredDeploymentEnvironments"] = protectionIncluded
	(*m)["includeRequiredStatusChecks"] = protectionIncluded
	(*m)["includeRequiresApprovingReviews"] = protectionIncluded
	(*m)["includeRequiresConversationResolution"] = protectionIncluded
	(*m)["includeRequiresCodeOwnerReviews"] = protectionIncluded
	(*m)["includeRequiresCommitSignatures"] = protectionIncluded
	(*m)["includeRequiresDeployments"] = protectionIncluded
	(*m)["includeRequiresLinearHistory"] = protectionIncluded
	(*m)["includeRequiresStatusChecks"] = protectionIncluded
	(*m)["includeRequiresStrictStatusChecks"] = protectionIncluded
	(*m)["includeRestrictsPushes"] = protectionIncluded
	(*m)["includeRestrictsReviewDismissals"] = protectionIncluded
	(*m)["includeMatchingBranches"] = protectionIncluded
}

func appendBranchProtectionRuleColumnIncludes(m *map[string]interface{}, cols []string) {
	(*m)["includeAllowsDeletions"] = githubv4.Boolean(slices.Contains(cols, "allows_deletions"))
	(*m)["includeAllowsForcePushes"] = githubv4.Boolean(slices.Contains(cols, "allows_force_pushes"))
	(*m)["includeBlocksCreations"] = githubv4.Boolean(slices.Contains(cols, "blocks_creations"))
	(*m)["includeCreator"] = githubv4.Boolean(slices.Contains(cols, "creator") || slices.Contains(cols, "creator_login"))
	(*m)["includeBranchProtectionRuleId"] = githubv4.Boolean(slices.Contains(cols, "id"))
	(*m)["includeDismissesStaleReviews"] = githubv4.Boolean(slices.Contains(cols, "dismisses_stale_reviews"))
	(*m)["includeIsAdminEnforced"] = githubv4.Boolean(slices.Contains(cols, "is_admin_enforced"))
	(*m)["includeLockAllowsFetchAndMerge"] = githubv4.Boolean(slices.Contains(cols, "lock_allows_fetch_and_merge"))
	(*m)["includeLockBranch"] = githubv4.Boolean(slices.Contains(cols, "lock_branch"))
	(*m)["includePattern"] = githubv4.Boolean(slices.Contains(cols, "pattern"))
	(*m)["includeRequireLastPushApproval"] = githubv4.Boolean(slices.Contains(cols, "require_last_push_approval"))
	(*m)["includeRequiredApprovingReviewCount"] = githubv4.Boolean(slices.Contains(cols, "required_approving_review_count"))
	(*m)["includeRequiredDeploymentEnvironments"] = githubv4.Boolean(slices.Contains(cols, "required_deployment_environments"))
	(*m)["includeRequiredStatusChecks"] = githubv4.Boolean(slices.Contains(cols, "required_status_checks"))
	(*m)["includeRequiresApprovingReviews"] = githubv4.Boolean(slices.Contains(cols, "requires_approving_reviews"))
	(*m)["includeRequiresConversationResolution"] = githubv4.Boolean(slices.Contains(cols, "requires_conversation_resolution"))
	(*m)["includeRequiresCodeOwnerReviews"] = githubv4.Boolean(slices.Contains(cols, "requires_code_owner_reviews"))
	(*m)["includeRequiresCommitSignatures"] = githubv4.Boolean(slices.Contains(cols, "requires_commit_signatures"))
	(*m)["includeRequiresDeployments"] = githubv4.Boolean(slices.Contains(cols, "requires_deployments"))
	(*m)["includeRequiresLinearHistory"] = githubv4.Boolean(slices.Contains(cols, "requires_linear_history"))
	(*m)["includeRequiresStatusChecks"] = githubv4.Boolean(slices.Contains(cols, "requires_status_checks"))
	(*m)["includeRequiresStrictStatusChecks"] = githubv4.Boolean(slices.Contains(cols, "requires_strict_status_checks"))
	(*m)["includeRestrictsPushes"] = githubv4.Boolean(slices.Contains(cols, "restricts_pushes"))
	(*m)["includeRestrictsReviewDismissals"] = githubv4.Boolean(slices.Contains(cols, "restricts_review_dismissals"))
	(*m)["includeMatchingBranches"] = githubv4.Boolean(slices.Contains(cols, "matching_branches"))
}

func appendCommitColumnIncludes(m *map[string]interface{}, cols []string) {
	// For BasicCommit struct
	(*m)["includeCommitShortSha"] = githubv4.Boolean(slices.Contains(cols, "short_sha"))
	(*m)["includeCommitAuthoredDate"] = githubv4.Boolean(slices.Contains(cols, "authored_date"))
	(*m)["includeCommitAuthor"] = githubv4.Boolean(slices.Contains(cols, "author") || slices.Contains(cols, "author_login"))
	(*m)["includeCommitCommittedDate"] = githubv4.Boolean(slices.Contains(cols, "committed_date"))
	(*m)["includeCommitCommitter"] = githubv4.Boolean(slices.Contains(cols, "committer") || slices.Contains(cols, "committer_login"))
	(*m)["includeCommitMessage"] = githubv4.Boolean(slices.Contains(cols, "message"))
	(*m)["includeCommitUrl"] = githubv4.Boolean(slices.Contains(cols, "url"))
	// For Commit struct
	(*m)["includeCommitAdditions"] = githubv4.Boolean(slices.Contains(cols, "additions"))
	(*m)["includeCommitAuthoredByCommitter"] = githubv4.Boolean(slices.Contains(cols, "authored_by_committer"))
	(*m)["includeCommitChangedFiles"] = githubv4.Boolean(slices.Contains(cols, "changed_files"))
	(*m)["includeCommitCommittedViaWeb"] = githubv4.Boolean(slices.Contains(cols, "committed_via_web"))
	(*m)["includeCommitCommitUrl"] = githubv4.Boolean(slices.Contains(cols, "commit_url"))
	(*m)["includeCommitDeletions"] = githubv4.Boolean(slices.Contains(cols, "deletions"))
	(*m)["includeCommitSignature"] = githubv4.Boolean(slices.Contains(cols, "signature"))
	(*m)["includeCommitTarballUrl"] = githubv4.Boolean(slices.Contains(cols, "tarball_url"))
	(*m)["includeCommitTreeUrl"] = githubv4.Boolean(slices.Contains(cols, "tree_url"))
	(*m)["includeCommitCanSubscribe"] = githubv4.Boolean(slices.Contains(cols, "can_subscribe"))
	(*m)["includeCommitSubscription"] = githubv4.Boolean(slices.Contains(cols, "subscription"))
	(*m)["includeCommitZipballUrl"] = githubv4.Boolean(slices.Contains(cols, "zipball_url"))
	(*m)["includeCommitMessageHeadline"] = githubv4.Boolean(slices.Contains(cols, "message_headline"))
	(*m)["includeCommitStatus"] = githubv4.Boolean(slices.Contains(cols, "status"))
	(*m)["includeCommitNodeId"] = githubv4.Boolean(slices.Contains(cols, "node_id"))
}

func appendCommunityProfileColumnIncludes(m *map[string]interface{}, cols []string) {
	(*m)["includeCPLicense"] = githubv4.Boolean(slices.Contains(cols, "license_info"))
	(*m)["includeCPCodeOfConduct"] = githubv4.Boolean(slices.Contains(cols, "code_of_conduct"))
	(*m)["includeCPIssueTemplates"] = githubv4.Boolean(slices.Contains(cols, "issue_templates"))
	(*m)["includeCPPullRequestTemplates"] = githubv4.Boolean(slices.Contains(cols, "pull_request_templates"))
	(*m)["includeCPReadme"] = githubv4.Boolean(slices.Contains(cols, "readme"))
	(*m)["includeCPContributing"] = githubv4.Boolean(slices.Contains(cols, "contributing"))
	(*m)["includeCPSecurity"] = githubv4.Boolean(slices.Contains(cols, "security"))
}

func appendOrganizationColumnIncludes(m *map[string]interface{}, cols []string) {
	(*m)["includeAnnouncement"] = githubv4.Boolean(slices.Contains(cols, "announcement"))
	(*m)["includeAnnouncementExpiresAt"] = githubv4.Boolean(slices.Contains(cols, "announcement_expires_at"))
	(*m)["includeAnnouncementUserDismissible"] = githubv4.Boolean(slices.Contains(cols, "announcement_user_dismissible"))
	(*m)["includeAnyPinnableItems"] = githubv4.Boolean(slices.Contains(cols, "any_pinnable_items"))
	(*m)["includeAvatarUrl"] = githubv4.Boolean(slices.Contains(cols, "avatar_url"))
	(*m)["includeEstimatedNextSponsorsPayoutInCents"] = githubv4.Boolean(slices.Contains(cols, "estimated_next_sponsors_payout_in_cents"))
	(*m)["includeHasSponsorsListing"] = githubv4.Boolean(slices.Contains(cols, "has_sponsors_listing"))
	(*m)["includeInteractionAbility"] = githubv4.Boolean(slices.Contains(cols, "interaction_ability"))
	(*m)["includeIsSponsoringYou"] = githubv4.Boolean(slices.Contains(cols, "is_sponsoring_you"))
	(*m)["includeIsVerified"] = githubv4.Boolean(slices.Contains(cols, "is_verified"))
	(*m)["includeLocation"] = githubv4.Boolean(slices.Contains(cols, "location"))
	(*m)["includeMonthlyEstimatedSponsorsIncomeInCents"] = githubv4.Boolean(slices.Contains(cols, "monthly_estimated_sponsors_income_in_cents"))
	(*m)["includeNewTeamUrl"] = githubv4.Boolean(slices.Contains(cols, "new_team_url"))
	(*m)["includePinnedItemsRemaining"] = githubv4.Boolean(slices.Contains(cols, "pinned_items_remaining"))
	(*m)["includeProjectsUrl"] = githubv4.Boolean(slices.Contains(cols, "projects_url"))
	(*m)["includeSamlIdentityProvider"] = githubv4.Boolean(slices.Contains(cols, "saml_identity_provider"))
	(*m)["includeSponsorsListing"] = githubv4.Boolean(slices.Contains(cols, "sponsors_listing"))
	(*m)["includeTeamsUrl"] = githubv4.Boolean(slices.Contains(cols, "teams_url"))
	(*m)["includeTotalSponsorshipAmountAsSponsorInCents"] = githubv4.Boolean(slices.Contains(cols, "total_sponsorship_amount_as_sponsor_in_cents"))
	(*m)["includeTwitterUsername"] = githubv4.Boolean(slices.Contains(cols, "twitter_username"))
	(*m)["includeOrgViewer"] = githubv4.Boolean(slices.Contains(cols, "can_administer") || slices.Contains(cols, "can_changed_pinned_items") || slices.Contains(cols, "can_create_projects") || slices.Contains(cols, "can_create_repositories") || slices.Contains(cols, "can_create_teams") || slices.Contains(cols, "can_sponsor"))
	(*m)["includeIsAMember"] = githubv4.Boolean(slices.Contains(cols, "is_a_member"))
	(*m)["includeIsFollowing"] = githubv4.Boolean(slices.Contains(cols, "is_following"))
	(*m)["includeIsSponsoring"] = githubv4.Boolean(slices.Contains(cols, "is_sponsoring"))
	(*m)["includeWebsiteUrl"] = githubv4.Boolean(slices.Contains(cols, "website_url"))
	(*m)["includeMembersWithRole"] = githubv4.Boolean(slices.Contains(cols, "members_with_role_total_count"))
	(*m)["includePackages"] = githubv4.Boolean(slices.Contains(cols, "packages_total_count"))
	(*m)["includePinnableItems"] = githubv4.Boolean(slices.Contains(cols, "pinnable_items_total_count"))
	(*m)["includePinnedItems"] = githubv4.Boolean(slices.Contains(cols, "pinned_items_total_count"))
	(*m)["includeProjects"] = githubv4.Boolean(slices.Contains(cols, "projects_total_count"))
	(*m)["includeProjectsV2"] = githubv4.Boolean(slices.Contains(cols, "projects_v2_total_count"))
	(*m)["includeSponsoring"] = githubv4.Boolean(slices.Contains(cols, "sponsoring_total_count"))
	(*m)["includeSponsors"] = githubv4.Boolean(slices.Contains(cols, "sponsors_total_count"))
	(*m)["includeTeams"] = githubv4.Boolean(slices.Contains(cols, "teams_total_count"))
	(*m)["includePrivateRepositories"] = githubv4.Boolean(slices.Contains(cols, "private_repositories_total_count"))
	(*m)["includePublicRepositories"] = githubv4.Boolean(slices.Contains(cols, "public_repositories_total_count"))
	(*m)["includeRepositories"] = githubv4.Boolean(slices.Contains(cols, "repositories_total_count"))
	(*m)["includeRepositories"] = githubv4.Boolean(slices.Contains(cols, "repositories_total_disk_usage"))
}

func appendStarColumnIncludes(m *map[string]interface{}, cols []string) {
	(*m)["includeStarNode"] = githubv4.Boolean(slices.Contains(cols, "repository_full_name") || slices.Contains(cols, "url"))
	(*m)["includeStarEdges"] = githubv4.Boolean(slices.Contains(cols, "starred_at"))
}

func repositoryCols() []string {
	return []string{
		"id",
		"node_id",
		"name",
		"allow_update_branch",
		"archived_at",
		"auto_merge_allowed",
		"code_of_conduct",
		"contact_links",
		"created_at",
		"default_branch_ref",
		"delete_branch_on_merge",
		"description",
		"disk_usage",
		"fork_count",
		"forking_allowed",
		"funding_links",
		"has_discussions_enabled",
		"has_issues_enabled",
		"has_projects_enabled",
		"has_vulnerability_alerts_enabled",
		"has_wiki_enabled",
		"homepage_url",
		"interaction_ability",
		"is_archived",
		"is_blank_issues_enabled",
		"is_disabled",
		"is_empty",
		"is_fork",
		"is_in_organization",
		"is_locked",
		"is_mirror",
		"is_private",
		"is_security_policy_enabled",
		"is_template",
		"is_user_configuration_repository",
		"issue_templates",
		"license_info",
		"lock_reason",
		"merge_commit_allowed",
		"merge_commit_message",
		"merge_commit_title",
		"mirror_url",
		"name_with_owner",
		"open_graph_image_url",
		"owner_login",
		"primary_language",
		"projects_url",
		"pull_request_templates",
		"pushed_at",
		"rebase_merge_allowed",
		"security_policy_url",
		"squash_merge_allowed",
		"squash_merge_commit_message",
		"squash_merge_commit_title",
		"ssh_url",
		"stargazer_count",
		"updated_at",
		"url",
		"uses_custom_open_graph_image",
		"can_administer",
		"can_create_projects",
		"can_subscribe",
		"can_update_topics",
		"has_starred",
		"possible_commit_emails",
		"subscription",
		"visibility",
		"your_permission",
		"web_commit_signoff_required",
		"repository_topics_total_count",
		"open_issues_total_count",
		"watchers_total_count",
		"hooks",
		"topics",
		"subscribers_count",
		"has_downloads",
		"has_pages",
		"network_count",
	}
}

func branchCols() []string {
	return []string{
		"repository_full_name",
		"name",
		"commit",
		"protected",
		"branch_protection_rule",
	}
}

func branchProtectionCols() []string {
	return []string{
		"repository_full_name",
		"id",
		"node_id",
		"matching_branches",
		"is_admin_enforced",
		"allows_deletions",
		"allows_force_pushes",
		"blocks_creations",
		"creator_login",
		"dismisses_stale_reviews",
		"lock_allows_fetch_and_merge",
		"lock_branch",
		"pattern",
		"require_last_push_approval",
		"requires_approving_reviews",
		"required_approving_review_count",
		"requires_conversation_resolution",
		"requires_code_owner_reviews",
		"requires_commit_signatures",
		"requires_deployments",
		"required_deployment_environments",
		"requires_linear_history",
		"requires_status_checks",
		"required_status_checks",
		"requires_strict_status_checks",
		"restricts_review_dismissals",
		"restricts_pushes",
		"push_allowance_apps",
		"push_allowance_teams",
		"push_allowance_users",
		"bypass_force_push_allowance_apps",
		"bypass_force_push_allowance_teams",
		"bypass_force_push_allowance_users",
		"bypass_pull_request_allowance_apps",
		"bypass_pull_request_allowance_teams",
		"bypass_pull_request_allowance_users",
		"repository_full_name",
		"name",
		"commit",
		"protected",
		"branch_protection_rule",
	}
}

func commitCols() []string {
	return []string{
		"repository_full_name",
		"sha",
		"short_sha",
		"message",
		"author_login",
		"authored_date",
		"author",
		"committer_login",
		"committed_date",
		"committer",
		"additions",
		"authored_by_committer",
		"deletions",
		"changed_files",
		"committed_via_web",
		"commit_url",
		"signature",
		"status",
		"tarball_url",
		"zipball_url",
		"tree_url",
		"can_subscribe",
		"subscription",
		"url",
		"node_id",
		"message_headline",
	}
}

func communityCols() []string {
	return []string{
		"repository_full_name",
		"code_of_conduct",
		"contributing",
		"issue_templates",
		"pull_request_templates",
		"license_info",
		"readme",
		"security",
	}
}

func organizationCols() []string {
	return []string{
		"login",
		"id",
		"node_id",
		"name",
		"created_at",
		"updated_at",
		"description",
		"email",
		"url",
		"announcement",
		"announcement_expires_at",
		"announcement_user_dismissible",
		"any_pinnable_items",
		"avatar_url",
		"estimated_next_sponsors_payout_in_cents",
		"has_sponsors_listing",
		"interaction_ability",
		"is_sponsoring_you",
		"is_verified",
		"location",
		"monthly_estimated_sponsors_income_in_cents",
		"new_team_url",
		"pinned_items_remaining",
		"projects_url",
		"saml_identity_provider",
		"sponsors_listing",
		"teams_url",
		"total_sponsorship_amount_as_sponsor_in_cents",
		"twitter_username",
		"can_administer",
		"can_changed_pinned_items",
		"can_create_projects",
		"can_create_repositories",
		"can_create_teams",
		"can_sponsor",
		"is_a_member",
		"is_following",
		"is_sponsoring",
		"website_url",
		"hooks",
		"billing_email",
		"two_factor_requirement_enabled",
		"default_repo_permission",
		"members_allowed_repository_creation_type",
		"members_can_create_internal_repos",
		"members_can_create_pages",
		"members_can_create_private_repos",
		"members_can_create_public_repos",
		"members_can_create_repos",
		"members_can_fork_private_repos",
		"plan_filled_seats",
		"plan_name",
		"plan_private_repos",
		"plan_seats",
		"plan_space",
		"followers",
		"following",
		"collaborators",
		"has_organization_projects",
		"has_repository_projects",
		"web_commit_signoff_required",
	}
}

func starCols() []string {
	return []string{
		"repository_full_name",
		"starred_at",
		"url",
	}
}

func getOwnerName(ctx context.Context, client *github.Client) (string, error) {
	owner, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return "", err
	}
	ownerName := *owner.Name
	return ownerName, err
}

func getRepositoriesName(ctx context.Context, client *github.Client, owner string) ([]string, error) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: maxPagesCount},
	}
	var repositories []string
	for {
		repos, resp, err := client.Repositories.List(ctx, owner, opt)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			repositories = append(repositories, repo.GetName())
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return repositories, nil
}

func formRepositoryFullName(owner, repo string) string {
	return fmt.Sprintf("%s/%s", owner, repo)
}
