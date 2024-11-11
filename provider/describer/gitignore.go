package describer

import (
	"context"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider"
	"github.com/opengovern/og-describer-template/provider/model"
)

func GetGitIgnoreTemplateList(ctx context.Context, githubClient provider.GitHubClient, stream *models.StreamSender) ([]models.Resource, error) {
	client := githubClient.RestClient
	gitIgnores, _, err := client.Gitignores.List(ctx)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, gitIgnore := range gitIgnores {
		repoValue, err := GetGitignoreTemplate(ctx, githubClient, gitIgnore)
		if err != nil {
			return nil, err
		}
		if stream != nil {
			if err := (*stream)(*repoValue); err != nil {
				return nil, err
			}
		} else {
			values = append(values, *repoValue)
		}
	}
	return values, nil
}

func GetGitignoreTemplate(ctx context.Context, githubClient provider.GitHubClient, gitIgnoreName string) (*models.Resource, error) {
	client := githubClient.RestClient
	gitIgnore, _, err := client.Gitignores.Get(ctx, gitIgnoreName)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   gitIgnoreName,
		Name: gitIgnoreName,
		Description: JSONAllFieldsMarshaller{
			Value: model.GitIgnore{
				Gitignore: *gitIgnore,
			},
		},
	}
	return &value, nil
}