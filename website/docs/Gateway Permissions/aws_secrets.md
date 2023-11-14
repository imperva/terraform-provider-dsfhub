---
subcategory: ""
layout: "imperva"
page_title: "DSF Agentless Gateway AWS IAM Kinesis Permissions"
description: |-
  AWS IAM permissions for the DSF Agentless Gateway to access logs via Kinesis. 
---

## Variable Reference

```
variable "aws_secretsmanager_access_tokens_arn" {
  description =  "The email address to notify about this asset"
  default = null 
}

variable "dsfhub_password_secret_aws_arn" {
  description =  "The arn of the secret for the dsfhub."
  default = null 
}

variable "db_password_secret_aws_arn" {
  description =  "The email address to notify about this asset"
  default = null 
}
```

## Example Usage

```
provider "aws" {
  region = var.region
}

data "aws_cloudwatch_log_group" "rds_log_group" {
    name = var.db_cloud_watch_log_group_name
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
          "${data.aws_cloudwatch_log_group.rds_log_group.arn}/*",
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "role_policy_attachment" {
  policy_arn = aws_iam_policy.log_group_policy.arn
  role       = var.agentless_gatway_iam_role_name
}
  
  inline_policy_secret = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
      {
        "Sid": "VisualEditor0",
        "Effect": "Allow",
        "Action": "secretsmanager:GetSecretValue",
        "Resource": concat([
            "${var.dsfhub_password_secret_aws_arn}",
            "${var.db_password_secret_aws_arn}"
          ],
          [
            var.aws_secretsmanager_access_tokens_arn
          ]
        )
      }
    ]
  }
  )
```