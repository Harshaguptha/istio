// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package constraint

import "encoding/json"

// Collection is a set of checks against a collection.
type Collection struct {
	Name  string  `json:"collection"`
	Check []Range `json:"check"`
}

var _ json.Unmarshaler = &Collection{}

// UnmarshalJSON implements json.Unmarshaler
func (c *Collection) UnmarshalJSON(b []byte) error {
	i := struct {
		Collection string            `json:"collection"`
		Check      []json.RawMessage `json:"check"`
	}{}

	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	c.Name = i.Collection

	for _, rb := range i.Check {
		rc, err := parseRange(rb)
		if err != nil {
			return err
		}
		c.Check = append(c.Check, rc)
	}
	return nil
}
