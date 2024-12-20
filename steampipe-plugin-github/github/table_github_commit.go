package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubCommit() *plugin.Table {
	return &plugin.Table{
		Name:        "github_commit",
		Description: "GitHub Commits bundle project files for download by users.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCommit,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AllColumns([]string{"repository_full_name", "id"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.GetCommit,
		},
		Columns: []*plugin.Column{
			{
				Name:        "repository_full_name",
				Type:        proto.ColumnType_STRING,
				Description: "Full name of the repository containing the commit.",
			},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "Unique identifier (SHA) of the commit.",
			},
			{
				Name:        "message",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Message"),
				Description: "Commit message.",
			},
			{
				Name:        "author_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Author.Name"),
				Description: "Name of the author of the commit.",
			},
			{
				Name:        "author_email",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Author.Email"),
				Description: "Email address of the author.",
			},
			{
				Name:        "additions",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Changes.Additions"),
				Description: "Number of lines added in the commit.",
			},
			{
				Name:        "deletions",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Changes.Deletions"),
				Description: "Number of lines deleted in the commit.",
			},
			{
				Name:        "total_changes",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Changes.Total"),
				Description: "Total number of changes (additions + deletions).",
			},
			{
				Name:        "is_verified",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.IsVerified"),
				Description: "Indicates if the commit is verified.",
			},
			{
				Name:        "comment_count",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CommentCount"),
				Description: "Number of comments on the commit.",
			},
			{
				Name:        "commit_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.HTMLURL"),
				Description: "URL of the commit on the repository.",
			},
			{
				Name:        "branch",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Target.Branch"),
				Description: "Branch where the commit resides.",
			},
			{
				Name:        "organization",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Target.Organization"),
				Description: "Organization that owns the repository.",
			},
			{
				Name:        "repository",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Target.Repository"),
				Description: "Repository where the commit resides.",
			},
			{
				Name:        "files",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Files"),
				Description: "List of files changed in the commit.",
			},
			{
				Name:        "node_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AdditionalDetails.NodeID"),
				Description: "Node ID of the commit.",
			},
			{
				Name:        "parents",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AdditionalDetails.Parents"),
				Description: "Parent commits of this commit.",
			},
			{
				Name:        "tree_sha",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AdditionalDetails.Tree.SHA"),
				Description: "SHA of the tree associated with the commit.",
			},
			{
				Name:        "verification_reason",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AdditionalDetails.VerificationDetails.Reason"),
				Description: "Reason for the verification status of the commit.",
			},
		},
	}
}
