package vinyldns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccVinylDNSGroupDataSource_basic(t *testing.T) {
	id := "global-acl-group-id"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
        Config: testAccCheckVinylDNSGroupDataSourceConfig(id),
				Check: resource.ComposeAggregateTestCheckFunc(
          resource.TestCheckResourceAttrSet("data.vinyldns_group.test", "id"),
					resource.TestCheckResourceAttr("data.vinyldns_group.test", "id", id),
					resource.TestCheckResourceAttr("data.vinyldns_group.test", "name", "globalACLGroup"),
					resource.TestCheckResourceAttr("data.vinyldns_group.test", "email", "email"),
				),
			},
		},
	})
}

func testAccCheckVinylDNSGroupDataSourceConfig(id string) string {
	return fmt.Sprintf(`
data "vinyldns_group" "test" {
	id = "%s"
}
`, id)
}
