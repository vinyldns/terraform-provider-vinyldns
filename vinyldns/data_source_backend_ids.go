package vinyldns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/smartystreets/go-aws-auth"
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

	ids, err := zoneBackendIDs(meta.(*vinyldns.Client))
	if err != nil {
		return err
	}

	if err := d.Set("backend_ids", ids); err != nil {
		return fmt.Errorf("error setting backend_ids: %s", err)
	}

	d.SetId("backend-ids")

	return nil
}

func zoneBackendIDs(client *vinyldns.Client) ([]string, error) {
	host := strings.TrimRight(client.Host, "/")
	endpoint := fmt.Sprintf("%s/zones/backendids", host)

	req, err := http.NewRequest("GET", endpoint, bytes.NewReader(nil))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", client.UserAgent)
	req.Header.Set("Content-Type", "application/json")

	awsauth.Sign4(req, awsauth.Credentials{
		AccessKeyID:     client.AccessKey,
		SecretAccessKey: client.SecretKey,
	})

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("unexpected response code from backend ids endpoint: %d", resp.StatusCode)
	}

	var ids []string
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}

	return ids, nil
}
