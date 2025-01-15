package integration

import (
	"encoding/json"
	"github.com/jackc/pgtype"
	"github.com/opengovern/og-describer-github/describer/pkg/wrapper"
	configs2 "github.com/opengovern/og-describer-github/describer/provider/configs"
	"github.com/opengovern/og-describer-github/plugin/integration/configs"
	"github.com/opengovern/og-describer-github/plugin/integration/discovery"
	"github.com/opengovern/og-describer-github/plugin/integration/healthcheck"
	"github.com/opengovern/og-util/pkg/integration"
	"github.com/opengovern/opencomply/services/integration/integration-type/interfaces"
	"github.com/opengovern/opencomply/services/integration/models"
	"strconv"
)

type Integration struct{}

func (i *Integration) GetConfiguration() interfaces.IntegrationConfiguration {
	return interfaces.IntegrationConfiguration{
		NatsScheduledJobsTopic:   configs2.JobQueueTopic,
		NatsManualJobsTopic:      configs2.JobQueueTopicManuals,
		NatsStreamName:           configs2.StreamName,
		NatsConsumerGroup:        configs2.ConsumerGroup,
		NatsConsumerGroupManuals: configs2.ConsumerGroupManuals,

		SteampipePluginName: "github",

		UISpec: configs.UISpec,

		DescriberDeploymentName: configs.DescriberDeploymentName,
		DescriberRunCommand:     configs.DescriberRunCommand,
	}
}

func (i *Integration) HealthCheck(jsonData []byte, providerId string, labels map[string]string, annotations map[string]string) (bool, error) {
	var credentials configs.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return false, err
	}

	var name string
	if v, ok := labels["OrganizationName"]; ok {
		name = v
	}
	isHealthy, err := healthcheck.GithubIntegrationHealthcheck(healthcheck.Config{
		Token:            credentials.PatToken,
		OrganizationName: name,
	})
	return isHealthy, err
}

func (i *Integration) DiscoverIntegrations(jsonData []byte) ([]models.Integration, error) {
	var credentials configs.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return nil, err
	}
	var integrations []models.Integration
	accounts, err := discovery.GithubIntegrationDiscovery(discovery.Config{
		Token: credentials.PatToken,
	})
	if err != nil {
		return nil, err
	}
	for _, a := range accounts {
		labels := map[string]string{
			"OrganizationName": a.Login,
		}
		labelsJsonData, err := json.Marshal(labels)
		if err != nil {
			return nil, err
		}
		integrationLabelsJsonb := pgtype.JSONB{}
		err = integrationLabelsJsonb.Set(labelsJsonData)
		if err != nil {
			return nil, err
		}
		integrations = append(integrations, models.Integration{
			ProviderID: strconv.FormatInt(a.ID, 10),
			Name:       a.Login,
			Labels:     integrationLabelsJsonb,
		})
	}
	return integrations, nil
}

func (i *Integration) GetResourceTypesByLabels(labels map[string]string) (map[string]interfaces.ResourceTypeConfiguration, error) {
	resourceTypesMap := make(map[string]interfaces.ResourceTypeConfiguration)
	for _, resourceType := range configs.ResourceTypesList {
		if v, ok := configs.ResourceTypeConfigs[resourceType]; ok {
			resourceTypesMap[resourceType] = *v
		} else {
			resourceTypesMap[resourceType] = interfaces.ResourceTypeConfiguration{}
		}
	}
	return resourceTypesMap, nil
}

func (i *Integration) GetResourceTypeFromTableName(tableName string) string {
	if v, ok := configs.TablesToResourceTypes[tableName]; ok {
		return v
	}

	return ""
}

func (i *Integration) GetIntegrationType() integration.Type {
	return configs.IntegrationTypeGithubAccount
}

func (i *Integration) ListAllTables() map[string][]string {
	plugin := wrapper.Plugin()
	tables := make(map[string][]string)
	for tableKey, table := range plugin.TableMap {
		columnNames := make([]string, 0, len(table.Columns))
		for _, column := range table.Columns {
			columnNames = append(columnNames, column.Name)
		}
		tables[tableKey] = columnNames
	}

	return tables
}
