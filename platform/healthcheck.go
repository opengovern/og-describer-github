package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/go-github/v66/github"
	"golang.org/x/oauth2"
)

type GraphQLResponse struct {
	Data   interface{}    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}
type GraphQLError struct {
	Message string `json:"message"`
}
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// Config represents the JSON input configuration
type Config struct {
	Token            string `json:"token"`
	OrganizationName string `json:"organization_name"`
}

// Define required permissions as constants
const (
	ReadPublicKey  = "read:public_key"
	ReadUser       = "read:user"
	ReadProject    = "read:project"
	RepoDeployment = "repo_deployment"
	ReadRepoHook   = "read:repo_hook"
	PublicRepo     = "public_repo"
	ReadOrg        = "read:org"
	RepoStatus     = "repo:status"
	ReadPackages   = "read:packages"
)

// HealthStatus represents the structure of the JSON output
type HealthStatus struct {
	Organization string  `json:"organization"`
	Healthy      bool    `json:"healthy"`
	Details      Details `json:"details"`
}

// Details contains required and missing permissions
type Details struct {
	RequiredPermissions []string `json:"required_permissions"`
	MissingPermissions  []string `json:"missing_permissions"`
}

// PermissionCheck represents a permission and its corresponding check function
type PermissionCheck struct {
	Name  string
	Check func(ctx context.Context, client *github.Client, org string) error
}

// IsHealthy checks if the PAT has read access to all required artifacts in the organization
func IsHealthy(ctx context.Context, client *github.Client, org string) error {
	// Define all required permissions and their corresponding checks
	permissions := []PermissionCheck{
		{
			Name: PublicRepo,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				// Attempt to list public repositories
				_, _, err := client.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
					Type: "public",
					ListOptions: github.ListOptions{
						PerPage: 1,
					},
				})
				return err
			},
		},
		{
			Name: ReadOrg,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				// Attempt to get organization details
				_, _, err := client.Organizations.Get(ctx, org)
				return err
			},
		},
		{
			Name: ReadPackages,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				// Define valid package types as per GitHub API
				packageTypes := []string{"container", "npm", "maven", "nuget", "rubygems", "docker", "composer"}

				// Iterate over each package type to check access
				for _, pkgType := range packageTypes {
					_, _, err := client.Organizations.ListPackages(ctx, org, &github.PackageListOptions{
						PackageType: github.String(pkgType),
						ListOptions: github.ListOptions{
							PerPage: 1,
						},
					})
					if err != nil {
						// If the error is a 422 Unprocessable Entity, it might mean no packages of this type exist
						// Skip to the next type
						if ghErr, ok := err.(*github.ErrorResponse); ok && ghErr.Response.StatusCode == 422 {
							continue
						}
						// For other errors, return immediately
						return err
					}
				}
				// If all package types resulted in 422, it implies no packages exist, but permissions might still be valid
				return nil
			},
		},
		{
			Name: ReadProject,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				graphqlQuery := `
						query CheckOrgProjectAccess($orgLogin: String!) {
						  organization(login: $orgLogin) {
							# Requesting projectsV2 itself requires read:project or read:org
							projectsV2(first: 1) {
							  # Requesting nodes requires read access
							  nodes {
								id # Requesting a simple field
							  }
							}
						  }
						}`

				variables := map[string]interface{}{
					"orgLogin": org,
				}

				requestBody := GraphQLRequest{
					Query:     graphqlQuery,
					Variables: variables,
				}
				requestBodyBytes, err := json.Marshal(requestBody)
				if err != nil {
					// Should not happen with this static structure, but handle defensively
					return fmt.Errorf("failed to marshal graphql request for project check: %w", err)
				}

				// Create POST request to GraphQL endpoint
				url := "https://api.github.com/graphql" // Correct GraphQL endpoint
				req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
				if err != nil {
					return fmt.Errorf("failed to create graphql request for project check: %w", err)
				}
				req.Header.Set("Content-Type", "application/json")

				// Use client.Do to execute the request with the client's auth
				graphqlResp := new(GraphQLResponse)
				// client.Do handles sending the request and basic response checking (non-2xx status codes)
				// It also decodes the response body into graphqlResp if successful
				_, err = client.Do(ctx, req, graphqlResp)
				if err != nil {
					// This err could be a transport error or a non-2xx status code error from CheckResponse
					// It might already indicate a 401/403 permission issue at the HTTP level
					return fmt.Errorf("graphql request for project check failed (HTTP level): %w", err)
				}

				// Even with 200 OK, check for errors reported within the GraphQL response body
				if len(graphqlResp.Errors) > 0 {
					errorMessages := ""
					for i, gqlErr := range graphqlResp.Errors {
						if i > 0 {
							errorMessages += "; "
						}
						errorMessages += gqlErr.Message
					}
					return fmt.Errorf("graphql query for project check returned errors: %s", errorMessages)
				}

				// If no error from client.Do and no errors in the response body, access is likely okay
				return nil
			},
		},
		{
			Name: ReadPublicKey,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				// Public keys are associated with users, not organizations.
				// Attempt to list public keys for the authenticated user
				_, _, err := client.Users.ListKeys(ctx, "", &github.ListOptions{
					PerPage: 1,
				})
				return err
			},
		},
		{
			Name: ReadRepoHook,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				// Attempt to list repository hooks for a sample repository
				repos, _, err := client.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
					Type: "public",
					ListOptions: github.ListOptions{
						PerPage: 10,
					},
				})
				if err != nil {
					return err
				}
				if len(repos) == 0 {
					return errors.New("no repositories found to check repo hooks")
				}
				_, _, err = client.Repositories.ListHooks(ctx, org, repos[0].GetName(), &github.ListOptions{
					PerPage: 1,
				})
				return err
			},
		},
		{
			Name: ReadUser,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				// Attempt to get the authenticated user
				user, _, err := client.Users.Get(ctx, "")
				if err != nil {
					return err
				}
				// Verify the user belongs to the organization
				_, _, err = client.Organizations.GetOrgMembership(ctx, user.GetLogin(), org)
				return err
			},
		},
		{
			Name: RepoStatus,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				// Attempt to list commit statuses for a sample commit
				repos, _, err := client.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
					Type: "public",
					ListOptions: github.ListOptions{
						PerPage: 10,
					},
				})
				if err != nil {
					return err
				}
				if len(repos) == 0 {
					return errors.New("no repositories found to check repo status")
				}
				// Get a commit SHA
				commits, _, err := client.Repositories.ListCommits(ctx, org, repos[0].GetName(), &github.CommitsListOptions{
					ListOptions: github.ListOptions{
						PerPage: 1,
					},
				})
				if err != nil {
					return err
				}
				if len(commits) == 0 {
					return errors.New("no commits found to check repo status")
				}
				// Use ListStatuses instead of the undefined ListCommitStatuses
				_, _, err = client.Repositories.ListStatuses(ctx, org, repos[0].GetName(), commits[0].GetSHA(), &github.ListOptions{
					PerPage: 1,
				})
				return err
			},
		},
		{
			Name: RepoDeployment,
			Check: func(ctx context.Context, client *github.Client, org string) error {
				// Attempt to list deployments for a sample repository
				repos, _, err := client.Repositories.ListByOrg(ctx, org, &github.RepositoryListByOrgOptions{
					Type: "public",
					ListOptions: github.ListOptions{
						PerPage: 10,
					},
				})
				if err != nil {
					return err
				}
				if len(repos) == 0 {
					return errors.New("no repositories found to check repo deployments")
				}
				_, _, err = client.Repositories.ListDeployments(ctx, org, repos[0].GetName(), &github.DeploymentsListOptions{
					ListOptions: github.ListOptions{
						PerPage: 1,
					},
				})
				return err
			},
		},
		// Add more permissions and their checks as needed
		// For brevity, not all permissions from the list are implemented here
	}

	requiredPermissions := []string{}
	missingPermissions := []string{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Channel to limit concurrency
	concurrencyLimit := 5
	sem := make(chan struct{}, concurrencyLimit)

	for _, perm := range permissions {
		wg.Add(1)
		go func(p PermissionCheck) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			err := p.Check(ctx, client, org)
			mu.Lock()
			defer mu.Unlock()
			requiredPermissions = append(requiredPermissions, p.Name)
			if err != nil {
				// Check if the error is due to permission issues
				if isPermissionError(err) {
					missingPermissions = append(missingPermissions, p.Name)
				} else {
					// Log other errors
					log.Printf("Error checking permission '%s': %v", p.Name, err)
					missingPermissions = append(missingPermissions, p.Name)
				}
			}
		}(perm)
	}

	wg.Wait()

	healthy := len(missingPermissions) == 0

	status := HealthStatus{
		Organization: org,
		Healthy:      healthy,
		Details: Details{
			RequiredPermissions: requiredPermissions,
			MissingPermissions:  missingPermissions,
		},
	}

	// Marshal to JSON and print
	output, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(output))

	if !healthy {
		errorMsg := fmt.Sprintf("Organization '%s' is not healthy. Missing permissions: %s", org, missingPermissions)
		return errors.New(errorMsg)
	}

	return nil
}

// isPermissionError determines if an error is due to insufficient permissions
func isPermissionError(err error) bool {
	if err == nil {
		return false
	}
	// Check if it's a GitHub API error
	if githubErr, ok := err.(*github.ErrorResponse); ok {
		if githubErr.Response != nil && githubErr.Response.StatusCode == 403 {
			return true
		}
	}
	return false
}

func GithubIntegrationHealthcheck(cfg Config) (bool, error) {
	token := cfg.Token
	if token == "" {
		return false, fmt.Errorf("no token provided")
	}

	// Read organization name
	orgName := cfg.OrganizationName
	if orgName == "" {
		return false, fmt.Errorf("organization name is required")
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Create an OAuth2 token source
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	// Create an OAuth2 client
	tc := oauth2.NewClient(ctx, ts)

	// Create a new GitHub client
	client := github.NewClient(tc)

	// Now process permissions for the specified organization
	fmt.Printf("\nChecking Access for Organization: %s\n", orgName)
	err := IsHealthy(ctx, client, orgName)
	if err != nil {
		return false, err
	}

	return true, nil
}
