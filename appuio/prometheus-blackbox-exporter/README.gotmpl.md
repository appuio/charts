<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Configuration

The following table lists the configurable parameters of the Blackbox-Exporter chart and their default values.

|               Parameter                |                   Description                   |            Default            |
| -------------------------------------- | ----------------------------------------------- | ----------------------------- |
| `config`                               | Prometheus blackbox configuration               | {}                            |
| `configmapReload.name`                 | configmap-reload container name                 | `configmap-reload`            |
| `configmapReload.image.repository`     | configmap-reload container image repository     | `jimmidyson/configmap-reload` |
| `configmapReload.image.tag`            | configmap-reload container image tag            | `v0.3.0`                      |
| `configmapReload.image.pullPolicy`     | configmap-reload container image pull policy    | `IfNotPresent`                |
| `configmapReload.extraArgs`            | Additional configmap-reload container arguments | `{}`                          |
| `configmapReload.extraConfigmapMounts` | Additional configmap-reload configMap mounts    | `[]`                          |
| `configmapReload.resources`            | configmap-reload pod resource requests & limits | `{}`                          |
| `extraArgs`                            | Optional flags for blackbox                     | `[]`                          |
| `image.repository`                     | container image repository                      | `prom/blackbox-exporter`      |
| `image.tag`                            | container image tag                             | `v0.16.0`                     |
| `image.pullPolicy`                     | container image pull policy                     | `IfNotPresent`                |
| `ingress.annotations`                  | Ingress annotations                             | None                          |
| `ingress.enabled`                      | Enables Ingress                                 | `false`                       |
| `ingress.hosts`                        | Ingress accepted hostnames                      | None                          |
| `ingress.tls`                          | Ingress TLS configuration                       | None                          |
| `nodeSelector`                         | node labels for pod assignment                  | `{}`                          |
| `podAnnotations`                       | annotations to add to each pod                  | `{}`                          |
| `resources`                            | pod resource requests & limits                  | `{}`                          |
| `restartPolicy`                        | container restart policy                        | `Always`                      |
| `service.type`                         | type of service to create                       | `ClusterIP`                   |
| `service.port`                         | port for the blackbox http service              | `9115`                        |
| `service.externalIPs`                  | list of external ips                            | []                            |
