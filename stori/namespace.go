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

// Namespace defines methods backends must implement for working with registry
// namespaces.
type Namespace interface {

	// CreateNamespace creates a new namespace in the registry.
	CreateNamespace(string) (*NamespaceInfo, error)

	// LookupNamespaceByName performs a lookup on the registry for a namespace
	// with a name matching the string provided.
	LookupNamespaceByName(string) (*NamespaceInfo, error)

	// LookupNamespaceByID performs a lookup on the registry for a namespace
	// with an ID matching the string provided.
	LookupNamespaceByID(string) (*NamespaceInfo, error)
}

// NamespaceInfo defines operations for stori namespaces. A namespace is
// a logical grouping of repositories and a fundamental building block for
// policy enforcement.
type NamespaceInfo struct {

	// Name is a name that uniquely identifies a namespace among all namespaces.
	Name string

	// NID is a unique value for a particular namespace that will change if the
	// namespace is removed from the system and another namespace is added with
	// the same name.
	ID int32

	// Repositories is a list of repositories that reside logically within
	// a namespace.
	Repositories []Repository
}
