package vinyldns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSRecordSetsDataSource_basic(t *testing.T) {
	zoneName := "terraformdatasourcezone."
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
	group, err := ensureTestGroup(client, groupName)
	if err != nil {
		return err
	}

	zone, err := ensureTestZone(client, zoneName, group.ID)
	if err != nil {
		return err
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
