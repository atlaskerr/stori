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

// Backend is an interface for Stori's persistence layer.
type Backend interface {
	Namespace

	// GetSchema retrieves a valid jsonschema in bytes from the backend. Used to
	// validate the config files.
	GetJSONSchema() []byte

	// Setup passes validated configuration data to the backend for
	// initialization.
	Setup(interface{}) error
}
