package dsfhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceDSFDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDSFDataSourceRead,
		Description: "Provides DSFDataSource from a unique asset_id.",

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

func dataSourceDSFDataSourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	curDSFDataSourceId := d.Get("asset_id").(string)
	log.Printf("[INFO] DataSource - Reading DSFDataSource with dsfDataSourceId: %s", curDSFDataSourceId)

	dsfDataSourceReadResponse, err := client.ReadDSFDataSource(curDSFDataSourceId)
	if dsfDataSourceReadResponse != nil {
		log.Printf("[INFO] Reading DSFDataSource with dsfDataSourceId: %s | err: %s\n", curDSFDataSourceId, err)
	}
	dsfDataSourceId := dsfDataSourceReadResponse.Data.AssetData.AssetID
	d.Set("asset_id", dsfDataSourceId)
	d.SetId(dsfDataSourceId)

	log.Printf("[INFO] Finished reading DataSource DSFDataSource with dsfDataSourceId: %s\n", dsfDataSourceId)
	return nil
}
