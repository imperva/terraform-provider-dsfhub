---
subcategory: "Agentless Gateway Permissions"
layout: "dsfhub"
page_title: "AWS IAM - Log Groups"
description: |-
  AWS IAM permissions for the DSF Agentless Gateway to access logs via Cloud Watch Log Groups.
---

# DSF Agentless Gateway Required IAM Permissions - Log Group

The DSF Agentless Gateway requires the following IAM permissions to access an [AWS Cloud Watch Log Group](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html).

## Variable Reference

```
# Variables for the DSF Agentless-Gateway IAM permissions granting access to cloudwatch logs
variable "agentless_gatway_iam_role_name" {
	description = "Name of the DSF agentless gateway role to grant permissions to access db logs via cloudwatch."
	type = string
	default =  "your-iam-gw-role-name"
}

variable "db_cloud_watch_log_group_name" {
	description = "Name of the Cloudwatch log group. Example: /aws/rds/instance/your-db-identifier-here/audit"
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
          "${data.aws_cloudwatch_log_group.rds_log_group.arn}:*",
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

## Troubleshooting

To validate that the DSF agentless gateway has the appropriate permissions to access the desired cloudwatch log group, you can [SSH](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/connect-linux-inst-ssh.html) to the gateway and run the following commands in a terminal:

```console
export $(cat /etc/sysconfig/jsonar)
$JSONAR_BASEDIR/bin/aws logs describe-log-streams --log-group-name "/aws/rds/instance/customappmysqlprod/audit" --region=us-east-2
```