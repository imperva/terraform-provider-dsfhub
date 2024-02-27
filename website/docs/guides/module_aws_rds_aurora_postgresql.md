---
subcategory: "Example Assets"
layout: "dsfhub"
page_title: "AWS RDS AURORA POSTGRESQL CLUSTER - Log Group"
description: |-
  Provides an combined example of creating an AWS RDS AURORA POSTGRESQL CLUSTER database, associated option groups enabling audit logs, onboarding to the DSFHUB with IAM permissions for the DSF Agentless Gateway to access.
---

# AWS RDS AURORA POSTGRESQL CLUSTER Onboarding Template - Multi-region

Provides a module template for creating an AWS RDS AURORA POSTGRESQL CLUSTER database spanning multiple regions, the associated option groups enabling audit logs, creating the [dsfhub_data_source](../r/data_source.md) resource to onboard to the DSFHUB with IAM permissions for the DSF Agentless Gateways.

<details>
<summary>AWS RDS AURORA POSTGRESQL CLUSTER Variables</summary>

### AWS RDS AURORA POSTGRESQL CLUSTER Variables

```hcl
# DSFHUB Provider Required Variables
variable "dsfhub_token" {} # TF_VAR_dsfhub_token env variable
variable "dsfhub_host" {} # TF_VAR_dsfhub_host env variable

# AWS Provider Required Variables
variable "region_primary" {
  description = "Cluster primary AWS region"
  type = string
  default = "us-east-1"
}

variable "region_secondary" {
  description = "Cluster secondary AWS region"
  type = string
  default = "us-east-2"
}

# DSFHUB Asset Variables
variable "admin_email" {
  description = "The email address to notify about this asset"
  type = string
  default = "your@email.com"
}

variable "dsf_cloud_account_asset_id" {
  description =  "DSFHUB Cloud Account Asset ID"
  type = string
  default = "arn:aws:iam::1234567890:user/your-user"
}

variable "gateway_id_primary" {
  description =  "The jsonarUid unique identifier of the agentless gateway in the primary region. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'"
  type = string
  default = "12345abcde-12345-abcde-12345-12345abcde"
}

variable "gateway_id_secondary" {
  description =  "The jsonarUid unique identifier of the agentless gateway in the secondary region. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'"
  type = string
  default = "12345abcde-12345-abcde-12345-12345abcde"
}

# RDS-DB Variables
variable "deployment_name" {
  description = "The name of the database deployment. i.e. 'aurora-demo-db-cluster-pg'"
  type = string
  default = "aurora-demo-db-cluster-pg"
}

variable "db_allocated_storage" {
  description = "The allocated storage in gibibytes. If max_allocated_storage is configured, this argument represents the initial storage allocation and differences from the configuration will be ignored automatically when Storage Autoscaling occurs. If replicate_source_db is set, the value is ignored during the creation of the instance."
  type = number
  default = 20
}

variable "db_audit_type" {
  description = "Example Values: LOG_GROUP, AGGREGATED, UNIFIED_AGGREGATED, UNIFIED. Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details"
  type = string
  default = "UNIFIED"
}

variable "db_cluster_instance_class" {
  description = "The instance type of the RDS instance. Example: 'db.t2'. Reference: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html"
  type = string
  default = "db.r5.large"
}

variable "db_cluster_instance_name" {
  description = "Identifier for the RDS instance"
  type = string
  default = "yourinstancename"
}

variable "db_engine_version" {
  description = "Database engine version, i.e. \"11.9\""
  type = string
  default = "11.9"
}

variable "db_family" {
  description = "The family of the DB parameter group. Example: 'aurora-postgresql13'."
  type        = string
  default     = "aurora-postgresql11"
}

variable "db_identifier" {
  type        = string
  description = "Identifier of RDS instance"
  default = "aurora-demo-db-cluster-pg"
}

variable "db_instance_class" {
  description = "The instance type of the RDS instance. Example: 'db.t2'. Reference: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html"
  type = string
  default = "db.t3.medium"
}

variable "db_license_model" {
  description = "(Optional, but required for some DB engines, i.e., Oracle SE1) License model information for this DB instance. Valid values for this field are as follows: RDS for Oracle: bring-your-own-license | license-included"
  type = string
  default = "bring-your-own-license"
}

variable "db_master_username" {
  description = "Username for the master DB user, must not use rdsamin as that is reserved. Cannot be specified for a replica."
  type = string
  default = "admin"
}

variable "db_master_password" {
  description = "Password for the master DB user. Note that this may show up in logs, and it will be stored in the state file."
  type = string
  default = "your-password-here"
}

variable "db_name" {
  description = "The database name (must begin with a letter and contain only alphanumeric characters)."
  type = string
  default = "YOURPOSTGRESQLCLUSTER"
}

variable "db_parameter_group_family" {
  description = "Database engine version, i.e. \"aurora-postgresql13\""
  type = string
  default = "aurora-postgresql13"
}

variable "db_parameter_group_name" {
  type        = string
  description = "Name of rds parameter group"
  default = "postgresql-parameter-group"
}

variable "db_subnet_group_name" {
  description = "Name of DB subnet group. DB instance will be created in the VPC associated with the DB subnet group. If unspecified, will be created in the default VPC, or in EC2 Classic, if available."
  type = string
  default = "your-db-subnet-group"
}

variable "vpc_security_group_ids" {
  description = "List of VPC security group IDs needed for TLC access"
  type        = list(string)
  default = [ "sg-12345abcde" ]
}

variable "bastion_host" {
  description = "Host of the bastion to connect to for executing remote database init script."
  type        = string
  default     = "1.2.3.4"
}

variable "dsf_gateway_private_key" {
  description = "Certificate used to connect to the bastion host."
  type        = string
  default     = "~/.ssh/your-key.pem"
}

variable "bastion_ssh_user" {
  description = "Username of the user connecting to the bastion host."
  type        = string
  default     = "ec2-user"
}

variable "dsf_gateway_host_primary" {
  description = "Host of the dsf gateway to connect to for executing remote database init script."
  type        = string
  default     = "172.1.2.3"
}

variable "dsf_gateway_ssh_user" {
  description = "Username of the user connecting to the DSF gateway host."
  type        = string
  default     = "ec2-user"
}

variable "bastion_private_key" {
  description = "Certificate used to connect to the DSF gateway."
  type        = string
  default     = "/path/to/your-key.pem"
}

# DSF Agentless-Gateway Variables for IAM permissions granting access logs
variable "agentless_gateway_iam_role_name" {
  description = "Name of the DSF agentless gateway role to add permissions to access db logs."
  type = string
  default = "your-gw-role-here"
}

```
</details>

## Providers and Resources

```hcl
## Providers ###
provider "aws" {
  alias  = "primary"
  region = var.region_primary
}

provider "aws" {
  alias  = "secondary"
  region = var.region_secondary
}

provider "dsfhub" {
  dsfhub_token = var.dsfhub_token
  dsfhub_host = var.dsfhub_host
}

### AWS Resources ###
resource "aws_rds_cluster_parameter_group" "rds_aurora_postgresql_cluster_pg_primary" {
  provider    = aws.primary
  name        = "${var.deployment_name}-parameter-group-primary"
  description = "${var.deployment_name}-parameter-group-primary"
  family      = var.db_family

  parameter {
    apply_method = "immediate"
    name         = "log_connections"
    value        = "1"
  }
  parameter {
    apply_method = "immediate"
    name         = "log_disconnections"
    value        = "1"
  }
  parameter {
    apply_method = "immediate"
    name         = "log_error_verbosity"
    value        = "verbose"
  }
  parameter {
    apply_method = "immediate"
    name         = "pgaudit.log"
    value        = "all"
  }
  parameter {
    apply_method = "immediate"
    name         = "pgaudit.role"
    value        = "rds_pgaudit"
  }
  parameter {
    apply_method = "pending-reboot"
    name         = "shared_preload_libraries"
    value        = "pgaudit,pg_stat_statements"
  }
}

resource "aws_rds_global_cluster" "rds_aurora_postgresql_cluster" {
  global_cluster_identifier = lower(var.db_name)
  engine                    = "aurora-postgresql"
  engine_version            = var.db_engine_version
  database_name             = var.db_name
}

resource "aws_rds_cluster" "primary" {
  provider                  = aws.primary
  engine                    = aws_rds_global_cluster.rds_aurora_postgresql_cluster.engine
  engine_version            = aws_rds_global_cluster.rds_aurora_postgresql_cluster.engine_version
  cluster_identifier        = "primary-cluster-${lower(var.db_name)}"
  master_username           = var.db_master_username
  master_password           = var.db_master_password
  database_name             = lower(var.db_name)
  global_cluster_identifier = aws_rds_global_cluster.rds_aurora_postgresql_cluster.id
  db_subnet_group_name      = var.db_subnet_group_name
  skip_final_snapshot       = true
  # audit
  enabled_cloudwatch_logs_exports = ["postgresql"]
  db_cluster_parameter_group_name = aws_rds_cluster_parameter_group.rds_aurora_postgresql_cluster_pg_primary.name
}

resource "aws_rds_cluster_instance" "primary" {
  provider             = aws.primary
  engine               = aws_rds_global_cluster.rds_aurora_postgresql_cluster.engine
  engine_version       = aws_rds_global_cluster.rds_aurora_postgresql_cluster.engine_version
  identifier           = "primary-instance-${lower(var.db_name)}"
  cluster_identifier   = aws_rds_cluster.primary.id
  instance_class       = "db.r5.large"
  db_subnet_group_name = var.db_subnet_group_name
}

# Secondary AWS cluster resources
resource "aws_rds_cluster_parameter_group" "rds_aurora_postgresql_cluster_pg_secondary" {
  provider    = aws.secondary
  name        = "${var.deployment_name}-parameter-group-secondary"
  description = "${var.deployment_name}-parameter-group-secondary"
  family      = var.db_family

  parameter {
    apply_method = "immediate"
    name         = "log_connections"
    value        = "1"
  }
  parameter {
    apply_method = "immediate"
    name         = "log_disconnections"
    value        = "1"
  }
  parameter {
    apply_method = "immediate"
    name         = "log_error_verbosity"
    value        = "verbose"
  }
  parameter {
    apply_method = "immediate"
    name         = "pgaudit.log"
    value        = "all"
  }
  parameter {
    apply_method = "immediate"
    name         = "pgaudit.role"
    value        = "rds_pgaudit"
  }
  parameter {
    apply_method = "pending-reboot"
    name         = "shared_preload_libraries"
    value        = "pgaudit,pg_stat_statements"
  }
}

resource "aws_rds_cluster" "secondary" {
  provider                  = aws.secondary
  engine                    = aws_rds_global_cluster.rds_aurora_postgresql_cluster.engine
  engine_version            = aws_rds_global_cluster.rds_aurora_postgresql_cluster.engine_version
  cluster_identifier        = "secondary-cluster-${lower(var.db_name)}"
  global_cluster_identifier = aws_rds_global_cluster.rds_aurora_postgresql_cluster.id
  db_subnet_group_name      = "default"
  skip_final_snapshot       = true
  # audit configurations
  enabled_cloudwatch_logs_exports = ["postgresql"]
  db_cluster_parameter_group_name = aws_rds_cluster_parameter_group.rds_aurora_postgresql_cluster_pg_secondary.name
  
  lifecycle {
    ignore_changes = [
      replication_source_identifier
    ]
  }
  
  depends_on = [
    aws_rds_cluster_instance.primary
  ]
}

resource "aws_rds_cluster_instance" "secondary" {
  provider             = aws.secondary
  engine               = aws_rds_global_cluster.rds_aurora_postgresql_cluster.engine
  engine_version       = aws_rds_global_cluster.rds_aurora_postgresql_cluster.engine_version
  identifier           = "secondary-instance-${lower(var.db_name)}"
  cluster_identifier   = aws_rds_cluster.secondary.id
  instance_class       = "db.r5.large"
  db_subnet_group_name = "default"
}
```

<details>
<summary>Granting Agentless Gateway LogGroup Access IAM Permission</summary>

```hcl
data "aws_cloudwatch_log_group" "postgresql_cluster_log_group_primary" {
  provider    = aws.primary
  depends_on  = [aws_rds_cluster.primary]
  name        = "/aws/rds/cluster/${aws_rds_cluster.primary.cluster_identifier}/postgresql"
}

data "aws_cloudwatch_log_group" "postgresql_cluster_log_group_secondary" {
  provider    = aws.secondary
  depends_on  = [aws_rds_cluster.secondary]
  name        = "/aws/rds/cluster/${aws_rds_cluster.secondary.cluster_identifier}/postgresql"
}

data "aws_iam_role" "agentless_gateway" {
  name = var.agentless_gateway_iam_role_name
}

resource "aws_iam_policy" "log_group_policy" {
  name        = "DSFAgentlessGatewayLogGroupPolicy-${var.deployment_name}"
  description = "DSF Agentless Gateway Log Group Policy for ${var.deployment_name}"

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
      {
        "Sid": "VisualEditor0",
        "Effect": "Allow",
        "Action": [
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:FilterLogEvents",
          "logs:GetLogEvents"
        ]
        "Resource": [
          "${data.aws_cloudwatch_log_group.postgresql_cluster_log_group_primary.arn}:*",
          "${data.aws_cloudwatch_log_group.postgresql_cluster_log_group_secondary.arn}:*",
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "log_group_policy_attachment" {
  policy_arn = aws_iam_policy.log_group_policy.arn
  role       = data.aws_iam_role.agentless_gateway.name
}
```
</details>

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
          "${aws_rds_cluster_instance.primary.arn}",
          "${aws_rds_cluster_instance.secondary.arn}",
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

## dsfhub_data_source resource for AWS RDS AURORA POSTGRESQL CLUSTER (multi-region)

The following is an example of the dsfhub_data_source resource used to onboard the Aurora Postgresql Cluster database to the DSFHUB.  For a multi-region deployment, repeat using the region specific cluster parameters and appropriate associated gateway_id for that region.

```hcl
### DSFHUB Resources - Primary Cluster Region Example ###
resource "dsfhub_data_source" "rds-postgresql-cluster-primary" {
  depends_on = [aws_rds_cluster.primary]
  server_type = "AWS RDS AURORA POSTGRESQL CLUSTER"

  admin_email = var.admin_email
  asset_display_name  = aws_rds_cluster.primary.cluster_identifier
  asset_id            = aws_rds_cluster.primary.arn
  gateway_id          = var.gateway_id_primary
  server_host_name    = aws_rds_cluster.primary.endpoint
  region              = var.region
  server_port         = aws_rds_cluster.primary.port
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

# ### DSFHUB Resources - Secondary Cluster Region Example ###
resource "dsfhub_data_source" "rds-postgresql-cluster-secondary" {
  depends_on = [aws_db_instance.postgresql_db]
  server_type = "AWS RDS AURORA POSTGRESQL  CLUSTER"

  admin_email = var.admin_email
  asset_display_name  = aws_rds_cluster.secondary.cluster_identifier
  asset_id            = aws_rds_cluster.secondary.arn
  gateway_id          = var.gateway_id_secondary
  server_host_name    = aws_rds_cluster.secondary.endpoint
  region              = var.region
  server_port         = aws_rds_cluster.secondary.port
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
export DB_INSTANCE_IDENTIFIER="$DB_INSTANCE_IDENTIFIER"
export JSONAR_BASEDIR="$JSONAR_BASEDIR"
export REGION="$REGION"
export ACTION="$ACTION"

if [ "$ACTION" == "REBOOT" ]; then
	export $(cat /etc/sysconfig/jsonar)
	${JSONAR_BASEDIR}/bin/aws rds reboot-db-instance --db-instance-identifier $DB_INSTANCE_IDENTIFIER --region $REGION
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
  depends_on = [aws_rds_cluster_instance.primary]

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
      "export DB_HOST='${aws_rds_cluster_instance.primary.endpoint}'",
      "export DB_INSTANCE_IDENTIFIER='${aws_rds_cluster_instance.primary.cluster_identifier}'",
      "export DB_PORT='${aws_rds_cluster_instance.primary.port}'",
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
resource "null_resource" "configure-database" {
  depends_on = [null_resource.reboot-database]
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
      "export DB_HOST='${aws_rds_cluster_instance.primary.endpoint}'",
      "export DB_INSTANCE_IDENTIFIER='${aws_rds_cluster_instance.primary.cluster_identifier}'",
      "export DB_PORT='${aws_rds_cluster_instance.primary.port}'",
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