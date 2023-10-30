Terraform `DSFHUB` Provider
=========================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by the team at [Imperva](https://www.imperva.com/).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-dsf`

```sh
git clone git@github.com:imperva/terraform-provider-dsfhub.git $GOPATH/src/github.com/terraform-providers/terraform-provider-dsfhub
```

Enter the provider directory and build the provider

```sh
cd $GOPATH/src/github.com/terraform-providers/terraform-provider-dsfhub
make build
```

Using the Provider
---------------------------
Set the following 2 environment variables: 
```
export TF_VAR_dsfhub_host="https://1.2.3.4:8443"
export TF_VAR_dsfhub_token="12345-654-4fda-9fc0-somekeyhere"
```
Sample terraform files can be found in the examples folder in this repository.  You can initialize and run these terraform files with the following commands:
```
terraform init
terraform plan
terraform apply --auto-approve
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
make install
...
$GOPATH/bin/terraform-provider-dsfhub
...
```

In order to test the provider, you can simply run `make test`.

```sh
make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
make testacc
```