package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/imperva/terraform-provider-dsfhub/dsfhub"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dsfhub.Provider})
}
