package vinyldns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSGroupsDataSource_basic(t *testing.T) {
	groupName := "terraformdatasourcegroup"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			if err := testAccVinylDNSGroupsDataSourcePreCheck(t, groupName); err != nil {
				t.Fatalf("precheck failed: %s", err)
			}
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVinylDNSGroupsDataSourceConfig(groupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.vinyldns_groups.test", "groups.0.id"),
					resource.TestCheckResourceAttr("data.vinyldns_groups.test", "groups.0.name", groupName),
				),
			},
		},
	})
}

func testAccVinylDNSGroupsDataSourcePreCheck(t *testing.T, name string) error {
	client := vinyldns.NewClientFromEnv()
	_, err := ensureTestGroup(client, name)
	return err
}

func testAccCheckVinylDNSGroupsDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "vinyldns_groups" "test" {
	name_filter = "%s"
}
`, name)
}
