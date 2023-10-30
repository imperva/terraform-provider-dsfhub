package dsfhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceLogAggregator() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLogAggregatorRead,
		Description: "Provides LogAggregator from a unique asset_id.",

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

func dataSourceLogAggregatorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	curLogAggregatorId := d.Get("asset_id").(string)
	log.Printf("[INFO] DataSource - Reading LogAggregator with secretManagerId: %s", curLogAggregatorId)

	logAggregatorReadResponse, err := client.ReadLogAggregator(curLogAggregatorId)
	if logAggregatorReadResponse != nil {
		log.Printf("[INFO] Reading LogAggregator with logAggregatorId: %s | err: %s\n", curLogAggregatorId, err)
	}
	logAggregatorId := logAggregatorReadResponse.Data.AssetData.AssetID
	d.Set("asset_id", logAggregatorId)
	d.SetId(logAggregatorId)

	log.Printf("[INFO] Finished reading DataSource LogAggregator with logAggregatorId: %s\n", logAggregatorId)
	return nil
}
