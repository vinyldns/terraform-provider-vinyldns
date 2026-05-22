package vinyldns

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func dataSourceVinylDNSGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVinylDNSGroupsRead,

		Schema: map[string]*schema.Schema{
			"name_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceVinylDNSGroupsRead(d *schema.ResourceData, meta interface{}) error {
	nameFilter := d.Get("name_filter").(string)

	log.Printf("[INFO] Reading VinylDNS groups (name_filter=%s)", nameFilter)

	groups, err := meta.(*vinyldns.Client).GroupsListAll(vinyldns.ListFilter{
		NameFilter: nameFilter,
	})
	if err != nil {
		return err
	}

	flattened := make([]interface{}, 0, len(groups))
	for _, group := range groups {
		memberIDs := make([]interface{}, 0, len(group.Members))
		for _, member := range group.Members {
			memberIDs = append(memberIDs, member.ID)
		}

		adminIDs := make([]interface{}, 0, len(group.Admins))
		for _, admin := range group.Admins {
			adminIDs = append(adminIDs, admin.ID)
		}

		flattened = append(flattened, map[string]interface{}{
			"id":          group.ID,
			"name":        group.Name,
			"email":       group.Email,
			"description": group.Description,
			"member_ids":  schema.NewSet(schema.HashString, memberIDs),
			"admin_ids":   schema.NewSet(schema.HashString, adminIDs),
		})
	}

	if err := d.Set("groups", flattened); err != nil {
		return fmt.Errorf("error setting groups: %s", err)
	}

	if nameFilter == "" {
		d.SetId("groups:all")
	} else {
		d.SetId(fmt.Sprintf("groups:%s", nameFilter))
	}

	return nil
}
