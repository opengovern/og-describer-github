package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func sharedPullRequestColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "repository_full_name", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepoFullName"),
			Description: "The full name of the repository the pull request belongs to."},
		{
			Name:        "organization",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Organization"),
			Description: "organization name",
		},
		{
			Name:        "repository_name",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepositoryName"),
			Description: "repository name",
		},
		{Name: "number", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Number"),
			Description: "The number of the pull request."},
		{Name: "id", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Id"),
			Description: "The ID of the pull request."},
		{Name: "node_id", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.NodeId"),
			Description: "The node ID of the pull request."},
		{Name: "active_lock_reason", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.ActiveLockReason"),
			Description: "Reason that the conversation was locked."},
		{Name: "additions", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Additions"),
			Description: "The number of additions in this pull request."},
		{Name: "author", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Author"),
			Description: "The author of the pull request."},
		{Name: "author_association", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.AuthorAssociation"),
			Description: "Author's association with the pull request."},
		{Name: "base_ref_name", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.BaseRefName"),
			Description: "Identifies the name of the base Ref associated with the pull request, even if the ref has been deleted."},
		{Name: "body", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Body"),
			Description: "The body as Markdown."},
		{Name: "changed_files", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.ChangedFiles"),
			Description: "The number of files changed in this pull request."},
		{Name: "checks_url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.ChecksUrl"),
			Description: "URL for the checks of this pull request."},
		{Name: "closed", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.Closed"),
			Description: "If true, pull request is closed."},
		{Name: "closed_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.ClosedAt").NullIfZero().Transform(convertTimestamp),
			Description: "Timestamp when the pull request was closed."},
		{Name: "created_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.CreatedAt").NullIfZero().Transform(convertTimestamp),
			Description: "Timestamp when the pull request was created."},
		{Name: "created_via_email", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CreatedViaEmail"),
			Description: "If true, pull request comment was created via email."},
		{Name: "deletions", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Deletions"),
			Description: "The number of deletions in this pull request."},
		{Name: "editor", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Editor"),
			Description: "The actor who edited the pull request's body."},
		{Name: "head_ref_name", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.HeadRefName"),
			Description: "Identifies the name of the head Ref associated with the pull request, even if the ref has been deleted."},
		{Name: "head_ref_oid", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.HeadRefOid"),
			Description: "Identifies the oid/sha of the head ref associated with the pull request, even if the ref has been deleted."},
		{Name: "includes_created_edit", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IncludesCreatedEdit"),
			Description: "If true, this pull request was edited and includes an edit with the creation data."},
		{Name: "is_cross_repository", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsCrossRepository"),
			Description: "If true, head and base repositories are different."},
		{Name: "is_draft", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsDraft"),
			Description: "If true, the pull request is a draft."},
		{Name: "is_read_by_user", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.IsReadByUser"),
			Description: "If true, this pull request was read by the current user."},
		{Name: "last_edited_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.LastEditedAt").NullIfZero().Transform(convertTimestamp),
			Description: "Timestamp the editor made the last edit."},
		{Name: "locked", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.Locked"),
			Description: "If true, the pull request is locked."},
		{Name: "maintainer_can_modify", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.MaintainerCanModify"),
			Description: "If true, maintainers can modify the pull request."},
		{Name: "mergeable", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Mergeable"),
			Description: "Whether or not the pull request can be merged based on the existence of merge conflicts."},
		{Name: "merged", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.Merged"),
			Description: "If true, the pull request was merged."},
		{Name: "merged_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.MergedAt").NullIfZero().Transform(convertTimestamp),
			Description: "Timestamp when pull request was merged."},
		{Name: "merged_by", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.MergedBy"),
			Description: "The actor who merged the pull request."},
		{Name: "milestone", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Milestone"),
			Description: "The milestone associated with the pull request."},
		{Name: "permalink", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Permalink"),
			Description: "Permanent URL for the pull request."},
		{Name: "published_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.PublishedAt").NullIfZero().Transform(convertTimestamp),
			Description: "Timestamp the pull request was published."},
		{Name: "revert_url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RevertUrl"),
			Description: "URL to revert the pull request."},
		{Name: "review_decision", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.ReviewDecision"),
			Description: "The current status of this pull request with respect to code review."},
		{Name: "state", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.State"),
			Description: "The current state of the pull request."},
		{Name: "title", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Title"),
			Description: "The title of the pull request."},
		{Name: "total_comments_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.TotalCommentsCount"),
			Description: "The number of comments on the pull request."},
		{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Description.UpdatedAt").NullIfZero().Transform(convertTimestamp),
			Description: "Timestamp when the pull request was last updated."},
		{Name: "url", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Url"),
			Description: "URL of the pull request."},
		{Name: "assignees", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.Assignees"),
			Description: "A list of Users assigned to the pull request."},
	}
}

func gitHubPullRequestColumns() []*plugin.Column {
	tableCols := []*plugin.Column{
		{Name: "base_ref", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.BaseRef"),
			Description: "The base ref associated with the pull request."},
		{Name: "head_ref", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.HeadRef"),
			Description: "The head ref associated with the pull request."},
		{Name: "merge_commit", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.MergeCommit"),
			Description: "The merge commit associated the pull request, null if not merged."},
		{Name: "suggested_reviewers", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.SuggestedReviewers"),
			Description: "Suggested reviewers for the pull request."},
		{Name: "can_apply_suggestion", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanApplySuggestion"),
			Description: "If true, current user can apply suggestions."},
		{Name: "can_close", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanClose"),
			Description: "If true, current user can close the pull request."},
		{Name: "can_delete_head_ref", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanDeleteHeadRef"),
			Description: "If true, current user can delete/restore head ref."},
		{Name: "can_disable_auto_merge", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanDisableAutoMerge"),
			Description: "If true, current user can disable auto-merge."},
		{Name: "can_edit_files", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanEditFiles"),
			Description: "If true, current user can edit files within this pull request."},
		{Name: "can_enable_auto_merge", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanEnableAutoMerge"),
			Description: "If true, current user can enable auto-merge."},
		{Name: "can_merge_as_admin", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanMergeAsAdmin"),
			Description: "If true, current user can bypass branch protections and merge the pull request immediately."},
		{Name: "can_react", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanReact"),
			Description: "If true, current user can react to the pull request."},
		{Name: "can_reopen", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanReopen"),
			Description: "If true, current user can reopen the pull request."},
		{Name: "can_subscribe", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanSubscribe"),
			Description: "If true, current user can subscribe to the pull request."},
		{Name: "can_update", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanUpdate"),
			Description: "If true, current user can update the pull request."},
		{Name: "can_update_branch", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.CanUpdateBranch"),
			Description: "If true, current user can update the head ref of the pull request by merging or rebasing the base ref."},
		{Name: "did_author", Type: proto.ColumnType_BOOL,
			Transform:   transform.FromField("Description.DidAuthor"),
			Description: "If true, current user authored the pull request."},
		{Name: "cannot_update_reasons", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.CannotUpdateReasons"),
			Description: "Reasons why the current user cannot update the pull request, if applicable."},
		{Name: "subscription", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Subscription"),
			Description: "Status of current users subscription to the pull request."},
		{Name: "labels_src", Type: proto.ColumnType_JSON,
			Transform:   transform.FromField("Description.LabelsSrc"),
			Description: "The first 100 labels associated to the pull request."},
		{Name: "labels", Type: proto.ColumnType_JSON,
			Description: "A map of labels for the pull request.",
			Transform:   transform.FromField("Description.Labels"),
		},
		{Name: "assignees_total_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.AssigneesTotalCount"),
			Description: "A count of users assigned to the pull request."},
		{Name: "labels_total_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.LabelsTotalCount"),
			Description: "A count of labels applied to the pull request."},
		{Name: "commits_total_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.CommitsTotalCount"),
			Description: "A count of commits in the pull request."},
		{Name: "review_requests_total_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.ReviewRequestsTotalCount"),
			Description: "A count of reviews requested on the pull request."},
		{Name: "reviews_total_count", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.ReviewsTotalCount"),
			Description: "A count of completed reviews on the pull request."},
	}

	return append(sharedPullRequestColumns(), tableCols...)
}

func tableGitHubPullRequest() *plugin.Table {
	return &plugin.Table{
		Name:        "github_pull_request",
		Description: "GitHub Pull requests let you tell others about changes you've pushed to a branch in a repository on GitHub. Once a pull request is opened, you can discuss and review the potential changes with collaborators and add follow-up commits before your changes are merged into the base branch.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPullRequest,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AllColumns([]string{"repository_full_name", "number"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.GetPullRequest,
		},
		Columns: commonColumns(gitHubPullRequestColumns()),
	}
}
