package vinyldns

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func dataSourceVinylDNSGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVinylDNSGroupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"member_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"admin_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceVinylDNSGroupRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Printf("[INFO] Reading VinylDNS group %s", name)

	groups, err := meta.(*vinyldns.Client).GroupsListAll(vinyldns.ListFilter{
		NameFilter: name,
	})
	if err != nil {
		return err
	}

	var match *vinyldns.Group
	for i, g := range groups {
		if g.Name == name {
			if match != nil {
				return fmt.Errorf("found multiple groups with name %s", name)
			}
			match = &groups[i]
		}
	}

	if match == nil {
		return fmt.Errorf("no group found with name %s", name)
	}

	d.SetId(match.ID)
	d.Set("name", match.Name)
	d.Set("email", match.Email)
	d.Set("description", match.Description)

	memIDs := make([]interface{}, 0, len(match.Members))
	for _, m := range match.Members {
		memIDs = append(memIDs, m.ID)
	}
	if err := d.Set("member_ids", schema.NewSet(schema.HashString, memIDs)); err != nil {
		return fmt.Errorf("error setting member_ids for group %s: %s", match.ID, err)
	}

	adminIDs := make([]interface{}, 0, len(match.Admins))
	for _, a := range match.Admins {
		adminIDs = append(adminIDs, a.ID)
	}
	if err := d.Set("admin_ids", schema.NewSet(schema.HashString, adminIDs)); err != nil {
		return fmt.Errorf("error setting admin_ids for group %s: %s", match.ID, err)
	}

	return nil
}
