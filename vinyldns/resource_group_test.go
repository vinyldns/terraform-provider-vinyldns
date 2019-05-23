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

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccVinylDNSGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccVinylDNSGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVinylDNSGroupConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSGroupExists("vinyldns_group.test_group"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "name", "terraformtestgroup"),
				),
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
			resource.TestStep{
				Config: testAccVinylDNSGroupConfigWithoutDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVinylDNSGroupExists("vinyldns_group.test_group"),
					resource.TestCheckResourceAttr("vinyldns_group.test_group", "description", "Managed by Terraform"),
				),
			},
		},
	})
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

func testAccCheckVinylDNSGroupExists(n string) resource.TestCheckFunc {
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
		if g.Description == "" {
			return fmt.Errorf("Group 'description' not set")
		}

		return nil
	}
}

const testAccVinylDNSGroupConfigBasic = `
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
}`

const testAccVinylDNSGroupConfigWithoutDescription = `
resource "vinyldns_group" "test_group" {
	name = "terraformtestgroup"
	email = "tftest@tf.com"
	member {
	  id = "ok"
	}
	admin {
	  id = "ok"
	}
}`
