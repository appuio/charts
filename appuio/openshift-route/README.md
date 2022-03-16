# openshift-route

![Version: 1.1.4](https://img.shields.io/badge/Version-1.1.4-informational?style=flat-square) ![AppVersion: v1](https://img.shields.io/badge/AppVersion-v1-informational?style=flat-square)

A Helm chart for OpenShift that simply creates a Route object

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install openshift-route appuio/openshift-route
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

This chart requires that `service.name` is set (`--set service.name=my-backend-service`).

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| alternateBackends | list | `[]` | Up to 4 additional backends can be enabled to support the route. Consult the documentation. |
| annotations | object | `{}` | Annotations on the route object. |
| fullnameOverride | string | `""` | If you need to give a specific name to the route, override this, as it affects the hostname generated if `host` is left empty. |
| host | string | `""` | The hostname that should be used. If left empty, OpenShift will generate one for you with defaults. |
| labels | object | `{}` | Additional labels on the route object. |
| nameOverride | string | `""` | Overrides the name after the Chart name. Affects route name and hostname generation. |
| path | string | `""` | Subpath of the route. |
| service.name | string | `""` | The backend service name of the route. **Required** |
| service.targetPort | string | `"http"` | The port of the backend service. The port is aware of the named ports of the service, so it can be a name too. |
| service.weight | int | `100` | Weight of the service endpoint. |
| tls.caCertificate | string | `""` | CA (chain) of the certificate in PEM format. |
| tls.certificate | string | `""` | Certificate in PEM format. |
| tls.destinationCACertificate | string | `""` | CA of the backend in PEM format. Only relevant when `termination=reencrypt`. |
| tls.enabled | bool | `false` | Whether to secure the Route with TLS. |
| tls.insecureEdgeTerminationPolicy | string | `"Redirect"` | What to do with non-HTTPS traffic (usually port 80). |
| tls.key | string | `""` | Certificate private key in PEM format. |
| tls.termination | string | `"edge"` | TLS termination of the route. |
| wildcardPolicy | string | `"None"` | Wildcard Policy of the route. |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
