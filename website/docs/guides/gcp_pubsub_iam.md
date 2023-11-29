---
subcategory: "Agentless Gateway Permissions"
layout: "dsfhub"
page_title: "GCP IAM - PubSub"
description: |-
  GCP IAM - PubSub
---

# Agentless Gateway IAM Permissions - PubSub Topic

The DSF Agentless Gateway requires the following IAM permissions to access a GCP PubSub Topic.

## Variable Reference

```hcl
# Variables for the creating GCP pubsub topic, project sink, and iam binding to grant the DSF Agentless-Gateway access database logs
variable "project" {
  description =  "The project field should be your personal project id. The project indicates the default GCP project all of your resources will be created in. Most Terraform resources will have a project field."
  type = string
  default = "Your_project_name_here"
}

variable "region" {
  description =  "The region will be used to choose the default location for regional resources. Regional resources are spread across several zones."
  type = string
  default = "us-east1"
}

variable "db_name" {
  description =  "The name of the instance. If the name is left blank, Terraform will randomly generate one when the instance is first created. This is done because after a name is used, it cannot be reused for up to one week."
  type = string
  default = "mydb-gcp"
}
```

## Example Usage

```hcl
### Google Provider ###
provider "google" {
  project = var.project
  region  = var.region
}

# Create Pubsub topic
resource "google_pubsub_topic" "mysql_pubsub_topic" {
  name = "${var.db_name}-pubsub-topic"
}

# Create Pubsub sink and attach a topic to it
resource "google_logging_project_sink" "mysql_pubsub_sink" {
    depends_on = [google_pubsub_topic.mysql_pubsub_topic]
    name = "${var.db_name}-pubsub-sink"
    destination = "pubsub.googleapis.com/projects/${var.project}/topics/${google_pubsub_topic.mysql_pubsub_topic.name}"
    filter = "resource.type = \"cloudsql_database\" resource.labels.database_id = \"${var.project}:${google_sql_database_instance.mysql_from_terraform.name}\" logName = \"projects/${var.project}/logs/cloudsql.googleapis.com%2Fmysql-general.log\""
}

# Attach pubsub publisher role to the topic
locals {
   role_name = regex("(service-\\d+@[a-z0-9\\-.]+)",google_logging_project_sink.mysql_pubsub_sink.writer_identity)
}

resource "google_pubsub_topic_iam_binding" "binding" {
  project = var.project
  topic = google_pubsub_topic.mysql_pubsub_topic.name
  role = "roles/pubsub.publisher"
  members = ["serviceAccount:${local.role_name[0]}"]
}
```

## Troubleshooting

To validate that the DSF agentless gateway has the appropriate permissions to access the desired pubsub topic in GCP, you can [SSH](https://cloud.google.com/compute/docs/instances/ssh) to the gateway and run the following commands in a terminal.

### Export the environment variables

```console
export $(cat /etc/sysconfig/jsonar)
```

### Validate access to pubsub topic

Use the local gcloud client and service-account to validate access and [describe the pubsub topic](https://cloud.google.com/sdk/gcloud/reference/pubsub/topics/describe)

```console
$JSONAR_BASEDIR/bin/gcloud auth activate-service-account --key-file=/path/to/your/key.json
$JSONAR_BASEDIR/bin/gcloud pubsub topics --project your-project-name-here describe your-pubsub-topic-here
```

### Describe subscription and pull single record

Use the local gcloud client to [describe the subscription](https://cloud.google.com/sdk/gcloud/reference/pubsub/subscriptions/describe), and check [pull a message](https://cloud.google.com/sdk/gcloud/reference/pubsub/subscriptions/pull):

```console
$JSONAR_BASEDIR/bin/gcloud pubsub subscriptions --project your-project-name-here describe your-subscription-name-here
$JSONAR_BASEDIR/bin/gcloud pubsub subscriptions --project your-project-name-here pull your-subscription-name-here --no-auto-ack
```