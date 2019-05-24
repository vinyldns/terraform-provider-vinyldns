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
					resource.TestCheckResourceAttr("vinyldns_record_set.test_a_record_set", "type", "A"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_a_record_set", "ttl", "6000"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_cname_record_set", "name", "cname-terraformtestrecordset"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_cname_record_set", "type", "CNAME"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_cname_record_set", "ttl", "6000"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_txt_record_set", "name", "txt-terraformtestrecordset"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_txt_record_set", "type", "TXT"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_txt_record_set", "ttl", "6000"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_ns_record_set", "name", "ns-terraformtestrecordset"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_ns_record_set", "type", "NS"),
					resource.TestCheckResourceAttr("vinyldns_record_set.test_ns_record_set", "ttl", "6000"),
				),
			},
			resource.TestStep{
				ResourceName:      "vinyldns_record_set.test_a_record_set",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSRecordSetImportARecordStateCheck,
			},
			resource.TestStep{
				ResourceName:      "vinyldns_record_set.test_cname_record_set",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSRecordSetImportCNAMERecordStateCheck,
			},
			resource.TestStep{
				ResourceName:      "vinyldns_record_set.test_txt_record_set",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSRecordSetImportTXTRecordStateCheck,
			},
			resource.TestStep{
				ResourceName:      "vinyldns_record_set.test_ns_record_set",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSRecordSetImportNSRecordStateCheck,
			},
		},
	})
}

func TestAccVinylDNSRecordSetMoveZones(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSRecordSetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVinylDNSRecordSetConfigMoveZones,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSRecordSetMoveZonesExists("vinyldns_record_set.test_a_record_set_2"),
				),
			},
			resource.TestStep{
				Config: testAccVinylDNSRecordSetConfigMoveZonesUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSRecordSetMoveZonesUpdatedExistsInNewZone("vinyldns_record_set.test_a_record_set_2"),
				),
			},
		},
	})
}

func testAccVinylDNSRecordSetImportARecordStateCheck(s []*terraform.InstanceState) error {
	if len(s) != 1 {
		return fmt.Errorf("expected 1 state: %#v", s)
	}

	rs := s[0]

	expName := "terraformtestrecordset"
	name := rs.Attributes["name"]
	if name != expName {
		return fmt.Errorf("expected name attribute to be %s, received %s", expName, name)
	}

	expType := "A"
	aType := rs.Attributes["type"]
	if aType != expType {
		return fmt.Errorf("expected type attribute to be %s, received %s", expType, aType)
	}

	expTTL := "6000"
	ttl := rs.Attributes["ttl"]
	if ttl != expTTL {
		return fmt.Errorf("expected ttl attribute to be %s, received %s", expTTL, ttl)
	}

	return nil
}

func testAccVinylDNSRecordSetImportCNAMERecordStateCheck(s []*terraform.InstanceState) error {
	if len(s) != 1 {
		return fmt.Errorf("expected 1 state: %#v", s)
	}

	rs := s[0]

	expName := "cname-terraformtestrecordset"
	name := rs.Attributes["name"]
	if name != expName {
		return fmt.Errorf("expected name attribute to be %s, received %s", expName, name)
	}

	expType := "CNAME"
	aType := rs.Attributes["type"]
	if aType != expType {
		return fmt.Errorf("expected type attribute to be %s, received %s", expType, aType)
	}

	expTTL := "6000"
	ttl := rs.Attributes["ttl"]
	if ttl != expTTL {
		return fmt.Errorf("expected ttl attribute to be %s, received %s", expTTL, ttl)
	}

	expRecordCNAME := "terraformtestrecordset.system-test."
	recordCNAME := rs.Attributes["record_cname"]
	if recordCNAME != expRecordCNAME {
		return fmt.Errorf("expected record_cname attribute to be %s, received %s", expRecordCNAME, recordCNAME)
	}

	return nil
}

func testAccVinylDNSRecordSetImportTXTRecordStateCheck(s []*terraform.InstanceState) error {
	if len(s) != 1 {
		return fmt.Errorf("expected 1 state: %#v", s)
	}

	rs := s[0]

	expName := "txt-terraformtestrecordset"
	name := rs.Attributes["name"]
	if name != expName {
		return fmt.Errorf("expected name attribute to be %s, received %s", expName, name)
	}

	expType := "TXT"
	aType := rs.Attributes["type"]
	if aType != expType {
		return fmt.Errorf("expected type attribute to be %s, received %s", expType, aType)
	}

	expTTL := "6000"
	ttl := rs.Attributes["ttl"]
	if ttl != expTTL {
		return fmt.Errorf("expected ttl attribute to be %s, received %s", expTTL, ttl)
	}

	return nil
}

func testAccVinylDNSRecordSetImportNSRecordStateCheck(s []*terraform.InstanceState) error {
	if len(s) != 1 {
		return fmt.Errorf("expected 1 state: %#v", s)
	}

	rs := s[0]

	expName := "ns-terraformtestrecordset"
	name := rs.Attributes["name"]
	if name != expName {
		return fmt.Errorf("expected name attribute to be %s, received %s", expName, name)
	}

	expType := "NS"
	aType := rs.Attributes["type"]
	if aType != expType {
		return fmt.Errorf("expected type attribute to be %s, received %s", expType, aType)
	}

	expTTL := "6000"
	ttl := rs.Attributes["ttl"]
	if ttl != expTTL {
		return fmt.Errorf("expected ttl attribute to be %s, received %s", expTTL, ttl)
	}

	return nil
}

func testAccVinylDNSRecordSetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*vinyldns.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vinyldns_record_set" {
			continue
		}
		zID, rsID := parseTwoPartID(rs.Primary.ID)

		// Try to find the record set
		_, err := client.RecordSet(zID, rsID)
		if err == nil {
			return fmt.Errorf("RecordSet %s still exists in zone %s", rsID, zID)
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
		zID, rsID := parseTwoPartID(rs.Primary.ID)
		readRs, err := client.RecordSet(zID, rsID)
		if err != nil {
			return err
		}

		if readRs.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

func testAccCheckVinylDNSRecordSetMoveZonesExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", rs)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RecordSet ID is set")
		}

		client := testAccProvider.Meta().(*vinyldns.Client)
		zID, rsID := parseTwoPartID(rs.Primary.ID)
		readRs, err := client.RecordSet(zID, rsID)
		if err != nil {
			return err
		}

		if readRs.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Record not found")
		}

		z, err := client.Zone(zID)
		if err != nil {
			return err
		}

		// confirm that the record set exists in the correct zone
		expectedZName := "system-test-two."
		if z.Name != expectedZName {
			fmt.Errorf("expected record set to exist in zone %s; it exists in %s", expectedZName, z.Name)
		}

		return nil
	}
}

func testAccCheckVinylDNSRecordSetMoveZonesUpdatedExistsInNewZone(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", rs)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No RecordSet ID is set")
		}

		client := testAccProvider.Meta().(*vinyldns.Client)
		zID, rsID := parseTwoPartID(rs.Primary.ID)
		readRs, err := client.RecordSet(zID, rsID)
		if err != nil {
			return err
		}

		if readRs.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Record not found")
		}

		z, err := client.Zone(zID)
		if err != nil {
			return err
		}

		// confirm that the record set exists in the correct zone after zone_id has changed
		expectedZName := "system-test-three."
		if z.Name != expectedZName {
			fmt.Errorf("expected record set to exist in zone %s; it exists in %s", expectedZName, z.Name)
		}

		return nil
	}
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

resource "vinyldns_record_set" "test_ns_record_set" {
	name = "ns-terraformtestrecordset"
	zone_id = "${vinyldns_zone.test_zone.id}"
	type = "NS"
	ttl = 6000
	record_nsdnames = ["ns1.parent.com."]
	depends_on = [
		"vinyldns_zone.test_zone"
	]
}`

const testAccVinylDNSRecordSetConfigMoveZones = `
resource "vinyldns_group" "test_group_2" {
	name = "terraformtestgrouptwo"
	description = "some description"
	email = "tftest@tf.com"
	member {
	  id = "ok"
	}
	admin {
	  id = "ok"
	}
}

resource "vinyldns_zone" "test_zone_2" {
	name = "system-test-two."
	email = "foo@bar.com"
	admin_group_id = "${vinyldns_group.test_group_2.id}"
	depends_on = [
		"vinyldns_group.test_group_2"
	]
}

resource "vinyldns_zone" "test_zone_3" {
	name = "system-test-three."
	email = "foo@bar.com"
	admin_group_id = "${vinyldns_group.test_group_2.id}"
	depends_on = [
		"vinyldns_group.test_group_2"
	]
}

resource "vinyldns_record_set" "test_a_record_set_2" {
	name = "terraformtestrecordsettwo"
	zone_id = "${vinyldns_zone.test_zone_2.id}"
	owner_group_id = "${vinyldns_group.test_group_2.id}"
	type = "A"
	ttl = 6000
	record_addresses = ["127.0.0.1", "127.0.0.1"]
	depends_on = [
		"vinyldns_zone.test_zone_2"
	]
}`

const testAccVinylDNSRecordSetConfigMoveZonesUpdated = `
resource "vinyldns_group" "test_group_2" {
	name = "terraformtestgrouptwo"
	description = "some description"
	email = "tftest@tf.com"
	member {
	  id = "ok"
	}
	admin {
	  id = "ok"
	}
}

resource "vinyldns_zone" "test_zone_2" {
	name = "system-test-two."
	email = "foo@bar.com"
	admin_group_id = "${vinyldns_group.test_group_2.id}"
	depends_on = [
		"vinyldns_group.test_group_2"
	]
}

resource "vinyldns_zone" "test_zone_3" {
	name = "system-test-three."
	email = "foo@bar.com"
	admin_group_id = "${vinyldns_group.test_group_2.id}"
	depends_on = [
		"vinyldns_group.test_group_2"
	]
}

resource "vinyldns_record_set" "test_a_record_set_2" {
	name = "terraformtestrecordsettwo"
	zone_id = "${vinyldns_zone.test_zone_3.id}"
	owner_group_id = "${vinyldns_group.test_group_2.id}"
	type = "A"
	ttl = 6000
	record_addresses = ["127.0.0.1", "127.0.0.1"]
	depends_on = [
		"vinyldns_zone.test_zone_2"
	]
}`
