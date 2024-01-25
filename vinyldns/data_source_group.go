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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVinylDNSGroupRead(d *schema.ResourceData, meta interface{}) error {
	var id string
	if n, ok := d.GetOk("id"); ok {
		id = n.(string)
	}
	if id == "" {
		return fmt.Errorf("%s must be provided", "id")
	}
	log.Printf("[INFO] Reading VinylDNS Group %s", id)
	z, err := meta.(*vinyldns.Client).Group(id)
	if err != nil {
		return err
	}
	d.SetId(z.ID)
	d.Set("name", z.Name)
	d.Set("email", z.Email)
	d.Set("description", z.Description)
	return nil
}
