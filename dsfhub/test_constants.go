package dsfhub

const (
	// Client acceptance tests constants
	testAdminEmail           = "test@email.com"
	testArn                  = "arn:aws:rds:us-east-2:123456789:db:your-db"
	testAssetDisplayName     = "arn:aws:rds:us-east-2:123456789:db:your-db"
	testDSServerType         = "AWS RDS MYSQL"
	testServerHostName       = "your-db-name.abcde12345.us-east-2.rds.amazonaws.com"
	testOnPremServerHostName = "server.company.com"
	testServerIP             = "1.2.3.4"
	testServerPort           = "8200"

	testInvalidDSFHUBHost = "https://invalid.host.com"
	testGatewayId         = "e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd"

	// Secret Managers tests constants
	testSMConnectionReason = "default"
	testSMRoleName         = "vault-role-for-ec2"
	testSMAuthMechanism    = "ec2"
	testSMAssetId          = "your-host-name-here"
	testSMServerType       = "HASHICORP"

	// Resource acceptance tests constants
	// AWS
	testAwsAccountId        = "123456789012"
	testAwsRegion           = "us-east-1"
	testAwsAccountArnPrefix = "arn:aws:iam::" + testAwsAccountId + ":role/"
	testAwsDynamoDbPrefix   = "arn:aws:dynamodb:us-east-1:" + testAwsAccountId + ":table/"
	testAwsKinesisPrefix    = "arn:aws:kinesis:us-east-1:" + testAwsAccountId + ":stream/"
	testAwsLogGroupPrefix   = "arn:aws:logs:us-east-1:" + testAwsAccountId + ":log-group:"
	testAwsRdsPrefix        = "arn:aws:rds:us-east-2:123456789012"
	testAwsRdsClusterPrefix = testAwsRdsPrefix + ":cluster:"
	testAwsRdsDbPrefix      = testAwsRdsPrefix + ":db:"
	testAwsRedshiftPrefix   = "arn:aws:redshift:us-east-1:" + testAwsAccountId + ":cluster:"
	testAwsS3BucketPrefix   = "arn:aws:s3:::"

	// Azure
	testAzurePrefix    = "/subscriptions/my-subscription-id/resourceGroups/my-resource-group/providers/"
	testCosmosPrefix   = testAzurePrefix + "Microsoft.DocumentDB/databaseAccounts/"
	testEventhubPrefix = testAzurePrefix + "Microsoft.EventHub/namespaces/my-namespace/eventhubs/"

	// GCP
	testGcpPrefix                = "my-project:us-west-1:"
	testGcpMsSqlServerPrefix     = "projects:us-west-1:sql-server-instance"
	testGcpSpannerPrefix         = "projects/my-project/locations/us-west2/clusters/my-cluster/instances/"
	testGcpServiceAccount        = "my_service_account_1@my-project.iam.gserviceaccount.com:my-project"
	testPubsubSubscriptionPrefix = "projects/my-project/subscriptions/"
)
