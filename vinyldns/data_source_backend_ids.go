package vinyldns

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func dataSourceVinylDNSBackendIDs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVinylDNSBackendIDsRead,

		Schema: map[string]*schema.Schema{
			"backend_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceVinylDNSBackendIDsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading VinylDNS backend IDs")

	ids, err := meta.(*vinyldns.Client).ZoneBackendIDs()
	if err != nil {
		return err
	}

	if err := d.Set("backend_ids", ids); err != nil {
		return fmt.Errorf("error setting backend_ids: %s", err)
	}

	d.SetId("backend-ids")

	return nil
}
