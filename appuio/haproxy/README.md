# haproxy

![Version: 1.7.2](https://img.shields.io/badge/Version-1.7.2-informational?style=flat-square) ![AppVersion: 2.5.1](https://img.shields.io/badge/AppVersion-2.5.1-informational?style=flat-square)

A Helm chart for HAProxy which can be customized by a config map.

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install haproxy appuio/haproxy
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Introduction

The Debian based image was choosen because it already contains the ca-certificates which makes it
easy to use HAProxy as a reverse TLS Proxy.

## Configuration

The configuration of HAProxy is derived through a configmap which in turn can be configured by adjusting the corresponding values in values.yaml.
The present example is a configmap to configure HAProxy as a TLS reverse proxy to communicate with an LDAP backend. For different purposes you can
create your own configmap and values.

### General parameters
| Parameter               | Description            | Default
|---                      | ---                    | ---
| `replicaCount`          | Number of replicas for the proxy. | `1`
| `image.registry`        | Registry of the image | `docker.io`
| `image.repository`      | Repository of the image | `haproxytech/haproxy-debian`
| `image.tag`             | Tag of the image | `2.3.5`
| `image.pullPolicy`      | Determines when the image is pulled | `IfNotPresent`
| `ingress.enabled`       | If an ingress object should be created | `false`
| `ingress.annotations`   | Annotations to set on the ingress object |
| `ingress.hosts`         | Array of host names which the ingress should resolve |
| `ingress.tls.enabled`   | If TLS should be enabled on the ingress |
| `ingress.tls.secretName`| Name of the secret containing the TLS certificate and key |
| `service.port` | Kubernetes Service port (haproxy) | `30636`
| `service.type` | Kubernetes Service type (haproxy) | `ClusterIP`
| `service.additionalPorts` | Additional Ports to be added to the haproxy service (e.g. for sidecars) | `[]`
| `sidecarContainers` | Extra sidecar containers to add to the deployment | `[]`
| `affinity`          | Affinity for pod assignment | `{}` (evaluated as a template)
| `haproxy.frontendPort` | Port for the Proxy to listen on the frontend | `30636`
| `haproxy.config`       | Suffix of the used config-map | `ldap-tls`
| `haproxy.check.existingSecret.name` | Use existing secret with authentication credentials for check commands. |
| `haproxy.check.existingSecret.key` | Key in secret which contains the authentication credentials for the check command | `auth-credentials`
| `haproxy.securityContext` | Security context for the pod to be used | `{}`
| `nodeSelector` | Pod node selector | `{}`
| `tolerations` | Pod tolerations | `[]`
| `affinity` | Pod affinity rules | `{}`
| `podDisruptionBudget.create` | If true, create a pod disruption budget for pods. | `true`
| `podDisruptionBudget.minAvailable` | Minimum number / percentage of pods that should remain scheduled | `nil`
| `podDisruptionBudget.maxUnavailable` | Maximum number / percentage of pods that may be made unavailable | `1`
| `resources` | HAProxy resource requests/limits | `{ limits: { cpu: 100m, memory: 128Mi }, requests: { cpu: 100m, memory 128Mi } }`

### Metrics
| Parameter              | Description            | Default
|---                     | ---                    | ---
| `metrics.enabled`      | Enables /stats and /metrics endpoint in haproxy | `true`
| `metrics.serviceMonitor.enabled` | If true, creates a Prometheus Operator ServiceMonitor (also requires `metrics.enabled` to be true) | `false`
| `metrics.serviceMonitor.namespace` | Optional namespace which Prometheus is running in |
| `metrics.serviceMonitor.interval` | How frequently to scrape metrics (use by default, falling back to Prometheus' default) |
| `metrics.serviceMonitor.selector` | Default to kube-prometheus install (CoreOS recommended), but should be set according to Prometheus install | `{ prometheus: kube-prometheus }`
| `metrics.serviceMonitor.relabelings` | ServiceMonitor relabelings. Value is evaluated as a template | `[]`
| `metrics.serviceMonitor.metricRelabelings` | ServiceMonitor metricRelabelings. Value is evaluated as a template | `[]`
| `metrics.service.type` | Kubernetes Service type (haproxy metrics) | `ClusterIP`
| `metrics.service.annotations` | Annotations for the metrics service | `{}`
| `metrics.service.labels` | Additional labels for the metrics service | `{}`
| `metrics.service.loadBalancerIP` | LoadBalancer IP if haproxy metrics service type is LoadBalancer | `{}`
| `metrics.prometheusRule.enabled` | Set this to true to create prometheusRules for Prometheus operator | `false`
| `metrics.prometheusRule.additionalLabels` | Additional labels that can be used so prometheusRules will be discovered by Prometheus | `{}`
| `metrics.prometheusRule.namespace` | Namespace where prometheusRules resource should be created | Same namespace as haproxy
| `metrics.prometheusRule.rules` | Rules to be created, check values for an example. | `[]`

### ldap-tls

Set `haproxy.config` to `ldap-tls` to use the LDAP TLS configuration (default).

| Parameter              | Description            | Default
|---                     | ---                    | ---
| `haproxy.ldapTls.backend` | Service the proxy should connect to |
| `haproxy.ldapTls.certificatePath` | Path to the corresponding root-certificate | `/etc/ssl/certs/ca-certificates.crt`

### galera

Set `haproxy.config` to `galera` to use the Galera configuration.

| Parameter              | Description            | Default
|---                     | ---                    | ---
| `haproxy.galera.balance` | What balance mode HAProxy should use | source
| `haproxy.galera.timeout.connect` | Set the maximum time to wait for a connection attempt to a server to succeed. | `5s`
| `haproxy.galera.timeout.client` | The inactivity timeout applies when the client is expected to acknowledge or send data | `10800s`
| `haproxy.galera.timeout.server` | The inactivity timeout applies when the server is expected to acknowledge or send data | `10800s`
| `haproxy.galera.check.enabled` | If check should be enabled | `true`
| `haproxy.galera.check.mysql.enabled` | If mysql-check should be enabled (requires check.enabled) | `true`
| `haproxy.galera.check.mysql.user` | The database user to use for mysql-check | `haproxy`
| `haproxy.galera.nodes` | List of Galera nodes | `[]`
| `haproxy.galera.nodes.address` | Address of the Galera node |
| `haproxy.galera.nodes.port` | Port of the Galera node | `3306`
| `haproxy.galera.nodes.backup` | If this node should be used as a backup node | `true`
| `haproxy.galera.metrics.enabled`| If the metric endpoint of the Galera backends should be exposed in haproxy | `false`
| `haproxy.galera.metrics.exposeLoadbalancer`| If the metric endpoint of the Galera backends should be exposed in the haproxy service | `true`

To use `mysql-check`, configure a user `haproxy` in the database (see: http://cbonte.github.io/haproxy-dconv/1.9/configuration.html#4-option%20mysql-check for more information).

### galerak8s

Set `haproxy.config` to `galerak8s` to use the Galera configuration with DNS service discovery. See https://www.haproxy.com/de/blog/dns-service-discovery-haproxy/

| Parameter              | Description            | Default
|---                     | ---                    | ---
| `haproxy.galerak8s.balance` | What balance mode HAProxy should use | `first`
| `haproxy.galerak8s.timeout.connect` | Set the maximum time to wait for a connection attempt to a server to succeed. | `5s`
| `haproxy.galerak8s.timeout.client` | The inactivity timeout applies when the client is expected to acknowledge or send data | `10800s`
| `haproxy.galerak8s.timeout.server` | The inactivity timeout applies when the server is expected to acknowledge or send data | `10800s`
| `haproxy.galerak8s.check.enabled` | If check should be enabled | `true`
| `haproxy.galerak8s.check.mysql.enabled` | If mysql-check should be enabled (requires check.enabled) | `true`
| `haproxy.galerak8s.check.mysql.user` | The database user to use for mysql-check | `haproxy`
| `haproxy.galerak8s.dnsservicename` | The DNS Record for service discovery | `mycluster-mariadb-galera-headless`
| `haproxy.galerak8s.nodeCount` | Max number of nodes in the backend | `3`
| `haproxy.galerak8s.port` | Port of the Galera node | `3306`
| `haproxy.galerak8s.metrics.enabled`| If the metric endpoint of the Galera backends should be exposed in haproxy | `false`
| `haproxy.galerak8s.metrics.exposeLoadbalancer`| If the metric endpoint of the Galera backends should be exposed in the haproxy service | `true`

### redisk8s

Set `haproxy.config` to `redisk8s` to use the Redis configuration with DNS service discovery. See https://www.haproxy.com/de/blog/dns-service-discovery-haproxy/

| Parameter              | Description            | Default
|---                     | ---                    | ---
| `haproxy.redisk8s.timeout.connect` | Set the maximum time to wait for a connection attempt to a server to succeed. | `5s`
| `haproxy.redisk8s.timeout.client` | The inactivity timeout applies when the client is expected to acknowledge or send data | `60s`
| `haproxy.redisk8s.timeout.server` | The inactivity timeout applies when the server is expected to acknowledge or send data | `60s`
| `haproxy.redisk8s.check.enabled` | If check should be enabled | `true`
| `haproxy.redisk8s.check.redis.auth` | Authentication credentials for Redis. `haproxy.check.existingSecret` overwrites this |
| `haproxy.redisk8s.dnsservicename` | The DNS Record for service discovery | `redis-access-headless`
| `haproxy.redisk8s.nodeCount` | Max number of nodes in the backend | `3`
| `haproxy.redisk8s.port` | Port of the Galera node | `6379`
| `haproxy.redisk8s.metrics.enabled`| If the metric endpoint of the Redis backends should be exposed in haproxy | `false`
| `haproxy.redisk8s.metrics.exposeLoadbalancer`| If the metric endpoint of the Redis backends should be exposed in the haproxy service | `true`

## Galera and Redis metrics

By setting the `haproxy.galerak8s.metrics.enabled` or `haproxy.redisk8s.metrics.enabled` parameters, you can expose the metrics of the backend pods.

For Galera the endpoints are:
* `<haproxy_service_ip>:9090/metrics/mariadb-0/metrics`
* `<haproxy_service_ip>:9090/metrics/mariadb-1/metrics`
* `...`

For Redis the endpoints are:
* `<haproxy_service_ip>:9090/metrics/redis-0/metrics`
* `<haproxy_service_ip>:9090/metrics/redis-1/metrics`
* `...`

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
