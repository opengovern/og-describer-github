package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubOrganizationMemberColumns() []*plugin.Column {
	tableCols := []*plugin.Column{
		{
			Name:        "organization",
			Type:        proto.ColumnType_STRING,
			Description: "The organization the member is associated with.",
			Transform:   transform.FromField("Description.Organization")},
		{
			Name:        "role",
			Type:        proto.ColumnType_STRING,
			Description: "The role this user has in the organization. Returns null if information is not available to viewer.",
			Transform:   transform.FromField("Description.Role")},
		{
			Name:      "has_two_factor_enabled",
			Type:      proto.ColumnType_BOOL,
			Transform: transform.FromField("Description.HasTwoFactorEnabled")},
		{
			Name:        "login",
			Type:        proto.ColumnType_STRING,
			Description: "",
			Transform:   transform.FromField("Description.Login")},
		{
			Name:        "login_id",
			Type:        proto.ColumnType_STRING,
			Description: "",
			Transform:   transform.FromField("Description.LoginID")},
		{
			Name:        "url",
			Type:        proto.ColumnType_STRING,
			Description: "",
			Transform:   transform.FromField("Description.URL")},
		{
			Name:        "email",
			Type:        proto.ColumnType_STRING,
			Description: "",
			Transform:   transform.FromField("Description.Email")},
		{
			Name:        "created_at",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "",
			Transform:   transform.FromField("Description.CreatedAt")},
		{
			Name:        "company",
			Type:        proto.ColumnType_STRING,
			Description: "",
			Transform:   transform.FromField("Description.Company")},
		{
			Name:        "status",
			Type:        proto.ColumnType_STRING,
			Description: "",
			Transform:   transform.FromField("Description.Status")},
	}

	return append(tableCols, sharedUserColumns()...)
}

func tableGitHubOrganizationMember() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_member",
		Description: "GitHub members for a given organization. GitHub Users are user accounts in GitHub.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrgMembers,
		},
		Columns: commonColumns(gitHubOrganizationMemberColumns()),
	}
}
