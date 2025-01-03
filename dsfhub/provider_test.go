package dsfhub

import (
	"context"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testAccProviderConfigure sync.Once

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"dsfhub": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[DEBUG] Running test TestProvider")
	if err := Provider().InternalValidate(); err != nil {
		log.Printf("[INFO] err: %s \n", err)
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test testAccPreCheck \n")
	testAccProviderConfigure.Do(func() {
		if v := os.Getenv("DSFHUB_TOKEN"); v == "" {
			t.Fatal("DSFHUB_TOKEN must be set for acceptance tests")
		}

		if v := os.Getenv("DSFHUB_HOST"); v == "" {
			t.Fatal("DSFHUB_HOST must be set for acceptance tests")
		}

		err := testAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
		if err != nil {
			t.Fatal(err)
		}
	})
}
