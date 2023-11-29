---
subcategory: "Example Assets"
layout: "dsfhub"
page_title: "GCP MYSQL - PubSub"
description: |-
  Provides a combined example of creating an MYSQL database on GCP, associated configurations for audit logs in pub sub, onboarding to the DSFHUB with necessary configs for the DSF Agentless Gateway to access.
---

# GCP MYSQL Onboarding Template

Provides a combined example of creating an MYSQL database on GCP, the associated pubsub topic and sink, publisher role and iam permission binding, creating the [dsfhub_data_source](../r/data_source.md) and [dsfhub_log_aggregator](../r/log_aggregator.md) records to onboard to the DSFHUB with necessary access for the DSF Agentless Gateway.

## Before you begin

1. [Configure the GOOGLE Provider](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/getting_started) to get started. This will walk through create a project in the Google Cloud Console and set up billing on that project. Any examples in this guide will be part of the GCP "always free" tier.  
2. [Enable data access audit logs](https://cloud.google.com/logging/docs/audit/configure-data-access) in your Google Cloud projects, billing accounts, folders, and organizations by using the Google Cloud console or the API.
3. [Create pubSub topic, project sink, and iam permission binding](gcp_pubsub_iam.md) to grant the [DSF Agentless Gateway](https://registry.terraform.io/modules/imperva/dsf-agentless-gw/aws/latest) access to the desired pubsub topic.

## Example Usage

<details>
<summary>GCP MYSQL Variables</summary>

### GCP MYSQL Variables

```hcl
# DSFHUB Provider Required Variables
variable "dsfhub_token" {} # TF_VAR_dsfhub_token env variable
variable "dsfhub_host" {} # TF_VAR_dsfhub_host env variable

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

variable "key_file" {
  description = "Location on disk for the key to be used to by the DSF Agentless Gateway to authenticate"
  type = string
  default = "/tmp/keyfile"
}

# GCP Variables
variable "db_authorized_ip" {
  description = "List of whitelisted IPs to access the database"
  type = list(object({
    name = string
    ip = string
  }))
  default = [{
    name = "local"
    ip = "127.0.0.1"
  }]
}

variable "db_name" {
  description =  " The name of the instance. If the name is left blank, Terraform will randomly generate one when the instance is first created. This is done because after a name is used, it cannot be reused for up to one week."
  type = string
  default = "mysql-gcp"
}

variable "db_password" {
  description =  "The password for the user. Can be updated. For Postgres instances this is a Required field, unless type is set to either CLOUD_IAM_USER or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance."
  type = string
  default = "mypassword"
}

variable "db_user" {
  description =  " The name of the user"
  type = string
  default = "myusername"
}

variable "db_version" {
  description =  " The MySQL, PostgreSQL or SQL Server version to use. Supported values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, POSTGRES_9_6,POSTGRES_10, POSTGRES_11, POSTGRES_12, POSTGRES_13, POSTGRES_14, POSTGRES_15, SQLSERVER_2017_STANDARD, SQLSERVER_2017_ENTERPRISE, SQLSERVER_2017_EXPRESS, SQLSERVER_2017_WEB. SQLSERVER_2019_STANDARD, SQLSERVER_2019_ENTERPRISE, SQLSERVER_2019_EXPRESS, SQLSERVER_2019_WEB"
  type = string
  default = "MYSQL_8_0"
}

variable "project" {
  description =  " The project field should be your personal project id. The project indicates the default GCP project all of your resources will be created in. Most Terraform resources will have a project field."
  type = string
  default = "My_project"
}

variable "region" {
  description =  " The region will be used to choose the default location for regional resources. Regional resources are spread across several zones."
  type = string
  default = "us-east1"
}

variable "tier" {
  description =  "The machine type to use. See tiers for more details and supported versions. Postgres supports only shared-core machine types, and custom machine types such as db-custom-2-13312."
  type = string
  default = "db-f1-micro"
}
```
</details>

### Providers and Resources

```hcl
# Specify path for provider
terraform {
  required_providers {
    dsfhub = {
      source = "imperva/dsfhub"
    }
  }
}

### DSF Provider ###
provider "dsfhub" {
  dsfhub_token = var.dsfhub_token
  dsfhub_host = var.dsfhub_host
}

### Google Provider ###
provider "google" {
  project = var.project
  region  = var.region
}

# Create MYSQL database instance
resource "google_sql_database_instance" "mysql_db" {
  name             = var.db_name
  database_version = var.db_version
  region           = var.region

  settings {
    tier = var.tier
    database_flags {
      name = "general_log"
      value = "On"
    }

    database_flags {
      name = "log_output"
      value = "FILE"
    }

    ip_configuration {
      dynamic "authorized_networks" {
        for_each = var.db_authorized_ip
        content {
          name = authorized_networks.value.name
          value = authorized_networks.value.ip
        }


      }
    }

  }
}

# Create MYSQL user
resource "google_sql_user" "users" {
  name     = var.db_user
  instance = google_sql_database_instance.mysql_db.name
  password = var.db_password
}

# ### Resource example for GCP MYSQL ###
data "google_pubsub_topic" "mysql_pubsub_topic_data" {
  name = "${var.db_name}-pubsub-topic"
}

resource "google_pubsub_subscription" "mysql_pubsub_subscription" {
  name  = "${var.db_name}-pubsub_subscription"
  topic = data.google_pubsub_topic.mysql_pubsub_topic_data.name
}

resource "dsfhub_data_source" "gcp_mysql" {
  server_type = "GCP MYSQL"
  admin_email = var.admin_email
  asset_display_name = var.db_name
  asset_id = google_sql_database_instance.mysql_db.connection_name
  gateway_id = var.gateway_id
  server_host_name = google_sql_database_instance.mysql_db.public_ip_address
  server_ip = google_sql_database_instance.mysql_db.public_ip_address
  logs_destination_asset_id = google_pubsub_subscription.mysql_pubsub_subscription.id
  audit_pull_enabled = true
  asset_connection {
    auth_mechanism = "password"
    password = var.db_password
    reason = "default"
    username = var.db_user
  }
}

# ### Resource example for GCP PUBSUB ###
resource "dsfhub_log_aggregator" "gcp_mysql_pubsub" {
  server_type = "GCP PUBSUB"
  admin_email = var.admin_email
  asset_display_name = "${var.db_name}-pubsub_subscription"
  asset_id = google_pubsub_subscription.mysql_pubsub_subscription.id
  gateway_id = var.gateway_id
  pubsub_subscription = google_pubsub_subscription.mysql_pubsub_subscription.id
  server_host_name = "pubsub.googleapis.com"
  server_ip = "pubsub.googleapis.com"
  server_port = 443
  audit_type =  "MYSQL"
  asset_connection {
    auth_mechanism = "service_account"
    key_file = var.key_file
    reason = "default"
  }
}
```

## Agentless Gateway Permission Dependencies:

The [DSF Agentless Gateway](https://registry.terraform.io/modules/imperva/dsf-agentless-gw/aws/latest) is required to have the following [IAM policy for Cloud Pub/Sub Topic](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/pubsub_topic_iam) access to database audit.

<ul>
<li><a target="_blank" href="gcp_pubsub_iam.md">GCP IAM Permissions for PubSub Topics</a></li>
</ul>