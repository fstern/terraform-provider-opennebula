package main

import (
	"github.com/hashicorp/terraform/plugin"
  "github.com/ImmoweltGroup/terraform-provider-opennebula/opennebula"
  //"./opennebula"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opennebula.Provider,
	})
}
