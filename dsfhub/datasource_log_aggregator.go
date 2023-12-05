package dsfhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"regexp"
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

func dataSourceLogAggregators() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLogAggregatorsRead,
		Description: "Provides a list of LogAggregator filtering for asset_id values by regex.",

		Schema: map[string]*schema.Schema{
			// Computed Attributes
			"asset_id_regex": {
				Type:        schema.TypeString,
				Description: "Regex pattern for asset IDs",
				Optional:    true,
				Default:     nil,
			},
			"asset_ids": {
				Type:        schema.TypeList,
				Description: "List of asset IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: false,
				Computed: true,
				Default:  nil,
			},
		},
	}
}

func dataSourceLogAggregatorsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	assetIdRegex := d.Get("asset_id_regex").(string)
	log.Printf("[INFO] Data Source - Reading LogAggregator filtering for asset_ids with assetIdRegex: %s", assetIdRegex)

	logAggregatorsReadResponse, err := client.ReadLogAggregators()
	if logAggregatorsReadResponse != nil {
		log.Printf("[INFO] Data Source - Reading LogAggregators filtering for asset_ids with assetIdRegex: %s | err: %s\n", assetIdRegex, err)
	}
	var assetIds []string
	for _, ca := range logAggregatorsReadResponse.Data {
		match, _ := regexp.MatchString(assetIdRegex, ca.ID)
		log.Printf("[INFO] Checking asset_id: %v against regex: %v match:%v\n", ca.ID, assetIdRegex, match)
		if match {
			log.Printf("[INFO] Matched asset_id: %v against regex: %v match:%v\n", ca.ID, assetIdRegex, match)
			assetIds = append(assetIds, ca.ID)
		} else {
			log.Printf("[INFO] Did not match asset_id: %v against regex: %v match:%v\n", ca.ID, assetIdRegex, match)
		}
	}

	d.Set("asset_ids", assetIds)
	d.SetId(client.config.DSFHUBHost + "_logAggregators")

	log.Printf("[INFO] Finished reading Data Source logAggregators with assetIdRegex: %s\n", assetIdRegex)
	return nil
}
