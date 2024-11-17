package dsfhub

import "fmt"

const awsLogGroupConnectionDefault = `
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
    region         = "us-east-2"
  }
`

// Output a terraform config for an AWS LOG GROUP log aggregator resource.
func testAccDSFLogAggregatorConfig_AwsLogGroup(resourceName string, gatewayId string, assetId string, parentAssetId string, auditPullEnabled bool, auditType string, dependsOn string) string {
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceAttributeReference(parentAssetId)

	return fmt.Sprintf(`
resource "`+dsfLogAggregatorResourceType+`" "%[1]s" {
  depends_on = [`+dependsOn+`]
  server_type        = "AWS LOG GROUP"

  admin_email        = "`+testAdminEmail+`"
  arn                = "%[3]s"
  asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  audit_pull_enabled = %[5]t
  audit_type         = "%[6]s"
  gateway_id         = "%[2]s"
  parent_asset_id    = `+parentAssetIdVal+`

  `+awsLogGroupConnectionDefault+`
}`, resourceName, gatewayId, assetId, parentAssetId, auditPullEnabled, auditType)
}

const gcpConnectionServiceAccount = `
  asset_connection {
    auth_mechanism = "service_account"
    reason         = "default"
    key_file       = "/data/jsonar/local/credentials/gcp_service_account.json"
  }
`

const gcpConnectDefault = `
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
`

// Output a terraform config for an GCP PUBSUB log aggregator resource.
func testAccDSFLogAggregatorConfig_GcpPubsub(resourceName string, gatewayId string, assetId string, authMechanism string, parentAssetId string, auditPullEnabled string, auditType string, contentType string) string {
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceAttributeReference(parentAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	// handle different asset_connection blocks
	var assetConnectionBlock string
	switch authMechanism {
	case "service_account":
		assetConnectionBlock = gcpConnectionServiceAccount
	default:
		assetConnectionBlock = gcpConnectDefault
	}

	return fmt.Sprintf(`
resource "`+dsfLogAggregatorResourceType+`" "%[1]s" {
  server_type         = "GCP PUBSUB"

  admin_email         = "`+testAdminEmail+`"
  asset_display_name  = "%[3]s"
  asset_id            = "%[3]s"
  audit_pull_enabled  = %[4]s
  audit_type          = "%[5]s"
  content_type        = "%[6]s"
  gateway_id          = "%[2]s"
  parent_asset_id     = `+parentAssetIdVal+`
  pubsub_subscription = "%[3]s"
  server_host_name    = "pubsub.googleapis.com"
  server_ip           = "pubsub.googleapis.com"
  server_port         = "443"

  `+assetConnectionBlock+`
}`, resourceName, gatewayId, assetId, auditPullEnabled, auditType, contentType)
}
