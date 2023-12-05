package dsfhub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"regexp"
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

func dataSourceSecretManagers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSecretManagersRead,
		Description: "Provides a list of SecretManagers filtering for asset_id values by regex.",

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

func dataSourceSecretManagersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	assetIdRegex := d.Get("asset_id_regex").(string)
	log.Printf("[INFO] Data Source - Reading SecretManagers filtering for asset_ids with assetIdRegex: %s", assetIdRegex)

	secretManagersReadResponse, err := client.ReadSecretManagers()
	if secretManagersReadResponse != nil {
		log.Printf("[INFO] Data Source - Reading SecretManagers filtering for asset_ids with assetIdRegex: %s | err: %s\n", assetIdRegex, err)
	}
	var assetIds []string
	for _, ca := range secretManagersReadResponse.Data {
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
	d.SetId(client.config.DSFHUBHost + "_secretManagers")

	log.Printf("[INFO] Finished reading Data Source SecretManagers with assetIdRegex: %s\n", assetIdRegex)
	return nil
}
