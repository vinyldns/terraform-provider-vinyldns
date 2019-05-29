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
)

func Test_parseTwoPartID(t *testing.T) {
	one, two := parseTwoPartID("123:456")
	if one != "123" {
		t.Fatalf("expected parseTwoPartID to return ID part 1 as '123'")
	}
	if two != "456" {
		t.Fatalf("expected parseTwoPartID to return ID part 2 as '456'")
	}
}

func Test_removeBrackets(t *testing.T) {
	r := removeBrackets("[123]")
	if r != "123" {
		t.Fatalf("expected removeBrackets to remove '[]' from a string; got %s", r)
	}
}
