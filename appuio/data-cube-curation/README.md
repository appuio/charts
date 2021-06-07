# data-cube-curation

![Version: 0.3.1](https://img.shields.io/badge/Version-0.3.1-informational?style=flat-square) ![AppVersion: 0.6.1](https://img.shields.io/badge/AppVersion-0.6.1-informational?style=flat-square)

RDF Data Cube curation service

**Homepage:** <https://github.com/zazuko/data-cube-curation>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install data-cube-curation appuio/data-cube-curation
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| dataCubeCuration.auth.audience | string | `""` |  |
| dataCubeCuration.auth.clientId | string | `""` |  |
| dataCubeCuration.auth.issuer | string | `""` |  |
| dataCubeCuration.gitlab.branch | string | `""` |  |
| dataCubeCuration.gitlab.existingSecret | string | `nil` |  |
| dataCubeCuration.gitlab.pipeline | string | `""` |  |
| dataCubeCuration.gitlab.token | string | `""` |  |
| dataCubeCuration.s3.accessKey | string | `""` |  |
| dataCubeCuration.s3.bucket | string | `""` |  |
| dataCubeCuration.s3.endpoint | string | `""` |  |
| dataCubeCuration.s3.existingSecret | string | `nil` |  |
| dataCubeCuration.s3.secretKey | string | `""` |  |
| dataCubeCuration.sparql.endpoint | string | `""` |  |
| dataCubeCuration.sparql.existingSecret | string | `nil` |  |
| dataCubeCuration.sparql.graphStoreEndpoint | string | `""` |  |
| dataCubeCuration.sparql.readModelEndpoint | string | `""` |  |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"docker.io/zazuko/datacube"` |  |
| image.tag | string | `"{{ .Chart.AppVersion }}"` |  |
| imagePullSecrets | list | `[]` |  |
| ingress.annotations | object | `{}` |  |
| ingress.enabled | bool | `true` |  |
| ingress.hosts[0].host | string | `"data-cube-curation.localhost"` |  |
| ingress.hosts[0].paths[0] | string | `"/"` |  |
| ingress.tls | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources | object | `{}` |  |
| securityContext.runAsUser | int | `65534` |  |
| service.port | int | `80` |  |
| service.type | string | `"ClusterIP"` |  |
| tolerations | list | `[]` |  |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
