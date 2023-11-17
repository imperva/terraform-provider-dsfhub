---
subcategory: "AWS IAM Permissions"
layout: "dsfhub"
page_title: "S3 Bucket"
description: |-
  AWS IAM permissions for the DSF Agentless Gateway to access logs via AWS S3 Buckets.
---

# DSF Agentless Gateway Required IAM Permissions - S3 Bucket

The DSF Agentless Gateway requires the following IAM permissions to access [AWS S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/GetStartedWithS3.html).

## Variable Reference

```
# Variables for the DSF Agentless-Gateway IAM permissions granting access to s3 buckets
variable "agentless_gatway_iam_role_name" {
	description = "Name of the DSF agentless gateway role to grant permissions to access db logs via s3 buckets."
	type = string
	default =  "your-iam-gw-role-name"
}

variable "db_s3_bucket_name" {
	description = "Name of the s3 bucket. Example: /aws/rds/instance/your-db-identifier-here/audit"
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

data "aws_s3_bucket" "db_audit_bucket" {
  bucket = var.db_s3_bucket_name
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
          "s3:ListAllMyBuckets", 
          "s3:GetBucketLocation",
          "s3:ListBucket",
          "s3:GetObject"
        ]
        "Resource": [
          "${data.aws_s3_bucket.db_audit_bucket.arn}/*",
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "role_policy_attachment" {
  policy_arn = aws_iam_policy.log_group_policy.arn
  role       = var.agentless_gatway_iam_role_name
}
```