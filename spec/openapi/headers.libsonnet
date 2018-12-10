local schemas = import 'schemas/schemas.libsonnet';
local string = schemas.common.string,

{
  wwwAuth:: {
    'WWW-Authenticate': {
      description: 'Defines the authentication method that should be used to gain access to a resource.',
      schema: string,
    },
  },

  dockerApiVersion:: {
    'Docker-Distribution-API-Version': {
      description: 'Used by clients to determine if the registry implements the distribution specification.',
      schema: string,
    },
  },

  dockerContentDigest:: {
    'Docker-Content-Digest': {
      description: 'Includes the digest of the target entity returned in the response. Clients can use this header to verify the integrity of downloaded content',
      schema: string,
    },
  },

  dockerUploadUUID:: {
    'Docker-Upload-UUID': {
      description: 'A unique identifier used to correlate local state with remote state. Clients use this to implement resumable uploads.',
      schema: string,
    },
  },

  contentLength:: {
    'Content-Length': {
      description: 'The byte length of the data returned .',
      schema: string,
    },
  },

  contentType:: {
    'Content-Type': {
      description: 'The MIME type of the data.',
      schema: string,
    },
  },

  contentTypeOptions:: {
    'X-Content-Type-Options': {
      description: 'Content type options.',
      schema: string,
    },
  },

  contentRange:: {
    'Content-Range': {
      description: 'The byte range of the data returned.',
      schema: string,
    },
  },

  date:: {
    Date: {
      description: 'The date and time the message was originated.',
      schema: string,
    },
  },

}
