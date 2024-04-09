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
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func resourceVinylDNSRecordSetMigrateState(
	v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	switch v {
	case 0:
		log.Println("[INFO] Found VinylDNS RecordSet State v0; migrating to v1")
		return migrateVinylDNSRecordSetStateV0toV1(is)
	default:
		return is, fmt.Errorf("Unexpected schema version: %d", v)
	}
}

func migrateVinylDNSRecordSetStateV0toV1(is *terraform.InstanceState) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Println("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}

	log.Printf("[DEBUG] Attributes before migration for record set %s: %#v", is.ID, is.Attributes)

	newID := is.Attributes["zone_id"] + ":" + is.ID
	is.ID = newID

	rText := is.Attributes["record_text"]
	if rText != "" {
		h := schema.HashString(rText)
		is.Attributes["record_texts.#"] = "1"
		is.Attributes[fmt.Sprintf("record_texts.%v", h)] = rText
	}

	log.Printf("[DEBUG] Attributes after migration: %#v, new id: %s", is.Attributes, newID)

	return is, nil
}
