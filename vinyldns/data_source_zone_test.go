package vinyldns

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSZoneDataSource_basic(t *testing.T) {
	name := "ok."
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
		Members: []vinyldns.User{
			vinyldns.User{
				ID: "ok",
			}},
		Admins: []vinyldns.User{
			vinyldns.User{
				ID: "ok",
			}},
	})
	if err != nil {
		log.Printf("[INFO] Error creating VinylDNS group %s", err)
		return err
	}
	log.Printf("[INFO] Created VinylDNS group %s", g.Name)

	z, err := client.ZoneCreate(&vinyldns.Zone{
		Name:         "ok.",
		Email:        "foo@email.com",
		AdminGroupID: g.ID,
		Connection: &vinyldns.ZoneConnection{
			Name:          "vinyldns.",
			Key:           "nzisn+4G2ldMn0q1CV3vsg==",
			KeyName:       "vinyldns.",
			PrimaryServer: "vinyldns-bind9",
		},
	})
	if err != nil {
		log.Printf("[INFO] Error creating VinylDNS zone %s", err)
		return err
	}

	createdZoneID := z.Zone.ID
	limit := 10
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		zg, err := client.Zone(createdZoneID)
		if err == nil && zg.ID != createdZoneID {
			log.Printf("[INFO] unable to get VinylDNS zone %s", createdZoneID)
			return err
		}
		if err == nil && zg.ID == createdZoneID {
			break
		}

		if i == (limit - 1) {
			log.Printf("[INFO] %d retries reached in polling VinylDNS zone %s", limit, createdZoneID)
			return err
		}
	}

	log.Printf("[INFO] Created VinylDNS zone %s", z.Zone.Name)

	return nil
}

func testAccCheckVinylDNSZoneDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "vinyldns_zone" "test" {
	name = "%s"
}
`, name)
}
