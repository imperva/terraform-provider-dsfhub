package dsfhub

import "fmt"

const awsConnectionDefault = `
  asset_connection {
    auth_mechanism  = "default"
    reason          = "default"
    region          = "us-west-1"
  }
`

var awsConnectionKey = fmt.Sprintf(`
  asset_connection {
    access_id       = "my-access-id"
    auth_mechanism  = "key"
    reason          = "default"
    region          = "us-east-1"
    secret_key      = "my-secret-key"
  }

  %[1]s
`, ignoreAssetConnectionChangesBlock())

const awsConnectionProfile = `
  asset_connection {
    auth_mechanism  = "profile"
    reason          = "default"
    region          = "us-east-2"
    username        = "dsfhubuser"
  }
`

// Output a terraform config for an AWS cloud account resource.
//
// Supports all authentication mechanisms: "key", "profile", "iam_role", and
// "default".
func testAccDSFCloudAccountConfig_Aws(resourceName string, gatewayId string, assetId string, authMechanism string) string {
	var assetConnectionBlock string

	// build the asset_connection block
	switch authMechanism {
	case "key":
		assetConnectionBlock = awsConnectionKey
	case "profile":
		assetConnectionBlock = awsConnectionProfile
	default:
		assetConnectionBlock = awsConnectionDefault
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type = "AWS"
  
  admin_email        = "%[3]s"
  arn                = "%[4]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  gateway_id         = "%[5]s"

  %[6]s
}`, dsfCloudAccountResourceType, resourceName, testAdminEmail, assetId, gatewayId, assetConnectionBlock)
}

var azureConnectionClientSecret = fmt.Sprintf(`
  asset_connection {
    auth_mechanism  = "client_secret"
    application_id  = "12345678-1234-1234-1234-123456789012" 
    client_secret   = "secret"
    directory_id    = "11111111-2222-3333-4444-123456789012"
    reason          = "default"
    subscription_id = "87654321-4321-4321-4321-210987654321"
  }

  %[1]s
`, ignoreAssetConnectionChangesBlock())

const azureConnectionAuthFile = `
  asset_connection {
    auth_mechanism  = "auth_file"
    key_file        = "/path/to/credentials/azure_auth_file.json"
    reason          = "default"
  }
`

const azureConnectionManagedIdentity = `
  asset_connection {
    auth_mechanism  = "managed_identity"
    reason          = "default"
    subscription_id = "87654321-4321-4321-4321-210987654321"
  }
`

// Output a terraform config for an Azure cloud account resource.
//
// Supports all authentication mechanisms: "client_secret", "auth_file"," and
// "managed_identity".
func testAccDSFCloudAccountConfig_Azure(resourceName string, gatewayId string, assetId string, authMechanism string) string {
	// build the asset_connection block
	var assetConnectionBlock string
	switch authMechanism {
	case "client_secret":
		assetConnectionBlock = azureConnectionClientSecret
	case "auth_file":
		assetConnectionBlock = azureConnectionAuthFile
	case "managed_identity":
		assetConnectionBlock = azureConnectionManagedIdentity
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type = "AZURE"
  
  admin_email        = "%[3]s"
  asset_display_name = "%[5]s"
  asset_id           = "%[5]s"
  gateway_id         = "%[4]s"

  %[6]s
}`, dsfCloudAccountResourceType, resourceName, testAdminEmail, gatewayId, assetId, assetConnectionBlock)
}

const gcpConnectionDefault = `
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
`

const gcpConnectionServiceAccount = `
  asset_connection {
    auth_mechanism = "service_account"
    key_file       = "/path/to/gcp/credentials/service_account.json" 
    reason         = "default"
  }
`

// Output a terraform config for a GCP cloud account resource.
//
// Supports all authentication mechanisms: "default", "service_account"
func testAccDSFCloudAccountConfig_Gcp(resourceName string, gatewayId string, assetId string, authMechanism string) string {
	var assetConnectionBlock string

	switch authMechanism {
	case "default":
		assetConnectionBlock = gcpConnectionDefault
	case "service_account":
		assetConnectionBlock = gcpConnectionServiceAccount
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type = "GCP"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  gateway_id         = "%[5]s"

  %[6]s  
}`, dsfCloudAccountResourceType, resourceName, testAdminEmail, assetId, gatewayId, assetConnectionBlock)
}
