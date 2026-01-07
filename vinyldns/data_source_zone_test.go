package vinyldns

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSZoneDataSource_basic(t *testing.T) {
	name := testZoneName()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			if err := testAccVinylDNSZoneDataSourcePreCheck(t, name); err != nil {
				t.Fatalf("precheck failed: %s", err)
			}
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVinylDNSZoneDataSourceConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.vinyldns_zone.test", "name"),
					resource.TestCheckResourceAttrSet("data.vinyldns_zone.test", "admin_group_id"),
					resource.TestCheckResourceAttr("data.vinyldns_zone.test", "name", name),
					resource.TestCheckResourceAttr("data.vinyldns_zone.test", "email", "foo@email.com"),
				),
			},
		},
	})
}

// ensure the zone exists in the target environment
func testAccVinylDNSZoneDataSourcePreCheck(t *testing.T, name string) error {
	client := vinyldns.NewClientFromEnv()
	group, err := ensureTestGroup(client, "terraformdatasourcetestgroup")
	if err != nil {
		log.Printf("[INFO] Error creating VinylDNS group %s", err)
		return err
	}

	_, err = ensureTestZone(client, name, group.ID)
	if err != nil {
		log.Printf("[INFO] Error ensuring VinylDNS zone %s", err)
		return err
	}

	return nil
}

func testAccCheckVinylDNSZoneDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "vinyldns_zone" "test" {
	name = "%s"
}
`, name)
}
