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
local resp = oapi.response;
local mt = import 'media-types.libsonnet';
local schemas = import '../schema/schema.libsonnet';
local h = import 'headers.libsonnet';

local err = {

  local content = mt.err,

  local badRequest =
    resp.new('400', 'Bad Request', content)
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local unauthorized =
    resp.new('401', 'Unauthorized', content)
    .addHeader(h.common.wwwAuth)
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local forbidden =
    resp.new('403', 'Forbidden', content)
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local notFound =
    resp.new('404', 'Not Found', content)
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local rangeNotSatisfiable =
    resp.new('416', 'Not Found', content)
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
    .addHeader(h.docker.uploadUUID)
  ,

  local tooManyRequests =
    resp.new('429', 'Too Many Requests', content)
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,


  unauthorized: unauthorized,
  forbidden: forbidden,
  notFound: notFound,
  tooManyRequests: tooManyRequests,
  badRequest: badRequest,
  rangeNotSatisfiable: rangeNotSatisfiable,
};

local common = {

  local ok =
    resp.new('200', 'OK')
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local accepted =
    resp.new('202', 'Accepted')
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local created =
    resp.new('201', 'Created')
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local noContent =
    resp.new('204', 'No Content')
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local partialContent =
    resp.new('206', 'Partial Content')
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  local temporaryRedirect =
    resp.new('307', 'Temporary Redirect')
    .addHeader(h.common.contentLength)
    .addHeader(h.common.date)
    .addHeader(h.common.contentType)
    .addHeader(h.common.contentTypeOptions)
  ,

  ok: ok,
  accepted: accepted,
  created: created,
  noContent: noContent,
  partialContent: partialContent,
  temporaryRedirect: temporaryRedirect,
};

{
  err:: err,
  common:: common,
}
