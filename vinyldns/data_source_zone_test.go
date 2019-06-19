package vinyldns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSZoneDataSource_basic(t *testing.T) {
	name := "test-zone"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccVinylDNSZoneDataSourcePreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVinylDNSZoneDataSourceConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.vinyldns_zone.test", "name"),
					resource.TestCheckResourceAttr("data.vinyldns_zone.test", "name", name),
					resource.TestCheckResourceAttr("data.vinyldns_zone.test", "email", "foo@email.com"),
				),
			},
		},
	})
}

// create a group and zone such that the zone can be imported
func testAccVinylDNSZoneDataSourcePreCheck(t *testing.T) error {
	client := vinyldns.NewClientFromEnv()
	g, err := client.GroupCreate(&vinyldns.Group{
		Name:  "terraformdatasourcetestgroup",
		Email: "foo@email.com",
	})
	if err != nil {
		return err
	}

	_, err = client.ZoneCreate(&vinyldns.Zone{
		Name:         "test-zone",
		Email:        "foo@email.com",
		AdminGroupID: g.ID,
	})

	if err != nil {
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
