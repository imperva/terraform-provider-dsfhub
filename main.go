package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/imperva/terraform-provider-dsfhub/dsfhub"
)

// Using a debugger: https://developer.hashicorp.com/terraform/plugin/sdkv2/debugging
func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		Debug:        debug,
		ProviderFunc: dsfhub.Provider})
}
