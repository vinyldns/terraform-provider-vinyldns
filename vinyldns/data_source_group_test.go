package vinyldns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSGroupDataSource_basic(t *testing.T) {
	name := "terraformdatasourcegroup"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			if err := testAccVinylDNSGroupDataSourcePreCheck(t, name); err != nil {
				t.Fatalf("precheck failed: %s", err)
			}
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVinylDNSGroupDataSourceConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.vinyldns_group.test", "name"),
					resource.TestCheckResourceAttrSet("data.vinyldns_group.test", "email"),
					resource.TestCheckResourceAttr("data.vinyldns_group.test", "name", name),
				),
			},
		},
	})
}

func testAccVinylDNSGroupDataSourcePreCheck(t *testing.T, name string) error {
	client := vinyldns.NewClientFromEnv()
	_, err := ensureTestGroup(client, name)
	if err != nil {
		return err
	}

	return nil
}

func testAccCheckVinylDNSGroupDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "vinyldns_group" "test" {
	name = "%s"
}
`, name)
}
