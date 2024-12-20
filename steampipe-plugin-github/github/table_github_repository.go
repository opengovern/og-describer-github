package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubRepositoryColumns() []*plugin.Column {
	repoColumns := []*plugin.Column{
		{
			Name:        "full_name",
			Type:        proto.ColumnType_STRING,
			Description: "The full name of the repository, including the owner and repo name.",
			Transform:   transform.FromField("Description.NameWithOwner"),
		},
	}
	return append(repoColumns, sharedRepositoryColumns()...)
}

func sharedRepositoryColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_INT,
			Description: "The numeric ID of the repository.",
			Transform:   transform.FromField("Description.GitHubRepoID"),
		},
		{
			Name:        "node_id",
			Type:        proto.ColumnType_STRING,
			Description: "The node ID of the repository.",
			Transform:   transform.FromField("Description.NodeID"),
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the repository.",
			Transform:   transform.FromField("Description.Name"),
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "The description of the repository.",
			Transform:   transform.FromField("Description.Description"),
		},
		{
			Name:        "created_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the repository was created.",
			Transform:   transform.FromField("Description.CreatedAt").NullIfZero().Transform(convertTimestamp),
		},
		{
			Name:        "updated_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the repository was last updated.",
			Transform:   transform.FromField("Description.UpdatedAt").NullIfZero().Transform(convertTimestamp),
		},
		{
			Name:        "pushed_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the repository was last pushed to.",
			Transform:   transform.FromField("Description.PushedAt").NullIfZero().Transform(convertTimestamp),
		},
		{
			Name:        "is_active",
			Type:        proto.ColumnType_BOOL,
			Description: "If true, the repository is active.",
			Transform:   transform.FromField("Description.IsActive"),
		},
		{
			Name:        "is_empty",
			Type:        proto.ColumnType_BOOL,
			Description: "If true, the repository is empty.",
			Transform:   transform.FromField("Description.IsEmpty"),
		},
		{
			Name:        "is_fork",
			Type:        proto.ColumnType_BOOL,
			Description: "If true, the repository is a fork.",
			Transform:   transform.FromField("Description.IsFork"),
		},
		{
			Name:        "is_security_policy_enabled",
			Type:        proto.ColumnType_BOOL,
			Description: "If true, the repository has a security policy enabled.",
			Transform:   transform.FromField("Description.IsSecurityPolicyEnabled"),
		},
		{
			Name:        "owner",
			Type:        proto.ColumnType_JSON,
			Description: "The owner of the repository.",
			Transform:   transform.FromField("Description.Owner"),
		},
		{
			Name:        "homepage_url",
			Type:        proto.ColumnType_STRING,
			Description: "The external homepage URL of the repository, if set.",
			Transform:   transform.FromField("Description.HomepageURL"),
		},
		{
			Name:        "license_info",
			Type:        proto.ColumnType_JSON,
			Description: "The license associated with the repository.",
			Transform:   transform.FromField("Description.LicenseInfo"),
		},
		{
			Name:        "topics",
			Type:        proto.ColumnType_JSON,
			Description: "A list of topics associated with the repository.",
			Transform:   transform.FromField("Description.Topics"),
		},
		{
			Name:        "visibility",
			Type:        proto.ColumnType_STRING,
			Description: "The visibility level of the repository.",
			Transform:   transform.FromField("Description.Visibility"),
		},
		{
			Name:        "default_branch_ref",
			Type:        proto.ColumnType_JSON,
			Description: "Default branch reference information.",
			Transform:   transform.FromField("Description.DefaultBranchRef"),
		},
		{
			Name:        "permissions",
			Type:        proto.ColumnType_JSON,
			Description: "The permissions associated with the repository.",
			Transform:   transform.FromField("Description.Permissions"),
		},
		{
			Name:        "organization",
			Type:        proto.ColumnType_JSON,
			Description: "The organization associated with the repository.",
			Transform:   transform.FromField("Description.Organization"),
		},
		{
			Name:        "parent",
			Type:        proto.ColumnType_JSON,
			Description: "The parent repository of this fork, if applicable.",
			Transform:   transform.FromField("Description.Parent"),
		},
		{
			Name:        "source",
			Type:        proto.ColumnType_JSON,
			Description: "The source repository of this fork, if applicable.",
			Transform:   transform.FromField("Description.Source"),
		},
		{
			Name:        "languages",
			Type:        proto.ColumnType_JSON,
			Description: "Languages and their usage in the repository.",
			Transform:   transform.FromField("Description.Languages"),
		},
		{
			Name:        "repository_settings",
			Type:        proto.ColumnType_JSON,
			Description: "Settings of the repository.",
			Transform:   transform.FromField("Description.RepositorySettings"),
		},
		{
			Name:        "security_settings",
			Type:        proto.ColumnType_JSON,
			Description: "Security settings of the repository.",
			Transform:   transform.FromField("Description.SecuritySettings"),
		},
		{
			Name:        "repo_urls",
			Type:        proto.ColumnType_JSON,
			Description: "Various URLs associated with the repository.",
			Transform:   transform.FromField("Description.RepoURLs"),
		},
		{
			Name:        "metrics",
			Type:        proto.ColumnType_JSON,
			Description: "Metrics of the repository.",
			Transform:   transform.FromField("Description.Metrics"),
		},
	}
}

func tableGitHubRepository() *plugin.Table {
	return &plugin.Table{
		Name:        "github_repository",
		Description: "GitHub Repositories contain all of your project's files and each file's revision history.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRepository,
		},
		Columns: commonColumns(gitHubRepositoryColumns()),
	}
}
