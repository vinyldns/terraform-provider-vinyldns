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

	"github.com/vinyldns/go-vinyldns/vinyldns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccVinylDNSGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVinylDNSGroupConfigBasic("description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSGroupExists("vinyldns_group.test_group", "description"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "name", "terraformtestgroup"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "description", "description"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "member_ids.#", "1"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "member_ids.0", "ok"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "admin_ids.#", "1"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "admin_ids.0", "ok"),
				),
			},
			{
				Config: testAccVinylDNSGroupConfigBasic("updated description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSGroupExists("vinyldns_group.test_group", "updated description"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "name", "terraformtestgroup"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "description", "updated description"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "member_ids.#", "1"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "member_ids.0", "ok"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "admin_ids.#", "1"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "admin_ids.0", "ok"),
				),
			},
			{
				ResourceName:      "vinyldns_group.test_group",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateCheck:  testAccVinylDNSGroupImportStateCheck,
			},
		},
	})
}

func TestAccVinylDNSGroupWithoutDescription(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVinylDNSGroupConfigWithoutDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSGroupExists("vinyldns_group.test_group", "Managed by Terraform"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "description", "Managed by Terraform"),
				),
			},
		},
	})
}

func testAccVinylDNSGroupImportStateCheck(s []*terraform.InstanceState) error {
	if len(s) != 1 {
		return fmt.Errorf("expected 1 state: %#v", s)
	}

	rs := s[0]

	expName := "terraformtestgroup"
	name := rs.Attributes["name"]
	if name != expName {
		return fmt.Errorf("expected name attribute to be %s, received %s", expName, name)
	}

	expDesc := "updated description"
	desc := rs.Attributes["description"]
	if desc != expDesc {
		return fmt.Errorf("expected description attribute to be %s, received %s", expDesc, desc)
	}

	expEmail := "tftest@test.com"
	email := rs.Attributes["email"]
	if email != expEmail {
		return fmt.Errorf("expected email attribute to be %s, received %s", expEmail, email)
	}

	return nil
}

func testAccVinylDNSGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*vinyldns.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vinyldns_group" {
			continue
		}

		// Try to find the group
		_, err := client.Group(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Group %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckVinylDNSGroupExists(n, desc string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", rs)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Group ID is set")
		}

		client := testAccProvider.Meta().(*vinyldns.Client)

		g, err := client.Group(rs.Primary.ID)
		if err != nil {
			return err
		}

		if g.Name != "terraformtestgroup" {
			return fmt.Errorf("Group not found")
		}
		if g.Description != desc {
			return fmt.Errorf("Group 'description' not properly set")
		}

		return nil
	}
}

func testAccVinylDNSGroupConfigBasic(desc string) string {
	const t = `
resource "vinyldns_group" "test_group" {
	name = "terraformtestgroup"
	description = "%s"
	email = "tftest@test.com"
	member_ids = ["ok"]
	admin_ids = ["ok"]
}`

	return fmt.Sprintf(t, desc)
}

const testAccVinylDNSGroupConfigWithoutDescription = `
resource "vinyldns_group" "test_group" {
	name = "terraformtestgroup"
	email = "tftest@test.com"
	member_ids = ["ok"]
	admin_ids = ["ok"]
}`
