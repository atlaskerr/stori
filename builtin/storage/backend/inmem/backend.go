// Copyright © 2018 Atlas Kerr atlaskerr@gmail.com
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

package inmem

import (
	"encoding/json"

	"github.com/atlaskerr/stori/schema/stori/storage/backend"
)

// Backend implements stori.Backend.
type Backend struct {
	schema []byte
}

// New returns an new in-memory backend.
func New() (*Backend, error) {
	b := new(Backend)

	// Load jsonschema
	source := backend.InmemConfigLoaderFunc()
	rawJSON, err := json.Marshal(source.JsonSource())
	if err != nil {
		return nil, err
	}

	b.schema = rawJSON

	return b, nil
}

// GetSchema retrieves a valid jsonschema in bytes from the backend. Used to
// validate the config files.
func (b *Backend) GetSchema() []byte {
	return b.schema
}

// Setup passes validated configuration data to the backend for
// initialization.
func (b *Backend) Setup(interface{}) error {
	return nil
}
