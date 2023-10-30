package dsfhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceSecretManager() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSecretManagerRead,
		Description: "Provides SecretManager from a unique asset_id.",

		Schema: map[string]*schema.Schema{
			// Computed Attributes
			"asset_id": {
				Type:        schema.TypeString,
				Description: "Current asset ID",
				Required:    true,
				Optional:    false,
			},
		},
	}
}

func dataSourceSecretManagerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	curSecretManagerId := d.Get("asset_id").(string)
	log.Printf("[INFO] DataSource - Reading SecretManager with secretManagerId: %s", curSecretManagerId)

	secretManagerReadResponse, err := client.ReadSecretManager(curSecretManagerId)
	if secretManagerReadResponse != nil {
		log.Printf("[INFO] Reading SecretManager with secretManagerId: %s | err: %s\n", curSecretManagerId, err)
	}
	secretManagerId := secretManagerReadResponse.Data.AssetData.AssetID
	d.Set("asset_id", secretManagerId)
	d.SetId(secretManagerId)

	log.Printf("[INFO] Finished reading DataSource SecretManager with secretManagerId: %s\n", secretManagerId)
	return nil
}
