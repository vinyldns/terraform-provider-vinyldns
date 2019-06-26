package vinyldns

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func dataSourceVinylDNSZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVinylDNSZoneRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVinylDNSZoneRead(d *schema.ResourceData, meta interface{}) error {
	var name string
	if n, ok := d.GetOk("name"); ok {
		name = n.(string)
	}

	if name == "" {
		return fmt.Errorf("%s must be provided", "name")
	}

	log.Printf("[INFO] Reading VinylDNS zone %s", name)

	z, err := meta.(*vinyldns.Client).ZoneByName(name)
	if err != nil {
		return err
	}

	d.SetId(z.ID)

	d.Set("name", z.Name)
	d.Set("email", z.Email)
	d.Set("admin_group_id", z.AdminGroupID)

	return nil
}
