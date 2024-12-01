package dsfhub

import "fmt"

const awsCommonConnectionDefault = `
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
    region         = "us-east-2"
  }
`

// Output a terraform config for an AWS KINESIS log aggregator resource.
func testAccDSFLogAggregatorConfig_AwsKinesis(resourceName string, gatewayId string, assetId string, parentAssetId string, auditPullEnabled bool, auditType string) string {
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceAttributeReference(parentAssetId)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS KINESIS"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  audit_pull_enabled = %[5]t
  audit_type         = "%[6]s"
  gateway_id         = "%[7]s"
  parent_asset_id    = %[8]s

  %[9]s
}
  `, dsfLogAggregatorResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, auditType, gatewayId, parentAssetIdVal, awsCommonConnectionDefault)
}

// Output a terraform config for an AWS LOG GROUP log aggregator resource.
func testAccDSFLogAggregatorConfig_AwsLogGroup(resourceName string, gatewayId string, assetId string, parentAssetId string, auditPullEnabled bool, auditType string, dependsOn string) string {
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceAttributeReference(parentAssetId)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  depends_on         = [%[3]s]
  server_type        = "AWS LOG GROUP"

  admin_email        = "%[4]s"
  arn                = "%[6]s"
  asset_display_name = "%[6]s"
  asset_id           = "%[6]s"
  audit_pull_enabled = %[8]t
  audit_type         = "%[9]s"
  gateway_id         = "%[5]s"
  parent_asset_id    = %[10]s

  %[11]s
}`, dsfLogAggregatorResourceType, resourceName, dependsOn, testAdminEmail, gatewayId, assetId, parentAssetId, auditPullEnabled, auditType, parentAssetIdVal, awsCommonConnectionDefault)
}

// Output a terraform config for an AWS S3 log aggregator resource.
func testAccDSFLogAggregatorConfig_AwsS3(resourceName string, gatewayId string, assetId string, parentAssetId string, auditPullEnabled string, auditType string) string {
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceAttributeReference(parentAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS S3"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  audit_pull_enabled = %[5]s
  audit_type         = "%[6]s"
  bucket_account_id  = "%[7]s"
  gateway_id         = "%[8]s"
  parent_asset_id    = %[9]s
  region             = "us-east-1"
  server_host_name   = "%[4]s"

  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
}
  `, dsfLogAggregatorResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, auditType, testAwsAccountId, gatewayId, parentAssetIdVal)
}

// Output an asset_connection block for an AZURE EVENTHUB log aggregator resource.
func azureEventhubConnectionBlock(authMechanism string, format string) string {
	var output string

	switch authMechanism {
	case "azure_ad":
		output = fmt.Sprintf(`
  asset_connection {
    auth_mechanism          = "azure_ad"
    azure_storage_account   = "mystorageaccount"
    azure_storage_container = "mystoragecontainer"
    eventhub_name           = "myeventhub"
    eventhub_namespace      = "myeventhubnamespace"
    format                  = "%[1]s"
    reason                  = "default"
  }`, format)
	case "client_secret":
		output = fmt.Sprintf(`
  asset_connection {
    application_id          = "a1b2c3de-123c-1234-ab12-ab12c2de3fg4"
    auth_mechanism          = "client_secret"
    azure_storage_account   = "mystorageaccount"
    azure_storage_container = "mystoragecontainer"
    client_secret           = "secret"
    directory_id            = "a1b2c3de-123c-1234-ab12-ab12c2de3fg4"
    eventhub_name           = "myeventhub"
    eventhub_namespace      = "myeventhubnamespace"
    format                  = "%[1]s"
    subscription_id         = "a1b2c3de-123c-1234-ab12-ab12c2de3fg4"
    reason                  = "default"
  }
  
  %[2]s
  `, format, ignoreAssetConnectionChangesBlock())
	case "default":
		output = fmt.Sprintf(`
  asset_connection {
    auth_mechanism           = "default"
    azure_storage_account    = "mystorageaccount"
    azure_storage_container  = "mystoragecontainer"
    azure_storage_secret_key = "storage-secret"
    eventhub_access_key      = "eventhub-secret"
    eventhub_access_policy   = "RootManageSharedAccessKey"
    eventhub_name            = "myeventhub"
    eventhub_namespace       = "myeventhubnamespace"
    format                   = "%[1]s"
    reason                   = "default"
  }
  
  %[2]s
  `, format, ignoreAssetConnectionChangesBlock())
	}

	return output
}

// Output a terraform config for an AZURE EVENTHUB log aggregator resource.
func testAccDSFLogAggregatorConfig_AzureEventhub(resourceName string, gatewayId string, assetId string, authMechanism string, parentAssetId string, auditPullEnabled string, contentType string, format string) string {
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceAttributeReference(parentAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type = "AZURE EVENTHUB"

  admin_email        = "%[3]s"
  asset_id           = "%[5]s"
  asset_display_name = "%[5]s"
  audit_pull_enabled = %[6]s
  content_type       = "%[7]s"
  gateway_id         = "%[4]s"
  parent_asset_id    = %[8]s
  server_host_name   = "my-namespace.servicebus.windows.net"
  server_port        = "443"

  %[9]s
}`, dsfLogAggregatorResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditPullEnabled, contentType, parentAssetIdVal, azureEventhubConnectionBlock(authMechanism, format))
}

const gcpPubsubConnectionServiceAccount = `
  asset_connection {
    auth_mechanism = "service_account"
    reason         = "default"
    key_file       = "/data/jsonar/local/credentials/gcp_service_account.json"
  }
`

const commonGcpConnectionDefault = `
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
`

// Output a terraform config for an GCP CLOUD STORAGE BUCKET log aggregator resource.
func testAccDSFLogAggregatorConfig_GcpCloudStorageBucket(resourceName string, gatewayId string, assetId string, parentAssetId string, auditPullEnabled string, contentType string) string {
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceAttributeReference(parentAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "GCP CLOUD STORAGE BUCKET"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  audit_pull_enabled = %[5]s
  audit_type         = "BUCKET"
  content_type       = "%[6]s"
  gateway_id         = "%[7]s"
  parent_asset_id    = %[8]s
  server_host_name   = "storage.googleapis.com"
  server_ip          = "storage.googleapis.com"
  server_port        = "443" 

  %[9]s
}`, dsfLogAggregatorResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, contentType, gatewayId, parentAssetIdVal, commonGcpConnectionDefault)

}

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
		assetConnectionBlock = gcpPubsubConnectionServiceAccount
	default:
		assetConnectionBlock = commonGcpConnectionDefault
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type         = "GCP PUBSUB"

  admin_email         = "%[3]s"
  asset_display_name  = "%[5]s"
  asset_id            = "%[5]s"
  audit_pull_enabled  = %[6]s
  audit_type          = "%[7]s"
  content_type        = "%[8]s"
  gateway_id          = "%[4]s"
  parent_asset_id     = %[9]s
  pubsub_subscription = "%[5]s"
  server_host_name    = "pubsub.googleapis.com"
  server_ip           = "pubsub.googleapis.com"
  server_port         = "443"

  %[10]s
}`, dsfLogAggregatorResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditPullEnabled, auditType, contentType, parentAssetIdVal, assetConnectionBlock)
}
