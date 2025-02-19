package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubRepositoryEnvironmentColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "repository_full_name", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepoFullName"),
			Description: "The full name of the repository (login/repo-name)."},
		{Name: "id", Type: proto.ColumnType_INT,
			Transform:   transform.FromField("Description.Id"),
			Description: "The ID of the environment."},
		{Name: "node_id", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.NodeId"),
			Description: "The node ID of the environment."},
		{Name: "name", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Name"),
			Description: "The name of the environment."},
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
	}
}

func tableGitHubRepositoryEnvironment() *plugin.Table {
	return &plugin.Table{
		Name:        "github_repository_environment",
		Description: "GitHub Environments are named deployment targets, usually isolated for usage such as test, prod, staging, etc.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRepoEnvironment,
		},
		Columns: commonColumns(gitHubRepositoryEnvironmentColumns()),
	}
}
