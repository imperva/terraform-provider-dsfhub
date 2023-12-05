package dsfhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"regexp"
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

func dataSourceDSFDataSources() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDSFDataSourcesRead,
		Description: "Provides a list of DSFDataSources filtering for asset_id values by regex.",

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

func dataSourceDSFDataSourcesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	assetIdRegex := d.Get("asset_id_regex").(string)
	log.Printf("[INFO] Data Source - Reading DSFDataSources filtering for asset_ids with assetIdRegex: %s", assetIdRegex)

	cloudAccountsReadResponse, err := client.ReadDSFDataSources()
	if cloudAccountsReadResponse != nil {
		log.Printf("[INFO] Data Source - Reading DSFDataSources filtering for asset_ids with assetIdRegex: %s | err: %s\n", assetIdRegex, err)
	}
	var assetIds []string
	for _, ca := range cloudAccountsReadResponse.Data {
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
	d.SetId(client.config.DSFHUBHost + "_dsfDataSources")

	log.Printf("[INFO] Finished reading Data Source DSFDataSources with assetIdRegex: %s\n", assetIdRegex)
	return nil
}
