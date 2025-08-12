package dsfhub

import (
	"fmt"
)

// ******************* Connections *******************
// var commonBasicConnectionPassword = fmt.Sprintf(`
//   asset_connection {
//     auth_mechanism = "password"
//     password       = "password"
//     reason         = "default"
//     username       = "username"
//   }

//   %[1]s
// `, ignoreChangesBlock([]string{"asset_connection"}))

func createBasicPasswordConnection(database_name string) string {
	database_name = nullIfEmpty(database_name)

	connection := fmt.Sprintf(`
  asset_connection {
    auth_mechanism = "password"
    database_name  = "%[1]s"
    password       = "password"
    reason         = "default"
    username       = "username"
  }
  
  %[2]s
`, database_name, ignoreChangesBlock([]string{"asset_connection"}))
	return connection
}

const awsDynamodbConnectionDefault = `
  asset_connection {
    auth_mechanism  = "default"
    reason          = "default"
  }
`

const awsDynamodbConnectionIamRole = `
  asset_connection {
    auth_mechanism  = "iam_role"
    reason          = "default"
  }
`

const awsDynamodbConnectionProfile = `
  asset_connection {
    auth_mechanism  = "profile"
    reason          = "default"
    username        = "dsfhubuser"
  }
`

var awsDynamodbConnectionKey = fmt.Sprintf(`
  asset_connection {
    access_id       = "my-access-id"
    auth_mechanism  = "key"
    reason          = "default"
    secret_key      = "my-secret-key"
  }

  %[1]s
`, ignoreChangesBlock([]string{"asset_connection"}))

func awsDynamodbConnectionBlock(authMechanism string) string {
	var assetConnectionBlock string

	switch authMechanism {
	case "key":
		assetConnectionBlock = awsDynamodbConnectionKey
	case "profile":
		assetConnectionBlock = awsDynamodbConnectionProfile
	case "iam_role":
		assetConnectionBlock = awsDynamodbConnectionIamRole
	default:
		assetConnectionBlock = awsDynamodbConnectionDefault
	}

	return assetConnectionBlock
}

// var awsRedshiftConnectionPassword = fmt.Sprintf(`
//   asset_connection {
//     auth_mechanism = "password"
//     database_name  = "dev"
//     password       = "password"
//     reason         = "default"
//     username       = "username"
//   }

//   %[1]s
// `, ignoreChangesBlock([]string{"asset_connection"}))

const awsNeptuneClusterConnectionEc2 = `
  asset_connection {
    auth_mechanism = "ec2"
    reason         = "default"
  }
`

const awsKeyFileConnection = `
  asset_connection {
    auth_mechanism = "key_file"
    reason         = "default"
    username       = "username"
    key_file       = "/some/path"
  }
`

const awsKerberosConnection = `
  asset_connection {
    auth_mechanism = "kerberos"
    reason         = "default"
  }
`

const awsAwsCredentialsConnection = `
  asset_connection {
    auth_mechanism    = "aws_credentials"
    reason            = "default"
    username          = "username"
    database_name     = "database"
    access_id         = "access_id"
    aws_connection_id = "aws_connection_id"
  }
`

// ******************* Data Source Configs *******************
// Output a terraform config for a basic dsfhub_data_source resource
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

// Output a terraform config for a basic data source resource with a basic (password) connection.
func testAccDSFDataSourceConfig_Basic_Connection(resourceName string, adminEmail string, assetId string, gatewayId string, serverHostName string, serverType string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  admin_email        = "%[3]s"
  asset_id           = "%[4]s"
  asset_display_name = "%[4]s"
  gateway_id         = "%[5]s"
  server_host_name   = "%[6]s"
  server_type        = "%[7]s"

  %[8]s
}`,
		dsfDataSourceResourceType,
		resourceName,
		adminEmail,
		assetId,
		gatewayId,
		serverHostName,
		serverType,
		createBasicPasswordConnection(""))
}

// Output a terraform config for an AWS DOCUMENTDB CLUSTER dsfhub_data_source resource
func testAccDSFDataSourceConfig_AwsDocumentdbCluster(resourceName string, gatewayId string, assetId string, auditPullEnabled string, parentAssetId string) string {
	// handle reference to other assets
	parentAssetIdVal := parseResourceAttributeReference(parentAssetId)

	// convert audit_pull_enabled to "null" if empty
	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS DOCUMENTDB CLUSTER"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  audit_pull_enabled = %[5]s
  audit_type         = "LOG_GROUP"
  gateway_id         = "%[6]s"
  parent_asset_id    = %[8]s
  region             = "us-east-2"
  server_host_name   = "my-docdb-cluster.cp9pk8rsfzja.us-east-1.docdb.amazonaws.com"
  server_ip          = "%[4]s"
  server_port        = "27017"

  %[7]s
}
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		assetId,
		auditPullEnabled,
		gatewayId,
		createBasicPasswordConnection(""),
		parentAssetIdVal)
}

// Output a terraform config for an AWS DYNAMODB dsfhub_data_source resource
func testAccDSFDataSourceConfig_AwsDynamodb(resourceName string, gatewayId string, assetId string, auditPullEnabled string, authMechanism string) string {
	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	assetConnectionBlock := awsDynamodbConnectionBlock(authMechanism)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS DYNAMODB"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  arn                = "%[4]s"
  audit_pull_enabled = %[5]s
  available_regions  = ["us-west-1", "us-east-1"]
  gateway_id         = "%[6]s"
  region             = "us-east-1"
  server_port        = "27017"

  %[7]s
}
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		assetId,
		auditPullEnabled,
		gatewayId,
		assetConnectionBlock)
}

// Output a terraform config for an AWS NEPTUNE CLUSTER dsfhub_data_source resource
func testAccDSFDataSourceConfig_AwsNeptuneCluster(resourceName string, gatewayId string, assetId string, auditType string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS NEPTUNE CLUSTER"

  admin_email	       = "%[3]s"
  asset_display_name = "%[5]s"
  asset_id           = "%[5]s"
  audit_type         = "%[6]s"
  gateway_id         = "%[4]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "5432"

  %[7]s
}	
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		auditType,
		awsNeptuneClusterConnectionEc2)
}

// Output a terraform config for an AWS RDS ORACLE dsfhub_data_source resource
//
// TODO: determine if depends_on is needed
//
//	depends_on               = [%[9]s]
func testAccDSFDataSourceConfig_AwsRdsOracle(resourceName string, gatewayId string, assetId string, auditType string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)
	// dependsOn = nullIfEmpty(dependsOn)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS ORACLE"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[7]s
  audit_type                = "%[6]s"
  gateway_id                = "%[4]s"
  logs_destination_asset_id = %[8]s
  server_host_name          = "test.com"
  server_port	              = "1521"
  service_name              = "ORCL"

  %[9]s
}
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		auditType,
		auditPullEnabled,
		logsDestinationAssetIdVal,
		// dependsOn,
		createBasicPasswordConnection(""))
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
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		auditType,
		clusterId,
		createBasicPasswordConnection(""))
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
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		clusterId,
		createBasicPasswordConnection(""))
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
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		auditType,
		clusterId,
		createBasicPasswordConnection(""))
}

// Output a terraform config for an AWS RDS AURORA MYSQL dsfhub_data_source resource
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
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		clusterId,
		createBasicPasswordConnection(""))
}

// Output a terraform config for an AWS RDS MARIADB dsfhub_data_source resource
func testAccDSFDataSourceConfig_AwsRdsMariadb(resourceName string, gatewayId string, assetId string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS MARIADB"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  database_name      = "my-database"
  gateway_id         = "%[5]s"
  #parent_asset_id    = "todo"
  region             = "us-east-2"
  server_host_name   = "my-database.xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "3306"
}  
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		assetId,
		gatewayId)
}

// Output a terraform config for an AWS RDS MS SQL SERVER dsfhub_data_source resource
func testAccDSFDataSourceConfig_AwsRdsMsSqlServer(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "AWS RDS MS SQL SERVER"

  admin_email               = "%[3]s"
  asset_display_name        = "%[4]s"
  asset_id                  = "%[4]s"
  audit_pull_enabled        = %[5]s
  database_name             = "master"
  gateway_id                = "%[6]s"
  logs_destination_asset_id = %[7]s
  #parent_asset_id           = "todo"
  region                    = "us-east-2"
  server_host_name          = "my-database.xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port               = "3306"
}  
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		assetId,
		auditPullEnabled,
		gatewayId,
		logsDestinationAssetIdVal)
}

// Output a terraform config for an AWS RDS MYSQL dsfhub_data_source resource
func testAccDSFDataSourceConfig_AwsRdsMysql(resourceName string, gatewayId string, assetId string, auditType string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS MYSQL"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  audit_type         = "%[5]s"
  database_name      = "master"
  gateway_id         = "%[6]s"
  #parent_asset_id    = "todo"
  region             = "us-east-2"
  server_host_name   = "my-rds-name.cp9pk8xxxxxx.<region>.rds.amazonaws.com"
  server_port        = "3306"
}  
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		assetId,
		auditType,
		gatewayId)
}

// Output a terraform config for an AWS RDS POSTGRESQL dsfhub_data_source resource
func testAccDSFDataSourceConfig_AwsRdsPostgresql(resourceName string, gatewayId string, assetId string, auditType string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "AWS RDS POSTGRESQL"

  admin_email        = "%[3]s"
  asset_display_name = "%[4]s"
  asset_id           = "%[4]s"
  audit_type         = "%[5]s"
  database_name      = "postgresql"
  gateway_id         = "%[6]s"
  #parent_asset_id    = "todo"
  region             = "us-east-2"
  server_host_name   = "my-database.xxxxr5ierus0.us-east-1.rds.amazonaws.com"
  server_port        = "3306"
}  
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		assetId,
		auditType,
		gatewayId)
}

// Output a terraform config for an AWS REDSHIFT dsfhub_data_source resource
func testAccDSFDataSourceConfig_AwsRedshift(resourceName string, gatewayId string, assetId string, auditType string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "AWS REDSHIFT"

  admin_email               = "%[3]s"
  asset_display_name        = "%[4]s"
  asset_id                  = "%[4]s"
  audit_type                = "%[5]s"
  audit_pull_enabled        = %[6]s
  gateway_id                = "%[7]s"
  logs_destination_asset_id = %[8]s
  # parent_asset_id           = "todo"
  region                    = "us-east-2"
  server_host_name          = "my-database.xxxxr5ierus0.us-east-1.rds.amazonaws.com"
  server_ip                 = "1.2.3.4"
  server_port               = "5439"

  %[9]s
}  
`,
		dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		assetId,
		auditType,
		auditPullEnabled,
		gatewayId,
		logsDestinationAssetIdVal,
		createBasicPasswordConnection("dev"), // TODO: support aws_credentials
	)
}

// Output a terraform config for an AZURE COSMOSDB (NoSQL) dsfhub_data_source resource
func testAccDSFDataSourceConfig_AzureCosmosDBNosql(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

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

// Output a terraform config for an AZURE COSMOSDB MONGO dsfhub_data_source resource
func testAccDSFDataSourceConfig_AzureCosmosDBMongo(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

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

// Output a terraform config for an AZURE COSMOSDB TABLE dsfhub_data_source resource
func testAccDSFDataSourceConfig_AzureCosmosDBTable(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

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

// Output a terraform config for an AZURE MS SQL SERVER dsfhub_data_source resource
func testAccDSFDataSourceConfig_AzureMsSqlServer(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "AZURE MS SQL SERVER"

  admin_email               = "%[3]s"
  asset_display_name        = "%[4]s"
  asset_id                  = "%[4]s"
  audit_pull_enabled        = %[5]s
  database_name             = "master"
  gateway_id                = "%[6]s"
  location                  = "us-east1"
  logs_destination_asset_id = %[7]s
  server_host_name          = "my-sql-server.database.windows.net"
  server_ip                 = "1.2.3.4"
  server_port               = "1433"

  %[8]s
}`, dsfDataSourceResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, gatewayId, logsDestinationAssetIdVal, createBasicPasswordConnection(""))
}

// Output a terraform config for an AZURE MYSQL FLEXIBLE dsfhub_data_source resource
func testAccDSFDataSourceConfig_AzureMysqlFlexible(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "AZURE MYSQL FLEXIBLE"

  admin_email               = "%[3]s"
  asset_display_name        = "%[4]s"
  asset_id                  = "%[4]s"
  audit_pull_enabled        = %[5]s
  gateway_id                = "%[6]s"
  location                  = "us-east2"
  logs_destination_asset_id = %[7]s
  server_host_name          = "someflexdatabase.azure.com"
  server_port               = "3306"

  %[8]s
}`, dsfDataSourceResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, gatewayId, logsDestinationAssetIdVal, createBasicPasswordConnection(""))
}

// Output a terraform config for an AZURE POSTGRESQL FLEXIBLE dsfhub_data_source resource
func testAccDSFDataSourceConfig_AzurePostgresqlFlexible(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "AZURE POSTGRESQL FLEXIBLE"

  admin_email               = "%[3]s"
  asset_display_name        = "%[4]s"
  asset_id                  = "%[4]s"
  audit_pull_enabled        = %[5]s
  database_name             = "postgres"
  gateway_id                = "%[6]s"
  location                  = "us-east2"
  logs_destination_asset_id = %[7]s
  server_host_name          = "someflexdatabase.azure.com"
  server_port               = "5432"

  %[8]s
}`, dsfDataSourceResourceType, resourceName, testAdminEmail, assetId, auditPullEnabled, gatewayId, logsDestinationAssetIdVal, createBasicPasswordConnection(""))
}

// Output a terraform config for a GCP ALLOYDB dsfhub_data_source resource
func testAccDSFDataSourceConfig_GcpAlloydbCluster(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string, clusterMemberId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "GCP ALLOYDB POSTGRESQL CLUSTER"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[6]s
  cluster_id                = "my-cluster"
  cluster_member_id         = "%[8]s"
  gateway_id                = "%[4]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "4.3.2.1"
  server_ip                 = "1.2.3.4"
  server_port               = "5432"
}
  `, dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		auditPullEnabled,
		logsDestinationAssetIdVal,
		clusterMemberId)
}

// Output a terraform config for a GCP ALLOYDB dsfhub_data_source resource
func testAccDSFDataSourceConfig_GcpAlloydb(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string, clusterMemberId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "GCP ALLOYDB POSTGRESQL"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[6]s
  cluster_id                = "my-cluster"
  cluster_member_id         = "%[8]s"
  gateway_id                = "%[4]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "4.3.2.1"
  server_ip                 = "1.2.3.4"
  server_port               = "5432"
}
  `, dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		auditPullEnabled,
		logsDestinationAssetIdVal,
		clusterMemberId)
}

// Output a terraform config for an AZURE SQL MANAGED INSTANCE dsfhub_data_source resource
func testAccDSFDataSourceConfig_AzureSqlManagedInstance(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "AZURE SQL MANAGED INSTANCE"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[6]s
  gateway_id                = "%[4]s"
  location                  = "us-west2"
  logs_destination_asset_id = %[7]s
  server_host_name          = "4.3.2.1"
  server_ip                 = "1.2.3.4"
  server_port               = "1433"
}
`, dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		auditPullEnabled,
		logsDestinationAssetIdVal)
}

// Output a terraform config for a GCP BIGQUERY dsfhub_data_source resource
func testAccDSFDataSourceConfig_GcpBigQuery(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

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
  server_ip                 = "bigquery.googleapis.com"
  server_port               = "443"
}
  `, dsfDataSourceResourceType,
		resourceName,
		testAdminEmail,
		gatewayId,
		assetId,
		auditPullEnabled,
		logsDestinationAssetIdVal)
}

// Output a terraform config for a GCP MS SQL SERVER dsfhub_data_source resource
func testAccDSFDataSourceConfig_GcpMsSqlServer(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string, auditType string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

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

// Output a terraform config for a GCP MYSQL dsfhub_data_source resource
func testAccDSFDataSourceConfig_GcpMysql(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

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

// Output a terraform config for a GCP POSTGRESQL dsfhub_data_source resource
func testAccDSFDataSourceConfig_GcpPostgresql(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)

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

// Output a terraform config for a GCP SPANNER dsfhub_data_source resource
func testAccDSFDataSourceConfig_GcpSpanner(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string, durationThreshold string) string {
	logsDestinationAssetIdVal := parseResourceAttributeReference(logsDestinationAssetId)

	auditPullEnabled = nullIfEmpty(auditPullEnabled)
	durationThreshold = nullIfEmpty(durationThreshold)

	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type               = "GCP SPANNER"

  admin_email               = "%[3]s"
  asset_display_name        = "%[5]s"
  asset_id                  = "%[5]s"
  audit_pull_enabled        = %[6]s
  duration_threshold        = %[8]s
  gateway_id                = "%[4]s"
  logs_destination_asset_id = %[7]s
  server_host_name          = "my-spanner-instance"
  server_ip                 = "1.2.3.4"
  server_port               = "443"
}
  `, dsfDataSourceResourceType, resourceName, testAdminEmail, gatewayId, assetId, auditPullEnabled, logsDestinationAssetIdVal, durationThreshold)
}
