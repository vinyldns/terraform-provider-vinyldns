package vinyldns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSZonesDataSource_basic(t *testing.T) {
	zoneName := testZoneName()
	groupName := "terraformdatasourcezonegroup"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			if err := testAccVinylDNSZonesDataSourcePreCheck(t, groupName, zoneName); err != nil {
				t.Fatalf("precheck failed: %s", err)
			}
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVinylDNSZonesDataSourceConfig(zoneName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.vinyldns_zones.test", "zones.0.name", zoneName),
					resource.TestCheckResourceAttrSet("data.vinyldns_zones.test", "zones.0.admin_group_id"),
				),
			},
		},
	})
}

func testAccVinylDNSZonesDataSourcePreCheck(t *testing.T, groupName, zoneName string) error {
	client := vinyldns.NewClientFromEnv()
	if _, err := client.ZoneByName(zoneName); err == nil {
		return nil
	}
	group, err := ensureTestGroup(client, groupName)
	if err != nil {
		return err
	}

	_, err = ensureTestZone(client, zoneName, group.ID)
	if err != nil {
		return err
	}

	return nil
}

func testAccCheckVinylDNSZonesDataSourceConfig(zoneName string) string {
	return fmt.Sprintf(`
data "vinyldns_zones" "test" {
	name_filter = "%s"
}
`, zoneName)
}
