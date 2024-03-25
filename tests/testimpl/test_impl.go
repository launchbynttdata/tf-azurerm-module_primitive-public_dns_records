package common

import (
	"context"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/dns/mgmt/dns"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/lib/azure/configure"
	internalDns "github.com/launchbynttdata/lcaf-component-terratest/lib/azure/dns"
	"github.com/launchbynttdata/lcaf-component-terratest/lib/azure/login"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

const terraformDir string = "../../examples/dns_zone_records"
const varFile string = "test.tfvars"

func TestDnsZoneRecords(t *testing.T, ctx types.TestContext) {
	envVarMap := login.GetEnvironmentVariables()
	clientID := envVarMap["clientID"]
	clientSecret := envVarMap["clientSecret"]
	tenantID := envVarMap["tenantID"]
	subscriptionID := envVarMap["subscriptionID"]

	spt, err := login.GetServicePrincipalToken(clientID, clientSecret, tenantID)
	if err != nil {
		t.Fatalf("Error getting Service Principal Token: %v", err)
	}

	dnsZoneRecordSetsClient := internalDns.GetDNSZoneRecordSetsClient(spt, subscriptionID)

	terraformOptions := configure.ConfigureTerraform(terraformDir, []string{terraformDir + "/" + varFile})

	t.Run("dnsZoneRecords", func(t *testing.T) {
		checkDNSZoneRecordSets(t, dnsZoneRecordSetsClient, terraformOptions)
	})
}

func checkDNSZoneRecordSets(t *testing.T, dnsZoneRecordSetsClient dns.RecordSetsClient, terraformOptions *terraform.Options) {
	checkRecord(t, dnsZoneRecordSetsClient, terraformOptions, "a_record_ids", "a_record_fqdns", dns.A)
	checkRecord(t, dnsZoneRecordSetsClient, terraformOptions, "cname_record_ids", "cname_record_fqdns", dns.CNAME)
	checkRecord(t, dnsZoneRecordSetsClient, terraformOptions, "ns_record_ids", "ns_record_fqdns", dns.NS)
	checkRecord(t, dnsZoneRecordSetsClient, terraformOptions, "txt_record_ids", "txt_record_fqdns", dns.TXT)
}

func checkRecord(t *testing.T, dnsZoneRecordSetsClient dns.RecordSetsClient, terraformOptions *terraform.Options, recordSetIdsKey string, recordSetFQDNsKey string, recordType dns.RecordType) {
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	zoneNames := terraform.OutputList(t, terraformOptions, "public_dns_zone_names")
	recordSetIds := terraform.OutputMap(t, terraformOptions, recordSetIdsKey)
	recordSetFQDNs := terraform.OutputMap(t, terraformOptions, recordSetFQDNsKey)

	for _, zoneName := range zoneNames {
		for key, recordSetId := range recordSetIds {
			relativeRecordSetName := getSubdomain(recordSetFQDNs[key])

			dnsZoneRecordSets, err := dnsZoneRecordSetsClient.Get(context.Background(), resourceGroupName, zoneName, relativeRecordSetName, recordType)
			if err != nil {
				t.Fatalf("Error getting DNS Zone: %v", err)
			}

			assert.Equal(t, strings.ToLower(recordSetId), strings.ToLower(*dnsZoneRecordSets.ID))
			assert.Equal(t, strings.ToLower(relativeRecordSetName), strings.ToLower(*dnsZoneRecordSets.Name))
			assert.Equal(t, strings.ToLower((string(recordType))), getRecordType(strings.ToLower(*dnsZoneRecordSets.Type)))
			assert.NotEmpty(t, dnsZoneRecordSets.RecordSetProperties)
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
