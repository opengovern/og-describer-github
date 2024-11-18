package describer

import (
	"context"
	"fmt"
	"github.com/opengovern/og-describer-github/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/provider/model"
	"strconv"
	"strings"

	goPipeline "github.com/buildkite/go-pipeline"

	"github.com/google/go-github/v55/github"
)

func GetAllWorkflows(ctx context.Context, githubClient GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient

	repositories, err := getRepositories(ctx, client, organizationName)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, repo := range repositories {
		repoValues, err := GetRepositoryWorkflows(ctx, githubClient, stream, organizationName, repo.GetName())
		if err != nil {
			return nil, err
		}
		values = append(values, repoValues...)
	}
	return values, nil
}

type FileContent struct {
	Repository string
	FilePath   string
	Content    string
}

func GetRepositoryWorkflows(ctx context.Context, githubClient GitHubClient, stream *models.StreamSender, owner, repo string) ([]models.Resource, error) {
	client := githubClient.RestClient
	opts := &github.ListOptions{PerPage: pageSize}
	repoFullName := formRepositoryFullName(owner, repo)
	var values []models.Resource
	for {
		workflows, resp, err := client.Actions.ListWorkflows(ctx, owner, repo, opts)
		if err != nil {
			return nil, err
		}
		for _, workflow := range workflows.Workflows {
			fileContent, err := getWorkflowFileContent(ctx, client, workflow, owner, repo)
			if err != nil {
				return nil, err
			}
			content, err := fileContent.GetContent()
			if err != nil {
				return nil, err
			}
			fileContentBasic := FileContent{
				Repository: repo,
				FilePath:   fileContent.GetPath(),
				Content:    content,
			}
			pipeline, err := decodeFileContentToPipeline(fileContentBasic)
			if err != nil {
				return nil, err
			}
			value := models.Resource{
				ID:   strconv.Itoa(int(workflow.GetID())),
				Name: workflow.GetName(),
				Description: JSONAllFieldsMarshaller{
					Value: model.WorkflowDescription{
						Workflow:                workflow,
						RepositoryFullName:      repoFullName,
						WorkFlowFileContent:     content,
						WorkFlowFileContentJson: fileContent,
						Pipeline:                pipeline,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return values, nil
}

func getWorkflowFileContent(ctx context.Context, client *github.Client, workflow *github.Workflow, owner, repo string) (*github.RepositoryContent, error) {
	if workflow.Path == nil {
		return nil, nil
	}
	workflowUrlParts := strings.Split(*workflow.HTMLURL, "/")
	defaultBranch := "main"
	if len(workflowUrlParts) > 6 {
		defaultBranch = workflowUrlParts[6]
	}
	content, _, _, err := client.Repositories.GetContents(ctx, owner, repo, workflow.GetPath(), &github.RepositoryContentGetOptions{Ref: defaultBranch})
	if err != nil {
		if strings.Contains(err.Error(), "404 Not Found") || strings.Contains(err.Error(), "404 No commit found") {
			return nil, nil
		}
		return nil, err
	}
	return content, nil
}

func decodeFileContentToPipeline(contentDetails FileContent) (*goPipeline.Pipeline, error) {
	pipeline, err := goPipeline.Parse(strings.NewReader(contentDetails.Content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the workflow file '%s', %v", contentDetails.FilePath, err)
	}
	return pipeline, nil
}
