# exoip

![Version: 1.0.4](https://img.shields.io/badge/Version-1.0.4-informational?style=flat-square) ![AppVersion: v0.4.3](https://img.shields.io/badge/AppVersion-v0.4.3-informational?style=flat-square)

Deploy exoip as a DaemonSet

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install exoip appuio/exoip
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| exoscaleApiEndpoint | string | `"https://api.exoscale.ch/compute"` |  |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"docker.io/exoscale/exoip"` |  |
| image.tag | string | `"v0.4.3"` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| peerGroup | string | `"kubernetes1-ingress"` |  |
| resources | object | `{}` |  |
| tolerations | list | `[]` |  |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
