local schemas = import 'schemas/schemas.libsonnet';
local string = schemas.types.string;

{
  namespace:: {
    name: 'namespace',
    'in': 'path',
    description: 'A logical grouping of repositories.',
    required: true,
    allowEmptyValue: false,
    schema: string,
  },

  repository:: {
    name: 'repository',
    'in': 'path',
    description: 'A place to store an image.',
    required: true,
    allowEmptyValue: false,
    schema: string,
  },

  reference:: {
    name: 'reference',
    'in': 'path',
    description: 'Can either be a tag name or digest.',
    required: true,
    allowEmptyValue: false,
    schema: string,
  },

  digest:: {
    name: 'digest',
    'in': 'path',
    description: 'A content addressable identifier.',
    required: true,
    allowEmptyValue: false,
    schema: string,
  },

  tag:: {
    name: 'tag',
    'in': 'path',
    description: 'An image tag.',
    required: true,
    allowEmptyValue: false,
    schema: string,
  },

  id:: {
    name: 'id',
    'in': 'path',
    description: 'A unique id.',
    required: true,
    allowEmptyValue: false,
    schema: string,
  },
}
