package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-guardians/guardians"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return guardians.Provider()
		},
	})
}
