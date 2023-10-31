---
layout: "dsfhub"
page_title: "Provider: DSFHUB"
sidebar_current: "docs-dsfhub-index"
description: |-
  The DSFHUB provider is used to interact with Data Security Fabric Hub resources supported by Imperva. The provider needs to be configured with the endpoint and authentication token before it can be used.
---

# DSFHUB Provider

The DSFHUB provider is used to interact with Data Security Fabric Hub resources supported by Imperva. The provider needs to be configured with the endpoint and authentication token before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage - Resources

```hcl
# Configure the DSFHUB provider
provider "dsfhub" {
  dsfhub_host = "${var.dsfhub_host}"
  dsfhub_token = "${var.dsfhub_token}"
}

# Create a AWS dsfhub_cloud_account
resource "dsfhub_cloud_account" "aws_cloud_account_key" {
	server_type = "AWS"
	admin_email         = "your@email.com"	
	asset_display_name  = "my-aws-cloud-account"	
	asset_id            = "arn:aws:iam::12345678"
	gateway_id          = "12345-abcde-12345-abcde"
	asset_connection {
		auth_mechanism = "key"
		access_id = var.aws_key_access_id
		reason = "default" 
		region = "us-east-2" 
		secret_key = var.aws_key_secret_key
	}
}

# Create an AWS RDS MYSQL dsfhub_data_source with password authentication
resource "dsfhub_data_source" "example-aws-rds-mysql-password" {
	admin_email = var.admin_email
	asset_display_name  = aws_db_instance.stats_demo_mysql_db.identifier
	asset_id            = aws_db_instance.stats_demo_mysql_db.arn
	gateway_id          = local.gw_ec2_01.jsonar_uid
	server_host_name    = "my-rds-mysql-endpoint-here"
	region              = "us-east-2"
	server_port         = "3306"
	version             = 8
	parent_asset_id     = dsf_cloud_account.aws_cloud_account_key.asset_id
	asset_connection {
		auth_mechanism  = "password"
		password        = var.password
		reason          = "default" 
		username        = var.username
	}
}

# Create an AWS Log Group dsfhub_log_aggregator
resource "dsfhub_log_aggregator" "rds-mysql-stats-demo-log-group" {
	server_type = "AWS LOG GROUP"
	admin_email         = var.admin_email	
	asset_display_name  = var.log_group_name
	asset_id            = var.log_group_arn
	gateway_id          = var.gateway_jsonar_uid
	parent_asset_id     = dsf_data_source.example-aws-rds-mysql-password.asset_id	
	asset_connection {
		auth_mechanism = "default"
		reason = "default" 
		region = "us-east-2" 
	}
}
```

## Example Usage - Resources and Data Sources

```hcl
# Configure the DSFHUB provider
provider "dsfhub" {
  dsfhub_host = "${var.dsfhub_host}"
  dsfhub_token = "${var.dsfhub_token}"
}

# Reference an AWS dsfhub_cloud_account (Data Source)
data "dsfhub_cloud_account" "example_aws_cloud_account" {
  asset_id = "arn:partition:service:region:account-id" # The value of the arn for aws resources
}

# ### Resource example for ORACLE with password auth_mechanism ###
resource "dsfhub_data_source" "example_oracle_password_and_wallet" {
	server_type = "ORACLE"
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
	server_ip = var.server_ip	# IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
	service_name = var.service_name	# Service Name or SID used to connect
	parent_asset_id = dsf_cloud_account.example_aws_cloud_account.asset_id
	asset_connection {
		auth_mechanism = "password"
		password = null # password description: "The password of the user being used to authenticate"
		reason = null # Example Values: "default", "sonargateway", "SDM", "audit management", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: "The user to use when connecting. For Oracle assets, the username should be in exact case as defined in table \"dba_users\", otherwise Oracle normally converts everything to all-caps."
	}
	asset_connection {
		auth_mechanism = "oracle_wallet"
		dsn = null # dsn description: "Data Source Name"
		password = null # password description: "The password of the user being used to authenticate"
		reason = null # Example Values: "default", "sonargateway", "SDM", "audit management", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		username = null # username description: "The user to use when connecting. For Oracle assets, the username should be in exact case as defined in table \"dba_users\", otherwise Oracle normally converts everything to all-caps."
		wallet_dir = null # wallet_dir description: "Path to the Oracle wallet directory"
	}
}

# Create an AWS Log Group log_aggregator
resource "dsfhub_log_aggregator" "rds-mysql-stats-demo-log-group" {
	server_type = "AWS LOG GROUP"
	admin_email         = var.admin_email	
	asset_display_name  = var.log_group_name
	asset_id            = var.log_group_arn
	gateway_id          = var.gateway_jsonar_uid
	parent_asset_id     = dsf_data_source.example_oracle_password_and_wallet.asset_id	
	asset_connection {
		auth_mechanism = "default"
		reason = "default" 
		region = "us-east-2" 
	}
}
```

## Argument Reference

The following arguments are supported:

* `dsfhub_host` - (Required) The DSF host endpoint for [DSF HUB API](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/84552.htm) operations. Example: 'https://1.2.3.4:8443'. Can be set via `TF_VAR_dsfhub_host` shell [environment variable](https://en.wikipedia.org/wiki/Environment_variable).
* `dsfhub_token` - (Required) The [DSF API Token](https://docs.imperva.com/bundle/v4.13-sonar-user-guide/page/84555.htm) for API operations. You can retrieve this from the DSF management hub console. Can be set via `TF_VAR_dsfhub_token` shell [environment variable](https://en.wikipedia.org/wiki/Environment_variable).  
* `insecure_ssl` - (Optional) The boolean flag that instructs the provider to allow for insecure SSL API calls to a DSF Hub instance to support tests against instances with self-signed certificates.
