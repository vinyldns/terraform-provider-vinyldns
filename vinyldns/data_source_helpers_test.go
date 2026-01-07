package vinyldns

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func ensureTestGroup(client *vinyldns.Client, name string) (*vinyldns.Group, error) {
	groups, err := client.GroupsListAll(vinyldns.ListFilter{NameFilter: name})
	if err != nil {
		return nil, err
	}

	for i, g := range groups {
		if g.Name == name {
			return &groups[i], nil
		}
	}

	return client.GroupCreate(&vinyldns.Group{
		Name:  name,
		Email: "foo@email.com",
		Members: []vinyldns.User{
			{ID: "ok"},
		},
		Admins: []vinyldns.User{
			{ID: "ok"},
		},
	})
}

func ensureTestZone(client *vinyldns.Client, name, adminGroupID string) (*vinyldns.Zone, error) {
	zone, err := client.ZoneByName(name)
	if err == nil && zone.ID != "" {
		return &zone, nil
	}
	if err != nil {
		if vErr, ok := err.(*vinyldns.Error); ok {
			if vErr.ResponseCode != http.StatusNotFound {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	change, err := client.ZoneCreate(&vinyldns.Zone{
		Name:         name,
		Email:        "foo@email.com",
		AdminGroupID: adminGroupID,
		Connection: &vinyldns.ZoneConnection{
			Name:          "vinyldns.",
			Key:           "nzisn+4G2ldMn0q1CV3vsg==",
			KeyName:       "vinyldns.",
			PrimaryServer: "localhost:19001",
		},
	})
	if err != nil {
		return nil, err
	}

	createdZoneID := change.Zone.ID
	limit := 10
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		zg, err := client.Zone(createdZoneID)
		if err == nil && zg.ID == createdZoneID {
			return &zg, nil
		}

		if i == (limit - 1) {
			log.Printf("[INFO] %d retries reached in polling VinylDNS zone %s", limit, createdZoneID)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("zone %s not available after retries", createdZoneID)
		}
	}

	return nil, fmt.Errorf("zone %s not available", createdZoneID)
}

func ensureTestRecordSet(client *vinyldns.Client, zoneID, name string) error {
	existing, err := client.RecordSetsListAll(zoneID, vinyldns.ListFilter{NameFilter: name})
	if err != nil {
		return err
	}

	for _, rs := range existing {
		if rs.Name == name {
			return nil
		}
	}

	_, err = client.RecordSetCreate(&vinyldns.RecordSet{
		Name:   name,
		ZoneID: zoneID,
		Type:   "A",
		TTL:    300,
		Records: []vinyldns.Record{
			{Address: "127.0.0.1"},
		},
	})
	if err != nil {
		return err
	}

	limit := 10
	for i := 0; i < limit; time.Sleep(5 * time.Second) {
		i++

		records, err := client.RecordSetsListAll(zoneID, vinyldns.ListFilter{NameFilter: name})
		if err == nil {
			for _, rs := range records {
				if rs.Name == name {
					return nil
				}
			}
		}

		if i == (limit - 1) {
			if err != nil {
				return err
			}
			return fmt.Errorf("record set %s not available after retries", name)
		}
	}

	return fmt.Errorf("record set %s not available", name)
}
