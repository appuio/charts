# redis

![Version: 1.3.1](https://img.shields.io/badge/Version-1.3.1-informational?style=flat-square) ![AppVersion: 6.2.1](https://img.shields.io/badge/AppVersion-6.2.1-informational?style=flat-square)

Open source, advanced key-value store. It is often referred to as a data structure server since keys can contain strings, hashes, lists, sets and sorted sets.

**Homepage:** <https://github.com/appuio/charts/tree/master/appuio/redis>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install redis appuio/redis
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cluster.enabled | bool | `true` | Enable clustering |
| cluster.slaveCount | int | `2` | Number of slaves |
| clusterDomain | string | `"cluster.local"` | Specifies the Kubernetes Cluster's Domain Name. |
| configmap | string | `"# Enable AOF https://redis.io/topics/persistence#append-only-file\nappendonly yes\n# Disable RDB persistence, AOF persistence already enabled.\nsave \"\""` | Redis(TM) config file |
| containerSecurityContext | object | `{"enabled":true,"runAsUser":1001}` | Container Security Context |
| global.redis | object | `{}` |  |
| image.pullPolicy | string | `"IfNotPresent"` | Redis image pull policy Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent' |
| image.registry | string | `"docker.io"` | Redis image registry |
| image.repository | string | `"bitnami/redis"` | Redis image repository |
| image.tag | string | `"6.2.1-debian-10-r36"` | Redis image tag (immutable tags are recommended) |
| master.affinity | object | `{}` |  |
| master.command | string | `"/run.sh"` | Redis(TM) command arguments |
| master.configmap | string | `nil` | Additional Redis(TM) configuration for the master nodes |
| master.customLivenessProbe | object | `{}` |  |
| master.customReadinessProbe | object | `{}` |  |
| master.disableCommands | list | `["FLUSHDB","FLUSHALL"]` | Comma-separated list of Redis(TM) commands to disable |
| master.extraEnvVars | list | `[]` | An array to add extra env vars |
| master.extraEnvVarsCM | list | `[]` | ConfigMap with extra env vars: |
| master.extraEnvVarsSecret | list | `[]` | Secret with extra env vars: |
| master.extraFlags | list | `[]` | Redis(TM) additional command line flags |
| master.hostAliases | list | `[]` | Deployment pod host aliases |
| master.livenessProbe.enabled | bool | `true` |  |
| master.livenessProbe.failureThreshold | int | `5` |  |
| master.livenessProbe.initialDelaySeconds | int | `5` |  |
| master.livenessProbe.periodSeconds | int | `5` |  |
| master.livenessProbe.successThreshold | int | `1` |  |
| master.livenessProbe.timeoutSeconds | int | `5` |  |
| master.persistence.accessModes[0] | string | `"ReadWriteOnce"` |  |
| master.persistence.enabled | bool | `true` | Enable persistence using Persistent Volume Claims |
| master.persistence.matchExpressions | object | `{}` |  |
| master.persistence.matchLabels | object | `{}` | Persistent Volume selectors |
| master.persistence.path | string | `"/data"` | The path the volume will be mounted at, useful when using different Redis(TM) images. |
| master.persistence.size | string | `"8Gi"` |  |
| master.persistence.subPath | string | `""` | The subdirectory of the volume to mount to. |
| master.persistence.volumes | string | `nil` |  |
| master.podAnnotations | object | `{}` | Redis(TM) Master additional pod annotations |
| master.podLabels | object | `{}` | Redis(TM) Master additional pod labels |
| master.preExecCmds | string | `""` | Additional commands to run prior to starting Redis(TM) |
| master.priorityClassName | string | `nil` | Redis(TM) Master pod priorityClassName |
| master.readinessProbe.enabled | bool | `true` |  |
| master.readinessProbe.failureThreshold | int | `5` |  |
| master.readinessProbe.initialDelaySeconds | int | `5` |  |
| master.readinessProbe.periodSeconds | int | `5` |  |
| master.readinessProbe.successThreshold | int | `1` |  |
| master.readinessProbe.timeoutSeconds | int | `1` |  |
| master.service.annotations | object | `{}` |  |
| master.service.externalTrafficPolicy | string | `"Cluster"` | External traffic policy (when service type is LoadBalancer) |
| master.service.labels | object | `{}` | Provide any additional labels |
| master.service.loadBalancerIP | string | `nil` |  |
| master.service.port | int | `6379` | Redis(TM) Master Service port |
| master.service.type | string | `"ClusterIP"` | Redis(TM) Master Service type |
| master.shareProcessNamespace | bool | `false` |  |
| master.statefulset.annotations | object | `{}` |  |
| master.statefulset.labels | object | `{}` |  |
| master.statefulset.updateStrategy | string | `"RollingUpdate"` |  |
| master.statefulset.volumeClaimTemplates.annotations | object | `{}` |  |
| master.statefulset.volumeClaimTemplates.labels | object | `{}` |  |
| metrics.enabled | bool | `false` | Enable Prometheus Exporter  |
| metrics.image.pullPolicy | string | `"IfNotPresent"` | Exporter pull policy |
| metrics.image.registry | string | `"docker.io"` | Exporter image registry |
| metrics.image.repository | string | `"bitnami/redis-exporter"` | Exporter image repository |
| metrics.image.tag | string | `"1.20.0-debian-10-r16"` | Exporter image tag |
| metrics.podAnnotations."prometheus.io/port" | string | `"9121"` |  |
| metrics.podAnnotations."prometheus.io/scrape" | string | `"true"` |  |
| metrics.priorityClassName | string | `nil` | Metrics exporter pod priorityClassName |
| metrics.prometheusRule.additionalLabels | object | `{}` |  |
| metrics.prometheusRule.enabled | bool | `false` |  |
| metrics.prometheusRule.namespace | string | `""` |  |
| metrics.prometheusRule.rules | list | `[]` |  |
| metrics.redisTargetHost | string | `"localhost"` | A way to specify an alternative redis hostname |
| metrics.service.annotations | object | `{}` |  |
| metrics.service.externalTrafficPolicy | string | `"Cluster"` | External traffic policy (when service type is LoadBalancer) |
| metrics.service.labels | object | `{}` |  |
| metrics.service.type | string | `"ClusterIP"` |  |
| metrics.serviceMonitor.enabled | bool | `false` |  |
| metrics.serviceMonitor.metricRelabelings | list | `[]` |  |
| metrics.serviceMonitor.relabelings | list | `[]` |  |
| metrics.serviceMonitor.selector.prometheus | string | `"kube-prometheus"` |  |
| networkPolicy.enabled | bool | `false` | Specifies whether a NetworkPolicy should be created |
| networkPolicy.ingressNSMatchLabels | object | `{}` |  |
| networkPolicy.ingressNSPodMatchLabels | object | `{}` |  |
| password | string | `""` | Redis(TM) password (both master and slave). Defaults to a random 10-character alphanumeric string if not set and usePassword is true. |
| persistence.existingClaim | string | `nil` |  |
| podDisruptionBudget.enabled | bool | `false` | Create a disruption budget |
| podDisruptionBudget.minAvailable | int | `1` |  |
| podSecurityPolicy.create | bool | `false` | Specifies whether a PodSecurityPolicy should be created |
| rbac.create | bool | `false` | Specifies whether RBAC resources should be created |
| rbac.role.rules | list | `[]` | Rules to create. It follows the role specification |
| redisPort | int | `6379` | Redis(TM) port |
| securityContext | object | `{"enabled":true,"fsGroup":1001}` | Redis(TM) pod Security Context |
| sentinel.cleanDelaySeconds | int | `5` | Delay seconds when cleaning nodes IPs. When starting it will clean the sentinels IP (RESET "*") in all the nodes. This is the delay time before sending the command to the next node. |
| sentinel.configmap | string | `nil` | Additional Redis(TM) configuration for the sentinel nodes |
| sentinel.customLivenessProbe | object | `{}` |  |
| sentinel.customReadinessProbe | object | `{}` |  |
| sentinel.downAfterMilliseconds | int | `20000` |  |
| sentinel.enabled | bool | `false` | Enable sentinel This will disable the master and slave services and create one redis service with ports to the sentinel and the redis instances |
| sentinel.extraEnvVars | list | `[]` | An array to add extra env var to the sentinel node configurations |
| sentinel.extraEnvVarsCM | list | `[]` | ConfigMap with extra env vars: |
| sentinel.extraEnvVarsSecret | list | `[]` | Secret with extra env vars: |
| sentinel.failoverTimeout | int | `18000` |  |
| sentinel.forceUpdate | bool | `false` | Force update the StatefulSet. If enabled the chart will recreate the StatefulSet without touching the Pods (cascade orphan), allowing you to update locked field, such as PVC size. Currently only supports changes in PVC size. |
| sentinel.forceUpdateKubectlImage | string | `"quay.io/bitnami/kubectl"` | Image used to recreate the StatefulSet. Needs to have `kubectl` and `jq` in its `$PATH`. |
| sentinel.image.pullPolicy | string | `"IfNotPresent"` | Sentinel image pull policy Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent' |
| sentinel.image.registry | string | `"docker.io"` | Sentinel image registry |
| sentinel.image.repository | string | `"bitnami/redis-sentinel"` | Sentinel image repository |
| sentinel.image.tag | string | `"6.2.1-debian-10-r35"` | Sentinel image tag |
| sentinel.initialCheckTimeout | int | `5` |  |
| sentinel.livenessProbe.enabled | bool | `true` |  |
| sentinel.livenessProbe.failureThreshold | int | `5` |  |
| sentinel.livenessProbe.initialDelaySeconds | int | `5` |  |
| sentinel.livenessProbe.periodSeconds | int | `5` |  |
| sentinel.livenessProbe.successThreshold | int | `1` |  |
| sentinel.livenessProbe.timeoutSeconds | int | `5` |  |
| sentinel.masterSet | string | `"mymaster"` |  |
| sentinel.metrics.enabled | bool | `false` | Enable Prometheus Exporter / Metrics for Redis Sentinel Exporter |
| sentinel.metrics.image.pullPolicy | string | `"IfNotPresent"` | Sentinel Exporter pull policy |
| sentinel.metrics.image.registry | string | `"docker.io"` | Sentinel Exporter image registry |
| sentinel.metrics.image.repository | string | `"bitnami/redis-sentinel-exporter"` | Sentinel Exporter image repository |
| sentinel.metrics.image.tag | string | `"1.7.1-debian-10-r109"` | Sentinel Exporter image tag |
| sentinel.metrics.resources | object | `{"limits":{},"requests":{}}` | Sentinel metrics exporter resource requests and limits |
| sentinel.metrics.service.annotations | object | `{}` |  |
| sentinel.metrics.service.externalTrafficPolicy | string | `"Cluster"` | External traffic policy (when service type is LoadBalancer) |
| sentinel.metrics.service.labels | object | `{}` |  |
| sentinel.metrics.service.port | int | `9355` | Metrics port |
| sentinel.metrics.service.type | string | `"ClusterIP"` | Metrics Service type |
| sentinel.metrics.serviceMonitor.enabled | bool | `false` |  |
| sentinel.metrics.serviceMonitor.selector.prometheus | string | `"kube-prometheus"` |  |
| sentinel.parallelSyncs | int | `1` |  |
| sentinel.port | int | `26379` |  |
| sentinel.preExecCmds | string | `""` | Additional commands to run prior to starting Redis(TM) node with sentinel |
| sentinel.quorum | int | `2` |  |
| sentinel.readinessProbe.enabled | bool | `true` |  |
| sentinel.readinessProbe.failureThreshold | int | `5` |  |
| sentinel.readinessProbe.initialDelaySeconds | int | `5` |  |
| sentinel.readinessProbe.periodSeconds | int | `5` |  |
| sentinel.readinessProbe.successThreshold | int | `1` |  |
| sentinel.readinessProbe.timeoutSeconds | int | `1` |  |
| sentinel.service.annotations | object | `{}` |  |
| sentinel.service.externalTrafficPolicy | string | `"Cluster"` | External traffic policy (when service type is LoadBalancer) |
| sentinel.service.labels | object | `{}` |  |
| sentinel.service.loadBalancerIP | string | `nil` |  |
| sentinel.service.redisPort | int | `6379` |  |
| sentinel.service.sentinelPort | int | `26379` |  |
| sentinel.service.type | string | `"ClusterIP"` | Redis(TM) Sentinel Service type |
| sentinel.staticID | bool | `false` |  |
| sentinel.usePassword | bool | `true` | Require password authentication on the sentinel itself |
| serviceAccount.create | bool | `false` | Specifies whether a ServiceAccount should be created |
| serviceAccount.name | string | `nil` | The name of the ServiceAccount to use. If not set and create is true, a name is generated using the fullname template |
| slave.affinity | object | `{}` | Redis(TM) Slave pod/node affinity/anti-affinity |
| slave.command | string | `"/run.sh"` | Can be used to specify command line arguments |
| slave.configmap | string | `nil` | Additional Redis(TM) configuration for the slave nodes |
| slave.customLivenessProbe | object | `{}` |  |
| slave.customReadinessProbe | object | `{}` |  |
| slave.disableCommands | list | `["FLUSHDB","FLUSHALL"]` | List of Redis(TM) commands to disable |
| slave.extraEnvVars | list | `[]` | An array to add extra env vars |
| slave.extraEnvVarsCM | list | `[]` | ConfigMap with extra env vars: |
| slave.extraEnvVarsSecret | list | `[]` | Secret with extra env vars: |
| slave.extraFlags | list | `[]` | Redis(TM) extra flags |
| slave.hostAliases | list | `[]` | Deployment pod host aliases |
| slave.livenessProbe.enabled | bool | `true` |  |
| slave.livenessProbe.failureThreshold | int | `5` |  |
| slave.livenessProbe.initialDelaySeconds | int | `30` |  |
| slave.livenessProbe.periodSeconds | int | `10` |  |
| slave.livenessProbe.successThreshold | int | `1` |  |
| slave.livenessProbe.timeoutSeconds | int | `5` |  |
| slave.persistence.accessModes[0] | string | `"ReadWriteOnce"` |  |
| slave.persistence.enabled | bool | `true` | Enable persistence using Persistent Volume Claims |
| slave.persistence.matchExpressions | object | `{}` |  |
| slave.persistence.matchLabels | object | `{}` |  |
| slave.persistence.path | string | `"/data"` | The path the volume will be mounted at. |
| slave.persistence.size | string | `"8Gi"` |  |
| slave.persistence.subPath | string | `""` | The subdirectory of the volume to mount to |
| slave.podAnnotations | object | `{}` | Redis(TM) slave pod Annotations |
| slave.podLabels | object | `{}` |  |
| slave.port | int | `6379` | Redis(TM) slave port |
| slave.preExecCmds | string | `""` | Additional commands to run prior to starting Redis(TM) |
| slave.priorityClassName | string | `nil` | Redis(TM) slave pod priorityClassName |
| slave.readinessProbe.enabled | bool | `true` |  |
| slave.readinessProbe.failureThreshold | int | `5` |  |
| slave.readinessProbe.initialDelaySeconds | int | `5` |  |
| slave.readinessProbe.periodSeconds | int | `10` |  |
| slave.readinessProbe.successThreshold | int | `1` |  |
| slave.readinessProbe.timeoutSeconds | int | `10` |  |
| slave.service.annotations | object | `{}` |  |
| slave.service.externalTrafficPolicy | string | `"Cluster"` | External traffic policy (when service type is LoadBalancer) |
| slave.service.labels | object | `{}` |  |
| slave.service.loadBalancerIP | string | `nil` |  |
| slave.service.port | int | `6379` | Redis(TM) port |
| slave.service.type | string | `"ClusterIP"` | Redis(TM) Slave Service type |
| slave.shareProcessNamespace | bool | `false` | Enable shared process namespace in a pod. |
| slave.spreadConstraints | object | `{}` | Kubernetes Spread Constraints for pod assignment |
| slave.statefulset.annotations | object | `{}` |  |
| slave.statefulset.labels | object | `{}` |  |
| slave.statefulset.updateStrategy | string | `"RollingUpdate"` |  |
| slave.statefulset.volumeClaimTemplates.annotations | object | `{}` |  |
| slave.statefulset.volumeClaimTemplates.labels | object | `{}` |  |
| sysctlImage.command | list | `[]` |  |
| sysctlImage.enabled | bool | `false` |  |
| sysctlImage.mountHostSys | bool | `false` |  |
| sysctlImage.pullPolicy | string | `"Always"` |  |
| sysctlImage.registry | string | `"docker.io"` |  |
| sysctlImage.repository | string | `"bitnami/bitnami-shell"` |  |
| sysctlImage.resources | object | `{}` |  |
| sysctlImage.tag | string | `"10"` |  |
| tls.authClients | bool | `true` | Whether to require clients to authenticate or not. |
| tls.certCAFilename | string | `nil` | CA Certificate filename |
| tls.certFilename | string | `nil` | Certificate filename |
| tls.certKeyFilename | string | `nil` | Certificate Key filename |
| tls.certificatesSecret | string | `nil` | Name of the Secret that contains the certificates |
| tls.enabled | bool | `false` | Enable TLS traffic |
| usePassword | bool | `true` | Use password authentication |
| usePasswordFile | bool | `false` | Mount secrets as files instead of environment variables |
| volumePermissions.enabled | bool | `false` |  |
| volumePermissions.image.pullPolicy | string | `"Always"` |  |
| volumePermissions.image.registry | string | `"docker.io"` |  |
| volumePermissions.image.repository | string | `"bitnami/bitnami-shell"` |  |
| volumePermissions.image.tag | string | `"10"` |  |
| volumePermissions.resources | object | `{}` |  |
| volumePermissions.securityContext.runAsUser | int | `0` |  |

## Source Code

* <https://github.com/bitnami/bitnami-docker-redis>
* <http://redis.io/>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | common | 1.x.x |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
