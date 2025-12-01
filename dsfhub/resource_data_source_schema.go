package dsfhub

var ignoreDataSourceParamsByServerType = map[string]map[string]bool{
	"AWS ATHENA":                        {"arn": true, "asset_display_name": true},
	"AWS DOCUMENTDB CLUSTER":            {"arn": true, "asset_display_name": true},
	"AWS DOCUMENTDB":                    {"arn": true, "asset_display_name": true},
	"AWS DYNAMODB":                      {"arn": true, "asset_display_name": true},
	"AWS GLUE":                          {"arn": true, "asset_display_name": true},
	"AWS LAKE FORMATION":                {"arn": true, "asset_display_name": true},
	"AWS NEPTUNE CLUSTER":               {"arn": true, "asset_display_name": true},
	"AWS NEPTUNE":                       {"arn": true, "asset_display_name": true},
	"AWS OPENSEARCH":                    {"arn": true, "asset_display_name": true},
	"AWS RDS AURORA MYSQL CLUSTER":      {"arn": true, "asset_display_name": true},
	"AWS RDS AURORA MYSQL":              {"arn": true, "asset_display_name": true},
	"AWS RDS AURORA POSTGRESQL CLUSTER": {"arn": true, "asset_display_name": true},
	"AWS RDS AURORA POSTGRESQL":         {"arn": true, "asset_display_name": true},
	"AWS RDS DB2":                       {"arn": true, "asset_display_name": true},
	"AWS RDS MARIADB":                   {"arn": true, "asset_display_name": true},
	"AWS RDS MS SQL SERVER":             {"arn": true, "asset_display_name": true},
	"AWS RDS MYSQL":                     {"arn": true, "asset_display_name": true},
	"AWS RDS ORACLE":                    {"arn": true, "asset_display_name": true},
	"AWS RDS POSTGRESQL":                {"arn": true, "asset_display_name": true},
	"AWS RDS POSTGRESQL CLUSTER":        {"arn": true, "asset_display_name": true},
	"AWS REDSHIFT":                      {"arn": true, "asset_display_name": true},
	"AWS REDSHIFT SERVERLESS":           {"arn": true, "asset_display_name": true},
	"AWS S3":                            {"arn": true, "asset_display_name": true},
	"AZURE COSMOSDB MONGO":              {"arn": true, "asset_display_name": true},
	"AZURE COSMOSDB TABLE":              {"arn": true, "asset_display_name": true},
	"AZURE COSMOSDB":                    {"arn": true, "asset_display_name": true},
	"AZURE DATABRICKS WORKSPACE":        {"arn": true, "asset_display_name": true},
	"AZURE DATA EXPLORER":               {"arn": true, "asset_display_name": true},
	"AZURE MARIADB":                     {"arn": true, "asset_display_name": true},
	"AZURE MS SQL SERVER":               {"arn": true, "asset_display_name": true},
	"AZURE MYSQL":                       {"arn": true, "asset_display_name": true},
	"AZURE MYSQL FLEXIBLE":              {"arn": true, "asset_display_name": true},
	"AZURE POSTGRESQL":                  {"arn": true, "asset_display_name": true},
	"AZURE POSTGRESQL FLEXIBLE":         {"arn": true, "asset_display_name": true},
	"AZURE SQL MANAGED INSTANCE":        {"arn": true, "asset_display_name": true},
	"AZURE STORAGE ACCOUNT":             {"arn": true, "asset_display_name": true},
}

var requiredDataSourceFieldsJson = `{
    "ServerTypes": {
        "AEROSPIKE": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "ALIBABA APSARA MONGODB": {
            "auth_mechanisms": {
                "default": [
                    "reason",
                    "server_port"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "ALIBABA APSARA RDS MYSQL": {
            "auth_mechanisms": {
                "default": [
                    "reason",
                    "server_port"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "ALIBABA APSARA RDS POSTGRESQL": {
            "auth_mechanisms": {
                "default": [
                    "reason",
                    "server_port"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "ALIBABA MAX COMPUTE": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "ALIBABA OSS": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "AMBARI": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "AWS ATHENA": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "admin_email"
            ]
        },
        "AWS DOCUMENTDB": {
            "auth_mechanisms": {
                "key_file": [
                    "reason",
                    "username",
                    "key_file"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "arn",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "AWS DOCUMENTDB CLUSTER": {
            "auth_mechanisms": {
                "key_file": [
                    "reason",
                    "username",
                    "key_file"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "arn",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "AWS DYNAMODB": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ],
                "iam_role": [
                    "reason"
                ],
                "key": [
                    "reason",
                    "access_id",
                    "secret_key"
                ],
                "profile": [
                    "reason",
                    "username"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "arn",
                "admin_email",
                "region"
            ]
        },
        "AWS GLUE": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "AWS LAKE FORMATION": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "AWS NEPTUNE": {
            "auth_mechanisms": {
                "ec2": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "admin_email"
            ]
        },
        "AWS NEPTUNE CLUSTER": {
            "auth_mechanisms": {
                "ec2": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "admin_email"
            ]
        },
        "AWS OPENSEARCH": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS AURORA MYSQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS AURORA MYSQL CLUSTER": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS AURORA POSTGRESQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS AURORA POSTGRESQL CLUSTER": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS DB2": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "resource_id",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS MARIADB": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS MS SQL SERVER": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS MYSQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS ORACLE": {
            "auth_mechanisms": {
                "oracle_wallet": [
                    "reason",
                    "username",
                    "password",
                    "dsn",
                    "wallet_dir"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS POSTGRESQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS RDS POSTGRESQL CLUSTER": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AWS REDSHIFT": {
            "auth_mechanisms": {
                "aws_credentials": [
                    "reason",
                    "username",
                    "database_name",
                    "access_id",
                    "aws_connection_id"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "AWS REDSHIFT SERVERLESS": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "AWS S3": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "arn",
                "server_host_name",
                "admin_email"
            ]
        },
        "AZURE COSMOSDB": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_port",
                "admin_email"
            ]
        },
        "AZURE COSMOSDB MONGO": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_port",
                "admin_email"
            ]
        },
        "AZURE COSMOSDB TABLE": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_port",
                "admin_email"
            ]
        },
        "AZURE DATABRICKS WORKSPACE": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_port",
                "admin_email"
            ]
        },
        "AZURE DATA EXPLORER": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_port",
                "admin_email"
            ]
        },
        "AZURE MARIADB": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "admin_email",
                "location"
            ]
        },
        "AZURE MS SQL SERVER": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email",
                "location"
            ]
        },
        "AZURE MYSQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "admin_email",
                "location"
            ]
        },
        "AZURE MYSQL FLEXIBLE": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "admin_email",
                "location"
            ]
        },
        "AZURE POSTGRESQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "admin_email",
                "location"
            ]
        },
        "AZURE POSTGRESQL FLEXIBLE": {
            "auth_mechanisms": {
                "password": [
                "reason",
                "password",
                "username"
                ]
            },
            "required": [
                "server_host_name",
                "admin_email",
                "asset_display_name",
                "asset_id",
                "location",
                "gateway_id"
            ]
        },
        "AZURE SQL MANAGED INSTANCE": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email",
                "location"
            ]
        },
        "AZURE STORAGE ACCOUNT": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_port",
                "admin_email"
            ]
        },
        "CASSANDRA": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "CLICKHOUSE": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "database_name",
                "admin_email"
            ]
        },
        "CLOUDANT": {
            "auth_mechanisms": {
                "iam_role": [
                    "reason",
                    "region",
                    "account_name",
                    "api_key",
                    "crn",
                    "service_key"
                ],
                "password": [
                    "reason",
                    "username",
                    "password",
                    "region",
                    "crn",
                    "service_key"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "CLOUDANT LOCAL": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "CLOUDERA": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "COCKROACHDB": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "COUCHBASE": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "DATASTAX": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "DB2": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ],
                "ssl": [
                    "reason",
                    "database_name",
                    "ssl_server_cert"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "DRUID CLUSTER": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "DRUID": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email",
                "cluster_id",
                "cluster_name"
            ]
        },
        "EDB POSTGRESQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "ELASTICSEARCH": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "ELOQUENCE": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "EMR": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GAUSSDB": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GCP ALLOYDB POSTGRESQL": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ],
                "service_account": [
                    "reason",
                    "key_file"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GCP ALLOYDB POSTGRESQL CLUSTER": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ],
                "service_account": [
                    "reason",
                    "key_file"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GCP BIGQUERY": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ],
                "service_account": [
                    "reason",
                    "key_file"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GCP BIGTABLE": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ],
                "service_account": [
                    "reason",
                    "key_file"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GCP FIRESTORE": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GCP MS SQL SERVER": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "GCP MYSQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "GCP POSTGRESQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "GCP SPANNER": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GEMFIRE": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GRAINITE": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "GRIDGAIN IGNITE": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "HBASE": {
            "auth_mechanisms": {
                "kerberos": [
                    "reason"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "HDFS": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "HIVE": {
            "auth_mechanisms": {
                "kerberos": [
                    "reason",
                    "database_name"
                ],
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "IMPALA": {
            "auth_mechanisms": {
                "key_file": [
                    "reason",
                    "username",
                    "key_file"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "INFORMIX": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "IRIS": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "MAPR FS": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "MAPR HBASE": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "MARIADB": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "MARKLOGIC": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "MONGODB": {
            "auth_mechanisms": {
                "key_file": [
                    "reason",
                    "username",
                    "key_file"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "MONGODB ATLAS": {
            "auth_mechanisms": {
                "default": [
                    "reason",
                    "access_id",
                    "secret_key"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email",
                "database_name"
            ]
        },
        "MS SQL SERVER": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "MYSQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "NEO4J": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "NETEZZA": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "ORACLE": {
            "auth_mechanisms": {
                "kerberos": [
                    "reason"
                ],
                "oracle_wallet": [
                    "reason",
                    "username",
                    "password",
                    "dsn",
                    "wallet_dir"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "service_name",
                "admin_email"
            ]
        },
        "PERCONA MONGODB": {
            "auth_mechanisms": {
                "key_file": [
                    "reason",
                    "username",
                    "key_file"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email",
                "database_name"
            ]
        },
        "PERCONA MYSQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "POSTGRESQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "PROGRESS OPENEDGE": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "REDIS": {
            "auth_mechanisms": {
                "default": [
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "SAP HANA": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "SAP IQ": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "SCYLLADB": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "SINGLESTORE": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "SNOWFLAKE": {
            "auth_mechanisms": {
                "oauth": [
                    "reason",
                    "username",
                    "access_id",
                    "token"
                ],
                "oauth-azure-ad": [
                    "reason",
                    "password",
                    "client_secret",
                    "access_id",
                    "client_id",
                    "principal",
                    "resource_id",
                    "snowflake_role",
                    "tenant_id"
                ],
                "oauth2": [
                    "reason",
                    "access_id",
                    "oauth_parameters",
                    "principal",
                    "snowflake_role",
                    "token_endpoint"
                ],
                "key_file": [
                    "reason",
                    "access_id",
                    "username",
                    "key_file"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "SPLUNK": {
            "auth_mechanisms": {
                "key": [
                    "reason",
                    "access_id",
                    "secret_key"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "SYBASE": {
            "auth_mechanisms": {
                "kerberos": [
                    "reason"
                ],
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "TERADATA": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "TIGERGRAPH": {
            "auth_mechanisms": {},
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "VERTICA": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password",
                    "database_name"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "admin_email"
            ]
        },
        "YUGABYTE CQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        },
        "YUGABYTE SQL": {
            "auth_mechanisms": {
                "password": [
                    "reason",
                    "username",
                    "password"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "server_ip",
                "server_port",
                "admin_email"
            ]
        }
    }
}`
