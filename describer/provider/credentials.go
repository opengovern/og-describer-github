package provider

import (
	"encoding/json"
	model "github.com/opengovern/og-describer-github/describer/pkg/sdk/models"
	"github.com/opengovern/og-describer-github/global"
	"github.com/opengovern/og-util/pkg/describe"
)

// AccountCredentialsFromMap TODO: converts a map to a configs.IntegrationCredentials.
func AccountCredentialsFromMap(m map[string]any) (global.IntegrationCredentials, error) {
	mj, err := json.Marshal(m)
	if err != nil {
		return global.IntegrationCredentials{}, err
	}

	var c global.IntegrationCredentials
	err = json.Unmarshal(mj, &c)
	if err != nil {
		return global.IntegrationCredentials{}, err
	}

	return c, nil
}

// GetResourceMetadata TODO: Get metadata as a map to add to the resources
func GetResourceMetadata(job describe.DescribeJob, resource model.Resource) (map[string]string, error) {
	metadata := make(map[string]string)

	return metadata, nil
}

// AdjustResource TODO: Do any needed adjustment on resource object before storing
func AdjustResource(job describe.DescribeJob, resource *model.Resource) error {
	return nil
}

// GetAdditionalParameters TODO: pass additional parameters needed in describer wrappers in /provider/describer_wrapper.go
func GetAdditionalParameters(job describe.DescribeJob) (map[string]string, error) {
	additionalParameters := make(map[string]string)

	if _, ok := job.IntegrationLabels["OrganizationName"]; ok {
		additionalParameters["OrganizationName"] = job.IntegrationLabels["OrganizationName"]
	}

	return additionalParameters, nil
}
