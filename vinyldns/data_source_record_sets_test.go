package vinyldns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSRecordSetsDataSource_basic(t *testing.T) {
	zoneName := testZoneName()
	groupName := "terraformdatasourcezonegroup"
	recordSetName := "terraformdatasourcerecordset"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			if err := testAccVinylDNSRecordSetsDataSourcePreCheck(t, groupName, zoneName, recordSetName); err != nil {
				t.Fatalf("precheck failed: %s", err)
			}
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckVinylDNSRecordSetsDataSourceConfig(zoneName, recordSetName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.vinyldns_record_sets.test", "record_sets.0.name", recordSetName),
					resource.TestCheckResourceAttr("data.vinyldns_record_sets.test", "record_sets.0.type", "A"),
				),
			},
		},
	})
}

func testAccVinylDNSRecordSetsDataSourcePreCheck(t *testing.T, groupName, zoneName, recordSetName string) error {
	client := vinyldns.NewClientFromEnv()
	zone, err := client.ZoneByName(zoneName)
	if err != nil || zone.ID == "" {
		group, gErr := ensureTestGroup(client, groupName)
		if gErr != nil {
			return gErr
		}

		zonePtr, zErr := ensureTestZone(client, zoneName, group.ID)
		if zErr != nil {
			return zErr
		}
		zone = *zonePtr
	}

	if err := ensureTestRecordSet(client, zone.ID, recordSetName); err != nil {
		return err
	}

	return nil
}

func testAccCheckVinylDNSRecordSetsDataSourceConfig(zoneName, recordSetName string) string {
	return fmt.Sprintf(`
data "vinyldns_zone" "test" {
	name = "%s"
}

data "vinyldns_record_sets" "test" {
	zone_id = "${data.vinyldns_zone.test.id}"
	name_filter = "%s"
}
`, zoneName, recordSetName)
}
