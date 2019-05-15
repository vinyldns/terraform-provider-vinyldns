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

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func resourceVinylDNSGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVinylDNSGroupCreate,
		Read:   resourceVinylDNSGroupRead,
		Update: resourceVinylDNSGroupUpdate,
		Delete: resourceVinylDNSGroupDelete,

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
			},
			"member": userSchema(),
			"admin":  userSchema(),
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
		Members:     usersToVinylDNSUser("member", d),
		Admins:      usersToVinylDNSUser("admin", d),
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
		return err
	}

	d.Set("name", g.Name)
	d.Set("email", g.Email)
	d.Set("description", g.Description)

	if g.Members != nil {
		mems := usersToSchema(g.Members)

		if err := d.Set("members", mems); err != nil {
			log.Printf("[WARN] Error setting members for (%s): %s", d.Id(), err)
		}
	}

	if g.Admins != nil {
		admins := usersToSchema(g.Admins)

		if err := d.Set("admins", admins); err != nil {
			log.Printf("[WARN] Error setting admins for (%s): %s", d.Id(), err)
		}
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
		Members:     usersToVinylDNSUser("member", d),
		Admins:      usersToVinylDNSUser("admin", d),
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
		return err
	}

	d.SetId("")

	return nil
}

func userSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"user_name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"first_name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"last_name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true},
				"email": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"created": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		Set: func(v interface{}) int {
			return hashcode.String(v.(string))
		},
	}
}

func usersToVinylDNSUser(userType string, d *schema.ResourceData) []vinyldns.User {
	users := []vinyldns.User{}
	usersCount := d.Get(fmt.Sprintf("%s.#", userType)).(int)

	for i := 0; i < usersCount; i++ {
		prefix := fmt.Sprintf("%s.%d", userType, i)

		users = append(users, vinyldns.User{
			UserName:  d.Get(prefix + ".user_name").(string),
			FirstName: d.Get(prefix + ".first_name").(string),
			LastName:  d.Get(prefix + ".last_name").(string),
			Email:     d.Get(prefix + ".email").(string),
			Created:   d.Get(prefix + ".created").(string),
			ID:        d.Get(prefix + ".id").(string),
		})
	}

	return users
}

func usersToSchema(users []vinyldns.User) []map[string]interface{} {
	var saves []map[string]interface{}

	for _, user := range users {
		u := make(map[string]interface{})

		u["user_name"] = user.UserName
		u["first_name"] = user.FirstName
		u["last_name"] = user.LastName
		u["email"] = user.Email
		u["created"] = user.Created
		u["id"] = user.ID

		saves = append(saves, u)
	}

	return saves
}
