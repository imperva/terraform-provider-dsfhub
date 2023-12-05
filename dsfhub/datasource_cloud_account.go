package dsfhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"regexp"
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
	log.Printf("[INFO] Data Source - Reading CloudAccount with cloudAccountId: %s", curCloudAccountId)

	cloudAccountReadResponse, err := client.ReadCloudAccount(curCloudAccountId)
	if cloudAccountReadResponse != nil {
		log.Printf("[INFO] Reading CloudAcount with cloudAccountId: %s | err: %s\n", curCloudAccountId, err)
	}
	cloudAccountId := cloudAccountReadResponse.Data.AssetData.AssetID
	d.Set("asset_id", cloudAccountId)
	d.SetId(cloudAccountId)

	log.Printf("[INFO] Finished reading Data Source CloudAccount with cloudAccountId: %s\n", cloudAccountId)
	return nil
}

func dataSourceCloudAccounts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudAccountsRead,
		Description: "Provides a list of CloudAccounts filtering for asset_id values by regex.",

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

func dataSourceCloudAccountsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	assetIdRegex := d.Get("asset_id_regex").(string)
	log.Printf("[INFO] Data Source - Reading CloudAcounts filtering for asset_ids with assetIdRegex: %s", assetIdRegex)

	cloudAccountsReadResponse, err := client.ReadCloudAccounts()
	if cloudAccountsReadResponse != nil {
		log.Printf("[INFO] Data Source - Reading CloudAcounts filtering for asset_ids with assetIdRegex: %s | err: %s\n", assetIdRegex, err)
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
	d.SetId(client.config.DSFHUBHost + "_cloudAccounts")

	log.Printf("[INFO] Finished reading Data Source CloudAccounts with assetIdRegex: %s\n", assetIdRegex)
	return nil
}
