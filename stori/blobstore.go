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

package stori

import (
	"github.com/xeipuuv/gojsonschema"
)

// BlobStore is an interface for Stori's content-addressable storage engine.
type BlobStore interface {

	// GetSchema retrieves a gojsonschema.JSONLoader from the blobstore server
	// to validate the block in a config file referencing the blobstore.
	GetSchema() gojsonschema.JSONLoader

	// Setup passes an empty interface containing configuration data validated
	// by the JSONLoader recieved from `GetSchema()`
	Setup(interface{}) error
}
