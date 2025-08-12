package dsfhub

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDSFDataSource_basicNoConnection(t *testing.T) {
	gatewayId := getGatewayId(t)

	const resourceName = "basic_no_conn_test_data_source"
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	configNoConnection, _ := createBasicConfigs(resourceName, testArn, gatewayId)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps: []resource.TestStep{
			{
				Config: configNoConnection,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "asset_connection.#", "0"),
					resource.TestCheckNoResourceAttr(resourceTypeAndName, "asset_connection.0.auth_mechanism"),
				),
			},
			createValidateImportStep(resourceTypeAndName),
		},
	})
}

func TestAccDSFDataSource_addRemoveConnection(t *testing.T) {
	// TODO: stop skipping test when id is added to schema
	for _, v := range []string{"4.16", "4.17", "4.18", "4.19", "15.0", "15.1", "master"} {
		skipTestForKnownIssue(t, v, "https://onejira.imperva.com/browse/SR-3396", true)
	}

	gatewayId := getGatewayId(t)

	const resourceName = "add_remove_conn_test_data_source"
	const asset_id = testArn + "-2"
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	configNoConnection, configWithConnection := createBasicConfigs(resourceName, asset_id, gatewayId)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps: []resource.TestStep{
			{
				Config: configWithConnection,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "asset_connection.#", "1"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "asset_connection.0.auth_mechanism", "password"),
				),
			},
			createValidateImportStep(resourceTypeAndName),
			{
				Config: configNoConnection,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "asset_connection.#", "0"),
					resource.TestCheckNoResourceAttr(resourceTypeAndName, "asset_connection.0.auth_mechanism"),
				),
			},
			createValidateImportStep(resourceTypeAndName),
		},
	})
}

func TestAccDSFDataSource_AwsDocumentdbCluster(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		clusterId    = "my-docdb-cluster"
		resourceName = "aws_documentdb"
		assetId      = testAwsRdsClusterPrefix + clusterId

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/docdb/" + clusterId + "/audit:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsDocumentdbCluster(resourceName, gatewayId, assetId, "", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			logGroupResourceTypeAndName: {
				"gateway_service": "gateway-aws@docdb.service",
			},
		},
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsDynamodbAllConnections(t *testing.T) {
	t.Skipf("Skipping test %s, details: due to SR-3677.", t.Name())

	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_dynamodb_all_connections"
		assetId      = testAwsDynamoDbPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps: []resource.TestStep{
			{
				Config: testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "iam_role"),
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr(resourceTypeAndName, "asset_connection.0.auth_mechanism", "iam_role")),
			},
			{
				Config:             testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "key"),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "asset_connection.0.auth_mechanism", "key"),
				),
			},
			{
				Config:             testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "profile"),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "asset_connection.0.auth_mechanism", "profile"),
				),
			},
			{
				Config:             testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "default"),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "asset_connection.0.auth_mechanism", "default"),
				),
			},
			createValidateImportStep(resourceTypeAndName),
		},
	})
}

func TestAccDSFDataSource_AwsDynamodbCloudwatch(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_dynamodb_cloudwatch"
		assetId      = testAwsDynamoDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/events/Dynamodb:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "default"),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			logGroupResourceTypeAndName: {
				"gateway_service": "gateway-aws@dynamodb.service",
			},
		},
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsDynamodbS3(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_dynamodb_s3"
		assetId      = testAwsDynamoDbPrefix + resourceName

		s3BucketAssetId      = testAwsS3BucketPrefix + "dynamodb-s3-bucket"
		s3BucketResourceName = resourceName + "_bucket"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	s3BucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, s3BucketResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsDynamodb(resourceName, gatewayId, assetId, "", "default"),
		testAccDSFLogAggregatorConfig_AwsS3(s3BucketResourceName, gatewayId, s3BucketAssetId, resourceTypeAndName+".asset_id", "true", "DYNAMODB", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{s3BucketResourceTypeAndName: {}},
		map[string]map[string]string{resourceTypeAndName: {}},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, s3BucketResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsNeptuneClusterSlowQuery(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_neptune_cluster"
		assetId      = testAwsRdsClusterPrefix + "my-neptune-cluster"

		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/my-neptune-cluster/audit:*"
		logGroupResourceName = resourceName + "_log_group"

		slowLogGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/my-neptune-cluster/slowquery:*"
		slowLogGroupResourceName = resourceName + "_slow_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)
	slowLogGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowLogGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsNeptuneCluster(resourceName, gatewayId, assetId, ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName+".asset_id", true, "AWS_NEPTUNE_SLOW", logGroupResourceTypeAndName),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			logGroupResourceTypeAndName: {
				"gateway_service": "gateway-aws@neptune.service",
			},
			slowLogGroupResourceTypeAndName: {
				"gateway_service": "gateway-aws@neptune-slow-query.service",
			},
		},
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName, slowLogGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsAuroraPostgresqlClusterCloudWatch(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_aurora_postgresql_cloudwatch_cluster"
		assetId      = testAwsRdsClusterPrefix + "my-aurora-postgresql-cluster"

		instanceResourceName = resourceName + "_instance"
		instanceAssetId      = assetId + "-writer"

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/my-cluster/postgresql:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName, gatewayId, assetId, "LOG_GROUP", resourceName),
		testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			logGroupResourceTypeAndName: {
				"gateway_service": "gateway-aws@aurora-postgresql.service",
			},
		},
		map[string]map[string]string{
			resourceTypeAndName:         {},
			instanceResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, instanceResourceTypeAndName, logGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsAuroraPostgresqlClusterKinesis(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_aurora_postgresql_kinesis_cluster"
		assetId      = testAwsRdsClusterPrefix + "my-aurora-postgresql-kinesis-cluster"

		instanceResourceName = resourceName + "_instance"
		instanceAssetId      = assetId + "-writer"

		kinesisResourceName = resourceName + "_kinesis_stream"
		kinesisAssetId      = testAwsKinesisPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	kinesisResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, kinesisResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
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
			createValidateImportStep(resourceTypeAndName),
		},
	})
}

func TestAccDSFDataSource_AwsRdsAuroraMysqlClusterCloudWatchSlowQuery(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		clusterId    = "my-aurora-mysql-cloudwatch-cluster"
		assetId      = testAwsRdsClusterPrefix + clusterId
		resourceName = "aws_aurora_mysql_cloudwatch_cluster"

		instanceResourceName = resourceName + "_instance"
		instanceAssetId      = assetId + "-writer"

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/" + clusterId + "/audit:*"

		slowLogGroupResourceName = resourceName + "_slow_log_group"
		slowLogGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/cluster/" + clusterId + "/slowquery:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	// TODO: check that instance asset is connected once fixed: https://onejira.imperva.com/browse/SR-2046
	// instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)
	slowLogGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowLogGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName, gatewayId, assetId, "", resourceName),
		testAccDSFDataSourceConfig_AwsRdsAuroraMysql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName+".asset_id", true, "AWS_RDS_AURORA_MYSQL_SLOW", logGroupResourceTypeAndName),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			logGroupResourceTypeAndName: {
				"gateway_service": "gateway-aws@aurora-mysql.service",
			},
			slowLogGroupResourceTypeAndName: {
				"gateway_service": "gateway-aws@aurora-mysql-slow-query.service",
			},
		},
		map[string]map[string]string{
			resourceTypeAndName: {},
			// instanceResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName, slowLogGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsAuroraMysqlClusterKinesis(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_aurora_mysql_kinesis_cluster"
		assetId      = testAwsRdsClusterPrefix + "my-aurora-mysql-kinesis-cluster"

		instanceResourceName = resourceName + "_instance"
		instanceAssetId      = assetId + "-writer"

		kinesisResourceName = resourceName + "_kinesis_stream"
		kinesisAssetId      = testAwsKinesisPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	// instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	kinesisResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, kinesisResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName, gatewayId, assetId, "", resourceName),
		testAccDSFDataSourceConfig_AwsRdsAuroraMysql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
		testAccDSFLogAggregatorConfig_AwsKinesis(kinesisResourceName, gatewayId, kinesisAssetId, resourceTypeAndName+".asset_id", true, "KINESIS"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{kinesisResourceTypeAndName: {"gateway_service": "gateway-kinesis@aurora-mysql.service"}},
		map[string]map[string]string{
			resourceTypeAndName: {},
			// instanceResourceTypeAndName: {},
		})

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, kinesisResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsMariadb(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_rds_mariadb"
		assetId      = testAwsRdsDbPrefix + "my-rds-mariadb"

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/my-rds-mariadb/audit:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsMariadb(resourceName, gatewayId, assetId),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{logGroupResourceTypeAndName: {"gateway_service": "gateway-aws@mariadb.service"}},
		map[string]map[string]string{resourceTypeAndName: {}},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsMsSqlServer(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_rds_ms_sql_server"
		assetId      = testAwsRdsDbPrefix + resourceName

		s3BucketResourceName = resourceName + "_bucket"
		s3BucketAssetId      = testAwsS3BucketPrefix + resourceName

		cloudAccountResourceName = resourceName + "_role"
		cloudAccountAssetId      = testAwsAccountArnPrefix + cloudAccountResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	s3BucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, s3BucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Aws(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		testAccDSFDataSourceConfig_AwsRdsMsSqlServer(resourceName, gatewayId, assetId, "true", s3BucketResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AwsS3(s3BucketResourceName, gatewayId, s3BucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "", "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{resourceTypeAndName: {}},
		map[string]map[string]string{
			// s3BucketResourceTypeAndName: {"provider": "aws-rds-mssql"},
			s3BucketResourceTypeAndName: {},
		})

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, s3BucketResourceTypeAndName, cloudAccountResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsMysql(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_rds_mysql"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/my-rds-mysql/audit:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsMysql(resourceName, gatewayId, assetId, "LOG_GROUP"),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{logGroupResourceTypeAndName: {"gateway_service": "gateway-aws@mysql.service"}},
		map[string]map[string]string{
			resourceTypeAndName: {},
		})

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsMysqlSlowQuery(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_rds_mysql_slow_query"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/aws_rds_mysql_slow_query/audit:*"

		slowLogGroupResourceName = resourceName + "_slow_log_group"
		slowLogGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/aws_rds_mysql_slow_query/slowquery:*"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)
	slowLogGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowLogGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsMysql(resourceName, gatewayId, assetId, ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName+".asset_id", true, "AWS_RDS_MYSQL_SLOW", logGroupResourceTypeAndName),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			logGroupResourceTypeAndName:     {"gateway_service": "gateway-aws@mysql.service"},
			slowLogGroupResourceTypeAndName: {"gateway_service": "gateway-aws@mysql-slow-query.service"},
		},
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName, slowLogGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsOracleCloudwatch(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_rds_oracle_cloudwatch"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/" + resourceName + "/audit:*"
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "LOG_GROUP", "", ""),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			logGroupResourceTypeAndName: {"gateway_service": "gateway-aws@oracle-rds.service"},
		},
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsOracleS3(t *testing.T) {
	for _, v := range []string{"15.0", "15.1", "master"} {
		skipTestForKnownIssue(t, v, "https://onejira.imperva.com/browse/SR-2057", true)
	}

	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_rds_oracle_s3"
		assetId      = testAwsRdsDbPrefix + resourceName

		s3BucketResourceName = resourceName + "_bucket"
		s3BucketAssetId      = testAwsS3BucketPrefix + s3BucketResourceName

		cloudAccountResourceName = resourceName + "_bucket_role"
		cloudAccountAssetId      = testAwsAccountArnPrefix + cloudAccountResourceName
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	s3BucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, s3BucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Aws(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		// testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "", "true", s3BucketResourceTypeAndName+".asset_id", s3BucketResourceTypeAndName),
		testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "", "true", s3BucketResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AwsS3(s3BucketResourceName, gatewayId, s3BucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "true", "ORACLE", ""),
	)

	// TODO: define checks
	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {"gateway_service": ""},
		},
		map[string]map[string]string{
			s3BucketResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, s3BucketResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRdsOracleUnified(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_rds_oracle_unified"
		assetId      = testAwsRdsDbPrefix + resourceName
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps: []resource.TestStep{
			// onboard and connect to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "UNIFIED", "true", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-odbc@oracle_unified.service"),
				),
			},
			// update audit_type -> reconnect asset to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "UNIFIED_AGGREGATED", "true", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-odbc@oracle_unified_aggregated.service"),
				),
			},
			// disconnect asset
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, assetId, "UNIFIED_AGGREGATED", "false", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", ""),
				),
			},
			createValidateImportStep(resourceTypeAndName),
		},
	})
}

func TestAccDSFDataSource_AwsRdsPostgresql(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_rds_postgresql"
		assetId      = testAwsRdsDbPrefix + resourceName

		logGroupResourceName = resourceName + "_log_group"
		logGroupAssetId      = testAwsLogGroupPrefix + "/aws/rds/instance/aws_rds_postgresql/postgres:*"
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRdsPostgresql(resourceName, gatewayId, assetId, "LOG_GROUP"),
		testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			logGroupResourceTypeAndName: {"gateway_service": "gateway-aws@postgresql.service"},
		},
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, logGroupResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRedshiftS3(t *testing.T) {
	for _, v := range []string{"4.17", "4.18", "4.19", "15.0"} {
		skipTestForKnownIssue(t, v, "https://onejira.imperva.com/browse/USC-2396", true)
	}

	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_redshift_s3"
		assetId      = testAwsRedshiftPrefix + resourceName

		s3BucketResourceName = resourceName + "_bucket"
		s3BucketAssetId      = testAwsS3BucketPrefix + s3BucketResourceName

		cloudAccountResourceName = resourceName + "_bucket_role"
		cloudAccountAssetId      = testAwsAccountArnPrefix + cloudAccountResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	s3BucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, s3BucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Aws(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		testAccDSFDataSourceConfig_AwsRedshift(resourceName, gatewayId, assetId, "", "true", s3BucketResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AwsS3(s3BucketResourceName, gatewayId, s3BucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "", "REDSHIFT", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {"gateway_service": ""},
		},
		map[string]map[string]string{
			s3BucketResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, s3BucketResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AwsRedshiftTable(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "aws_redshift_table"
		assetId      = testAwsRedshiftPrefix + resourceName
	)
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AwsRedshift(resourceName, gatewayId, assetId, "TABLE", "true", ""),
	)

	initialChecks, _ := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {"gateway_service": "gateway-odbc@redshift.service"},
		}, nil,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, nil, true, []string{resourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AzureCosmosDBMongo(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "azure_cosmosdb_mongo"
		assetId      = testAzurePrefix + testCosmosPrefix + "my-cosmos-mongodb"

		eventhubResourceName = resourceName + "_eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureCosmosDBMongo(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Cosmos_Mongo"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			eventhubResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, eventhubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AzureCosmosDBNosql(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "azure_cosmosdb_nosql"
		assetId      = testAzurePrefix + testCosmosPrefix + "my-cosmos-nosql"

		eventhubResourceName = resourceName + "_eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureCosmosDBNosql(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Cosmos_SQL"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			eventhubResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, eventhubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AzureCosmosDBTable(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "azure_cosmosdb_table"
		assetId      = testAzurePrefix + testCosmosPrefix + "my-cosmos-table"

		eventhubResourceName = resourceName + "_eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureCosmosDBTable(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Cosmos_Table"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			eventhubResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, eventhubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AzureMsSqlServer(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "azure_ms_sql_server"
		assetId      = testAzurePrefix + "Microsoft.Sql/servers/my-sql-server"

		eventhubResourceName = resourceName + "_eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureMsSqlServer(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Sql"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			eventhubResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, eventhubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AzureMysqlFlexible(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "azure_mysql_flexible"
		assetId      = testAzurePrefix + "Microsoft.DBforMySQL/flexibleservers/my-mysql-flexible"

		eventhubResourceName = resourceName + "_eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureMysqlFlexible(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Mysql_Flexible"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			eventhubResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, eventhubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AzurePostgresqlFlexible(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "azure_postgresql_flexible"
		assetId      = testAzurePrefix + "Microsoft.DBforPostgreSQL/flexibleservers/my-postgresql-flexible"

		eventhubResourceName = resourceName + "_eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzurePostgresqlFlexible(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "Postgresql_Flexible"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			eventhubResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, eventhubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_AzureSqlManagedInstance(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "azure_sql_managed_instance"
		assetId      = testAzurePrefix + "Microsoft.Sql/managedinstances/my-managed-instance"

		eventhubResourceName = resourceName + "_eventhub"
		eventhubAssetId      = testEventhubPrefix + eventhubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	eventhubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, eventhubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_AzureSqlManagedInstance(resourceName, gatewayId, assetId, "true", eventhubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_AzureEventhub(eventhubResourceName, gatewayId, eventhubAssetId, "default", "", "", "", "AzureSQL_Managed"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			eventhubResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, eventhubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_GcpAlloydb(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		prefix               = "gcp_alloydb"
		clusterResourceName  = prefix + "_cluster"
		instanceResourceName = prefix + "_instance"
		clusterAssetId       = testGcpSpannerPrefix + "-primary-instance"
		instanceAssetId      = testGcpSpannerPrefix + "-reader"

		pubsubResourceName = prefix + "_pubsub"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	clusterResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, clusterResourceName)
	instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpAlloydbCluster(clusterResourceName, gatewayId, clusterAssetId, "true", pubsubResourceTypeAndName+".asset_id", "cluster"),
		testAccDSFDataSourceConfig_GcpAlloydb(instanceResourceName, gatewayId, instanceAssetId, "true", pubsubResourceTypeAndName+".asset_id", "instance"),
		testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "ALLOYDB_POSTGRESQL", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			clusterResourceTypeAndName:  {},
			instanceResourceTypeAndName: {},
		},
		map[string]map[string]string{
			pubsubResourceTypeAndName: {"gateway_service": "gateway-gcp@alloydb-postgresql.service"},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{clusterResourceTypeAndName, instanceResourceTypeAndName, pubsubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_GcpBigQuery(t *testing.T) {
	gatewayId := getGatewayId(t)

	getDSFDataSourceFields("GCP BIGQUERY")

	const (
		resourceName = "gcp_bigquery"
		assetId      = "projects/my-project/bigquery"

		pubsubResourceName = resourceName + "_pubsub"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpBigQuery(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "BIGQUERY", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			pubsubResourceTypeAndName: {"gateway_service": "gateway-gcp@bigquery.service"},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, pubsubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_GcpMsSqlServer(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "gcp_ms_sql_server"
		assetId      = testGcpMsSqlServerPrefix + "-1"

		bucketResourceName = resourceName + "_bucket"
		bucketAssetId      = "my-project:" + bucketResourceName

		cloudAccountResourceName = resourceName + "_cloud_account"
		cloudAccountAssetId      = testGcpServiceAccount
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	bucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, bucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Gcp(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		testAccDSFDataSourceConfig_GcpMsSqlServer(resourceName, gatewayId, assetId, "true", bucketResourceTypeAndName+".asset_id", ""),
		testAccDSFLogAggregatorConfig_GcpCloudStorageBucket(bucketResourceName, gatewayId, bucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			bucketResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, bucketResourceTypeAndName, cloudAccountResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_GcpMsSqlServerManyToOne(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		prefix           = "gcp_ms_sql_server_many_to_one"
		resourceNameProd = prefix + "_prod"
		assetIdProd      = testGcpMsSqlServerPrefix + "-2-prod"
		resourceNameDev  = prefix + "_dev"
		assetIdDev       = testGcpMsSqlServerPrefix + "-2-dev"

		bucketResourceName = prefix + "_bucket"
		bucketAssetId      = "my-project:" + bucketResourceName

		cloudAccountResourceName = prefix + "_cloud_account"
		cloudAccountAssetId      = "my_service_account_2@my-project.iam.gserviceaccount.com:my-project"
	)

	resourceTypeAndNameProd := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceNameProd)
	resourceTypeAndNameDev := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceNameDev)
	bucketResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, bucketResourceName)
	cloudAccountResourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, cloudAccountResourceName)

	config := ConfigCompose(
		testAccDSFCloudAccountConfig_Gcp(cloudAccountResourceName, gatewayId, cloudAccountAssetId, "default"),
		testAccDSFDataSourceConfig_GcpMsSqlServer(resourceNameProd, gatewayId, assetIdProd, "true", bucketResourceTypeAndName+".asset_id", ""),
		testAccDSFDataSourceConfig_GcpMsSqlServer(resourceNameDev, gatewayId, assetIdDev, "true", bucketResourceTypeAndName+".asset_id", ""),
		testAccDSFLogAggregatorConfig_GcpCloudStorageBucket(bucketResourceName, gatewayId, bucketAssetId, cloudAccountResourceTypeAndName+".asset_id", "", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndNameProd: {},
			resourceTypeAndNameDev:  {},
		},
		map[string]map[string]string{
			bucketResourceTypeAndName: {},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, false, []string{resourceTypeAndNameProd, resourceTypeAndNameDev, bucketResourceTypeAndName, cloudAccountResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_GcpMysql(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "gcp_mysql"
		assetId      = testGcpPrefix + "mysql-instance-1"

		pubsubResourceName = resourceName + "_pubsub"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpMysql(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "MYSQL", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			pubsubResourceTypeAndName: {"gateway_service": "gateway-gcp@mysql.service"},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, pubsubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_GcpMysqlSlowQuery(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "gcp_mysql_slow_query"
		assetId      = testGcpPrefix + "mysql-instance-2"

		auditPubsubResourceName = resourceName + "_audit_pubsub"
		auditPubsubAssetId      = testPubsubSubscriptionPrefix + auditPubsubResourceName

		slowQueryPubsubResourceName = resourceName + "_slow_query_pubsub"
		slowQueryPubsubAssetId      = testPubsubSubscriptionPrefix + slowQueryPubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	auditPubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, auditPubsubResourceName)
	slowQueryPubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowQueryPubsubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpMysql(resourceName, gatewayId, assetId, "true", auditPubsubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_GcpPubsub(auditPubsubResourceName, gatewayId, auditPubsubAssetId, "default", "", "", "MYSQL", ""),
		testAccDSFLogAggregatorConfig_GcpPubsub(slowQueryPubsubResourceName, gatewayId, slowQueryPubsubAssetId, "default", "", "true", "GCP_MYSQL_SLOW", "GCP MYSQL"),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			auditPubsubResourceTypeAndName:     {"gateway_service": "gateway-gcp@mysql.service"},
			slowQueryPubsubResourceTypeAndName: {"gateway_service": "gateway-gcp@mysql-slow-query.service"},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, auditPubsubResourceTypeAndName, slowQueryPubsubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_GcpPostgresql(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "gcp_postgresql"
		assetId      = testGcpPrefix + "postgresql-instance-1"

		pubsubResourceName = resourceName + "_pubsub"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpPostgresql(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id"),
		testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "POSTGRESQL", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			pubsubResourceTypeAndName: {"gateway_service": "gateway-gcp@postgresql.service"},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, pubsubResourceTypeAndName}),
	})
}

func TestAccDSFDataSource_GcpSpanner(t *testing.T) {
	gatewayId := getGatewayId(t)

	const (
		resourceName = "gcp_spanner"
		assetId      = "projects/my-project/spanner"

		pubsubResourceName = resourceName + "_pubsub"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	config := ConfigCompose(
		testAccDSFDataSourceConfig_GcpSpanner(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id", ""),
		testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "SPANNER", ""),
	)

	initialChecks, refreshChecks := buildConnectDisconnectGatewayChecks(
		map[string]map[string]string{
			resourceTypeAndName: {},
		},
		map[string]map[string]string{
			pubsubResourceTypeAndName: {"gateway_service": "gateway-gcp@spanner.service"},
		},
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyDSFResources,
		Steps:        createConnectDisconnectGatewayTestSteps(config, initialChecks, refreshChecks, true, []string{resourceTypeAndName, pubsubResourceTypeAndName}),
	})
}

// *****************************************
// **************** Helpers ****************
// *****************************************

// Returns the ID of the first DSF data source found in the state
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

// Confirm assets are destroyed after an acceptance test run on the warehouse
func testAccCheckDestroyDSFResources(state *terraform.State) error {
	log.Printf("[INFO] Running testAccCheckDestroyDSFResources")

	// allow "disableAsset" playbook enough time to run
	time.Sleep(15 + time.Second)

	client := testAccProvider.Meta().(*Client)

	type clientReadFunc func(string) (*ResourceWrapper, error)
	checks := map[string]struct {
		checkFunc clientReadFunc
		notFound  string
	}{
		dsfDataSourceResourceType:    {client.ReadDSFDataSource, "DSFDataSource not found"},
		dsfLogAggregatorResourceType: {client.ReadLogAggregator, "LogAggregator not found"},
		dsfCloudAccountResourceType:  {client.ReadCloudAccount, "CloudAccount not found"},
		dsfSecretManagerResourceType: {client.ReadSecretManager, "SecretManager not found"},
	}

	for _, rs := range state.RootModule().Resources {
		assetId := rs.Primary.ID
		check, ok := checks[rs.Type]
		if !ok {
			return fmt.Errorf("Error creating check function for %s with asset_id: %s", rs.Type, assetId)
		}
		_, err := check.checkFunc(assetId)
		if err == nil {
			return fmt.Errorf("Resource type %s with asset_id %s still exists in the DSFHub.", rs.Type, assetId)
		}
		if strings.Contains(fmt.Sprintf("%v", err), check.notFound) {
			continue
		}
	}
	return nil
}
