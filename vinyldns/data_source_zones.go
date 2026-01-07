package vinyldns

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func dataSourceVinylDNSZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVinylDNSZonesRead,

		Schema: map[string]*schema.Schema{
			"name_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zones": {
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
						"admin_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shared": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"backend_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVinylDNSZonesRead(d *schema.ResourceData, meta interface{}) error {
	nameFilter := d.Get("name_filter").(string)

	log.Printf("[INFO] Reading VinylDNS zones (name_filter=%s)", nameFilter)

	filter := vinyldns.ListFilter{
		NameFilter: nameFilter,
	}

	zones, err := meta.(*vinyldns.Client).ZonesListAll(filter)
	if err != nil {
		return err
	}

	flattened := make([]interface{}, 0, len(zones))
	for _, zone := range zones {
		flattened = append(flattened, map[string]interface{}{
			"id":             zone.ID,
			"name":           zone.Name,
			"email":          zone.Email,
			"admin_group_id": zone.AdminGroupID,
			"status":         zone.Status,
			"shared":         zone.Shared,
			"backend_id":     zone.BackendID,
		})
	}

	if err := d.Set("zones", flattened); err != nil {
		return fmt.Errorf("error setting zones: %s", err)
	}

	if nameFilter == "" {
		d.SetId("zones:all")
	} else {
		d.SetId(fmt.Sprintf("zones:%s", nameFilter))
	}

	return nil
}
