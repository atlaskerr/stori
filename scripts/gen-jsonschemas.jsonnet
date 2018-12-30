local oci = import '../libsonnet/schemas/oci.libsonnet';
local stori = import '../libsonnet/schemas/stori.libsonnet';

{
  'oci/config.json': oci.imageConfig(),
  'oci/image-manifest.json': oci.imageManifest(),
  'oci/image-index.json': oci.imageIndex(),
  'oci/error.json': oci.errors(),
  'oci/catalog.json': oci.catalog(),
  'oci/tag-list.json': oci.tagsList(),
  'stori/server-config.json': stori.serverConfig(),
}
