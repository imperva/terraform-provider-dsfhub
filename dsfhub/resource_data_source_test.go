package dsfhub

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDSFDataSource_Basic(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const resourceName = "basic_test_data_source"
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDSFDataSourceConfig_Basic(
					resourceName,
					testAdminEmail,
					testArn,
					gatewayId,
					testServerHostName,
					testDSServerType,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
				),
			},
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFDataSource_AwsDocumentdbCluster(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = testAwsRdsClusterPrefix + "my-docdb-cluster"
		resourceName = "aws_documentdb_onboarding"

		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/docdb/my-docdb-cluster/audit:*"
		logGroupResourceName = resourceName + "_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsDocumentdbCluster(resourceName, gatewayId, assetId, ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@docdb.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsDynamodbBasic(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = "aws_dynamodb_basic"
		resourceName = "aws_dynamodb_basic"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// check various auth mechanisms
			{Config: testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "iam_role")},
			{Config: testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "key")},
			{Config: testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "profile")},
			{Config: testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "default")},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFDataSource_AwsDynamodbCloudwatch(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = "aws_dynamodb_cloudwatch_onboarding"
		resourceName = "aws_dynamodb_cloudwatch_onboarding"

		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/events/Dynamodb:*"
		logGroupResourceName = resourceName + "_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "default"),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@dynamodb.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsDynamodbS3(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = "aws_dynamodb_s3_onboarding"
		resourceName = "aws_dynamodb_s3_onboarding"

		s3BucketAssetId      = testAwsS3BucketPrefix + "dynamodb-s3-bucket"
		s3BucketResourceName = resourceName + "_bucket"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	s3BucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, s3BucketResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "default"),
		testAccDSFLogAggregatorConfig_AwsS3(s3BucketResourceName, gatewayId, s3BucketAssetId, resourceTypeAndName+".asset_id", "true", "DYNAMODB"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[s3BucketResourceTypeAndName] = make(map[string]string)
	initialChecks[s3BucketResourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsNeptuneClusterSlowQuery(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = testAwsRdsClusterPrefix + "my-neptune-cluster"
		resourceName = "aurora_neptune_cluster_onboarding"

		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/my-neptune-cluster/audit:*"
		logGroupResourceName = resourceName + "_log_group"

		slowLogGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/my-neptune-cluster/slowquery:*"
		slowLogGroupResourceName = resourceName + "_slow_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)
	slowLogGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowLogGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsNeptuneCluster(resourceName, gatewayId, assetId, ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName+".asset_id", true, "AWS_NEPTUNE_SLOW", logGroupResourceTypeAndName),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@neptune.service"
	initialChecks[slowLogGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[slowLogGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[slowLogGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@neptune-slow-query.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsAuroraPostgresqlClusterCloudWatch(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = testAwsRdsClusterPrefix + "my-aurorapostgresql-cluster"
		resourceName = "aurora_postgresql_cluster_onboarding"

		instanceAssetId      = assetId + "-writer"
		instanceResourceName = resourceName + "_instance"

		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/my-cluster/postgresql:*"
		logGroupResourceName = resourceName + "_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName, gatewayId, assetId, "LOG_GROUP", resourceName),
		testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@aurora-postgresql.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"
	refreshChecks[instanceResourceTypeAndName] = make(map[string]string)
	refreshChecks[instanceResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsAuroraPostgresqlClusterKinesis(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = testAwsRdsClusterPrefix + "my-aurorapostgresql-kinesis-cluster"
		resourceName = "aurora_postgresql_kinesis_cluster_onboarding"

		instanceAssetId      = assetId + "-writer"
		instanceResourceName = resourceName + "_instance"

		kinesisAssetId      = testAwsKinesisPrefix + resourceName
		kinesisResourceName = resourceName + "_kinesis_stream"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	kinesisResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, kinesisResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway, check that the kinesis stream is connected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName, gatewayId, assetId, "", resourceName),
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
					testAccDSFLogAggregatorConfig_AwsKinesis(kinesisResourceName, gatewayId, kinesisAssetId, resourceTypeAndName+".asset_id", true, "KINESIS"),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(kinesisResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(kinesisResourceTypeAndName, "gateway_service", "gateway-kinesis@aurora-postgresql.service"),
				),
			},
			// refresh and verify DB assets are connected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(instanceResourceTypeAndName, "audit_pull_enabled", "true"),
				),
			},
			// update audit_type -> reconnect asset to gateway
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName, gatewayId, assetId, "", resourceName),
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
					testAccDSFLogAggregatorConfig_AwsKinesis(kinesisResourceName, gatewayId, kinesisAssetId, resourceTypeAndName+".asset_id", true, "KINESIS_AGGREGATED"),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(kinesisResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(kinesisResourceTypeAndName, "gateway_service", "gateway-kinesis@aurora-postgresql-aggregated.service"),
				),
			},
			// refresh and verify DB assets are still connected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(instanceResourceTypeAndName, "audit_pull_enabled", "true"),
				),
			},
			// disconnect gateway, check that the kinesis stream is disconnected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName, gatewayId, assetId, "", resourceName),
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
					testAccDSFLogAggregatorConfig_AwsKinesis(kinesisResourceName, gatewayId, kinesisAssetId, resourceTypeAndName+".asset_id", false, "KINESIS_AGGREGATED"),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(kinesisResourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(kinesisResourceTypeAndName, "gateway_service", ""),
				),
			},
			// refresh and verify DB assets are disconnected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(instanceResourceTypeAndName, "audit_pull_enabled", "false"),
				),
			},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFDataSource_AwsRdsAuroraMysqlClusterCloudWatchSlowQuery(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = testAwsRdsClusterPrefix + "my-auroramysql-cluster"
		resourceName = "aurora_mysql_cluster_onboarding"

		instanceAssetId      = assetId + "-writer"
		instanceResourceName = resourceName + "_instance"

		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/my-aurora-cluster/audit:*"
		logGroupResourceName = resourceName + "_log_group"

		slowLogGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/my-aurora-cluster/slowquery:*"
		slowLogGroupResourceName = resourceName + "_slow_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	//TODO: check that instance asset is connected once fixed: https://onejira.imperva.com/browse/SR-2046
	// instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)
	slowLogGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowLogGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName, gatewayId, assetId, "", resourceName),
		testAccDSFDataSourceConfig_AwsRdsAuroraMysql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName+".asset_id", true, "AWS_RDS_AURORA_MYSQL_SLOW", logGroupResourceTypeAndName),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@aurora-mysql.service"
	initialChecks[slowLogGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[slowLogGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[slowLogGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@aurora-mysql-slow-query.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"
	// refreshChecks[instanceResourceTypeAndName] = make(map[string]string)
	// refreshChecks[instanceResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsAuroraMysqlClusterKinesis(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = testAwsRdsClusterPrefix + "my-auroramysql-kinesis-cluster"
		resourceName = "aurora_mysql_kinesis_cluster_onboarding"

		instanceAssetId      = assetId + "-writer"
		instanceResourceName = resourceName + "_instance"

		kinesisAssetId      = testAwsKinesisPrefix + resourceName
		kinesisResourceName = resourceName + "_kinesis_stream"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	// instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	kinesisResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, kinesisResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName, gatewayId, assetId, "", resourceName),
		testAccDSFDataSourceConfig_AwsRdsAuroraMysql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
		testAccDSFLogAggregatorConfig_AwsKinesis(kinesisResourceName, gatewayId, kinesisAssetId, resourceTypeAndName+".asset_id", true, "KINESIS"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[kinesisResourceTypeAndName] = make(map[string]string)
	initialChecks[kinesisResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[kinesisResourceTypeAndName]["gateway_service"] = "gateway-kinesis@aurora-mysql.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"
	// refreshChecks[instanceResourceTypeAndName] = make(map[string]string)
	// refreshChecks[instanceResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsMariadb(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "rds_mariadb_cloudwatch"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/rds_mariadb_cloudwatch/audit:*"
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsMariadb(resourceName, gatewayId, assetId),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@mariadb.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsMsSqlServer(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "rds_ms_sql_server"
		assetId      = testAwsRdsDbPrefix + resourceName

		s3BucketResourceName = resourceName + "_bucket"
		s3BucketAssetId      = testAwsS3BucketPrefix + resourceName

		cloudAccountResourceName = "aws_rds_ms_sql_server_role"
		cloudAccountAssetId      = testAwsAccountArnPrefix + cloudAccountResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	s3BucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, s3BucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Aws(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		testAccDSFDataSourceConfig_AwsRdsMsSqlServer(resourceName, gatewayId, assetId, "true", s3BucketResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AwsS3(s3BucketResourceName, gatewayId, s3BucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[s3BucketResourceTypeAndName] = make(map[string]string)
	refreshChecks[s3BucketResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsMysql(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "rds_mysql_cloudwatch"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/rds_mysql_cloudwatch/audit:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsMysql(resourceName, gatewayId, assetId, "LOG_GROUP"),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@mysql.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsMysqlSlowQuery(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "rds_mysql_cloudwatch_slow_query"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/rds_mysql_cloudwatch_slow_query/audit:*"

		slowLogGroupResourceName = resourceName + "_slow_log_group"
		slowLogGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/rds_mysql_cloudwatch_slow_query/slowquery:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)
	slowLogGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowLogGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsMysql(resourceName, gatewayId, assetId, ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName+".asset_id", true, "AWS_RDS_MYSQL_SLOW", logGroupResourceTypeAndName),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@mysql.service"
	initialChecks[slowLogGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[slowLogGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[slowLogGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@mysql-slow-query.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsOracleCloudwatch(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "rds_oracle_cloudwatch"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/rds_oracle_cloudwatch/audit:*"
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "LOG_GROUP", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@oracle-rds.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRdsOracleUnified(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "rds_oracle_unified"
		assetId      = testAwsRdsDbPrefix + resourceName
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "UNIFIED", "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-odbc@oracle_unified.service"),
				),
			},
			// update audit_type -> reconnect asset to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "UNIFIED_AGGREGATED", "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-odbc@oracle_unified_aggregated.service"),
				),
			},
			// disconnect asset
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "UNIFIED_AGGREGATED", "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", ""),
				),
			},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFDataSource_AwsRdsPostgresql(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "rds_postgresql_cloudwatch"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/rds_postgresql_cloudwatch/postgres:*"
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsPostgresql(resourceName, gatewayId, assetId, "LOG_GROUP"),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[logGroupResourceTypeAndName] = make(map[string]string)
	initialChecks[logGroupResourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[logGroupResourceTypeAndName]["gateway_service"] = "gateway-aws@postgresql.service"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[resourceTypeAndName] = make(map[string]string)
	refreshChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

// TODO: determine why logs_destination_asset_id does not appear in USC and skip
// this test until fixed
func TestAccDSFDataSource_AwsRedshiftS3(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "aws_redshift_s3"
		assetId      = testAwsRedshiftPrefix + resourceName

		s3BucketResourceName = resourceName + "_bucket"
		s3BucketAssetId      = testAwsS3BucketPrefix + s3BucketResourceName

		cloudAccountResourceName = "aws_redshift_s3_bucket_role"
		cloudAccountAssetId      = testAwsAccountArnPrefix + cloudAccountResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	s3BucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, s3BucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Aws(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		testAccDSFDataSourceConfig_AwsRedshift(resourceName, gatewayId, assetId, "", "true", s3BucketResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AwsS3(s3BucketResourceName, gatewayId, s3BucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "", "REDSHIFT"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[resourceTypeAndName]["gateway_service"] = ""

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[s3BucketResourceTypeAndName] = make(map[string]string)
	refreshChecks[s3BucketResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, nil, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AwsRedshiftTable(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "aws_redshift_table"
		assetId      = testAwsRedshiftPrefix + resourceName
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRedshift(resourceName, gatewayId, assetId, "TABLE", "true", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"
	initialChecks[resourceTypeAndName]["gateway_service"] = "gateway-odbc@redshift.service"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, nil, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AzureCosmosDB(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "azure_cosmosdb_sql_connect_disconnect_gateway"
		assetId      = testAzurePrefix + "Microsoft.DocumentDB/databaseAccounts/my-cosmos-sql"

		eventhubResourceName = "azure-cosmosdb-sql-eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureCosmosDB(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Cosmos_SQL"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[eventhubResourceTypeAndName] = make(map[string]string)
	refreshChecks[eventhubResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AzureCosmosDBMongo(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "azure_cosmosdb_mongo_connect_disconnect_gateway"
		assetId      = testAzurePrefix + "Microsoft.DocumentDB/databaseAccounts/my-cosmos-mongodb"

		eventhubResourceName = "azure-cosmosdb-mongo-eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureCosmosDBMongo(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Cosmos_Mongo"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[eventhubResourceTypeAndName] = make(map[string]string)
	refreshChecks[eventhubResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AzureCosmosDBTable(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "azure_cosmosdb_table_connect_disconnect_gateway"
		assetId      = testAzurePrefix + "Microsoft.DocumentDB/databaseAccounts/my-cosmos-table"

		eventhubResourceName = "azure-cosmosdb-table-eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureCosmosDBTable(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Cosmos_Table"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[eventhubResourceTypeAndName] = make(map[string]string)
	refreshChecks[eventhubResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AzureMsSqlServer(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "azure_sql_server_connect_disconnect_gateway"
		assetId      = testAzurePrefix + "Microsoft.Sql/servers/my-sql-server"

		eventhubResourceName = "azure-sql-server-eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureMsSqlServer(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Sql"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[eventhubResourceTypeAndName] = make(map[string]string)
	refreshChecks[eventhubResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AzurePostgresqlFlexible(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "azure_postgresql_flexible_connect_disconnect_gateway"
		assetId      = testAzurePrefix + "Microsoft.DBforPostgreSQL/flexibleservers/someflexdatabase"

		eventhubResourceName = "azure-postgresql-flexible-eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzurePostgresqlFlexible(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Postgresql_Flexible"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[eventhubResourceTypeAndName] = make(map[string]string)
	refreshChecks[eventhubResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_AzureSqlManagedInstance(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "azure_sql_managed_instance_connect_disconnect_gateway"
		assetId      = testAzurePrefix + "Microsoft.Sql/managedinstances/my-managed-instance"

		eventhubResourceName = "azure-sql-managed-instance-eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureSqlManagedInstance(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "AzureSQL_Managed"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[eventhubResourceTypeAndName] = make(map[string]string)
	refreshChecks[eventhubResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_GcpBigQuery(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "gcp_bigquery_connect_disconnect_gateway"
		assetId      = "projects/my-project-name/bigquery"

		pubsubResourceName = "bigquery-subscription"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpBigQuery(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "BIGQUERY", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[pubsubResourceTypeAndName] = make(map[string]string)
	refreshChecks[pubsubResourceTypeAndName]["audit_pull_enabled"] = "true"
	refreshChecks[pubsubResourceTypeAndName]["gateway_service"] = "gateway-gcp@bigquery.service"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_GcpMsSqlServer(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "gcp_ms_sql_server_connect_disconnect_gateway"
		assetId      = "my-project:us-west-1:sql-server-instance-1"

		bucketResourceName = "sql-server-instance-1-bucket"
		bucketAssetId      = "my-project:" + bucketResourceName

		cloudAccountResourceName = "sql-server-cloud-account-1"
		cloudAccountAssetId      = "my_service_account_1@my-project.iam.gserviceaccount.com:my-project"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	bucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, bucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Gcp(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		testAccDSFDataSourceConfig_GcpMsSqlServer(resourceName, gatewayId, assetId, "true", bucketResourceTypeAndName+".asset_id", ""),
		testAccDSFLogAggregatorConfig_GcpCloudStorageBucket(bucketResourceName, gatewayId, bucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[bucketResourceTypeAndName] = make(map[string]string)
	refreshChecks[bucketResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_GcpMsSqlServerManyToOne(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceNameProd = "gcp_ms_sql_server_many_to_one_prod"
		assetIdProd      = "my-project:us-west-1:sql-server-instance-2-prod"
		resourceNameDev  = "gcp_ms_sql_server_many_to_one_dev"
		assetIdDev       = "my-project:us-west-1:sql-server-instance-2-dev"

		bucketResourceName = "sql-server-instance-2-bucket"
		bucketAssetId      = "my-project:" + bucketResourceName

		cloudAccountResourceName = "sql-server-cloud-account-2"
		cloudAccountAssetId      = "my_service_account_2@my-project.iam.gserviceaccount.com:my-project"
	)

	resourceTypeAndNameProd := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceNameProd)
	resourceTypeAndNameDev := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceNameDev)
	bucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, bucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Gcp(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		testAccDSFDataSourceConfig_GcpMsSqlServer(resourceNameProd, gatewayId, assetIdProd, "true", bucketResourceTypeAndName+".asset_id", ""),
		testAccDSFDataSourceConfig_GcpMsSqlServer(resourceNameDev, gatewayId, assetIdDev, "true", bucketResourceTypeAndName+".asset_id", ""),
		testAccDSFLogAggregatorConfig_GcpCloudStorageBucket(bucketResourceName, gatewayId, bucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndNameProd] = make(map[string]string)
	initialChecks[resourceTypeAndNameProd]["audit_pull_enabled"] = "true"
	initialChecks[resourceTypeAndNameDev] = make(map[string]string)
	initialChecks[resourceTypeAndNameDev]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[bucketResourceTypeAndName] = make(map[string]string)
	refreshChecks[bucketResourceTypeAndName]["audit_pull_enabled"] = "true"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, false, ""),
	})
}

func TestAccDSFDataSource_GcpMysql(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "gcp_mysql_connect_disconnect_gateway"
		assetId      = "my-project:us-west-1:mysql-instance-1"

		pubsubResourceName = "mysql-instance-1-subscription"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpMysql(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "MYSQL", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[pubsubResourceTypeAndName] = make(map[string]string)
	refreshChecks[pubsubResourceTypeAndName]["audit_pull_enabled"] = "true"
	refreshChecks[pubsubResourceTypeAndName]["gateway_service"] = "gateway-gcp@mysql.service"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_GcpMysqlSlowQuery(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "gcp_mysql_slow_query"
		assetId      = "my-project:us-west-1:mysql-instance-2"

		auditPubsubResourceName = "mysql-instance-2-audit-subscription"
		auditPubsubAssetId      = testPubsubSubscriptionPrefix + auditPubsubResourceName

		slowQueryPubsubResourceName = "mysql-instance-2-slow-query-subscription"
		slowQueryPubsubAssetId      = testPubsubSubscriptionPrefix + slowQueryPubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	auditPubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, auditPubsubResourceName)
	slowQueryPubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowQueryPubsubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpMysql(resourceName, gatewayId, assetId, "true", auditPubsubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_GcpPubsub(auditPubsubResourceName, gatewayId, auditPubsubAssetId, "default", "", "", "MYSQL", ""),
		testAccDSFLogAggregatorConfig_GcpPubsub(slowQueryPubsubResourceName, gatewayId, slowQueryPubsubAssetId, "default", "", "true", "GCP_MYSQL_SLOW", "GCP MYSQL"),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[auditPubsubResourceTypeAndName] = make(map[string]string)
	refreshChecks[auditPubsubResourceTypeAndName]["audit_pull_enabled"] = "true"
	refreshChecks[auditPubsubResourceTypeAndName]["gateway_service"] = "gateway-gcp@mysql.service"
	refreshChecks[slowQueryPubsubResourceTypeAndName] = make(map[string]string)
	refreshChecks[slowQueryPubsubResourceTypeAndName]["audit_pull_enabled"] = "true"
	refreshChecks[slowQueryPubsubResourceTypeAndName]["gateway_service"] = "gateway-gcp@mysql-slow-query.service"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func TestAccDSFDataSource_GcpPostgresql(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "gcp_postgresql_connect_disconnect_gateway"
		assetId      = "my-project:us-west-1:postgresql-instance-1"

		pubsubResourceName = "postgresql-instance-1-subscription"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	// initial config
	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpPostgresql(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "POSTGRESQL", ""),
	)

	// resource attribute checks
	var initialChecks = make(map[string]map[string]string)
	initialChecks[resourceTypeAndName] = make(map[string]string)
	initialChecks[resourceTypeAndName]["audit_pull_enabled"] = "true"

	var refreshChecks = make(map[string]map[string]string)
	refreshChecks[pubsubResourceTypeAndName] = make(map[string]string)
	refreshChecks[pubsubResourceTypeAndName]["audit_pull_enabled"] = "true"
	refreshChecks[pubsubResourceTypeAndName]["gateway_service"] = "gateway-gcp@postgresql.service"

	// run test
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     connectDisconnectGatewaySteps(config, initialChecks, refreshChecks, true, resourceTypeAndName),
	})
}

func testAccDSFDataSourceId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccDSFDataSourceId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != dsfDataSourceResourceType {
			continue
		}
		return fmt.Sprintf("%s", rs.Primary.ID), nil
	}
	return "", fmt.Errorf("error finding DSF dataSourceId")
}

// Confirm assets are destroyed after an acceptance test run
func testAccDSFDataSourceDestroy(state *terraform.State) error {
	log.Printf("[INFO] Running test testAccDSFDataSourceDestroy")
	// allow "disableAsset" playbook enough time to run
	time.Sleep(5 + time.Second)

	// check if asset still exists on hub
	client := testAccProvider.Meta().(*Client)
	for _, res := range state.RootModule().Resources {
		if res.Type != dsfDataSourceResourceType {
			continue
		}
		assetId := res.Primary.ID
		readDSFDataSourceResponse, err := client.ReadDSFDataSource(assetId)
		if readDSFDataSourceResponse.Errors == nil {
			return fmt.Errorf("DSF Data Source %s should have received an error in the response", assetId)
		}
		if err == nil {
			return fmt.Errorf("DSF Data Source %s still exists", assetId)
		}
	}
	return nil
}
