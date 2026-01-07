package vinyldns

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func dataSourceVinylDNSRecordSets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVinylDNSRecordSetsRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"record_sets": {
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
						"fqdn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ttl": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"owner_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVinylDNSRecordSetsRead(d *schema.ResourceData, meta interface{}) error {
	zoneID := d.Get("zone_id").(string)
	nameFilter := d.Get("name_filter").(string)

	log.Printf("[INFO] Reading VinylDNS record sets (zone_id=%s name_filter=%s)", zoneID, nameFilter)

	filter := vinyldns.ListFilter{
		NameFilter: nameFilter,
	}

	records, err := meta.(*vinyldns.Client).RecordSetsListAll(zoneID, filter)
	if err != nil {
		return err
	}

	flattened := make([]interface{}, 0, len(records))
	for _, rs := range records {
		flattened = append(flattened, map[string]interface{}{
			"id":             rs.ID,
			"name":           rs.Name,
			"fqdn":           rs.FQDN,
			"type":           rs.Type,
			"ttl":            rs.TTL,
			"owner_group_id": rs.OwnerGroupID,
			"status":         rs.Status,
		})
	}

	if err := d.Set("record_sets", flattened); err != nil {
		return fmt.Errorf("error setting record_sets for zone %s: %s", zoneID, err)
	}

	d.SetId(fmt.Sprintf("%s:%s", zoneID, nameFilter))

	return nil
}
