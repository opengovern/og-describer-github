package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func gitHubOrganizationExternalIdentityColumns() []*plugin.Column {
	return []*plugin.Column{
		{Name: "organization", Type: proto.ColumnType_STRING, Description: "The organization the external identity is associated with.",
			Transform: transform.FromField("Description.Organization")},
		{Name: "guid", Type: proto.ColumnType_STRING, Description: "Guid identifier for the external identity.",
			Transform: transform.FromField("Description.Guid")},
		{Name: "user_login", Type: proto.ColumnType_STRING, Description: "The GitHub user login.",
			Transform: transform.FromField("Description.User.Login")},
		{Name: "user_detail", Type: proto.ColumnType_JSON, Description: "The GitHub user details.",
			Transform: transform.FromField("Description.User")},
		{Name: "saml_identity", Type: proto.ColumnType_JSON, Description: "The external SAML identity.",
			Transform: transform.FromField("Description.SamlIdentity")},
		{Name: "scim_identity", Type: proto.ColumnType_JSON, Description: "The external SCIM identity.",
			Transform: transform.FromField("Description.ScimIdentity")},
		{Name: "organization_invitation", Type: proto.ColumnType_JSON, Description: "The invitation to the organization.",
			Transform: transform.FromField("Description.OrganizationInvitation")},
	}
}

func tableGitHubOrganizationExternalIdentity() *plugin.Table {
	return &plugin.Table{
		Name:        "github_organization_external_identity",
		Description: "GitHub members for a given organization. GitHub Users are user accounts in GitHub.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOrgExternalIdentity,
		},
		Columns: commonColumns(gitHubOrganizationExternalIdentityColumns()),
	}
}
