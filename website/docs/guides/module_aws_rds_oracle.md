---
subcategory: "Example Assets"
layout: "dsfhub"
page_title: "AWS RDS POSTGRESQL - Log Group"
description: |-
  Provides an combined example of creating an AWS RDS POSTGRESQL database, associated option groups enabling audit logs, onboarding to the DSFHUB with IAM permissions for the DSF Agentless Gateway to access.
---

# AWS RDS Oracle Onboarding Template

Provides a module template for creating an AWS RDS Oracle database, the associated option groups enabling audit logs, creating the [dsfhub_data_source](../r/data_source.md) resource to onboard to the DSFHUB with IAM permissions for the DSF Agentless Gateway.

<details>
<summary>AWS RDS Oracle Variables</summary>

## AWS RDS Oracle Variables

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

# RDS-DB Variables
variable "bastion_host" {
  description = "Host of the bastion to connect to for executing remote database init script."
  type        = string
  default     = "1.2.3.4"
}

variable "bastion_private_key" {
  description = "Certificate used to connect to the DSF gateway."
  type        = string
  default     = "/path/to/your/key.pem"
}

variable "bastion_ssh_user" {
  description = "Username of the user connecting to the bastion host."
  type        = string
  default     = "ec2-user"
}

variable "db_allocated_storage" {
  description = "The allocated storage in gibibytes. If max_allocated_storage is configured, this argument represents the initial storage allocation and differences from the configuration will be ignored automatically when Storage Autoscaling occurs. If replicate_source_db is set, the value is ignored during the creation of the instance."
  type = number
  default = 20
}

variable "db_engine_version" {
  description = "Database engine version, i.e. \"19\""
  type = string
  default = "19"
}

variable "db_engine_major_version" {
  description = "Database engine major version, i.e. \"19\""
  type = string
  default = "19"
}

variable "db_audit_type" {
  description = "Example Values: LOG_GROUP, AGGREGATED, UNIFIED_AGGREGATED, UNIFIED. Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details"
  type = string
  default = "UNIFIED"
}

variable "db_audit_policy_name" {
  description = "Name of the audit policy to configure in the database."
  type = string
  default = "imperva_audit_policy"
}

variable "db_audit_username" {
  description = "Username for the audit user dsf will use to connect to the database with."
  type = string
  default = "impvaudituser"
}

variable "db_audit_password" {
  description = "Password for the audit user dsf will use to connect to the database with."
  type = string
  default = ""
}

variable "dsf_gateway_private_key" {
  description = "Certificate used to connect to the bastion host."
  type        = string
  default     = "/path/to/your/key.pem"
}

variable "dsf_gateway_host" {
  description = "Host of the dsf gateway to connect to for executing remote database init script."
  type        = string
  default     = "5.6.7.8"
}

variable "dsf_gateway_ssh_user" {
  description = "Username of the user connecting to the DSF gateway host."
  type        = string
  default     = "ec2-user"
}

variable "db_identifier" {
  type        = string
  description = "Identifier of RDS instance"
  default = "rds-oracle"
}

variable "db_instance_class" {
  description = "The instance type of the RDS instance. Example: 'db.t2'. Reference: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html"
  type = string
  default = "db.t3.small"
}

variable "db_license_model" {
  description = "(Optional, but required for some DB engines, i.e., Oracle SE1) License model information for this DB instance. Valid values for this field are as follows: RDS for Oracle: bring-your-own-license | license-included"
  type = string
  default = "bring-your-own-license"
}

variable "db_master_username" {
  description = "Username for the master DB user, must not use rdsamin as that is reserved. Cannot be specified for a replica."
  type = string
  default = "yourusername"
}

variable "db_master_password" {
  description = "Password for the master DB user. Note that this may show up in logs, and it will be stored in the state file."
  type = string
  default = ""
}

variable "db_parameter_group_family" {
  description = "Database engine version, i.e. \"oracle-ee-19\""
  type = string
  default = "oracle-ee-19"
}

variable "db_parameter_group_name" {
  type        = string
  description = "Name of rds parameter group"
  default = "rds-oracle-parameter-group"
}

variable "db_subnet_group_name" {
  description = "Name of DB subnet group. DB instance will be created in the VPC associated with the DB subnet group. If unspecified, will be created in the default VPC, or in EC2 Classic, if available."
  type = string
  default = "your-db-subnet-group"
}

variable "db_vpc_security_groups" {
  description = "List of VPC security group IDs needed for TLC access"
  type        = list(string)
  default = ["sg-001575eb169c1dbcd"]
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
resource "aws_db_parameter_group" "rds_oracle_pg" {
  name   = var.db_parameter_group_name
  family = var.db_parameter_group_family
  # Enable auditing
  parameter {
    name            = "audit_trail"
    value           = "DB,EXTENDED"
    apply_method    = "pending-reboot"
  }
  parameter {
    name            = "audit_sys_operations"
    value           = "TRUE"
    apply_method    = "pending-reboot"
  }
}

# Database
resource "aws_db_instance" "rds-oracle" {
  allocated_storage         = var.db_allocated_storage
  apply_immediately         = true
  db_subnet_group_name      = var.db_subnet_group_name
  engine                    = "oracle-ee"
  engine_version            = var.db_engine_version
  identifier                = var.db_identifier
  instance_class            = var.db_instance_class
  license_model             = var.db_license_model
  parameter_group_name      = aws_db_parameter_group.rds_oracle_pg.name
  password                  = var.db_master_password
  publicly_accessible       = true
  skip_final_snapshot       = true
  username                  = var.db_master_username
  vpc_security_group_ids    = var.db_vpc_security_groups
}

resource "null_resource" "upload-script" {
  depends_on = [aws_db_instance.rds-oracle]

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
}

# ### null_resource used to run database script to create user for dsf to consume logs ###
resource "null_resource" "remote-execute" {
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
      "export ORACLE_HOSTNAME='${aws_db_instance.rds-oracle.address}'",
      "export ORACLE_SID='${aws_db_instance.rds-oracle.db_name}'",
      "export ADMIN_USER='${var.db_master_username}'",
      "export ADMIN_PASSWORD='${var.db_master_password}'",
      "export DB_USER='${var.db_audit_username}'",
      "export DB_PASSWORD='${var.db_audit_password}'",
      "export POLICY_NAME='${var.db_audit_policy_name}'",
      "chmod +x /tmp/configure_database.sh",
      "/tmp/configure_database.sh",
    ]
  }
}

# ### DSFHUB Resources ###
resource "dsfhub_data_source" "aws_rds_oracle" {
  depends_on = [null_resource.remote-execute]
  server_type = "AWS RDS ORACLE"
  admin_email = var.admin_email
  asset_display_name = aws_db_instance.rds-oracle.arn
  asset_id = aws_db_instance.rds-oracle.arn
  gateway_id = var.gateway_id

  server_host_name = aws_db_instance.rds-oracle.address
  service_name = aws_db_instance.rds-oracle.db_name
  server_port = aws_db_instance.rds-oracle.port
  audit_pull_enabled  = true
  audit_type = var.db_audit_type
  version = var.db_engine_major_version

  asset_connection {
    auth_mechanism = "password"
    password = var.db_audit_password
    reason = "default"
    username = var.db_audit_username
  }
}

```