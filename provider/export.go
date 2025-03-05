package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/provider"
)

type ClientBuilder clients.ClientBuilder

func AzureProvider() *schema.Provider {
	return provider.AzureProvider()
}
