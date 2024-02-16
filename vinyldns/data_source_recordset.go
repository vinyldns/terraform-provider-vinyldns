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
			"zoneid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recordset": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVinylDNSRecordSetRead(d *schema.ResourceData, meta interface{}) error {
	var zoneid string

	if n, ok := d.GetOk("zoneid"); ok {
		zoneid = n.(string)
	}

	if zoneid == "" {
		return fmt.Errorf("%s must be provided", "zoneid")
	}

	log.Printf("[INFO] Reading VinylDNS Recordset %s", zoneid)
	z, err := meta.(*vinyldns.Client).RecordSets(zoneid)
	if err != nil {
		return err
	}
	elementMap := make(map[int]vinyldns.RecordSet)
	//mapString := make(map[string]interface{})
	for i, num := range z {
		elementMap[i] = num
	}
	// for key, value := range elementMap {
	// 	strKey := fmt.Sprintf("%v", key)
	// 	strValue := fmt.Sprintf("%v", value)

	// 	mapString[strKey] = strValue
	// }

	s := fmt.Sprintf("%v", elementMap)

	d.SetId(zoneid)
	d.Set("zoneid", zoneid)
	d.Set("recordset", s)
	return nil
}
