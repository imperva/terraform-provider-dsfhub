package dsfhub

import "fmt"

// Output a terraform config for a basic data source resource.
func testAccDSFDataSourceConfig_Basic(resourceName string, adminEmail string, assetId string, gatewayId string, serverHostName string, serverType string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  admin_email        = "%[3]s"
  asset_id           = "%[4]s"
  asset_display_name = "%[4]s"
  gateway_id         = "%[5]s"
  server_host_name   = "%[6]s"
  server_type        = "%[7]s"
}`, dsfDataSourceResourceType, resourceName, adminEmail, assetId, gatewayId, serverHostName, serverType)
}

var commonBasicConnectionPassword = fmt.Sprintf(`
  asset_connection {
    auth_mechanism = "password"
    password       = "password"
    reason         = "default"
    username       = "username"
  }
  
  %[1]s
`, ignoreAssetConnectionChangesBlock())

// Output a terraform config for an AWS RDS ORACLE data source resource.
func testAccDSFDataSourceConfig_AwsRdsOracle(resourceName string, gatewayId string, assetId string, auditType string, auditPullEnabled string) string {
	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS ORACLE"

  admin_email        = "%[3]s"
  asset_display_name = "%[5]s"
  asset_id           = "%[5]s"
  audit_pull_enabled = %[7]s
  audit_type         = "%[6]s"
  gateway_id         = "%[4]s"
  server_host_name   = "test.com"
  server_port	       = "1521"
  service_name       = "ORCL"

  %[8]s
}
`, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditType, auditPullEnabled, commonBasicConnectionPassword)
}

// Output a terraform config for an AWS RDS AURORA POSTGRESQL CLUSTER data
// source resource.
func testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName string, gatewayId string, assetId string, auditType string, clusterId string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS AURORA POSTGRESQL CLUSTER"

  admin_email	       = "%[3]s"
  asset_display_name = "%[5]s"
  asset_id           = "%[5]s"
  audit_type         = "%[6]s"
  cluster_id         = "%[7]s"
  cluster_name       = "%[7]s"
  gateway_id         = "%[4]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "5432"

  %[8]s
}	
`, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditType, clusterId, commonBasicConnectionPassword)
}

// Output a terraform config for an AWS RDS AURORA POSTGRESQL data source
// reource.
func testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(resourceName string, gatewayId string, assetId string, clusterId string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS AURORA POSTGRESQL"

  admin_email        = "%[3]s"
  asset_display_name = "%[5]s"
  asset_id           = "%[5]s"
  cluster_id         = "%[6]s"
  cluster_name       = "%[6]s"
  gateway_id         = "%[4]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "5432"

  %[7]s
}	
`, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, clusterId, commonBasicConnectionPassword)
}

// Output a terraform config for an AWS RDS AURORA MYSQL CLUSTER data source
// resource.
func testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName string, gatewayId string, assetId string, auditType string, clusterId string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS AURORA MYSQL CLUSTER"

  admin_email        = "%[3]s"
  asset_display_name = "%[5]s"
  asset_id           = "%[5]s"
  audit_type         = "%[6]s"
  cluster_id         = "%[7]s"
  cluster_name       = "%[7]s"
  gateway_id         = "%[4]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "3306"

  %[8]s
}	
`, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditType, clusterId, commonBasicConnectionPassword)
}

// Output a terraform config for an AWS RDS AURORA MYSQL data source resource.
func testAccDSFDataSourceConfig_AwsRdsAuroraMysql(resourceName string, gatewayId string, assetId string, clusterId string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS AURORA MYSQL"

  admin_email        = "%[3]s"
  asset_display_name = "%[5]s"
  asset_id           = "%[5]s"
  #TODO: re-add cluster fields when supported by USC: https://onejira.imperva.com/browse/USC-2389
  #cluster_id         = "%[6]s"
  #cluster_name       = "%[6]s"
  gateway_id          = "%[4]s"
  region              = "us-east-2"
  server_host_name    = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port         = "5432"

  %[7]s
}
`, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, clusterId, commonBasicConnectionPassword)
}

// Output a terraform config for an AZURE COSMOSDB data source resource.
func testAccDSFDataSourceConfig_AzureCosmosDB(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	// handle reference to other assets
	logsDestinationAssetIdVal := testAccParseResourceAttributeReference(logsDestinationAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type = "AZURE COSMOSDB"

  admin_email               = "%[3]s"
  asset_display_name        = "%[4]s"
  asset_id                  = "%[4]s"
  audit_pull_enabled        = %[5]s
  gateway_id                = "%[6]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "my-cosmosdbsql.documents.azure.com"
  server_port               = "443"
}`, dsfDataSourceResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, gatewayId, logsDestinationAssetIdVal)
}

// Output a terraform config for an AZURE COSMOSDB MONGO data source resource.
func testAccDSFDataSourceConfig_AzureCosmosDBMongo(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	// handle reference to other assets
	logsDestinationAssetIdVal := testAccParseResourceAttributeReference(logsDestinationAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type = "AZURE COSMOSDB MONGO"

  admin_email               = "%[3]s"
  asset_display_name        = "%[4]s"
  asset_id                  = "%[4]s"
  audit_pull_enabled        = %[5]s
  gateway_id                = "%[6]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "my-cosmos-mongodb.mongo.cosmos.azure.com"
  server_port               = "443"
}`, dsfDataSourceResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, gatewayId, logsDestinationAssetIdVal)
}

// Output a terraform config for an AZURE COSMOSDB TABLE data source resource.
func testAccDSFDataSourceConfig_AzureCosmosDBTable(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	// handle reference to other assets
	logsDestinationAssetIdVal := testAccParseResourceAttributeReference(logsDestinationAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type = "AZURE COSMOSDB TABLE"

  admin_email               = "%[3]s"
  asset_display_name        = "%[4]s"
  asset_id                  = "%[4]s"
  audit_pull_enabled        = %[5]s
  gateway_id                = "%[6]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "my-cosmosdbtable.table.cosmos.azure.com"
  server_port               = "443"
}`, dsfDataSourceResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, gatewayId, logsDestinationAssetIdVal)
}

// Output a terraform config for a GCP BIGQUERY data source resource.
func testAccDSFDataSourceConfig_GcpBigQuery(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	// handle reference to other assets
	logsDestinationAssetIdVal := testAccParseResourceAttributeReference(logsDestinationAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "GCP BIGQUERY"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[6]s
  gateway_id                = "%[4]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "bigquery.googleapis.com"
  server_ip                 = "1.2.3.4"
  server_port               = "3306"
}
  `, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditPullEnabled, logsDestinationAssetIdVal)
}

// Output a terraform config for a GCP MS SQL SERVER data source resource.
func testAccDSFDataSourceConfig_GcpMsSqlServer(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string, auditType string) string {
	// handle reference to other assets
	logsDestinationAssetIdVal := testAccParseResourceAttributeReference(logsDestinationAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "GCP MS SQL SERVER"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[6]s
  audit_type                = "%[8]s"
  gateway_id                = "%[4]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "4.3.2.1"
  server_ip                 = "1.2.3.4"
  server_port               = "1433"
}
  `, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditPullEnabled, logsDestinationAssetIdVal, auditType)
}

// Output a terraform config for a GCP MYSQL data source resource.
func testAccDSFDataSourceConfig_GcpMysql(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	// handle reference to other assets
	logsDestinationAssetIdVal := testAccParseResourceAttributeReference(logsDestinationAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "GCP MYSQL"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[6]s
  gateway_id                = "%[4]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "4.3.2.1"
  server_ip                 = "1.2.3.4"
  server_port               = "3306"
}
  `, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditPullEnabled, logsDestinationAssetIdVal)
}

// Output a terraform config for a GCP POSTGRESQL data source resource.
func testAccDSFDataSourceConfig_GcpPostgresql(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	// handle reference to other assets
	logsDestinationAssetIdVal := testAccParseResourceAttributeReference(logsDestinationAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "GCP POSTGRESQL"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[6]s
  gateway_id                = "%[4]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "4.3.2.1"
  server_ip                 = "1.2.3.4"
  server_port               = "5432"
}
  `, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditPullEnabled, logsDestinationAssetIdVal)
}
