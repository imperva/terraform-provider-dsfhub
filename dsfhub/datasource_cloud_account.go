package dsfhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceCloudAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudAccountRead,
		Description: "Provides CloudAccount from a unique asset_id.",

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

func dataSourceCloudAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	curCloudAccountId := d.Get("asset_id").(string)
	log.Printf("[INFO] DataSource - Reading CloudAccount with cloudAccountId: %s", curCloudAccountId)

	cloudAccountReadResponse, err := client.ReadCloudAccount(curCloudAccountId)
	if cloudAccountReadResponse != nil {
		log.Printf("[INFO] Reading CloudAcount with cloudAccountId: %s | err: %s\n", curCloudAccountId, err)
	}
	cloudAccountId := cloudAccountReadResponse.Data.AssetData.AssetID
	d.Set("asset_id", cloudAccountId)
	d.SetId(cloudAccountId)

	log.Printf("[INFO] Finished reading DataSource CloudAccount with cloudAccountId: %s\n", cloudAccountId)
	return nil
}
