---
subcategory: "DSF Agentless Gateway Permissions"
layout: "dsfhub"
page_title: "DSF Agentless Gateway Required IAM Permissions - Kinesis"
description: |-
  AWS IAM permissions for the DSF Agentless Gateway to access logs via AWS Kinesis Data Streams.
---

# DSF Agentless Gateway Requied IAM Permissions - Kinesis

The DSF Agentless Gateway requires the following IAM permissions to access AWS Kinesis Data Streams.

## Variable Reference

```
# Variables for the DSF Agentless-Gateway IAM permissions granting access to cloudwatch logs
variable "agentless_gatway_iam_role_name" {
	description = "Name of the DSF agentless gateway role to grant permissions to access db logs via cloudwatch."
	type = string
	default =  "your-iam-gw-role-name"
}

variable "kinesis_data_stream_arn" {
	description = "ARN of the Kinesis Data Stream. Example: arn:aws:kinesis:us-east-2:1234567890:stream/your_receiver_name_here"
	type = string
	default =  null
}

variable "deployment_name" {
	description = "The name of the database deployment. i.e. 'custom-app-prod'"
	type = string
	default = null
}

variable "region" {
	description = "AWS region"
	type = string
	default = "us-east-2"
}
```

## Example Usage

```
provider "aws" {
  region = var.region
}

data "aws_iam_role" "agentless_gatway_role_name" {
  name = var.agentless_gatway_iam_role_name
}

resource "aws_iam_policy" "kinesis_policy" {
  name        = "DSFAgentlessGatewayKinesisPolicy-${var.deployment_name}"
  description = "DSF Agentless Gateway Kinesis Policy for ${var.deployment_name}"

  policy = jsonencode({
  "Version": "2012-10-17",
  "Statement": [
      {
        "Sid": "VisualEditor0",
        "Effect": "Allow",
        "Action": [
          "kinesis:GetShardIterator",
          "kinesis:GetRecords",
          "kinesis:DescribeStream",
          "kms:Decrypt"
        ]
        "Resource": [
          "${var.kinesis_data_stream_arn}",
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "role_policy_attachment" {
  policy_arn = aws_iam_policy.kinesis_policy.arn
  role       = aws_iam_role.assume_role_policy.name
}
```