package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubTeamRepositoryColumns() []*plugin.Column {
	teamColumns := []*plugin.Column{
		{Name: "permission", Type: proto.ColumnType_STRING, Description: "The permission level the team has on the repository.",
			Transform: transform.FromQual("Description.Permission")},
		{Name: "team_id", Type: proto.ColumnType_INT, Description: "",
			Transform: transform.FromQual("Description.TeamID")},
		{Name: "repository_full_name", Type: proto.ColumnType_STRING, Description: "",
			Transform: transform.FromQual("Description.RepositoryFullName")},
		{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "",
			Transform: transform.FromQual("Description.CreatedAt")},
		{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "",
			Transform: transform.FromQual("Description.UpdatedAt")},
	}

	return teamColumns
}

func tableGitHubTeamRepository() *plugin.Table {
	return &plugin.Table{
		Name:        "github_team_repository",
		Description: "GitHub Repositories that a given team is associated with. GitHub Repositories contain all of your project's files and each file's revision history.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListTeamRepository,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "repository_full_name", Require: plugin.Required},
			},
			Hydrate: opengovernance.GetTeamRepository,
		},
		Columns: commonColumns(gitHubTeamRepositoryColumns()),
	}
}
