package vinyldns

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSGroupDataSource_basic(t *testing.T) {
	id := "fdc7c6ed-44dc-4c75-9ec1-bb96e73489d5"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccVinylDNSGroupDataSourcePreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVinylDNSGroupDataSourceConfig(id),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.vinyldns_group.test", "name"),
					resource.TestCheckResourceAttrSet("data.vinyldns_group.test", "email"),
					resource.TestCheckResourceAttrSet("data.vinyldns_group.test", "description"),
					resource.TestCheckResourceAttr("data.vinyldns_group.test", "id", id),
				),
			},
		},
	})
}

// create a group and zone such that the zone can be imported
func testAccVinylDNSGroupDataSourcePreCheck(t *testing.T) error {
	client := vinyldns.NewClientFromEnv()
	g, err := client.GroupCreate(&vinyldns.Group{
		Name:  "terraform-provider-group_datasource_test",
		Email: "foo@bar.com",
		Members: []vinyldns.User{
			vinyldns.User{
				ID: "6e07183f-a68e-42cf-acce-044296ede753",
			}},
		Admins: []vinyldns.User{
			vinyldns.User{
				ID: "6e07183f-a68e-42cf-acce-044296ede753",
			}},
	})
	if err != nil {
		log.Printf("[INFO] Error creating VinylDNS group %s", err)
		return err
	}
	log.Printf("[INFO] Created VinylDNS group %s", g.Name)

	return nil
}

func testAccCheckVinylDNSGroupDataSourceConfig(id string) string {
	return fmt.Sprintf(`
data "vinyldns_group" "test" {
	id = "%s"
}
`, id)
}
