package common

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/lib/azure/configure"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

const terraformDir string = "../../examples/dns_zone_records"
const varFile string = "test.tfvars"

// TODO: Add support for other cloud configurations
const azEnvironment string = "AZUREPUBLICCLOUD"

func TestDnsZoneRecords(t *testing.T, ctx types.TestContext) {
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		t.Fatalf("ARM_SUBSCRIPTION_ID must be set for acceptance tests")
	}
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Error getting credentials: %e\n", err)
	}
	cloudCfg, err := getCloudConfiguration(azEnvironment)
	if err != nil {
		t.Fatalf("Error interpreting Azure Cloud type: %e\n", err)
	}

	clientOpt := policy.ClientOptions{Cloud: cloudCfg}
	dnsZoneRecordSetsClient, err := armdns.NewRecordSetsClient(subscriptionID, credential, &arm.ClientOptions{ClientOptions: clientOpt})
	if err != nil {
		t.Fatalf("Error creating DNS Zone Record Sets Client: %e\n", err)
	}

	terraformOptions := configure.ConfigureTerraform(terraformDir, []string{terraformDir + "/" + varFile})

	t.Run("dnsZoneRecords", func(t *testing.T) {
		checkDNSZoneRecordSets(t, dnsZoneRecordSetsClient, terraformOptions)
	})
}

func checkDNSZoneRecordSets(t *testing.T, dnsZoneRecordSetsClient *armdns.RecordSetsClient, terraformOptions *terraform.Options) {
	checkRecord(t, dnsZoneRecordSetsClient, terraformOptions, "a_record_ids", "a_record_fqdns", armdns.RecordTypeA)
	checkRecord(t, dnsZoneRecordSetsClient, terraformOptions, "cname_record_ids", "cname_record_fqdns", armdns.RecordTypeCNAME)
	checkRecord(t, dnsZoneRecordSetsClient, terraformOptions, "ns_record_ids", "ns_record_fqdns", armdns.RecordTypeNS)
	checkRecord(t, dnsZoneRecordSetsClient, terraformOptions, "txt_record_ids", "txt_record_fqdns", armdns.RecordTypeTXT)
}

func checkRecord(t *testing.T, dnsZoneRecordSetsClient *armdns.RecordSetsClient, terraformOptions *terraform.Options, recordSetIdsKey string, recordSetFQDNsKey string, recordType armdns.RecordType) {
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	zoneNames := terraform.OutputList(t, terraformOptions, "public_dns_zone_names")
	recordSetIds := terraform.OutputMap(t, terraformOptions, recordSetIdsKey)
	recordSetFQDNs := terraform.OutputMap(t, terraformOptions, recordSetFQDNsKey)
	options := armdns.RecordSetsClientGetOptions{}

	for _, zoneName := range zoneNames {
		for key, recordSetId := range recordSetIds {
			relativeRecordSetName := getSubdomain(recordSetFQDNs[key])

			dnsZoneRecordSets, err := dnsZoneRecordSetsClient.Get(context.Background(), resourceGroupName, zoneName, relativeRecordSetName, recordType, &options)
			if err != nil {
				t.Fatalf("Error getting DNS Zone: %v", err)
			}

			assert.Equal(t, strings.ToLower(recordSetId), strings.ToLower(*dnsZoneRecordSets.ID))
			assert.Equal(t, strings.ToLower(relativeRecordSetName), strings.ToLower(*dnsZoneRecordSets.Name))
			assert.Equal(t, strings.ToLower((string(recordType))), getRecordType(strings.ToLower(*dnsZoneRecordSets.Type)))
			assert.NotEmpty(t, dnsZoneRecordSets.Properties)
		}
	}
}

func getSubdomain(domain string) string {
	parts := strings.Split(domain, ".")
	return parts[0]
}

func getRecordType(input string) string {
	parts := strings.Split(input, "/")
	return parts[len(parts)-1]
}

func getCloudConfiguration(name string) (cloud.Configuration, error) {
	switch strings.ToUpper(name) {
	case "AZURECLOUD", "AZUREPUBLICCLOUD", "":
		return cloud.AzurePublic, nil
	case "AZUREUSGOVERNMENT", "AZUREUSGOVERNMENTCLOUD":
		return cloud.AzureGovernment, nil
	case "AZURECHINACLOUD":
		return cloud.AzureChina, nil
	}
	return cloud.Configuration{}, fmt.Errorf("unknown cloud configuration name: %s", name)
}
