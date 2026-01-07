package vinyldns

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccVinylDNSBackendIDsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVinylDNSBackendIDsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.vinyldns_backend_ids.test", "backend_ids.0"),
				),
			},
		},
	})
}

const testAccCheckVinylDNSBackendIDsDataSourceConfig = `
data "vinyldns_backend_ids" "test" {}
`
