---
subcategory: "Example Assets"
layout: "dsfhub"
page_title: "AWS RDS ORACLE - ODBC"
description: |-
  Provides an combined example of creating an AWS RDS ORACLE database, associated option groups enabling audit logs, and onboarding to the DSFHUB.
---

# AWS RDS Oracle Onboarding Template

Provides a module template for creating an AWS RDS Oracle database, the associated option groups enabling audit logs, creating the [dsfhub_data_source](../r/data_source.md) resource to onboard to the DSFHUB a SQL script to create a user for the DSF Agentless Gateway to retrieve logs from the database via ODBC connection.

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
```

## Script to create dsf database user, and remotely execute

The following script is remotely executed through the agentless gateway against the Oracle server to create the user that is used to retrieve database logs via ODBC connection.  Create the `configure_database.sh` file in the same directory that terraform is being executed from, or you can execute this separately o the database to create this user. 

```hcl
#!/usr/bin/env bash
#sudo su - sonarw
source /etc/sysconfig/jsonar
export LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:${JSONAR_BASEDIR}/lib:${JSONAR_BASEDIR}/lib64:${JSONAR_LOCALDIR}/lib

${JSONAR_BASEDIR}/bin/isql -v -k "Driver=${JSONAR_BASEDIR}/lib/libsqora.so; DBQ=$ORACLE_HOSTNAME/$ORACLE_SID; UID=$ADMIN_USER; Pwd=$ADMIN_PASSWORD" << EOF
CREATE USER $DB_USER IDENTIFIED BY $DB_PASSWORD;
GRANT CREATE SESSION TO $DB_USER;
CREATE AUDIT POLICY $POLICY_NAME ACTIONS ALL ONLY TOPLEVEL;
AUDIT POLICY $POLICY_NAME;
EOF
# ${JSONAR_BASEDIR}/bin/isql -v -k "Driver=${JSONAR_BASEDIR}/lib/libsqora.so; DBQ=$ORACLE_HOSTNAME/$ORACLE_SID; UID=$ADMIN_USER; Pwd=$ADMIN_PASSWORD" << EOF
# .i ./create_purge.sql
EOF
```

The following null_resource.remote-execute resource remote executes the script to create the db user.

```hcl
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
```

## dsfhub_data_source resource for AWS RDS ORACLE

The following is an example of the dsfhub_data_source resource used to onboard the Oracle database to the DSFHUB.

```hcl
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

<details>
<summary>Create Audit Purge Jobs on Oracle Server</summary>

The following SQL is an example of creating audit purge jobs on the oracle database server.

```hcl
/* Create audit purge job */
BEGIN
DBMS_SCHEDULER.create_job (
job_name => 'SET_LAST_ARCHIVE_TIME',
job_type => 'PLSQL_BLOCK',
job_action => 'BEGIN
DBMS_AUDIT_MGMT.SET_LAST_ARCHIVE_TIMESTAMP(
audit_trail_type => DBMS_AUDIT_MGMT.AUDIT_TRAIL_UNIFIED,
/* Change this value to set a different archive timestamp */
last_archive_time => sysdate-30
);
END;',
start_date => SYSTIMESTAMP,
/* Update frequency as needed */
repeat_interval => 'freq=daily; byhour=0; byminute=0; bysecond=0;',
end_date => NULL,
enabled => TRUE,
comments => 'last archive time.');
END;
/
BEGIN
DBMS_AUDIT_MGMT.SET_LAST_ARCHIVE_TIMESTAMP(
audit_trail_type => DBMS_AUDIT_MGMT.AUDIT_TRAIL_UNIFIED,
/* Change this value to set a different archive timestamp */
last_archive_time => sysdate-30
);
END;
/
BEGIN
DBMS_AUDIT_MGMT.CREATE_PURGE_JOB(
audit_trail_type            =>  DBMS_AUDIT_MGMT.AUDIT_TRAIL_UNIFIED,
audit_trail_purge_interval  =>  24 /* hours */,
audit_trail_purge_name      =>  'SCHEDULED_AUDIT_PURGE',
use_last_arch_timestamp     =>  TRUE
);
END;
/
```
</details>
