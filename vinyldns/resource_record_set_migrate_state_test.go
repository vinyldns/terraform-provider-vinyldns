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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestVinylDNSRecordSetMigrateStateID(t *testing.T) {
	cases := map[string]struct {
		StateVersion int
		ID           string
		Attributes   map[string]string
		Expected     string
		Meta         interface{}
	}{
		"v0_0": {
			StateVersion: 0,
			ID:           "id",
			Attributes: map[string]string{
				"zone_id": "zone-id",
			},
			Expected: "zone-id:id",
		},
	}

	for tn, tc := range cases {
		is := &terraform.InstanceState{
			ID:         tc.ID,
			Attributes: tc.Attributes,
		}
		is, err := resourceVinylDNSRecordSetMigrateState(
			tc.StateVersion, is, tc.Meta)

		if err != nil {
			t.Fatalf("bad: %s, err: %#v", tn, err)
		}

		if is.ID != tc.Expected {
			t.Fatalf("bad VinylDNS RecordSet Migrate: %s\n\n expected: %s", is.ID, tc.Expected)
		}
	}
}

func TestVinylDNSRecordSetMigrateStateRecordTexts(t *testing.T) {
	cases := map[string]struct {
		StateVersion int
		ID           string
		Attributes   map[string]string
		Expected     map[string]string
		Meta         interface{}
	}{
		"v0_0": {
			StateVersion: 0,
			ID:           "id",
			Attributes: map[string]string{
				"zone_id":     "zone-id",
				"record_text": "some-text",
			},
			Expected: map[string]string{
				"record_texts.#":          "1",
				"record_texts.3073014027": "some-text",
			},
		},
	}

	for tn, tc := range cases {
		is := &terraform.InstanceState{
			ID:         tc.ID,
			Attributes: tc.Attributes,
		}
		is, err := resourceVinylDNSRecordSetMigrateState(
			tc.StateVersion, is, tc.Meta)

		if err != nil {
			t.Fatalf("bad: %s, err: %#v", tn, err)
		}

		for k, v := range tc.Expected {
			if is.Attributes[k] != v {
				t.Fatalf(
					"bad: %s\n\n expected: %#v -> %#v\n got: %#v -> %#v\n in: %#v",
					tn, k, v, k, is.Attributes[k], is.Attributes)
			}
		}
	}
}
