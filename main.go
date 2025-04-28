package main

import (
	"github.com/terraform-provider-cloudportal/cloudportal/internal/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// main function is the entry point of the provider plugin
func main() {

	// Use the plugin library to start the provider
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider, // Provider is the function you defined in provider/provider.go
	})

	// If there's an error, log it and exit

}
