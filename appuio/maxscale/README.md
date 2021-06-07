# maxscale

![Version: 1.1.2](https://img.shields.io/badge/Version-1.1.2-informational?style=flat-square) ![AppVersion: 2.2.20](https://img.shields.io/badge/AppVersion-2.2.20-informational?style=flat-square)

MaxScale Helm Chart

**Homepage:** <https://mariadb.com/products/technology/maxscale>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install maxscale appuio/maxscale
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| db1Address | string | `"db1.example.com"` | IP or Hostname of the 1. Galera node |
| db2Address | string | `"db2.example.com"` | IP or Hostname of the 2. Galera node |
| db3Address | string | `"db3.example.com"` | IP or Hostname of the 3. Galera node |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"docker.io/appuio/maxscale"` |  |
| image.tag | string | `"2.2.20"` |  |
| metrics.enabled | bool | `false` |  |
| metrics.image.pullPolicy | string | `"IfNotPresent"` |  |
| metrics.image.repository | string | `"skord/maxscale_exporter"` |  |
| metrics.image.tag | string | `"latest"` |  |
| metrics.resources | object | `{}` |  |
| monitorPassword | string | `"secret"` | MaxScale monitor user password |
| monitorUser | string | `"maxscale"` | MaxScale monitor user name |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources | object | `{}` |  |
| service.port | int | `3306` |  |
| service.type | string | `"ClusterIP"` |  |
| servicePassword | string | `"secret"` | MaxScale service user password |
| serviceUser | string | `"maxscale"` | MaxScale service user name |
| tolerations | list | `[]` |  |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
