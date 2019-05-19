/*
Copyright 2018 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vinyldns

import (
	"errors"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/vinyldns/go-vinyldns/vinyldns"
)

func resourceVinylDNSZone() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceVinylDNSZoneCreate,
		Read:          resourceVinylDNSZoneRead,
		Update:        resourceVinylDNSZoneUpdate,
		Delete:        resourceVinylDNSZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"admin_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"shared": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"transfer_connection": connectionSchema(),
			"zone_connection":     connectionSchema(),
		},
	}
}

func resourceVinylDNSZoneCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	log.Printf("[INFO] Creating vinyldns zone: %s", name)
	change, err := meta.(*vinyldns.Client).ZoneCreate(zone(d))
	if err != nil {
		return err
	}

	log.Printf("[INFO] Setting *schema.ResourceData zone ID to: %s", change.Zone.ID)

	d.SetId(change.Zone.ID)

	log.Printf("[INFO] *schema.ResourceData ID: %s", d.Id())

	err = waitUntilZoneCreated(d, meta)
	if err != nil {
		return err
	}

	return resourceVinylDNSZoneRead(d, meta)
}

func resourceVinylDNSZoneRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading vinyldns zone: %s", d.Id())
	zone, err := meta.(*vinyldns.Client).Zone(d.Id())
	if err != nil {
		return err
	}

	d.Set("name", zone.Name)
	d.Set("email", zone.Email)
	d.Set("admin_group_id", zone.AdminGroupID)
	d.Set("status", zone.Status)
	d.Set("shared", zone.Shared)
	d.Set("created", zone.Created)

	if zone.Connection != nil {
		d.Set("zone_connection", []interface{}{
			map[string]interface{}{
				"name":           zone.Connection.Name,
				"key":            zone.Connection.Key,
				"key_name":       zone.Connection.KeyName,
				"primary_server": zone.Connection.PrimaryServer,
			},
		})
	}

	if zone.TransferConnection != nil {
		d.Set("transfer_connection", []interface{}{
			map[string]interface{}{
				"name":           zone.TransferConnection.Name,
				"key":            zone.TransferConnection.Key,
				"key_name":       zone.TransferConnection.KeyName,
				"primary_server": zone.TransferConnection.PrimaryServer,
			},
		})
	}

	return nil
}

func resourceVinylDNSZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Updating vinyldns zone: %s", d.Id())
	change, err := meta.(*vinyldns.Client).ZoneUpdate(d.Id(), zone(d))
	if err != nil {
		return err
	}

	err = waitUntilZoneChangeDeployed(d, meta, change.Zone.ID)
	if err != nil {
		return err
	}

	return resourceVinylDNSZoneRead(d, meta)
}

func zoneConnection(d *schema.ResourceData) *vinyldns.ZoneConnection {
	name := d.Get("zone_connection.0.name").(string)

	if name != "" {
		log.Printf("[INFO] setting zone connection: %s", d.Get("zone_connection.0.name"))
		return &vinyldns.ZoneConnection{
			Name:          name,
			Key:           d.Get("zone_connection.0.key").(string),
			KeyName:       d.Get("zone_connection.0.key_name").(string),
			PrimaryServer: d.Get("zone_connection.0.primary_server").(string),
		}
	}

	return &vinyldns.ZoneConnection{}
}

func transferConnection(d *schema.ResourceData) *vinyldns.ZoneConnection {
	name := d.Get("transfer_connection.0.name").(string)

	if name != "" {
		log.Printf("[INFO] setting transfer connection: %s", d.Get("transfer_connection.0.name"))
		return &vinyldns.ZoneConnection{
			Name:          name,
			Key:           d.Get("transfer_connection.0.key").(string),
			KeyName:       d.Get("transfer_connection.0.key_name").(string),
			PrimaryServer: d.Get("transfer_connection.0.primary_server").(string),
		}
	}

	return &vinyldns.ZoneConnection{}
}

func resourceVinylDNSZoneDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting vinyldns zone: %s", d.Id())

	_, err := meta.(*vinyldns.Client).ZoneDelete(d.Id())
	if err != nil {
		return err
	}

	err = waitUntilZoneDeleted(d, meta, d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func waitUntilZoneChangeDeployed(d *schema.ResourceData, meta interface{}, changeID string) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Pending", ""},
		Target:       []string{"Synced"},
		Refresh:      zoneStateRefreshFunc(d, meta, changeID),
		Timeout:      30 * time.Minute,
		Delay:        500 * time.Millisecond,
		MinTimeout:   15 * time.Second,
		PollInterval: 500 * time.Millisecond,
	}

	_, err := stateConf.WaitForState()
	return err
}

func zoneStateRefreshFunc(d *schema.ResourceData, meta interface{}, changeID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[INFO] waiting for Complete status of %v, %s", d.Get("name"), d.Id())
		zc, err := meta.(*vinyldns.Client).ZoneChange(d.Id(), changeID)
		if err != nil {
			log.Printf("[ERROR] %#v", err)
			return nil, "", err
		}
		if zc.Status == "Failed" {
			err = errors.New("zone status Failed")
			log.Printf("[ERROR] zone status Failed: %#v", err)
			return zc, zc.Status, err
		}

		return zc, zc.Status, nil
	}
}

func waitUntilZoneDeleted(d *schema.ResourceData, meta interface{}, zoneID string) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Pending"},
		Target:       []string{"Deleted"},
		Refresh:      zoneDeletedStateRefreshFunc(d, meta, zoneID),
		Timeout:      30 * time.Minute,
		Delay:        500 * time.Millisecond,
		MinTimeout:   15 * time.Second,
		PollInterval: 500 * time.Millisecond,
	}

	_, err := stateConf.WaitForState()
	return err
}

func zoneDeletedStateRefreshFunc(d *schema.ResourceData, meta interface{}, zoneID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		state := "Pending"

		log.Printf("[INFO] waiting for successful deletion of %v, %s", d.Get("name"), d.Id())
		exists, err := meta.(*vinyldns.Client).ZoneExists(d.Id())
		if err != nil {
			log.Printf("[ERROR] %#v", err)
			return nil, "", err
		}

		if !exists {
			state = "Deleted"
		}

		return &zoneState{State: state}, state, err
	}
}

func waitUntilZoneCreated(d *schema.ResourceData, meta interface{}) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Pending"},
		Target:       []string{"Created"},
		Refresh:      zoneCreatedStateRefreshFunc(d, meta),
		Timeout:      30 * time.Minute,
		Delay:        500 * time.Millisecond,
		MinTimeout:   15 * time.Second,
		PollInterval: 500 * time.Millisecond,
	}

	_, err := stateConf.WaitForState()
	return err
}

func zoneCreatedStateRefreshFunc(d *schema.ResourceData, meta interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		state := "Pending"

		log.Printf("[INFO] waiting for successful creation of %v, %s", d.Get("name"), d.Id())
		exists, err := meta.(*vinyldns.Client).ZoneExists(d.Id())
		if err != nil {
			log.Printf("[ERROR] %#v", err)
			return nil, "", err
		}

		if exists {
			state = "Created"
		}

		return &zoneState{State: state}, state, err
	}
}

type zoneState struct {
	State string
}

func connectionSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"key": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"key_name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"primary_server": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func zone(d *schema.ResourceData) *vinyldns.Zone {
	zone := &vinyldns.Zone{
		Name:         d.Get("name").(string),
		Email:        d.Get("email").(string),
		AdminGroupID: d.Get("admin_group_id").(string),
	}

	if d.Get("zone_connection.0.name").(string) != "" {
		zone.Connection = zoneConnection(d)
	}

	if d.Get("transfer_connection.0.name").(string) != "" {
		zone.TransferConnection = transferConnection(d)
	}

	shared := d.Get("shared").(bool)
	if shared == true || shared == false {
		zone.Shared = shared
	}

	return zone
}
