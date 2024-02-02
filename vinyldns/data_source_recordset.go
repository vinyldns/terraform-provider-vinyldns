package vinyldns

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func dataSourceVinylDNSRecordSet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVinylDNSRecordSetRead,
		Schema: map[string]*schema.Schema{
			"recordId": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zoneId": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
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
			"zoneName": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"records": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"isShared": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ownerGroupId": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVinylDNSRecordSetRead(d *schema.ResourceData, meta interface{}) error {
	var zoneId string
	var recordId string
	if n, ok := d.GetOk("zoneId"); ok {
		zoneId = n.(string)
	}
	if n, ok := d.GetOk("recordId"); ok {
		recordId = n.(string)
	}
	if zoneId == "" {
		return fmt.Errorf("%s must be provided", "id")
	}
	if recordId == "" {
		return fmt.Errorf("%s must be provided", "id")
	}
	log.Printf("[INFO] Reading VinylDNS Recordset %s,%s", zoneId, recordId)
	z, err := meta.(*vinyldns.Client).RecordSet(zoneId, recordId)
	if err != nil {
		return err
	}
	d.SetId(z.ID)
	d.Set("name", z.Name)
	d.Set("type", z.Type)
	d.Set("ttl", z.TTL)
	d.Set("type", z.Type)
	d.Set("zoneName", z.ZoneName)
	d.Set("records", z.Records)
	d.Set("isShared", z.IsShared)
	d.Set("ownerGroupId", z.OwnerGroupID)

	return nil
}
