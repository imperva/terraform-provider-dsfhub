package dsfhub

import "fmt"

// Output a terraform config for an AWS LOG GROUP log aggregator resource.
func testAccDSFLogAggregatorConfig_AwsLogGroup(resourceName string, gatewayId string, assetId string, parentAssetId string, auditPullEnabled bool, auditType string, dependsOn string) string {
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceAttributeReference(parentAssetId)

	return fmt.Sprintf(`
resource "`+dsfLogAggregatorResourceType+`" "%[1]s" {
  depends_on = [`+dependsOn+`]
  server_type = "AWS LOG GROUP"

  admin_email = "`+testAdminEmail+`"
  arn	= "%[3]s"
  asset_display_name = "%[3]s"
  asset_id = "%[3]s"
  audit_pull_enabled = %[5]t
  audit_type = "%[6]s"
  gateway_id = "%[2]s"
  parent_asset_id = `+parentAssetIdVal+`

  asset_connection {
    auth_mechanism = "default"
    reason = "default"
    region = "us-east-2"
  }
}`, resourceName, gatewayId, assetId, parentAssetId, auditPullEnabled, auditType)
}
