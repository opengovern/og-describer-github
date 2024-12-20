//go:generate go run ../../SDK/runnable/models/main.go --file $GOFILE --output ../../SDK/generated/resources_clients.go --type $PROVIDER

// Implement types for each resource

package model

import (
	"encoding/json"
	"time"

	goPipeline "github.com/buildkite/go-pipeline"
	"github.com/google/go-github/v55/github"
	"github.com/shurcooL/githubv4"
	steampipemodels "github.com/turbot/steampipe-plugin-github/github/models"
)

type Metadata struct{}

type ArtifactDescription struct {
	ID                 int64
	NodeID             string
	Name               string
	SizeInBytes        int64
	ArchiveDownloadURL string
	Expired            bool
	CreatedAt          github.Timestamp
	ExpiresAt          github.Timestamp
	RepoFullName       string
}

type RunnerDescription struct {
	*github.Runner
	RepoFullName string
}

type SecretDescription struct {
	*github.Secret
	RepoFullName string
}

type SimpleActor struct {
	Login  string `json:"login"`
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
	Type   string `json:"type"`
}

type SimpleRepo struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
}

type CommitRefWorkflow struct {
	ID string `json:"id"`
}

type WorkflowRunDescription struct {
	ID                  int                `json:"id"`
	Name                string             `json:"name"`
	HeadBranch          string             `json:"head_branch"`
	HeadSHA             string             `json:"head_sha"`
	Status              string             `json:"status"`
	Conclusion          string             `json:"conclusion"`
	HTMLURL             string             `json:"html_url"`
	WorkflowID          int                `json:"workflow_id"`
	RunNumber           int                `json:"run_number"`
	Event               string             `json:"event"`
	CreatedAt           string             `json:"created_at"`
	UpdatedAt           string             `json:"updated_at"`
	RunAttempt          int                `json:"run_attempt"`
	RunStartedAt        string             `json:"run_started_at"`
	Actor               *SimpleActor       `json:"triggering_actor,omitempty"`
	HeadCommit          *CommitRefWorkflow `json:"head_commit,omitempty"`
	Repository          *SimpleRepo        `json:"repository,omitempty"`
	HeadRepository      *SimpleRepo        `json:"head_repository,omitempty"`
	ReferencedWorkflows []interface{}      `json:"referenced_workflows,omitempty"`
	ArtifactCount       int                `json:"artifact_count"`
	Artifacts           []WorkflowArtifact `json:"artifacts,omitempty"`
}

type WorkflowRunsResponse struct {
	TotalCount   int                      `json:"total_count"`
	WorkflowRuns []WorkflowRunDescription `json:"workflow_runs"`
}

type WorkflowArtifact struct {
	ID                 int    `json:"id"`
	NodeID             string `json:"node_id"`
	Name               string `json:"name"`
	SizeInBytes        int    `json:"size_in_bytes"`
	URL                string `json:"url"`
	ArchiveDownloadURL string `json:"archive_download_url"`
	Expired            bool   `json:"expired"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	ExpiresAt          string `json:"expires_at"`
}

type ArtifactsResponse struct {
	TotalCount int                `json:"total_count"`
	Artifacts  []WorkflowArtifact `json:"artifacts"`
}

//type WorkflowRunDescription struct {
//	ID                 int64
//	Name               string
//	NodeID             string
//	HeadBranch         string
//	HeadSHA            string
//	RunNumber          int
//	RunAttempt         int
//	Event              string
//	DisplayTitle       string
//	Status             string
//	Conclusion         string
//	WorkflowID         int64
//	CheckSuiteID       int64
//	CheckSuiteNodeID   string
//	URL                string
//	HTMLURL            string
//	PullRequests       []*github.PullRequest
//	CreatedAt          github.Timestamp
//	UpdatedAt          github.Timestamp
//	RunStartedAt       github.Timestamp
//	JobsURL            string
//	LogsURL            string
//	CheckSuiteURL      string
//	ArtifactsURL       string
//	CancelURL          string
//	RerunURL           string
//	PreviousAttemptURL string
//	HeadCommit         *github.HeadCommit
//	WorkflowURL        string
//	Repository         *github.Repository
//	HeadRepository     *github.Repository
//	Actor              *github.User
//	TriggeringActor    *github.User
//	RepoFullName       string
//}

type AuditLogDescription struct {
	ID            string
	CreatedAt     github.Timestamp
	Organization  string
	Phrase        string
	Include       string
	Action        string
	Actor         string
	ActorLocation *github.ActorLocation
	Team          string
	UserLogin     string
	Repo          string
	Data          *github.AuditEntryData
}

type BlobDescription struct {
	*github.Blob
	RepoFullName string
}

type BranchDescription struct {
	RepoFullName         string
	Name                 string
	Commit               steampipemodels.BaseCommit
	BranchProtectionRule steampipemodels.BranchProtectionRule
	Protected            bool
}

type BranchApp struct {
	Name string
	Slug string
}

type BranchTeam struct {
	Name string
	Slug string
}

type BranchUser struct {
	Name  string
	Login string
}

type BranchProtectionDescription struct {
	steampipemodels.BranchProtectionRule
	RepoFullName                    string
	CreatorLogin                    string
	MatchingBranches                int
	PushAllowanceApps               []BranchApp
	PushAllowanceTeams              []BranchTeam
	PushAllowanceUsers              []BranchUser
	BypassForcePushAllowanceApps    []BranchApp
	BypassForcePushAllowanceTeams   []BranchTeam
	BypassForcePushAllowanceUsers   []BranchUser
	BypassPullRequestAllowanceApps  []BranchApp
	BypassPullRequestAllowanceTeams []BranchTeam
	BypassPullRequestAllowanceUsers []BranchUser
}

type Parents []struct {
	SHA string `json:"sha"`
}

type Tree struct {
	SHA string `json:"sha"`
}

type VerificationDetails struct {
	Reason     string  `json:"reason"`
	Signature  *string `json:"signature"`
	VerifiedAt *string `json:"verified_at"`
}

type AdditionalDetails struct {
	NodeID              string              `json:"node_id"`
	Parents             Parents             `json:"parents"`
	Tree                Tree                `json:"tree"`
	VerificationDetails VerificationDetails `json:"verification_details"`
}

type Author struct {
	Email   string `json:"email"`
	HTMLURL string `json:"html_url"`
	ID      int    `json:"id"`
	Login   string `json:"login"`
	Name    string `json:"name"`
	NodeID  string `json:"node_id"`
	Type    string `json:"type"`
}

type Changes struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}

type File struct {
	Additions int    `json:"additions"`
	Changes   int    `json:"changes"`
	Deletions int    `json:"deletions"`
	Filename  string `json:"filename"`
	SHA       string `json:"sha"`
	Status    string `json:"status"`
}

type Target struct {
	Branch       string `json:"branch"`
	Organization string `json:"organization"`
	Repository   string `json:"repository"`
}

type CommitDescription struct {
	AdditionalDetails AdditionalDetails `json:"additional_details"`
	Author            Author            `json:"author"`
	Changes           Changes           `json:"changes"`
	CommentCount      int               `json:"comment_count"`
	Date              string            `json:"date"`
	Files             []File            `json:"files"`
	HTMLURL           string            `json:"html_url"`
	ID                string            `json:"id"`
	IsVerified        bool              `json:"is_verified"`
	Message           string            `json:"message"`
	PullRequests      []int             `json:"pull_requests"`
	Target            Target            `json:"target"`
}

type IssueDescription struct {
	RepositoryFullName      string
	Id                      int
	NodeId                  string
	Number                  int
	ActiveLockReason        githubv4.LockReason
	Author                  steampipemodels.Actor
	AuthorLogin             string
	AuthorAssociation       githubv4.CommentAuthorAssociation
	Body                    string
	BodyUrl                 string
	Closed                  bool
	ClosedAt                steampipemodels.NullableTime
	CreatedAt               steampipemodels.NullableTime
	CreatedViaEmail         bool
	Editor                  steampipemodels.Actor
	FullDatabaseId          string
	IncludesCreatedEdit     bool
	IsPinned                bool
	IsReadByUser            bool
	LastEditedAt            steampipemodels.NullableTime
	Locked                  bool
	Milestone               steampipemodels.Milestone
	PublishedAt             steampipemodels.NullableTime
	State                   githubv4.IssueState
	StateReason             githubv4.IssueStateReason
	Title                   string
	UpdatedAt               steampipemodels.NullableTime
	Url                     string
	UserCanClose            bool
	UserCanReact            bool
	UserCanReopen           bool
	UserCanSubscribe        bool
	UserCanUpdate           bool
	UserCannotUpdateReasons []githubv4.CommentCannotUpdateReason
	UserDidAuthor           bool
	UserSubscription        githubv4.SubscriptionState
	CommentsTotalCount      int
	LabelsTotalCount        int
	LabelsSrc               []steampipemodels.Label
	Labels                  map[string]steampipemodels.Label
	AssigneesTotalCount     int
	Assignees               []steampipemodels.BaseUser
}

type IssueCommentDescription struct {
	steampipemodels.IssueComment
	RepoFullName string
	Number       int
	AuthorLogin  string
	EditorLogin  string
}

type LicenseDescription struct {
	steampipemodels.License
}

type OrganizationDescription struct {
	steampipemodels.Organization
	Hooks                                []*github.Hook
	BillingEmail                         string
	TwoFactorRequirementEnabled          bool
	DefaultRepoPermission                string
	MembersAllowedRepositoryCreationType string
	MembersCanCreateInternalRepos        bool
	MembersCanCreatePages                bool
	MembersCanCreatePrivateRepos         bool
	MembersCanCreatePublicRepos          bool
	MembersCanCreateRepos                bool
	MembersCanForkPrivateRepos           bool
	PlanFilledSeats                      int
	PlanName                             string
	PlanPrivateRepos                     int
	PlanSeats                            int
	PlanSpace                            int
	Followers                            int
	Following                            int
	Collaborators                        int
	HasOrganizationProjects              bool
	HasRepositoryProjects                bool
	WebCommitSignoffRequired             bool
	MembersWithRoleTotalCount            int
	PackagesTotalCount                   int
	PinnableItemsTotalCount              int
	PinnedItemsTotalCount                int
	ProjectsTotalCount                   int
	ProjectsV2TotalCount                 int
	SponsoringTotalCount                 int
	SponsorsTotalCount                   int
	TeamsTotalCount                      int
	PrivateRepositoriesTotalCount        int
	PublicRepositoriesTotalCount         int
	RepositoriesTotalCount               int
	RepositoriesTotalDiskUsage           int
}

type OrgCollaboratorsDescription struct {
	Organization   string
	Affiliation    string
	RepositoryName githubv4.String
	Permission     githubv4.RepositoryPermission
	UserLogin      steampipemodels.CollaboratorLogin
}

type OrgAlertDependabotDescription struct {
	AlertNumber                 int
	State                       string
	DependencyPackageEcosystem  string
	DependencyPackageName       string
	DependencyManifestPath      string
	DependencyScope             string
	SecurityAdvisoryGHSAID      string
	SecurityAdvisoryCVEID       string
	SecurityAdvisorySummary     string
	SecurityAdvisoryDescription string
	SecurityAdvisorySeverity    string
	SecurityAdvisoryCVSSScore   *float64
	SecurityAdvisoryCVSSVector  string
	SecurityAdvisoryCWEs        []string
	SecurityAdvisoryPublishedAt github.Timestamp
	SecurityAdvisoryUpdatedAt   github.Timestamp
	SecurityAdvisoryWithdrawnAt github.Timestamp
	URL                         string
	HTMLURL                     string
	CreatedAt                   github.Timestamp
	UpdatedAt                   github.Timestamp
	DismissedAt                 github.Timestamp
	DismissedReason             string
	DismissedComment            string
	FixedAt                     github.Timestamp
}

type OrgExternalIdentityDescription struct {
	steampipemodels.OrganizationExternalIdentity
	Organization string
	UserLogin    string
	UserDetail   steampipemodels.BasicUser
}

type OrgMembersDescription struct {
	steampipemodels.User
	Organization        string
	HasTwoFactorEnabled *bool
	Role                *string
}

type PullRequestDescription struct {
	RepoFullName             string
	Id                       int
	NodeId                   string
	Number                   int
	ActiveLockReason         githubv4.LockReason
	Additions                int
	Author                   steampipemodels.Actor
	AuthorAssociation        githubv4.CommentAuthorAssociation
	BaseRefName              string
	Body                     string
	ChangedFiles             int
	ChecksUrl                string
	Closed                   bool
	ClosedAt                 steampipemodels.NullableTime
	CreatedAt                steampipemodels.NullableTime
	CreatedViaEmail          bool
	Deletions                int
	Editor                   steampipemodels.Actor
	HeadRefName              string
	HeadRefOid               string
	IncludesCreatedEdit      bool
	IsCrossRepository        bool
	IsDraft                  bool
	IsReadByUser             bool
	LastEditedAt             steampipemodels.NullableTime
	Locked                   bool
	MaintainerCanModify      bool
	Mergeable                githubv4.MergeableState
	Merged                   bool
	MergedAt                 steampipemodels.NullableTime
	MergedBy                 steampipemodels.Actor
	Milestone                steampipemodels.Milestone
	Permalink                string
	PublishedAt              steampipemodels.NullableTime
	RevertUrl                string
	ReviewDecision           githubv4.PullRequestReviewDecision
	State                    githubv4.PullRequestState
	Title                    string
	TotalCommentsCount       int
	UpdatedAt                steampipemodels.NullableTime
	Url                      string
	Assignees                []steampipemodels.BaseUser
	BaseRef                  *steampipemodels.BasicRef
	HeadRef                  *steampipemodels.BasicRef
	MergeCommit              *steampipemodels.BasicCommit
	SuggestedReviewers       []steampipemodels.SuggestedReviewer
	CanApplySuggestion       bool
	CanClose                 bool
	CanDeleteHeadRef         bool
	CanDisableAutoMerge      bool
	CanEditFiles             bool
	CanEnableAutoMerge       bool
	CanMergeAsAdmin          bool
	CanReact                 bool
	CanReopen                bool
	CanSubscribe             bool
	CanUpdate                bool
	CanUpdateBranch          bool
	DidAuthor                bool
	CannotUpdateReasons      []githubv4.CommentCannotUpdateReason
	Subscription             githubv4.SubscriptionState
	LabelsSrc                []steampipemodels.Label
	Labels                   map[string]steampipemodels.Label
	CommitsTotalCount        int
	ReviewRequestsTotalCount int
	ReviewsTotalCount        int
	LabelsTotalCount         int
	AssigneesTotalCount      int
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SPDXID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

type Permissions struct {
	Admin    bool `json:"admin"`
	Maintain bool `json:"maintain"`
	Push     bool `json:"push"`
	Triage   bool `json:"triage"`
	Pull     bool `json:"pull"`
}

type StatusObj struct {
	Status string `json:"status"`
}

type RepoDetail struct {
	ID                        int                    `json:"id"`
	NodeID                    string                 `json:"node_id"`
	Name                      string                 `json:"name"`
	FullName                  string                 `json:"full_name"`
	Private                   bool                   `json:"private"`
	Owner                     *Owner                 `json:"owner"`
	HTMLURL                   string                 `json:"html_url"`
	Description               *string                `json:"description"`
	Fork                      bool                   `json:"fork"`
	CreatedAt                 string                 `json:"created_at"`
	UpdatedAt                 string                 `json:"updated_at"`
	PushedAt                  string                 `json:"pushed_at"`
	GitURL                    string                 `json:"git_url"`
	SSHURL                    string                 `json:"ssh_url"`
	CloneURL                  string                 `json:"clone_url"`
	SVNURL                    string                 `json:"svn_url"`
	Homepage                  *string                `json:"homepage"`
	Size                      int                    `json:"size"`
	StargazersCount           int                    `json:"stargazers_count"`
	WatchersCount             int                    `json:"watchers_count"`
	Languages                 *string                `json:"languages"` // original string language field
	HasIssues                 bool                   `json:"has_issues"`
	HasProjects               bool                   `json:"has_projects"`
	HasDownloads              bool                   `json:"has_downloads"`
	HasWiki                   bool                   `json:"has_wiki"`
	HasPages                  bool                   `json:"has_pages"`
	HasDiscussions            bool                   `json:"has_discussions"`
	ForksCount                int                    `json:"forks_count"`
	MirrorURL                 *string                `json:"mirror_url"`
	Archived                  bool                   `json:"archived"`
	Disabled                  bool                   `json:"disabled"`
	OpenIssuesCount           int                    `json:"open_issues_count"`
	License                   *License               `json:"license"`
	AllowForking              bool                   `json:"allow_forking"`
	IsTemplate                bool                   `json:"is_template"`
	WebCommitSignoffRequired  bool                   `json:"web_commit_signoff_required"`
	Topics                    []string               `json:"topics"`
	Visibility                string                 `json:"visibility"`
	DefaultBranch             string                 `json:"default_branch"`
	Permissions               *Permissions           `json:"permissions"`
	AllowSquashMerge          bool                   `json:"allow_squash_merge"`
	AllowMergeCommit          bool                   `json:"allow_merge_commit"`
	AllowRebaseMerge          bool                   `json:"allow_rebase_merge"`
	AllowAutoMerge            bool                   `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool                   `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool                   `json:"allow_update_branch"`
	UseSquashPRTitleAsDefault bool                   `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitMessage  string                 `json:"squash_merge_commit_message"`
	SquashMergeCommitTitle    string                 `json:"squash_merge_commit_title"`
	MergeCommitMessage        string                 `json:"merge_commit_message"`
	MergeCommitTitle          string                 `json:"merge_commit_title"`
	CustomProperties          map[string]interface{} `json:"custom_properties"`
	Organization              *Organization          `json:"organization"`
	Parent                    *RepoDetail            `json:"parent"`
	Source                    *RepoDetail            `json:"source"`
	NetworkCount              int                    `json:"network_count"`
	SubscribersCount          int                    `json:"subscribers_count"`
	BlankIssuesEnabled        bool                   `json:"blank_issues_enabled"`
	Locked                    bool                   `json:"locked"`

	SecurityAndAnalysis *struct {
		SecretScanning                    *StatusObj `json:"secret_scanning"`
		SecretScanningPushProtection      *StatusObj `json:"secret_scanning_push_protection"`
		DependabotSecurityUpdates         *StatusObj `json:"dependabot_security_updates"`
		SecretScanningNonProviderPatterns *StatusObj `json:"secret_scanning_non_provider_patterns"`
		SecretScanningValidityChecks      *StatusObj `json:"secret_scanning_validity_checks"`
	} `json:"security_and_analysis"`
}

type RepositorySettings struct {
	HasDiscussionsEnabled     bool                   `json:"has_discussions_enabled"`
	HasIssuesEnabled          bool                   `json:"has_issues_enabled"`
	HasProjectsEnabled        bool                   `json:"has_projects_enabled"`
	HasWikiEnabled            bool                   `json:"has_wiki_enabled"`
	MergeCommitAllowed        bool                   `json:"merge_commit_allowed"`
	MergeCommitMessage        string                 `json:"merge_commit_message"`
	MergeCommitTitle          string                 `json:"merge_commit_title"`
	SquashMergeAllowed        bool                   `json:"squash_merge_allowed"`
	SquashMergeCommitMessage  string                 `json:"squash_merge_commit_message"`
	SquashMergeCommitTitle    string                 `json:"squash_merge_commit_title"`
	HasDownloads              bool                   `json:"has_downloads"`
	HasPages                  bool                   `json:"has_pages"`
	WebCommitSignoffRequired  bool                   `json:"web_commit_signoff_required"`
	MirrorURL                 *string                `json:"mirror_url"`
	AllowAutoMerge            bool                   `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool                   `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool                   `json:"allow_update_branch"`
	UseSquashPRTitleAsDefault bool                   `json:"use_squash_pr_title_as_default"`
	CustomProperties          map[string]interface{} `json:"custom_properties"`
	ForkingAllowed            bool                   `json:"forking_allowed"`
	IsTemplate                bool                   `json:"is_template"`
	AllowRebaseMerge          bool                   `json:"allow_rebase_merge"`
	// Renamed fields:
	Archived bool `json:"archived"`
	Disabled bool `json:"disabled"`
	Locked   bool `json:"locked"`
}

type SecuritySettings struct {
	VulnerabilityAlertsEnabled               bool `json:"vulnerability_alerts_enabled"`
	SecretScanningEnabled                    bool `json:"secret_scanning_enabled"`
	SecretScanningPushProtectionEnabled      bool `json:"secret_scanning_push_protection_enabled"`
	DependabotSecurityUpdatesEnabled         bool `json:"dependabot_security_updates_enabled"`
	SecretScanningNonProviderPatternsEnabled bool `json:"secret_scanning_non_provider_patterns_enabled"`
	SecretScanningValidityChecksEnabled      bool `json:"secret_scanning_validity_checks_enabled"`

	// New field
	PrivateVulnerabilityReportingEnabled bool `json:"private_vulnerability_reporting_enabled"`
}

type RepoURLs struct {
	GitURL   string `json:"git_url"`
	SSHURL   string `json:"ssh_url"`
	CloneURL string `json:"clone_url"`
	SVNURL   string `json:"svn_url"`
	HTMLURL  string `json:"html_url"`
}
type Owner struct {
	Login   string `json:"login"`
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`
	Type    string `json:"type"`
}

type Organization struct {
	Login        string `json:"login"`
	ID           int    `json:"id"`
	NodeID       string `json:"node_id"`
	HTMLURL      string `json:"html_url"`
	Type         string `json:"type"`
	UserViewType string `json:"user_view_type"`
	SiteAdmin    bool   `json:"site_admin"`
}

type Metrics struct {
	Stargazers   int `json:"stargazers"`
	Forks        int `json:"forks"`
	Subscribers  int `json:"subscribers"`
	Size         int `json:"size"`
	Tags         int `json:"tags"`
	Commits      int `json:"commits"`
	Issues       int `json:"issues"`
	OpenIssues   int `json:"open_issues"`
	Branches     int `json:"branches"`
	PullRequests int `json:"pull_requests"`
	Releases     int `json:"releases"`
}

type RepositoryDescription struct {
	GitHubRepoID            int                    `json:"id"`
	NodeID                  string                 `json:"node_id"`
	Name                    string                 `json:"name"`
	NameWithOwner           string                 `json:"name_with_owner"`
	Description             *string                `json:"description"`
	CreatedAt               string                 `json:"created_at"`
	UpdatedAt               string                 `json:"updated_at"`
	PushedAt                string                 `json:"pushed_at"`
	IsActive                bool                   `json:"is_active"`
	IsEmpty                 bool                   `json:"is_empty"`
	IsFork                  bool                   `json:"is_fork"`
	IsSecurityPolicyEnabled bool                   `json:"is_security_policy_enabled"`
	Owner                   *Owner                 `json:"owner"`
	HomepageURL             *string                `json:"homepage_url"`
	LicenseInfo             json.RawMessage        `json:"license_info"`
	Topics                  []string               `json:"topics"`
	Visibility              string                 `json:"visibility"`
	DefaultBranchRef        json.RawMessage        `json:"default_branch_ref"`
	Permissions             *Permissions           `json:"permissions"`
	Organization            *Organization          `json:"organization"`
	Parent                  *RepositoryDescription `json:"parent"`
	Source                  *RepositoryDescription `json:"source"`
	Languages               map[string]int         `json:"language"`
	RepositorySettings      RepositorySettings     `json:"repo_settings"`
	SecuritySettings        SecuritySettings       `json:"security_settings"`
	RepoURLs                RepoURLs               `json:"repo_urls"`
	Metrics                 Metrics                `json:"metrics"`
}

type MinimalRepoInfo struct {
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
	Disabled bool   `json:"disabled"`
	Owner    struct {
		Login string `json:"login"`
	} `json:"owner"`
}

//type RepoDescription struct {
//	github.Repository
//	ID                            int
//	NodeID                        string
//	Name                          string
//	NameWithOwner                 string
//	Description                   string
//	IsArchived                    bool
//	IsEmpty                       bool
//	IsFork                        bool
//	IsSecurityPolicyEnabled       bool
//	OwnerLogin                    string
//	LicenseInfo                   steampipemodels.BasicLicense
//	Visibility                    githubv4.RepositoryVisibility
//	Topics                        []string
//	DefaultBranchRef              steampipemodels.BasicRefWithBranchProtectionRule
//	YourPermission                githubv4.RepositoryPermission
//	CreatedAt                     steampipemodels.NullableTime
//	UpdatedAt                     steampipemodels.NullableTime
//	PushedAt                      steampipemodels.NullableTime
//	AllowUpdateBranch             bool
//	ArchivedAt                    steampipemodels.NullableTime
//	AutoMergeAllowed              bool
//	CodeOfConduct                 steampipemodels.RepositoryCodeOfConduct
//	ContactLinks                  []steampipemodels.RepositoryContactLink
//	DeleteBranchOnMerge           bool
//	DiskUsage                     int
//	ForkCount                     int
//	ForkingAllowed                bool
//	FundingLinks                  []steampipemodels.RepositoryFundingLinks
//	HasDiscussionsEnabled         bool
//	HasIssuesEnabled              bool
//	HasProjectsEnabled            bool
//	HasVulnerabilityAlertsEnabled bool
//	HasWikiEnabled                bool
//	HomepageURL                   string
//	InteractionAbility            steampipemodels.RepositoryInteractionAbility
//	IsBlankIssuesEnabled          bool
//	IsDisabled                    bool
//	IsInOrganization              bool
//	IsLocked                      bool
//	IsMirror                      bool
//	IsPrivate                     bool
//	IsTemplate                    bool
//	IsUserConfigurationRepository bool
//	IssueTemplates                []steampipemodels.IssueTemplate
//	LockReason                    githubv4.LockReason
//	MergeCommitAllowed            bool
//	MergeCommitMessage            githubv4.MergeCommitMessage
//	MergeCommitTitle              githubv4.MergeCommitTitle
//	MirrorURL                     string
//	OpenGraphImageURL             string
//	PrimaryLanguage               steampipemodels.Language
//	ProjectsURL                   string
//	PullRequestTemplates          []steampipemodels.PullRequestTemplate
//	RebaseMergeAllowed            bool
//	SecurityPolicyURL             string
//	SquashMergeAllowed            bool
//	SquashMergeCommitMessage      githubv4.SquashMergeCommitMessage
//	SquashMergeCommitTitle        githubv4.SquashMergeCommitTitle
//	SSHURL                        string
//	StargazerCount                int
//	URL                           string
//	//UsesCustomOpenGraphImage      bool
//	//CanAdminister                 bool
//	//CanCreateProjects             bool
//	//CanSubscribe                  bool
//	//CanUpdateTopics               bool
//	//HasStarred                    bool
//	PossibleCommitEmails []string
//	//Subscription                  githubv4.SubscriptionState
//	WebCommitSignOffRequired   bool
//	RepositoryTopicsTotalCount int
//	OpenIssuesTotalCount       int
//	WatchersTotalCount         int
//	Hooks                      []*github.Hook
//	SubscribersCount           int
//	HasDownloads               bool
//	HasPages                   bool
//	NetworkCount               int
//}

type ReleaseDescription struct {
	github.RepositoryRelease
	RepositoryFullName string
}

type RepoCollaboratorsDescription struct {
	Affiliation  string
	RepoFullName string
	Permission   githubv4.RepositoryPermission
	UserLogin    string
}

type RepoAlertDependabotDescription struct {
	RepoFullName                string
	AlertNumber                 int
	State                       string
	DependencyPackageEcosystem  string
	DependencyPackageName       string
	DependencyManifestPath      string
	DependencyScope             string
	SecurityAdvisoryGHSAID      string
	SecurityAdvisoryCVEID       string
	SecurityAdvisorySummary     string
	SecurityAdvisoryDescription string
	SecurityAdvisorySeverity    string
	SecurityAdvisoryCVSSScore   *float64
	SecurityAdvisoryCVSSVector  string
	SecurityAdvisoryCWEs        []string
	SecurityAdvisoryPublishedAt github.Timestamp
	SecurityAdvisoryUpdatedAt   github.Timestamp
	SecurityAdvisoryWithdrawnAt github.Timestamp
	URL                         string
	HTMLURL                     string
	CreatedAt                   github.Timestamp
	UpdatedAt                   github.Timestamp
	DismissedAt                 github.Timestamp
	DismissedReason             string
	DismissedComment            string
	FixedAt                     github.Timestamp
}

type RepoDeploymentDescription struct {
	steampipemodels.Deployment
	RepoFullName string
}

type RepoEnvironmentDescription struct {
	steampipemodels.Environment
	RepoFullName string
}

type RepoRuleSetDescription struct {
	steampipemodels.Ruleset
	RepoFullName string
}

type RepoSBOMDescription struct {
	RepositoryFullName string
	SPDXID             string
	SPDXVersion        string
	CreationInfo       *github.CreationInfo
	Name               string
	DataLicense        string
	DocumentDescribes  []string
	DocumentNamespace  string
	Packages           []*github.RepoDependencies
}

type RepoVulnerabilityAlertDescription struct {
	RepositoryFullName         string
	Number                     int
	NodeID                     string
	AutoDismissedAt            steampipemodels.NullableTime
	CreatedAt                  steampipemodels.NullableTime
	DependencyScope            githubv4.RepositoryVulnerabilityAlertDependencyScope
	DismissComment             string
	DismissReason              string
	DismissedAt                steampipemodels.NullableTime
	Dismisser                  steampipemodels.BasicUser
	FixedAt                    steampipemodels.NullableTime
	State                      githubv4.RepositoryVulnerabilityAlertState
	SecurityAdvisory           steampipemodels.SecurityAdvisory
	SecurityVulnerability      steampipemodels.SecurityVulnerability
	VulnerableManifestFilename string
	VulnerableManifestPath     string
	VulnerableRequirements     string
	Severity                   githubv4.SecurityAdvisorySeverity
	CvssScore                  float64
}

type SearchCodeDescription struct {
	*github.CodeResult
	RepoFullName string
	Query        string
}

type SearchCommitDescription struct {
	*github.CommitResult
	RepoFullName string
	Query        string
}

type SearchIssueDescription struct {
	IssueDescription
	RepoFullName string
	Query        string
	TextMatches  []steampipemodels.TextMatch
}

type StarDescription struct {
	RepoFullName string
	StarredAt    steampipemodels.NullableTime
	Url          string
}

type StargazerDescription struct {
	RepoFullName string
	StarredAt    steampipemodels.NullableTime
	UserLogin    string
	UserDetail   steampipemodels.BasicUser
}

type TagDescription struct {
	RepositoryFullName string
	Name               string
	TaggerDate         time.Time
	TaggerName         string
	TaggerLogin        string
	Message            string
	Commit             steampipemodels.BaseCommit
}

type ParentTeam struct {
	Id     int
	NodeId string
	Name   string
	Slug   string
}

type TeamDescription struct {
	Organization           string
	Slug                   string
	Name                   string
	ID                     int
	NodeID                 string
	Description            string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	CombinedSlug           string
	ParentTeam             ParentTeam
	Privacy                string
	AncestorsTotalCount    int
	ChildTeamsTotalCount   int
	DiscussionsTotalCount  int
	InvitationsTotalCount  int
	MembersTotalCount      int
	ProjectsV2TotalCount   int
	RepositoriesTotalCount int
	URL                    string
	AvatarURL              string
	DiscussionsURL         string
	EditTeamURL            string
	MembersURL             string
	NewTeamURL             string
	RepositoriesURL        string
	TeamsURL               string
	CanAdminister          bool
	CanSubscribe           bool
	Subscription           string
}

type TeamMembersDescription struct {
	steampipemodels.User
	Organization string
	Slug         string
	Role         githubv4.TeamMemberRole
}

type TrafficViewDailyDescription struct {
	*github.TrafficData
	RepositoryFullName string
}

type TrafficViewWeeklyDescription struct {
	*github.TrafficData
	RepositoryFullName string
}

type TreeDescription struct {
	TreeSHA            string
	RepositoryFullName string
	Recursive          bool
	Truncated          bool
	SHA                string
	Path               string
	Mode               string
	Type               string
	Size               int
	URL                string
}

type UserDescription struct {
	steampipemodels.User
	RepositoriesTotalDiskUsage    int
	FollowersTotalCount           int
	FollowingTotalCount           int
	PublicRepositoriesTotalCount  int
	PrivateRepositoriesTotalCount int
	PublicGistsTotalCount         int
	IssuesTotalCount              int
	OrganizationsTotalCount       int
	PublicKeysTotalCount          int
	OpenPullRequestsTotalCount    int
	MergedPullRequestsTotalCount  int
	ClosedPullRequestsTotalCount  int
	PackagesTotalCount            int
	PinnedItemsTotalCount         int
	SponsoringTotalCount          int
	SponsorsTotalCount            int
	StarredRepositoriesTotalCount int
	WatchingTotalCount            int
}

type WorkflowDescription struct {
	ID                      *int64
	NodeID                  *string
	Name                    *string
	Path                    *string
	State                   *string
	CreatedAt               *github.Timestamp
	UpdatedAt               *github.Timestamp
	URL                     *string
	HTMLURL                 *string
	BadgeURL                *string
	RepositoryFullName      string
	WorkFlowFileContent     string
	WorkFlowFileContentJson *github.RepositoryContent
	Pipeline                *goPipeline.Pipeline
}

type CodeOwnerDescription struct {
	RepositoryFullName string
	LineNumber         int64
	Pattern            string
	Users              []string
	Teams              []string
	PreComments        []string
	LineComment        string
}

type OwnerLogin struct {
	Login string `json:"login"`
}

type Package struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	PackageType string     `json:"package_type"`
	Visibility  string     `json:"visibility"`
	HTMLURL     string     `json:"html_url"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	Owner       OwnerLogin `json:"owner"`
}

type ContainerMetadata struct {
	Container struct {
		Tags []string `json:"tags"`
	} `json:"container"`
}

type ContainerPackageDescription struct {
	ID             int               `json:"id"`
	Digest         string            `json:"digest"`
	PackageURI     string            `json:"package_uri"`
	PackageHTMLURL string            `json:"package_html_url"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	HTMLURL        string            `json:"html_url"`
	Name           string            `json:"name"`
	MediaType      string            `json:"media_type"`
	TotalSize      int64             `json:"total_size"`
	Metadata       ContainerMetadata `json:"metadata"`
	Manifest       interface{}       `json:"manifest"`

	// -- Add these new fields/methods: --

	// The GitHub "version.Name" or tag,
	// if you want to store that separately from the real Docker digest.
	GHVersionName string `json:"gh_version_name,omitempty"`

	// When deduplicating, any subsequent tags for the same (ID,digest)
	// can be appended here.
	AdditionalPackageURIs []string `json:"additional_package_uris,omitempty"`
}

// Provide a helper so that code calling `cpd.ActualDigest()` compiles:
func (c ContainerPackageDescription) ActualDigest() string {
	// If you want the *Docker* digest, we simply return `c.Digest`
	// because your code is already overwriting .Digest with the real Docker digest
	return c.Digest
}

type PackageVersion struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	PackageHTMLURL string            `json:"package_html_url"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	HTMLURL        string            `json:"html_url"`
	Metadata       ContainerMetadata `json:"metadata"`
}

type OwnerDetail struct {
	Login        string `json:"login"`
	ID           int    `json:"id,omitempty"`
	NodeID       string `json:"node_id,omitempty"`
	HTMLURL      string `json:"html_url,omitempty"`
	Type         string `json:"type,omitempty"`
	UserViewType string `json:"user_view_type,omitempty"`
	SiteAdmin    bool   `json:"site_admin,omitempty"`
}

type RepoOwnerDetail struct {
	Login     string `json:"login"`
	ID        int    `json:"id,omitempty"`
	NodeID    string `json:"node_id,omitempty"`
	HTMLURL   string `json:"html_url,omitempty"`
	Type      string `json:"type,omitempty"`
	SiteAdmin bool   `json:"site_admin,omitempty"`
}

type Repository struct {
	ID          int             `json:"id"`
	NodeID      string          `json:"node_id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Private     bool            `json:"private"`
	Owner       RepoOwnerDetail `json:"owner"`
	HTMLURL     string          `json:"html_url"`
	Description string          `json:"description"`
	Fork        bool            `json:"fork"`
	URL         string          `json:"url"`
}

type PackageListItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PackageType string `json:"package_type"`
	Visibility  string `json:"visibility"`
	HTMLURL     string `json:"html_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Owner       struct {
		Login string `json:"login"`
	} `json:"owner"`
	URL string `json:"url"`
}

type PackageDetailDescription struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	PackageType  string      `json:"package_type"`
	Owner        OwnerDetail `json:"owner"`
	VersionCount int         `json:"version_count"`
	Visibility   string      `json:"visibility"`
	URL          string      `json:"url"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	Repository   Repository  `json:"repository"`
	HTMLURL      string      `json:"html_url"`
}

type PackageDescription struct {
	ID         string
	RegistryID string
	Name       string
	URL        string
	CreatedAt  github.Timestamp
	UpdatedAt  github.Timestamp
}

type PackageVersionDescription struct {
	ID          int
	Name        string
	PackageName string
	VersionURI  string
	Digest      *string
	CreatedAt   github.Timestamp
	UpdatedAt   github.Timestamp
}

type CodeSearchResult struct {
	TotalCount        int             `json:"total_count"`
	IncompleteResults bool            `json:"incomplete_results"`
	Items             []CodeSearchHit `json:"items"`
}

type CodeSearchHit struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Sha        string `json:"sha"`
	URL        string `json:"url"`
	GitURL     string `json:"git_url"`
	HTMLURL    string `json:"html_url"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login   string `json:"login"`
			ID      int    `json:"id"`
			NodeID  string `json:"node_id"`
			URL     string `json:"url"`
			HTMLURL string `json:"html_url"`
			Type    string `json:"type"`
		} `json:"owner"`
		HTMLURL     string `json:"html_url"`
		Description string `json:"description"`
		Fork        bool   `json:"fork"`
	} `json:"repository"`
	Score float64 `json:"score"`
}

type ContentResponse struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Sha      string `json:"sha"`
	Size     int    `json:"size"`
	URL      string `json:"url"`
	HTMLURL  string `json:"html_url"`
	GitURL   string `json:"git_url"`
	Type     string `json:"type"`
	Content  string `json:"content"` // base64
	Encoding string `json:"encoding"`
}

type CommitResponse struct {
	Commit struct {
		Author struct {
			Date string `json:"date"`
		} `json:"author"`
		Committer struct {
			Date string `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}

type ArtifactDockerFileDescription struct {
	Sha                     string                 `json:"sha"`
	Name                    string                 `json:"name"`
	Path                    string                 `json:"path"`
	LastUpdatedAt           string                 `json:"last_updated_at"`
	GitURL                  string                 `json:"git_url"`
	HTMLURL                 string                 `json:"html_url"`
	URI                     string                 `json:"uri"` // Unique identifier
	DockerfileContent       string                 `json:"dockerfile_content"`
	DockerfileContentBase64 string                 `json:"dockerfile_content_base64"`
	Repository              map[string]interface{} `json:"repository"`
}
