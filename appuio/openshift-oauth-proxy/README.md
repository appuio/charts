# openshift-oauth-proxy

![Version: 0.2.3](https://img.shields.io/badge/Version-0.2.3-informational?style=flat-square) ![AppVersion: 1.1.0](https://img.shields.io/badge/AppVersion-1.1.0-informational?style=flat-square)

OpenShift oauth-proxy

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install openshift-oauth-proxy appuio/openshift-oauth-proxy
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| extraArgs | string | `nil` | Extra arguments to pass to the oauth-proxy |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"docker.io/openshift/oauth-proxy"` |  |
| image.tag | string | `"v1.1.0"` |  |
| proxy.openshiftSar | string | `"{\"resource\": \"namespaces\", \"verb\": \"get\"}"` | OpenShift Subject Access Review |
| proxy.upstream | string | `"http://localhost:9090"` | Upstream url |
| resources.limits.cpu | string | `"200m"` |  |
| resources.limits.memory | string | `"128Mi"` |  |
| resources.requests.cpu | string | `"50m"` |  |
| resources.requests.memory | string | `"32Mi"` |  |
| route.annotations | object | `{"kubernetes.io/tls-acme":"true"}` | Add annotations to route |
| route.host | string | `""` | Define route host |
| route.tls | bool | `true` | Enable tls for the route |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
