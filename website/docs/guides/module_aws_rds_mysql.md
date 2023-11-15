---
subcategory: "Example Modules"
layout: "dsfhub"
page_title: "Modules - AWS RDS MYSQL"
description: |-
  Provides an combined example of creating an AWS RDS MYSQL database, associated option groups enabling audit logs, onboarding to the DSFHUB with IAM permissions for the DSF Agentless Gateway to access.
---

# AWS RDS MYSQL Onboarding Template

Provides a module template for creating an AWS RDS MYSQL database, the associated option groups enabling audit logs, creating the [dsfhub_data_source](../r/data_source.md) and [dsfhub_log_aggregator](../r/log_aggregator.md) records to onboard to the DSFHUB with IAM permissions for the DSF Agentless Gateway.

## Variables

```hcl
# DSFHUB Provider Required Variables
variable "dsfhub_token" {} # TF_VAR_dsfhub_token env variable
variable "dsfhub_host" {} # TF_VAR_dsfhub_host env variable

# AWS Provider Required Variables
variable "region" {
  description = "AWS region"
  type = string
  default = "us-east-2"
}

# DSFHUB Asset Variables
variable "admin_email" {
  description = "The email address to notify about this asset"
  type = string
  default = "your@email.com"
}

variable "gateway_id" {
  description =  "The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'"
  type = string
  default = "12345abcde-12345-abcde-12345-12345abcde"
}

variable "dsf_cloud_account_asset_id" {
  description =  "DSFHUB Cloud Account Asset ID"
  type = string
  default = "arn:aws:iam::1234567890:user/your-user"
}

# RDS-DB Variables
variable "deployment_name" {
  description = "The name of the database deployment. i.e. 'custom-app-mysql-prod'"
  type = string
  default = "custom-app-mysql-prod"
}

variable "db_name" {
  description = "The database name (must begin with a letter and contain only alphanumeric characters)."
  type = string
  default = "CustomAppMySqlProd"
}

variable "db_allocated_storage" {
  description = "The allocated storage in gibibytes. If max_allocated_storage is configured, this argument represents the initial storage allocation and differences from the configuration will be ignored automatically when Storage Autoscaling occurs. If replicate_source_db is set, the value is ignored during the creation of the instance."
  type = number
  default = 10
}

variable "db_engine_version" {
  description = "Database engine version, i.e. \"8.0.33\""
  type = string
  default = "8.0.33"
}

variable "db_instance_class" {
  description = "The instance type of the RDS instance. Example: 'db.t2'. Reference: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html"
  type = string
  default = "db.t3.micro"
}

variable "db_major_engine_version" {
  description = "Specifies the major version of the engine that this option group should be associated with, i.e. \"8.0\""
  type = string
  default = "8.0"
}

variable "db_master_username" {
  description = "Username for the master DB user, must not use rdsamin as that is reserved. Cannot be specified for a replica."
  type = string
  default = "youradmin"
}

variable "db_master_password" {
  description = "Password for the master DB user. Note that this may show up in logs, and it will be stored in the state file."
  type = string
  default = ""
}

variable "db_subnet_group_name" {
  description = "Name of DB subnet group. DB instance will be created in the VPC associated with the DB subnet group. If unspecified, will be created in the default VPC, or in EC2 Classic, if available."
  type = string
  default = "isbt_db-db-subnet-group"
}

variable "server_audit_excluded_users" {
  description = "A comman seperated string of usernames to exclude activity from the audit feed. By default, activity is recorded for all users. Example: \"rdsadmin,etladmin\""
  type = string
  default = "rdsadmin"
}

variable "vpc_security_group_ids" {
  description = "List of VPC security groups to associate."
  type = list
  default = ["sg-abcde12345"]
}
```

## Providers and Resources

```hcl
## Providers ###
provider "aws" {
  region = var.region
}

provider "dsfhub" {
  dsfhub_token = var.dsfhub_token
  dsfhub_host = var.dsfhub_host
}

### AWS Resources ###
resource "aws_db_option_group" "mysql_option_group" {
  name                        = "${var.deployment_name}-option-group"
  option_group_description    = "${var.deployment_name}-option-group"
  engine_name                 = "mysql"
  major_engine_version        = var.db_major_engine_version

  option {
    option_name = "MARIADB_AUDIT_PLUGIN"
    option_settings {
      name    = "SERVER_AUDIT_EVENTS"
      value   = "CONNECT,QUERY,QUERY_DDL,QUERY_DML,QUERY_DCL,QUERY_DML_NO_SELECT"
    }
    option_settings {
      name    = "SERVER_AUDIT_EXCL_USERS"
      value   = var.server_audit_excluded_users
    }
  }
}

resource "aws_db_instance" "mysql_db" {
  allocated_storage    = var.db_allocated_storage
  db_name              = var.db_name
  engine               = "mysql"
  engine_version       = var.db_engine_version
  identifier           = lower(var.db_name)
  instance_class       = var.db_instance_class
  username             = var.db_master_username
  password             = var.db_master_password
  skip_final_snapshot  = true

  # network
  publicly_accessible       = true
  db_subnet_group_name      = var.db_subnet_group_name
  vpc_security_group_ids    = var.vpc_security_group_ids

  # audit
  enabled_cloudwatch_logs_exports = ["audit"]
  option_group_name    = "${aws_db_option_group.mysql_option_group.name}"
}

# ### DSFHUB Resources ###
data "aws_cloudwatch_log_group" "rds_mysql_log_group" {
  depends_on  = [aws_db_instance.mysql_db]
  name        = "/aws/rds/instance/${aws_db_instance.mysql_db.identifier}/audit"
}

resource "dsfhub_data_source" "rds-mysql-db" {
  server_type = "AWS RDS MYSQL"

  admin_email = var.admin_email
  asset_display_name  = aws_db_instance.mysql_db.identifier
  asset_id            = aws_db_instance.mysql_db.arn
  gateway_id          = var.gateway_id
  server_host_name    = aws_db_instance.mysql_db.arn
  region              = var.region
  server_port         = aws_db_instance.mysql_db.port
  version             = aws_db_option_group.mysql_option_group.major_engine_version
  parent_asset_id     = var.dsf_cloud_account_asset_id

  asset_connection {
    auth_mechanism  = "password"
    password        = var.db_master_password
    reason          = "default"
    username        = aws_db_instance.mysql_db.username
  }
}

resource "dsfhub_log_aggregator" "rds-mysql-db-log-group" {
  server_type = "AWS LOG GROUP"

  admin_email         = var.admin_email
  asset_display_name  = data.aws_cloudwatch_log_group.rds_mysql_log_group.name
  asset_id            = data.aws_cloudwatch_log_group.rds_mysql_log_group.arn
  gateway_id          = var.gateway_id
  parent_asset_id     = dsfhub_data_source.rds-mysql-db.asset_id

  asset_connection {
    auth_mechanism = "default"
    reason = "default"
    region = var.region
  }
}
```

## Agentless Gateway Permission Dependencies:

The [DSF Agentless Gateway](https://registry.terraform.io/modules/imperva/dsf-agentless-gw/aws/latest) is required to have [AWS IAM Role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) access to the AWS service the database is configured to publish logs to in order to consume audit.

<ul>
<li><a target="_blank" href="aws_iam_kinesis.md">AWS IAM Permissions for Kinesis Streams</a></li>
<li><a target="_blank" href="aws_iam_log_group.md">AWS IAM Permissions for CloudWatch Log Groups</a></li>
<li><a target="_blank" href="aws_iam_secrets.md">AWS IAM Permissions for Secret Manager</a></li>
</ul>