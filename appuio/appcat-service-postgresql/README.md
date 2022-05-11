# provider-postgresql

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

VSHN-opinionated PostgreSQL operator for AppCat

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install provider-postgresql appuio/provider-postgresql
```
```bash
kubectl apply -f https://github.com/vshn/appcat-service-postgresql/releases/download/provider-postgresql-0.1.0/crds.yaml
```

<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Handling CRDs

* Always upgrade the CRDs before upgrading the Helm release.
* Watch out for breaking changes in the Provider-Postgresql release notes.

<!---
The values below are generated with helm-docs!

Document your changes in values.yaml and let `make chart-docs` generate this section.
-->
## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` | Operator image pull policy If set to empty, then Kubernetes default behaviour applies. |
| image.registry | string | `"ghcr.io"` | Operator image registry |
| image.repository | string | `"vshn/appcat-service-postgresql"` | Operator image repository |
| image.tag | string | `"latest"` | Operator image tag |
| imagePullSecrets | list | `[]` | List of image pull secrets if custom image is behind authentication. |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| operator.args | list | `[]` | Overrides arguments passed to the entrypoint |
| podAnnotations | object | `{}` | Annotations to add to the Pod spec. |
| podSecurityContext | object | `{}` | Security context to add to the Pod spec. |
| replicaCount | int | `1` | How many operator pods should run. Follower pods reduce interruption time as they're on hot standby when leader is unresponsive. |
| resources | object | `{}` |  |
| securityContext | object | `{}` | Container security context |
| service.port | int | `80` | Service port number |
| service.type | string | `"ClusterIP"` | Service type |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| serviceAccount.name | string | `""` | The name of the service account to use. If not set and `.create` is `true`, a name is generated using the fullname template |
| tolerations | list | `[]` |  |
| webhook.caBundle | string | `""` | Certificate in PEM format for the ValidatingWebhookConfiguration. |
| webhook.certificate | string | `""` | Certificate in PEM format for the TLS secret. |
| webhook.privateKey | string | `""` | Private key in PEM format for the TLS secret. |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
