package vinyldns

import (
  "fmt"
  "log"

  "github.com/hashicorp/terraform/helper/schema"
  "github.com/vinyldns/go-vinyldns/vinyldns"
)

func dataSourceVinylDNSGroup() *schema.Resource {
  return &schema.Resource{
    Read: dataSourceVinylDNSGroupRead,

    Schema: map[string]*schema.Schema{
      "id": {
        Type: schema.TypeString,
        Optional: true,
      },
      "name": {
        Type: schema.TypeString,
        Computed: true,
      },
      "email": {
        Type: schema.TypeString,
        Computed: true,
      },
    },
  }
}

func dataSourceVinylDNSGroupRead(d *schema.ResourceData, meta interface{}) error {
  var id string
  if i, ok := d.GetOk("id"); ok {
    id = i.(string)
  }

  if id == "" {
    return fmt.Errorf("%s must be provided", "id")
  }

  log.Printf("[INFO] Reading VinylDNS group %s", id)

  g, err := meta.(*vinyldns.Client).Group(id)
  if err != nil {
    return err
  }

  d.SetId(g.ID)

  d.Set("name", g.Name)
  d.Set("email", g.Email)

  return nil
}
