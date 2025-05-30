package describers

import (
	"context"

	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
)

func GetNPMPackageList(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})
	packages, err := fetchAllPackages(sdk, organizationName, "npm")
	if err != nil {
		return nil, err
	}

	var values []models.Resource
	for _, p := range packages {
		packageValue, err := fetchPackageDetails(sdk, organizationName, "npm", p.Name, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, packageValue...)
	}
	return values, nil
}
