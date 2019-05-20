/*
Copyright 2018 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vinyldns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func TestAccVinylDNSRecordSetBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSRecordSetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVinylDNSRecordSetConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSRecordSetExists("vinyldns_record_set.test_a_record_set"),
					testAccCheckVinylDNSRecordSetExists("vinyldns_record_set.test_cname_record_set"),
					testAccCheckVinylDNSRecordSetExists("vinyldns_record_set.test_txt_record_set"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_a_record_set", "name", "terraformtestrecordset"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_cname_record_set", "name", "cname-terraformtestrecordset"),
				),
			},
			resource.TestStep{
				ResourceName:      "vinyldns_record_set.test_a_record_set",
				ImportState:       true,
				ImportStateVerify: true,
			},
			resource.TestStep{
				ResourceName:      "vinyldns_record_set.test_cname_record_set",
				ImportState:       true,
				ImportStateVerify: true,
			},
			resource.TestStep{
				ResourceName:      "vinyldns_record_set.test_txt_record_set",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVinylDNSRecordSetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*vinyldns.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vinyldns_record_set" {
			continue
		}
		id := rs.Primary.ID
		testZId, err := testZoneID()
		if err != nil {
			return fmt.Errorf("Error fetching system-test. zone ID")
		}

		// Try to find the record set
		_, err = client.RecordSet(testZId, id)
		if err == nil {
			return fmt.Errorf("RecordSet %s still exists", id)
		}
	}

	return nil
}

func testAccCheckVinylDNSRecordSetExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", rs)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RecordSet ID is set")
		}

		client := testAccProvider.Meta().(*vinyldns.Client)
		testZId, err := testZoneID()
		if err != nil {
			return fmt.Errorf("Error fetching system-test. zone ID")
		}
		if testZId == "" {
			return fmt.Errorf("Could not find system-test. zone ID")
		}

		readRs, err := client.RecordSet(testZId, rs.Primary.ID)
		if err != nil {
			return err
		}

		if readRs.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

func testZoneID() (string, error) {
	client := testAccProvider.Meta().(*vinyldns.Client)
	zones, err := client.ZonesListAll(vinyldns.ListFilter{})
	if err != nil {
		return "", err
	}

	for _, each := range zones {
		fmt.Println(each)
		if each.Name == "system-test." {
			return each.ID, nil
		}
	}

	return "", nil
}

const testAccVinylDNSRecordSetConfigBasic = `
resource "vinyldns_group" "test_group" {
	name = "terraformtestgroup"
	description = "some description"
	email = "tftest@tf.com"
	member {
	  id = "ok"
	}
	admin {
	  id = "ok"
	}
}

resource "vinyldns_zone" "test_zone" {
	name = "system-test."
	email = "foo@bar.com"
	admin_group_id = "${vinyldns_group.test_group.id}"
	depends_on = [
		"vinyldns_group.test_group"
	]
}

resource "vinyldns_record_set" "test_a_record_set" {
	name = "terraformtestrecordset"
	zone_id = "${vinyldns_zone.test_zone.id}"
	owner_group_id = "${vinyldns_group.test_group.id}"
	type = "A"
	ttl = 6000
	record_addresses = ["127.0.0.1", "127.0.0.1"]
	depends_on = [
		"vinyldns_zone.test_zone"
	]
}

resource "vinyldns_record_set" "test_cname_record_set" {
	name = "cname-terraformtestrecordset"
	zone_id = "${vinyldns_zone.test_zone.id}"
	type = "CNAME"
	ttl = 6000
	record_cname = "terraformtestrecordset.system-test."
	depends_on = [
		"vinyldns_record_set.test_a_record_set"
	]
}

resource "vinyldns_record_set" "test_txt_record_set" {
	name = "txt-terraformtestrecordset"
	zone_id = "${vinyldns_zone.test_zone.id}"
	type = "TXT"
	ttl = 6000
	record_texts = ["Lorem ipsum and all that jazz"]
	depends_on = [
		"vinyldns_zone.test_zone"
	]
}

resource "vinyldns_record_set" "test_nsd_record_set" {
	name = "nsd-terraformtestrecordset"
	zone_id = "${vinyldns_zone.test_zone.id}"
	type = "NS"
	ttl = 6000
	record_nsdnames = ["ns1.parent.com."]
	depends_on = [
		"vinyldns_zone.test_zone"
	]
}`
