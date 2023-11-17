---
subcategory: "Example Modules"
layout: "dsfhub"
page_title: "AWS RDS POSTGRESQL - Log Group"
description: |-
  Provides an combined example of creating an AWS RDS POSTGRESQL database, associated option groups enabling audit logs, onboarding to the DSFHUB with IAM permissions for the DSF Agentless Gateway to access.
---

# AWS RDS MSSQL Onboarding Template

Provides a module template for creating an AWS RDS POSTGRESQL database, the associated option groups enabling audit logs, creating the [dsfhub_data_source](../r/data_source.md) resource to onboard to the DSFHUB with IAM permissions for the DSF Agentless Gateway.

<details>
<summary>AWS RDS MSSQL Variables</summary>

## AWS RDS MSSQL Variables

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
</details>

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
esource "aws_db_parameter_group" "postgresql_param_group" {
  name   = var.deployment_name
  family = "postgres15"

  parameter {
    name  = "log_connections"
    value = "1"
  }
}

resource "aws_db_instance" "postgresql_db" {
  allocated_storage      = var.db_allocated_storage
  engine                 = "postgres"
  engine_version         = var.db_engine_version
  identifier             = lower(var.db_name)
  instance_class         = var.db_instance_class
  license_model          = "postgresql-license"
  skip_final_snapshot    = true

  # Credentials
  username               = var.db_master_username
  password               = var.db_master_password

  # network
  publicly_accessible    = true
  db_subnet_group_name   = var.db_subnet_group_name
  vpc_security_group_ids = var.vpc_security_group_ids

  # audit
  enable_cloudwatch_logs_exports = ["postgresql","upgrade"]
  parameter_group_name   = aws_db_parameter_group.postgresql_param_group.name
}

# ### DSFHUB Resources ###
resource "dsfhub_data_source" "rds-postgresql-db" {
  server_type = "AWS RDS POSTGRESQL"

  admin_email = var.admin_email
  asset_display_name  = aws_db_instance.postgresql_db.identifier
  asset_id            = aws_db_instance.postgresql_db.arn
  gateway_id          = var.gateway_id
  server_host_name    = aws_db_instance.postgresql_db.arn
  region              = var.region
  server_port         = aws_db_instance.postgresql_db.port
  version             = var.engine_version
  parent_asset_id     = var.dsf_cloud_account_asset_id
  audit_pull_enabled  = true

  asset_connection {
    auth_mechanism  = "password"
    password        = var.db_master_password
    reason          = "default"
    username        = var.db_master_username
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