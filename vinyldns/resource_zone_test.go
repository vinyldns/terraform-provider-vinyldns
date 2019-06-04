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
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

const (
	zName             = "system-test."
	zEmail            = "email@foo.com"
	zEmailUpdated     = "updated_email@foo.com"
	zConName          = "vinyldns."
	zConKey           = "nzisn+4G2ldMn0q1CV3vsg=="
	zConKeyName       = "vinyldns."
	zConPrimaryServer = "vinyldns-bind9"
)

func TestAccVinylDNSZoneBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVinylDNSZoneConfigBasic(zEmail),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSZoneBasicExists("vinyldns_zone.test_zone", zEmail),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "name", zName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "email", zEmail),
				),
			},
			resource.TestStep{
				Config: testAccVinylDNSZoneConfigBasic(zEmailUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSZoneBasicExists("vinyldns_zone.test_zone", zEmailUpdated),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "name", zName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "email", zEmailUpdated),
				),
			},
			resource.TestStep{
				ResourceName:      "vinyldns_zone.test_zone",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSZoneImportStateCheck,
			},
		},
	})
}

func TestAccVinylDNSZoneWithACL(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVinylDNSZoneConfigWithACL(zEmail, "TXT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSZoneWithACLExists("vinyldns_zone.test_zone", zEmail, "TXT"),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "name", zName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "email", zEmail),
				),
			},
			resource.TestStep{
				Config: testAccVinylDNSZoneConfigWithACL(zEmailUpdated, "A"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSZoneWithACLExists("vinyldns_zone.test_zone", zEmailUpdated, "A"),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "name", zName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "email", zEmailUpdated),
				),
			},
			resource.TestStep{
				ResourceName:      "vinyldns_zone.test_zone",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSZoneWithACLImportStateCheck,
			},
		},
	})
}

func TestAccVinylDNSZoneWithConnection(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVinylDNSZoneConfigWithConnection(zEmail, "zone_connection"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSZoneWithConnectionExists("vinyldns_zone.test_zone", zEmail),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "name", zName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "email", zEmail),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "zone_connection.0.name", zConName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "zone_connection.0.key", zConKey),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "zone_connection.0.key_name", zConKeyName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "zone_connection.0.primary_server", zConPrimaryServer),
				),
			},
			resource.TestStep{
				Config: testAccVinylDNSZoneConfigWithConnection(zEmailUpdated, "zone_connection"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSZoneWithConnectionExists("vinyldns_zone.test_zone", zEmailUpdated),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "name", zName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "email", zEmailUpdated),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "zone_connection.0.name", zConName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "zone_connection.0.key", zConKey),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "zone_connection.0.key_name", zConKeyName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "zone_connection.0.primary_server", zConPrimaryServer),
				),
			},
			resource.TestStep{
				ResourceName:      "vinyldns_zone.test_zone",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSZoneImportStateCheck,
			},
		},
	})
}

func TestAccVinylDNSZoneWithTransferConnection(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSZoneDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVinylDNSZoneConfigWithConnection(zEmail, "transfer_connection"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSZoneWithTransferConnectionExists("vinyldns_zone.test_zone", zEmail),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "name", zName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "email", zEmail),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "transfer_connection.0.name", zConName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "transfer_connection.0.key", zConKey),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "transfer_connection.0.key_name", zConKeyName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "transfer_connection.0.primary_server", zConPrimaryServer),
				),
			},
			resource.TestStep{
				Config: testAccVinylDNSZoneConfigWithConnection(zEmailUpdated, "transfer_connection"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSZoneWithTransferConnectionExists("vinyldns_zone.test_zone", zEmailUpdated),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "name", zName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "email", zEmailUpdated),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "transfer_connection.0.name", zConName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "transfer_connection.0.key", zConKey),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "transfer_connection.0.key_name", zConKeyName),
					resource.TestCheckResourceAttr("vinyldns_zone.test_zone", "transfer_connection.0.primary_server", zConPrimaryServer),
				),
			},
			resource.TestStep{
				ResourceName:      "vinyldns_zone.test_zone",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSZoneImportStateCheck,
			},
		},
	})
}

func testAccVinylDNSZoneImportStateCheck(s []*terraform.InstanceState) error {
	if len(s) != 1 {
		return fmt.Errorf("expected 1 state: %#v", s)
	}

	rs := s[0]

	expName := zName
	name := rs.Attributes["name"]
	if name != zName {
		return fmt.Errorf("expected name attribute to be %s, received %s", expName, name)
	}

	expEmail := zEmailUpdated
	email := rs.Attributes["email"]
	if email != expEmail {
		return fmt.Errorf("expected email attribute to be %s, received %s", expEmail, email)
	}

	if rs.Attributes["admin_group_id"] == "" {
		return fmt.Errorf("expected admin_group_id attribute to have value")
	}

	return nil
}

func testAccVinylDNSZoneWithACLImportStateCheck(s []*terraform.InstanceState) error {
	if len(s) != 1 {
		return fmt.Errorf("expected 1 state: %#v", s)
	}

	rs := s[0]

	expName := zName
	name := rs.Attributes["name"]
	if name != zName {
		return fmt.Errorf("expected name attribute to be %s, received %s", expName, name)
	}

	expEmail := zEmailUpdated
	email := rs.Attributes["email"]
	if email != expEmail {
		return fmt.Errorf("expected email attribute to be %s, received %s", expEmail, email)
	}

	if rs.Attributes["admin_group_id"] == "" {
		return fmt.Errorf("expected admin_group_id attribute to have value")
	}

	//TODO: why does acl_rule get a different index each time?
	// According to Terraform docs, "items are stored in state with an index value calculated by the hash of the attributes of the set."
	// https://www.terraform.io/docs/extend/schemas/schema-types.html
	/*
		accessLev := rs.Attributes["acl_rule.1878097405.access_level"]
		if accessLev != "Delete" {
			return fmt.Errorf("expected acl_rule attribute to have value; got %s", accessLev)
		}

			recTypes := rs.Attributes["acl_rule.0.record_types.430998943"]
			if recTypes != "A" {
				return fmt.Errorf("expected acl_rule record_types attribute to have value; got %s", recTypes)
			}
	*/

	return nil
}

func testAccVinylDNSZoneDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*vinyldns.Client)

	for _, rs := range s.RootModule().Resources {
		log.Printf("[INFO] testing zone destruction; rs.Type: %s", rs.Type)
		if rs.Type != "vinyldns_zone" {
			continue
		}
		id := rs.Primary.ID

		log.Printf("[INFO] testing zone destruction: %s", id)

		// Try to find the zone
		_, err := client.Zone(id)
		if err == nil {
			return fmt.Errorf("Zone still exists")
		}
	}

	return nil
}

func testAccCheckVinylDNSZoneBasicExists(n, email string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", rs)
		}
		log.Printf("[INFO] testing that zone exists: %s", rs.Primary.ID)

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		client := testAccProvider.Meta().(*vinyldns.Client)

		readZone, err := client.Zone(rs.Primary.ID)
		if err != nil {
			return err
		}

		if readZone.Name != zName {
			return fmt.Errorf("Zone %s not found", zName)
		}

		if readZone.Email != email {
			return fmt.Errorf("Zone %s with email %s not found", zName, email)
		}

		return nil
	}
}

func testAccCheckVinylDNSZoneWithACLExists(n, email, recType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", rs)
		}

		log.Printf("[INFO] testing that zone exists: %s", rs.Primary.ID)

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		client := testAccProvider.Meta().(*vinyldns.Client)

		readZone, err := client.Zone(rs.Primary.ID)
		if err != nil {
			return err
		}

		if readZone.Name != zName {
			return fmt.Errorf("Zone %s not found", zName)
		}

		if readZone.Email != email {
			return fmt.Errorf("Zone %s with email %s not found", zName, email)
		}

		acl := readZone.ACL.Rules[0]

		if acl.GroupID == "" {
			return fmt.Errorf("Zone %s ACL rule has an empty GroupID", zName)
		}

		if acl.AccessLevel != "Delete" {
			return fmt.Errorf("Expected Zone %s ACL rule AccessLevel to be 'Delete'; got %s", zName, acl.AccessLevel)
		}

		if acl.RecordTypes[0] != recType {
			return fmt.Errorf("Expected Zone %s ACL rule RecordTypes to include 'TXT'; got %s", zName, acl.RecordTypes[0])
		}

		return nil
	}
}

func testAccCheckVinylDNSZoneWithConnectionExists(n, email string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", rs)
		}
		log.Printf("[INFO] testing that zone exists: %s", rs.Primary.ID)

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		client := testAccProvider.Meta().(*vinyldns.Client)

		readZone, err := client.Zone(rs.Primary.ID)
		if err != nil {
			return err
		}

		if readZone.Name != zName {
			return fmt.Errorf("Zone %s not found", zName)
		}

		if readZone.Email != email {
			return fmt.Errorf("Zone %s with email %s not found", zName, email)
		}

		if readZone.Connection.Name != zConName {
			return fmt.Errorf("Zone %s with connection name %s not found", zName, zConName)
		}

		if readZone.Connection.KeyName != zConKeyName {
			return fmt.Errorf("Zone %s with connection key name %s not found", zName, zConKeyName)
		}

		if readZone.Connection.Key != zConKey {
			return fmt.Errorf("Zone %s with connection key %s not found", zName, zConKey)
		}

		if readZone.Connection.PrimaryServer != zConPrimaryServer {
			return fmt.Errorf("Zone %s with connection primary server %s not found", zName, zConPrimaryServer)
		}

		return nil
	}
}

func testAccCheckVinylDNSZoneWithTransferConnectionExists(n, email string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", rs)
		}
		log.Printf("[INFO] testing that zone exists: %s", rs.Primary.ID)

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Zone ID is set")
		}

		client := testAccProvider.Meta().(*vinyldns.Client)

		readZone, err := client.Zone(rs.Primary.ID)
		if err != nil {
			return err
		}

		if readZone.Name != zName {
			return fmt.Errorf("Zone %s not found", zName)
		}

		if readZone.Email != email {
			return fmt.Errorf("Zone %s with email %s not found", zName, email)
		}

		if readZone.TransferConnection.Name != zConName {
			return fmt.Errorf("Zone %s with transfer connection name %s not found", zName, zConName)
		}

		if readZone.TransferConnection.KeyName != zConKeyName {
			return fmt.Errorf("Zone %s with transfer connection key name %s not found", zName, zConKeyName)
		}

		if readZone.TransferConnection.Key != zConKey {
			return fmt.Errorf("Zone %s with transfer connection key %s not found", zName, zConKey)
		}

		if readZone.TransferConnection.PrimaryServer != zConPrimaryServer {
			return fmt.Errorf("Zone %s with transfer connection primary server %s not found", zName, zConPrimaryServer)
		}

		return nil
	}
}

func testAccVinylDNSZoneConfigBasic(email string) string {
	const t = `
resource "vinyldns_group" "test_group" {
	name = "terraformtestgroup"
	email = "tftest@tf.com"
	member_ids = ["ok"]
	admin_ids = ["ok"]
}

resource "vinyldns_zone" "test_zone" {
	name = "system-test."
	email = "%s"
	admin_group_id = "${vinyldns_group.test_group.id}"
	depends_on = [
		"vinyldns_group.test_group"
	]
}`

	return fmt.Sprintf(t, email)
}

func testAccVinylDNSZoneConfigWithACL(email, rTypes string) string {
	const t = `
resource "vinyldns_group" "test_group" {
	name = "terraformtestgroup"
	email = "tftest@tf.com"
	member_ids = ["ok"]
	admin_ids = ["ok"]
}

resource "vinyldns_zone" "test_zone" {
	name = "%s"
	email = "%s"
	admin_group_id = "${vinyldns_group.test_group.id}"
	acl_rule {
		access_level = "Delete"
		group_id = "${vinyldns_group.test_group.id}"
		record_types = ["%s"]
	}
	depends_on = [
		"vinyldns_group.test_group"
	]
}`

	return fmt.Sprintf(t, zName, email, rTypes)
}

func testAccVinylDNSZoneConfigWithConnection(email, conType string) string {
	const t = `
resource "vinyldns_group" "test_group" {
	name = "terraformtestgroup"
	email = "tftest@tf.com"
	member_ids = ["ok"]
	admin_ids = ["ok"]
}

resource "vinyldns_zone" "test_zone" {
	name = "%s"
	email = "%s"
	admin_group_id = "${vinyldns_group.test_group.id}"
	%s {
		name = "%s"
		key = "%s"
		key_name = "%s"
		primary_server = "%s"
	}
	depends_on = [
		"vinyldns_group.test_group"
	]
}`

	return fmt.Sprintf(t, zName, email, conType, zConName, zConKey, zConKeyName, zConPrimaryServer)
}
