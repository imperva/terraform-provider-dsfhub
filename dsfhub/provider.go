package dsfhub

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var baseAPIPrefix string
var descriptions map[string]string

func init() {
	baseAPIPrefix = "/dsf/api/v2"
	descriptions = map[string]string{
		"dsfhub_token": "The API token for API operations. You can retrieve this\n" +
			"from the DSF management hub console. Can be set via DSF_TOKEN " +
			"environment variable.",

		"dsfhub_host": "The DSF host endpoint for API operations.\n" +
			"Example: 'https://1.2.3.4:8443'. Can be set via DSF_HOST " +
			"environment variable.",

		"insecure_ssl": "The boolean flag that instructs the provider to allow for insecure SSL API calls to the DSF Hub, or support for self-signed certificates.\n" +
			"Example: 'true/false" +
			"environment variable.",
	}
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		DSFHUBToken: d.Get("dsfhub_token").(string),
		DSFHUBHost:  d.Get("dsfhub_host").(string),
		InsecureSSL: d.Get("insecure_ssl").(bool),
	}

	return config.Client()
}

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"dsf_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DSF_TOKEN", ""),
				Description: descriptions["dsf_token"],
			},
			"dsf_host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DSF_HOST", ""),
				Description: descriptions["dsf_host"],
			},
			"insecure_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INSECURE_SSL", true),
				Description: descriptions["insecure_ssl"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"dsfhub_cloud_account":  dataSourceCloudAccount(),
			"dsfhub_secret_manager": dataSourceSecretManager(),
			"dsfhub_log_aggregator": dataSourceLogAggregator(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"dsfhub_data_source":    resourceDSFDataSource(),
			"dsfhub_secret_manager": resourceSecretManager(),
			"dsfhub_cloud_account":  resourceCloudAccount(),
			"dsfhub_log_aggregator": resourceLogAggregator(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return provider
}
