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
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func resourceVinylDNSGroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceVinylDNSGroupCreate,
		Read:          resourceVinylDNSGroupRead,
		Update:        resourceVinylDNSGroupUpdate,
		Delete:        resourceVinylDNSGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Managed by Terraform",
			},
			"member_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode(v.(string))
				},
			},
			"admin_ids": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode(v.(string))
				},
			},
		},
	}
}

func resourceVinylDNSGroupCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	log.Printf("[INFO] Creating Group: %s", name)
	created, err := meta.(*vinyldns.Client).GroupCreate(&vinyldns.Group{
		Name:        d.Get("name").(string),
		Email:       d.Get("email").(string),
		Description: d.Get("description").(string),
		Members:     users("member_ids", d),
		Admins:      users("admin_ids", d),
	})
	if err != nil {
		return err
	}

	d.SetId(created.ID)

	return resourceVinylDNSGroupRead(d, meta)
}

func resourceVinylDNSGroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading vinyldns group: %s", d.Id())
	g, err := meta.(*vinyldns.Client).Group(d.Id())
	if err != nil {
		if vErr, ok := err.(*vinyldns.Error); ok {
			if vErr.ResponseCode == http.StatusNotFound {
				log.Printf("[WARN] group (%s) not found, error code (404)", d.Id())

				d.SetId("")

				return nil
			}

			return fmt.Errorf("error reading group (%s): %s", d.Id(), err)
		}

		return fmt.Errorf("error reading group (%s): %s", d.Id(), err)
	}

	d.Set("name", g.Name)
	d.Set("email", g.Email)
	d.Set("description", g.Description)

	memIDs := make([]interface{}, 0, len(g.Members))
	for _, m := range g.Members {
		memIDs = append(memIDs, m.ID)
	}

	if err := d.Set("member_ids", schema.NewSet(schema.HashString, memIDs)); err != nil {
		return fmt.Errorf("error setting member_ids for group %s: %s", d.Id(), err)
	}

	adminIDs := make([]interface{}, 0, len(g.Admins))
	for _, a := range g.Admins {
		adminIDs = append(adminIDs, a.ID)
	}

	if err := d.Set("admin_ids", schema.NewSet(schema.HashString, adminIDs)); err != nil {
		return fmt.Errorf("error setting admin_ids for group %s: %s", d.Id(), err)
	}

	return nil
}

func resourceVinylDNSGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Updating vinyldns group: %s", d.Id())
	_, err := meta.(*vinyldns.Client).GroupUpdate(d.Id(), &vinyldns.Group{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Email:       d.Get("email").(string),
		Description: d.Get("description").(string),
		Members:     users("member_ids", d),
		Admins:      users("admin_ids", d),
	})
	if err != nil {
		return err
	}

	return resourceVinylDNSGroupRead(d, meta)
}

func resourceVinylDNSGroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting vinyldns group: %s", d.Id())

	_, err := meta.(*vinyldns.Client).GroupDelete(d.Id())
	if err != nil {
		if vErr, ok := err.(*vinyldns.Error); ok {
			if vErr.ResponseCode == http.StatusNotFound {
				log.Printf("[WARN] group (%s) not found, error code (404)", d.Id())

				return nil
			}

			return fmt.Errorf("error deleting group (%s): %s", d.Id(), err)
		}

		return fmt.Errorf("error deleting group (%s): %s", d.Id(), err)
	}

	return nil
}

func users(userType string, d *schema.ResourceData) []vinyldns.User {
	resourceUsers := stringSetToStringSlice(d.Get(userType).(*schema.Set))
	users := []vinyldns.User{}
	count := len(resourceUsers)

	for i := 0; i < count; i++ {
		users = append(users, vinyldns.User{
			ID: resourceUsers[i],
		})
	}

	return users
}
