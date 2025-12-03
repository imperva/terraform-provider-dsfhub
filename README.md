# Terraform Provider for `DSFHUB` 

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/Terraform_Logo.svg/768px-Terraform_Logo.svg.png" width="600px">

The DSFHUB provider is a plugin that allows [Terraform](https://www.terraform.io) to manage resources on DSFHub. This provider plugin is maintained by [Thales Group](https://www.thalesgroup.com/en).

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.22 (to build the provider plugin)

## Using the provider

Please see the [instructions](https://registry.terraform.io/providers/imperva/dsfhub/latest/docs) on how to configure the DSFHUB provider.

For examples, please see the [DSFHub Agentless-Onboarding modules](https://github.com/imperva/terraform-dsfhub-agentless-onboarding).

### Upgrading the provider

The DSFHUB provider doesn't upgrade automatically once you've started using it. After a new release you can run

```
terraform init -upgrade
```

to upgrade to the latest stable version of the provider. See the Terraform website for more information on provider upgrades, and how to set version constraints on your provider.
