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

local oapi = import 'openapi-jsonnet/v3.0.0/openapi.libsonnet';
local ops = import 'operations.libsonnet';
local params = import 'parameters.libsonnet';
local pathItem = oapi.pathItem;

local endpoints = {
  stori:: {
    namespace:: '/images/namespaces',
  },
};

local items = {
  stori:: {
    local namespace =
      pathItem.new()
      .addOperation('PUT', ops.stori.namespaceCreate),

    namespace:: namespace,
  },
};

local storiPaths =
  oapi.paths.new()
  .addPath(endpoints.stori.namespace, items.stori.namespace)
;

storiPaths
