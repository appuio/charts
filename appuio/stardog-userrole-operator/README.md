# stardog-userrole-operator

![Version: 0.2.2](https://img.shields.io/badge/Version-0.2.2-informational?style=flat-square) ![AppVersion: 0.2.2](https://img.shields.io/badge/AppVersion-0.2.2-informational?style=flat-square)

Stardog User and Role Operator

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install stardog-userrole-operator appuio/stardog-userrole-operator
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| config.clusterAccess.enabled | bool | `true` | Enable cluster access to operator |
| config.env.DISABLED_ENVIRONMENTS | string | `""` | semicolon separated ex: "stardog-test;stardog-prod" |
| config.env.RECONCILIATION_FREQUENCY | string | `"30s"` | Reconciliation frequency of CRD |
| config.env.RECONCILIATION_FREQUENCY_ON_ERROR | string | `"5m"` | Reconciliation frequency of CRD on error |
| image.pullPolicy | string | `"Always"` |  |
| image.registry | string | `"ghcr.io"` |  |
| image.repository | string | `"vshn/stardog-userrole-operator"` |  |
| image.tag | string | `"v0.2.2"` |  |
| imagePullSecrets | list | `[]` |  |
| leaderElection.enabled | bool | `false` | Enable leader election for multiple replicas |
| metrics.proxy.enabled | bool | `false` | Enable metrics via service behind a authenticated proxy |
| metrics.proxy.image | string | `"gcr.io/kubebuilder/kube-rbac-proxy:v0.12.0"` | Proxy image |
| metrics.proxy.port | int | `8443` |  |
| metrics.service.enabled | bool | `true` |  |
| metrics.service.port | int | `8080` |  |
| replicaCount | int | `1` |  |
| resources.limits.cpu | string | `"100m"` |  |
| resources.limits.memory | string | `"100Mi"` |  |
| resources.requests.cpu | string | `"30m"` |  |
| resources.requests.memory | string | `"20Mi"` |  |
| serviceAccount.create | bool | `true` |  |

## Source Code

* <https://github.com/vshn/stardog-userrole-operator>

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
