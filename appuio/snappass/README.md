# snappass

![Version: 0.2.15](https://img.shields.io/badge/Version-0.2.15-informational?style=flat-square) ![AppVersion: v1.4.2](https://img.shields.io/badge/AppVersion-v1.4.2-informational?style=flat-square)

A Helm chart for SnapPass

**Homepage:** <https://github.com/pinterest/snappass>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install snappass appuio/snappass
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"docker.io/samueldg/snappass"` |  |
| image.tag | string | `"latest@sha256:3987195edbe623a9a7400fbb68318059d4fc8c20d48f73a6bd52598e56b3b9b3"` |  |
| imagePullSecrets | list | `[]` |  |
| ingress.annotations | object | `{}` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hosts[0].host | string | `"chart-example.local"` |  |
| ingress.hosts[0].paths | list | `[]` |  |
| ingress.tls | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| redis.enabled | bool | `true` | Enables the Redis Helm chart dependency |
| replicaCount | int | `1` |  |
| resources.limits.cpu | string | `"500m"` |  |
| resources.limits.memory | string | `"256Mi"` |  |
| resources.requests.cpu | string | `"100m"` |  |
| resources.requests.memory | string | `"128Mi"` |  |
| securityContext.capabilities.drop[0] | string | `"ALL"` |  |
| securityContext.readOnlyRootFilesystem | bool | `true` |  |
| securityContext.runAsNonRoot | bool | `true` |  |
| securityContext.runAsUser | int | `1000` |  |
| service.port | int | `80` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.create | bool | `false` |  |
| serviceAccount.name | string | `nil` |  |
| snappass.debug.enabled | bool | `false` | If the debug mode should be enabled |
| snappass.existingSecret | string | `""` | Existing secret with a `secretKey` field, used to sign cookies |
| snappass.redis.db | int | `0` | Existing secret with a `redis-password` field, used to authenticate against Redis |
| snappass.redis.host | string | `""` |  |
| snappass.redis.password | string | `""` |  |
| snappass.redis.port | int | `6379` |  |
| snappass.redis.prefix | string | `""` |  |
| snappass.secretKey | string | `""` | Secret key used to sign cookies |
| snappass.ssl.enabled | bool | `false` | Enables SSL on application level |
| snappass.staticUrl | string | `""` | Location of static assets |
| snappass.urlPrefix | string | `""` | URL Prefix (e.g. when running SnapPass behind a reverse proxy) |
| tolerations | list | `[]` |  |

## Source Code

* <https://github.com/pinterest/snappass>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://raw.githubusercontent.com/bitnami/charts/defb094c658024e4aa8245622dab202874880cbc/bitnami/ | redis | 10.7.4 |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
