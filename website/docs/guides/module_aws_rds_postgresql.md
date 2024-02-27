---
subcategory: "Example Assets"
layout: "dsfhub"
page_title: "AWS RDS POSTGRESQL - Log Group"
description: |-
  Provides an combined example of creating an AWS RDS POSTGRESQL database, associated option groups enabling audit logs, onboarding to the DSFHUB with IAM permissions for the DSF Agentless Gateway to access.
---

# AWS RDS POSTGRESQL Onboarding Template

Provides a module template for creating an AWS RDS POSTGRESQL database, the associated option groups enabling audit logs, creating the [dsfhub_data_source](../r/data_source.md) resource to onboard to the DSFHUB with IAM permissions for the DSF Agentless Gateway.

<details>
<summary>AWS RDS POSTGRESQL Variables</summary>

### AWS RDS POSTGRESQL Variables

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
resource "aws_db_parameter_group" "postgresql_param_group" {
  name   = var.deployment_name
  family = "postgres15"

  parameter {
    name  = "log_connections"
    value = "1"
    apply_method = "immediate"
  }

  parameter {
    name  = "log_disconnections"
    value = "1"
    apply_method = "immediate"
  }

  parameter {
    name  = "log_error_verbosity"
    value = "verbose"
    apply_method = "immediate"
  }

  parameter {
    name  = "log_min_duration_statement"
    value = "5000"
    apply_method = "immediate"
  }

  parameter {
    name  = "pgaudit.log"
    value = "all"
    apply_method = "immediate"
  }

  parameter {
    name  = "pgaudit.role"
    value = "rds_pgaudit"
    apply_method = "immediate"
  }

  parameter {
    name  = "shared_preload_libraries"
    value = "pgaudit,pg_stat_statements"
    apply_method = "pending-reboot"
  }
}

resource "aws_db_instance" "postgresql_db" {
  depends_on 			 = [aws_db_parameter_group.postgresql_param_group]
  allocated_storage      = var.db_allocated_storage
  engine                 = "postgres"
  engine_version         = var.db_engine_version
  identifier             = lower(var.db_name)
  instance_class         = var.db_instance_class
  license_model          = "postgresql-license"
  skip_final_snapshot    = true
  apply_immediately      = true

  # Credentials
  username               = var.db_master_username
  password               = var.db_master_password

  # network
  publicly_accessible    = true
  db_subnet_group_name   = var.db_subnet_group_name
  vpc_security_group_ids = var.vpc_security_group_ids

  # audit
  enabled_cloudwatch_logs_exports = ["postgresql","upgrade"]
  parameter_group_name            = aws_db_parameter_group.postgresql_param_group.name
}
```

## dsfhub_data_source resource for AWS RDS POSTGRESQL

The following is an example of the dsfhub_data_source resource used to onboard the RDS Postgresql database to the DSFHUB.

```hcl
# ### Resource example for AWS ###
resource "dsfhub_data_source" "rds-postgresql-db" {
  depends_on = [aws_db_instance.postgresql_db]
  server_type = "AWS RDS POSTGRESQL"

  admin_email = var.admin_email
  asset_display_name  = aws_db_instance.postgresql_db.identifier
  asset_id            = aws_db_instance.postgresql_db.arn
  gateway_id          = var.gateway_id
  server_host_name    = aws_db_instance.postgresql_db.arn
  region              = var.region
  server_port         = aws_db_instance.postgresql_db.port
  version             = var.db_engine_version
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

<details>
<summary>Granting Agentless Gateway rds:RebootDBInstance IAM Permission</summary>

```hcl
resource "aws_iam_role_policy_attachment" "log_group_policy_attachment" {
  policy_arn = aws_iam_policy.log_group_policy.arn
  role       = data.aws_iam_role.agentless_gateway.name
}

resource "aws_iam_policy" "db_reboot_policy" {
  name        = "DSFAgentlessGatewayDBRebootPolicy-${var.deployment_name}"
  description = "DSF Agentless Gateway DB Reboot Policy for ${var.deployment_name}"

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
      {
        "Sid": "VisualEditor0",
        "Effect": "Allow",
        "Action": [
          "rds:RebootDBInstance"
        ]
        "Resource": [
          "${aws_db_instance.postgresql_db.arn}",
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "db_reboot_policy_attachment" {
  policy_arn = aws_iam_policy.db_reboot_policy.arn
  role       = data.aws_iam_role.agentless_gateway.name
}
```
</details>


<details>
<summary>Script examples to reboot and configure the database</summary>

## Script examples to reboot and configure the database  

The following scripts are remotely executed through the agentless gateway using a `remote-execute` with a [null_resource](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) against the PostgreSQL server to reboot the database (required for the shared_preload_libraries param in the aws_db_parameter_group), and configure database granting the user `rds_superuser` permissions to the `rds_pgaudit` role, and creating the `pgaudit` extension.

### Create the following scripts in the same directory as the terraform configuration file

Script 1: `configure_database.sh`
```bash
source /etc/sysconfig/jsonar
export DB_HOST="$DB_HOST"
export DB_PORT="$DB_PORT"
export ADMIN_USER="$ADMIN_USER"
export ADMIN_PASSWORD="$ADMIN_PASSWORD"
export DB_NAME="$DB_NAME"
export JSONAR_BASEDIR="$JSONAR_BASEDIR"
export REGION="$REGION"
export ACTION="$ACTION"

if [ "$ACTION" == "REBOOT" ]; then
	export $(cat /etc/sysconfig/jsonar)
	${JSONAR_BASEDIR}/bin/aws rds reboot-db-instance --db-instance-identifier $DB_NAME --region $REGION
else:
	${JSONAR_BASEDIR}/bin/python3 /tmp/configure_database.py
fi
```

Script 2: `configure_database.py`
```python
import os
import ssl
import pg8000

host = os.environ.get("DB_HOST")
port = os.environ.get('DB_PORT')
user = os.environ.get('ADMIN_USER')
password = os.environ.get('ADMIN_PASSWORD')
database = "postgres"

ssl_context = ssl.SSLContext()

client = pg8000.connect(host=host, 
                        port=port, 
                        user=user, 
                        password=password, 
                        database=database,
                        ssl_context=ssl_context)

client.run("GRANT rds_superuser TO "+user)
client.run("ALTER USER "+user+" WITH CREATEROLE")

auditor_role_exists = client.run("SELECT rolname FROM pg_catalog.pg_roles WHERE rolname = 'rds_pgaudit'")
if auditor_role_exists:
	print('Auditor role "rds_pgaudit" already exists.')
else:
	print('Creating auditor role "rds_pgaudit".')
	client.run("CREATE ROLE rds_pgaudit")
	client.run("COMMIT")
	
audit_extension_exists = client.run("SELECT extname FROM pg_catalog.pg_extension WHERE extname = 'pgaudit'")
if audit_extension_exists:
	print('Audit extension "pgaudit" already exists.')
else:
	print('Creating audit extension "pgaudit".')
	client.run("CREATE EXTENSION pgaudit")
	client.run("COMMIT")
```

### null_resource examples to execute the scripts

```hcl
### Uploading scripts to the agentless gateway ###
resource "null_resource" "upload-script" {
    depends_on = [aws_db_instance.postgresql_db]

	connection {
		type        		= "ssh"
		user        		= var.dsf_gateway_ssh_user
		private_key 		= file(var.dsf_gateway_private_key)
		host        		= var.dsf_gateway_host
		bastion_host 		= var.bastion_host
		bastion_user 		= var.bastion_ssh_user
		bastion_private_key = file(var.bastion_private_key)
	}

	provisioner "file" {
		source      = "configure_database.sh"
		destination = "/tmp/configure_database.sh"
	}
	provisioner "file" {
		source      = "configure_database.py"
		destination = "/tmp/configure_database.py"
	}
}

### The shared_preload_libraries in the aws_db_parameter_group requires ###
### a reboot, this parameter is required for the pgaudit extension to   ###
### work, executed in the null_resource.remote-execute below            ###
resource "null_resource" "reboot-database" {
    depends_on = [null_resource.upload-script]
	connection {
		type        		= "ssh"
		user        		= var.dsf_gateway_ssh_user
		private_key 		= file(var.dsf_gateway_private_key)
		host        		= var.dsf_gateway_host
		bastion_host 		= var.bastion_host
		bastion_user 		= var.bastion_ssh_user
		bastion_private_key = file(var.bastion_private_key)
	}

	provisioner "remote-exec" {
    	inline = [
			"export DB_HOST='${aws_db_instance.postgresql_db.address}'",
			"export DB_NAME='${var.db_name}'",
			"export DB_PORT='${aws_db_instance.postgresql_db.port}'",
			"export ADMIN_USER='${var.db_master_username}'",
			"export ADMIN_PASSWORD='${var.db_master_password}'",
			"export ACTION='REBOOT'",
			"export REGION='${var.region}'",
			"source /etc/sysconfig/jsonar",
			"chmod +x /tmp/configure_database.sh",
			"/tmp/configure_database.sh",
		]
	}
}

### Execute the script to configure the database ###
resource "null_resource" "reboot-database" {
    depends_on = [null_resource.upload-script]
	connection {
		type        		= "ssh"
		user        		= var.dsf_gateway_ssh_user
		private_key 		= file(var.dsf_gateway_private_key)
		host        		= var.dsf_gateway_host
		bastion_host 		= var.bastion_host
		bastion_user 		= var.bastion_ssh_user
		bastion_private_key = file(var.bastion_private_key)
	}

	provisioner "remote-exec" {
    	inline = [
			"export DB_HOST='${aws_db_instance.postgresql_db.address}'",
			"export DB_NAME='postgres'",
			"export DB_PORT='${aws_db_instance.postgresql_db.port}'",
			"export ADMIN_USER='${var.db_master_username}'",
			"export ADMIN_PASSWORD='${var.db_master_password}'",
			"export ACTION='CONFIGURE'",
			"export REGION='${var.region}'",
			"source /etc/sysconfig/jsonar",
			"chmod +x /tmp/configure_database.sh",
			"/tmp/configure_database.sh",
		]
	}
}
```
</details>

## Agentless Gateway Permission Dependencies:

The [DSF Agentless Gateway](https://registry.terraform.io/modules/imperva/dsf-agentless-gw/aws/latest) is required to have [AWS IAM Role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) access to the AWS service the database is configured to publish logs to in order to consume audit.

<ul>
<li><a target="_blank" href="aws_iam_log_group.md">AWS IAM Permissions for CloudWatch Log Groups</a></li>
</ul>