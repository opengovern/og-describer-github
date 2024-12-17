package steampipe

import (
	"github.com/opengovern/og-describer-github/pkg/sdk/es"
)

var Map = map[string]string{
  "Github/Actions/Artifact": "github_actions_artifact",
  "Github/Actions/Repository/Runner": "github_actions_repository_runner",
  "Github/Actions/Repository/Secret": "github_actions_repository_secret",
  "Github/Actions/Repository/Workflow_run": "github_actions_repository_workflow_run",
  "Github/Blob": "github_blob",
  "Github/Branch": "github_branch",
  "Github/Branch/Protection": "github_branch_protection",
  "Github/Commit": "github_commit",
  "Github/Issue": "github_issue",
  "Github/License": "github_license",
  "Github/Organization": "github_organization",
  "Github/Organization/Collaborator": "github_organization_collaborators",
  "Github/Organization/DependabotAlert": "github_organization_dependabot_alert",
  "Github/Organization/ExternalIdentity": "github_organization_external_identity",
  "Github/Organization/Member": "github_organization_member",
  "Github/PullRequest": "github_pull_request",
  "Github/Release": "github_release",
  "Github/Repository": "github_repository",
  "Github/Repository/Collaborator": "github_repository_collaborator",
  "Github/Repository/DependabotAlert": "github_repository_dependabot_alert",
  "Github/Repository/Deployment": "github_repository_deployment",
  "Github/Repository/Environment": "github_repository_environment",
  "Github/Repository/Ruleset": "github_repository_ruleset",
  "Github/Repository/SBOM": "github_repository_sbom",
  "Github/Repository/VulnerabilityAlert": "github_repository_vulnerability_alert",
  "Github/Tag": "github_tag",
  "Github/Team": "github_team",
  "Github/Team/Member": "github_team_member",
  "Github/Team/Repository": "github_team_repository",
  "Github/Tree": "github_tree",
  "Github/User": "github_user",
  "Github/Workflow": "github_workflow",
  "Github/CodeOwner": "github_code_owner",
  "Github/Package/Container": "github_container_package",
  "Github/Package/Maven": "github_maven_package",
  "Github/Package/NPM": "github_npm_package",
  "Github/Package/Nuget": "github_nuget_package",
  "Github/ArtifactDockerFile": "github_artifact_dockerfile",
}

var DescriptionMap = map[string]interface{}{
  "Github/Actions/Artifact": opengovernance.Artifact{},
  "Github/Actions/Repository/Runner": opengovernance.Runner{},
  "Github/Actions/Repository/Secret": opengovernance.Secret{},
  "Github/Actions/Repository/Workflow_run": opengovernance.WorkflowRun{},
  "Github/Blob": opengovernance.Blob{},
  "Github/Branch": opengovernance.Branch{},
  "Github/Branch/Protection": opengovernance.BranchProtection{},
  "Github/Commit": opengovernance.Commit{},
  "Github/Issue": opengovernance.Issue{},
  "Github/License": opengovernance.License{},
  "Github/Organization": opengovernance.Organization{},
  "Github/Organization/Collaborator": opengovernance.OrgCollaborators{},
  "Github/Organization/DependabotAlert": opengovernance.OrgAlertDependabot{},
  "Github/Organization/ExternalIdentity": opengovernance.OrgExternalIdentity{},
  "Github/Organization/Member": opengovernance.OrgMembers{},
  "Github/PullRequest": opengovernance.PullRequest{},
  "Github/Release": opengovernance.Release{},
  "Github/Repository": opengovernance.Repository{},
  "Github/Repository/Collaborator": opengovernance.RepoCollaborators{},
  "Github/Repository/DependabotAlert": opengovernance.RepoAlertDependabot{},
  "Github/Repository/Deployment": opengovernance.RepoDeployment{},
  "Github/Repository/Environment": opengovernance.RepoEnvironment{},
  "Github/Repository/Ruleset": opengovernance.RepoRuleSet{},
  "Github/Repository/SBOM": opengovernance.RepoSBOM{},
  "Github/Repository/VulnerabilityAlert": opengovernance.RepoVulnerabilityAlert{},
  "Github/Tag": opengovernance.Tag{},
  "Github/Team": opengovernance.Team{},
  "Github/Team/Member": opengovernance.TeamMembers{},
  "Github/Team/Repository": opengovernance.TeamRepository{},
  "Github/Tree": opengovernance.Tree{},
  "Github/User": opengovernance.User{},
  "Github/Workflow": opengovernance.Workflow{},
  "Github/CodeOwner": opengovernance.CodeOwner{},
  "Github/Package/Container": opengovernance.ContainerPackage{},
  "Github/Package/Maven": opengovernance.PackageDetail{},
  "Github/Package/NPM": opengovernance.PackageDetail{},
  "Github/Package/Nuget": opengovernance.Package{},
  "Github/ArtifactDockerFile": opengovernance.ArtifactDockerFile{},
}

var ReverseMap = map[string]string{
  "github_actions_artifact": "Github/Actions/Artifact",
  "github_actions_repository_runner": "Github/Actions/Repository/Runner",
  "github_actions_repository_secret": "Github/Actions/Repository/Secret",
  "github_actions_repository_workflow_run": "Github/Actions/Repository/Workflow_run",
  "github_blob": "Github/Blob",
  "github_branch": "Github/Branch",
  "github_branch_protection": "Github/Branch/Protection",
  "github_commit": "Github/Commit",
  "github_issue": "Github/Issue",
  "github_license": "Github/License",
  "github_organization": "Github/Organization",
  "github_organization_collaborators": "Github/Organization/Collaborator",
  "github_organization_dependabot_alert": "Github/Organization/DependabotAlert",
  "github_organization_external_identity": "Github/Organization/ExternalIdentity",
  "github_organization_member": "Github/Organization/Member",
  "github_pull_request": "Github/PullRequest",
  "github_release": "Github/Release",
  "github_repository": "Github/Repository",
  "github_repository_collaborator": "Github/Repository/Collaborator",
  "github_repository_dependabot_alert": "Github/Repository/DependabotAlert",
  "github_repository_deployment": "Github/Repository/Deployment",
  "github_repository_environment": "Github/Repository/Environment",
  "github_repository_ruleset": "Github/Repository/Ruleset",
  "github_repository_sbom": "Github/Repository/SBOM",
  "github_repository_vulnerability_alert": "Github/Repository/VulnerabilityAlert",
  "github_tag": "Github/Tag",
  "github_team": "Github/Team",
  "github_team_member": "Github/Team/Member",
  "github_team_repository": "Github/Team/Repository",
  "github_tree": "Github/Tree",
  "github_user": "Github/User",
  "github_workflow": "Github/Workflow",
  "github_code_owner": "Github/CodeOwner",
  "github_container_package": "Github/Package/Container",
  "github_maven_package": "Github/Package/Maven",
  "github_npm_package": "Github/Package/NPM",
  "github_nuget_package": "Github/Package/Nuget",
  "github_artifact_dockerfile": "Github/ArtifactDockerFile",
}
