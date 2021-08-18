# redis

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![AppVersion: 6.2.1](https://img.shields.io/badge/AppVersion-6.2.1-informational?style=flat-square)

Open source, advanced key-value store. It is often referred to as a data structure server since keys can contain strings, hashes, lists, sets and sorted sets.

**Homepage:** <https://github.com/appuio/charts/tree/master/appuio/redis>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| APPUiO Team | info@appuio.ch |  |

## Source Code

* <https://github.com/bitnami/bitnami-docker-redis>
* <http://redis.io/>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | common | 1.x.x |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cluster.enabled | bool | `true` |  |
| cluster.slaveCount | int | `2` |  |
| clusterDomain | string | `"cluster.local"` |  |
| configmap | string | `"# Enable AOF https://redis.io/topics/persistence#append-only-file\nappendonly yes\n# Disable RDB persistence, AOF persistence already enabled.\nsave \"\""` |  |
| containerSecurityContext.enabled | bool | `true` |  |
| containerSecurityContext.runAsUser | int | `1001` |  |
| global.redis | object | `{}` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.registry | string | `"docker.io"` |  |
| image.repository | string | `"bitnami/redis"` |  |
| image.tag | string | `"6.2.1-debian-10-r36"` |  |
| master.affinity | object | `{}` |  |
| master.command | string | `"/run.sh"` |  |
| master.configmap | string | `nil` |  |
| master.customLivenessProbe | object | `{}` |  |
| master.customReadinessProbe | object | `{}` |  |
| master.disableCommands[0] | string | `"FLUSHDB"` |  |
| master.disableCommands[1] | string | `"FLUSHALL"` |  |
| master.extraEnvVars | list | `[]` |  |
| master.extraEnvVarsCM | list | `[]` |  |
| master.extraEnvVarsSecret | list | `[]` |  |
| master.extraFlags | list | `[]` |  |
| master.hostAliases | list | `[]` |  |
| master.livenessProbe.enabled | bool | `true` |  |
| master.livenessProbe.failureThreshold | int | `5` |  |
| master.livenessProbe.initialDelaySeconds | int | `5` |  |
| master.livenessProbe.periodSeconds | int | `5` |  |
| master.livenessProbe.successThreshold | int | `1` |  |
| master.livenessProbe.timeoutSeconds | int | `5` |  |
| master.persistence.accessModes[0] | string | `"ReadWriteOnce"` |  |
| master.persistence.enabled | bool | `true` |  |
| master.persistence.matchExpressions | object | `{}` |  |
| master.persistence.matchLabels | object | `{}` |  |
| master.persistence.path | string | `"/data"` |  |
| master.persistence.size | string | `"8Gi"` |  |
| master.persistence.subPath | string | `""` |  |
| master.persistence.volumes | string | `nil` |  |
| master.podAnnotations | object | `{}` |  |
| master.podLabels | object | `{}` |  |
| master.preExecCmds | string | `""` |  |
| master.priorityClassName | string | `nil` |  |
| master.readinessProbe.enabled | bool | `true` |  |
| master.readinessProbe.failureThreshold | int | `5` |  |
| master.readinessProbe.initialDelaySeconds | int | `5` |  |
| master.readinessProbe.periodSeconds | int | `5` |  |
| master.readinessProbe.successThreshold | int | `1` |  |
| master.readinessProbe.timeoutSeconds | int | `1` |  |
| master.service.annotations | object | `{}` |  |
| master.service.externalTrafficPolicy | string | `"Cluster"` |  |
| master.service.labels | object | `{}` |  |
| master.service.loadBalancerIP | string | `nil` |  |
| master.service.port | int | `6379` |  |
| master.service.type | string | `"ClusterIP"` |  |
| master.shareProcessNamespace | bool | `false` |  |
| master.statefulset.annotations | object | `{}` |  |
| master.statefulset.labels | object | `{}` |  |
| master.statefulset.updateStrategy | string | `"RollingUpdate"` |  |
| master.statefulset.volumeClaimTemplates.annotations | object | `{}` |  |
| master.statefulset.volumeClaimTemplates.labels | object | `{}` |  |
| metrics.enabled | bool | `false` |  |
| metrics.image.pullPolicy | string | `"IfNotPresent"` |  |
| metrics.image.registry | string | `"docker.io"` |  |
| metrics.image.repository | string | `"bitnami/redis-exporter"` |  |
| metrics.image.tag | string | `"1.20.0-debian-10-r16"` |  |
| metrics.podAnnotations."prometheus.io/port" | string | `"9121"` |  |
| metrics.podAnnotations."prometheus.io/scrape" | string | `"true"` |  |
| metrics.priorityClassName | string | `nil` |  |
| metrics.prometheusRule.additionalLabels | object | `{}` |  |
| metrics.prometheusRule.enabled | bool | `false` |  |
| metrics.prometheusRule.namespace | string | `""` |  |
| metrics.prometheusRule.rules | list | `[]` |  |
| metrics.redisTargetHost | string | `"localhost"` |  |
| metrics.service.annotations | object | `{}` |  |
| metrics.service.externalTrafficPolicy | string | `"Cluster"` |  |
| metrics.service.labels | object | `{}` |  |
| metrics.service.type | string | `"ClusterIP"` |  |
| metrics.serviceMonitor.enabled | bool | `false` |  |
| metrics.serviceMonitor.metricRelabelings | list | `[]` |  |
| metrics.serviceMonitor.relabelings | list | `[]` |  |
| metrics.serviceMonitor.selector.prometheus | string | `"kube-prometheus"` |  |
| networkPolicy.enabled | bool | `false` |  |
| networkPolicy.ingressNSMatchLabels | object | `{}` |  |
| networkPolicy.ingressNSPodMatchLabels | object | `{}` |  |
| password | string | `""` |  |
| persistence.existingClaim | string | `nil` |  |
| podDisruptionBudget.enabled | bool | `false` |  |
| podDisruptionBudget.minAvailable | int | `1` |  |
| podSecurityPolicy.create | bool | `false` |  |
| rbac.create | bool | `false` |  |
| rbac.role.rules | list | `[]` |  |
| redisPort | int | `6379` |  |
| securityContext.enabled | bool | `true` |  |
| securityContext.fsGroup | int | `1001` |  |
| sentinel.cleanDelaySeconds | int | `5` |  |
| sentinel.configmap | string | `nil` |  |
| sentinel.customLivenessProbe | object | `{}` |  |
| sentinel.customReadinessProbe | object | `{}` |  |
| sentinel.downAfterMilliseconds | int | `20000` |  |
| sentinel.enabled | bool | `false` |  |
| sentinel.extraEnvVars | list | `[]` |  |
| sentinel.extraEnvVarsCM | list | `[]` |  |
| sentinel.extraEnvVarsSecret | list | `[]` |  |
| sentinel.failoverTimeout | int | `18000` |  |
| sentinel.image.pullPolicy | string | `"IfNotPresent"` |  |
| sentinel.image.registry | string | `"docker.io"` |  |
| sentinel.image.repository | string | `"bitnami/redis-sentinel"` |  |
| sentinel.image.tag | string | `"6.2.1-debian-10-r35"` |  |
| sentinel.initialCheckTimeout | int | `5` |  |
| sentinel.livenessProbe.enabled | bool | `true` |  |
| sentinel.livenessProbe.failureThreshold | int | `5` |  |
| sentinel.livenessProbe.initialDelaySeconds | int | `5` |  |
| sentinel.livenessProbe.periodSeconds | int | `5` |  |
| sentinel.livenessProbe.successThreshold | int | `1` |  |
| sentinel.livenessProbe.timeoutSeconds | int | `5` |  |
| sentinel.masterSet | string | `"mymaster"` |  |
| sentinel.metrics.enabled | bool | `false` |  |
| sentinel.metrics.image.pullPolicy | string | `"IfNotPresent"` |  |
| sentinel.metrics.image.registry | string | `"docker.io"` |  |
| sentinel.metrics.image.repository | string | `"bitnami/redis-sentinel-exporter"` |  |
| sentinel.metrics.image.tag | string | `"1.7.1-debian-10-r109"` |  |
| sentinel.metrics.resources.limits | object | `{}` |  |
| sentinel.metrics.resources.requests | object | `{}` |  |
| sentinel.metrics.service.annotations | object | `{}` |  |
| sentinel.metrics.service.externalTrafficPolicy | string | `"Cluster"` |  |
| sentinel.metrics.service.labels | object | `{}` |  |
| sentinel.metrics.service.port | int | `9355` |  |
| sentinel.metrics.service.type | string | `"ClusterIP"` |  |
| sentinel.metrics.serviceMonitor.enabled | bool | `false` |  |
| sentinel.metrics.serviceMonitor.selector.prometheus | string | `"kube-prometheus"` |  |
| sentinel.parallelSyncs | int | `1` |  |
| sentinel.port | int | `26379` |  |
| sentinel.preExecCmds | string | `""` |  |
| sentinel.quorum | int | `2` |  |
| sentinel.readinessProbe.enabled | bool | `true` |  |
| sentinel.readinessProbe.failureThreshold | int | `5` |  |
| sentinel.readinessProbe.initialDelaySeconds | int | `5` |  |
| sentinel.readinessProbe.periodSeconds | int | `5` |  |
| sentinel.readinessProbe.successThreshold | int | `1` |  |
| sentinel.readinessProbe.timeoutSeconds | int | `1` |  |
| sentinel.service.annotations | object | `{}` |  |
| sentinel.service.externalTrafficPolicy | string | `"Cluster"` |  |
| sentinel.service.labels | object | `{}` |  |
| sentinel.service.loadBalancerIP | string | `nil` |  |
| sentinel.service.redisPort | int | `6379` |  |
| sentinel.service.sentinelPort | int | `26379` |  |
| sentinel.service.type | string | `"ClusterIP"` |  |
| sentinel.staticID | bool | `false` |  |
| sentinel.usePassword | bool | `true` |  |
| serviceAccount.create | bool | `false` |  |
| serviceAccount.name | string | `nil` |  |
| slave.affinity | object | `{}` |  |
| slave.command | string | `"/run.sh"` |  |
| slave.configmap | string | `nil` |  |
| slave.customLivenessProbe | object | `{}` |  |
| slave.customReadinessProbe | object | `{}` |  |
| slave.disableCommands[0] | string | `"FLUSHDB"` |  |
| slave.disableCommands[1] | string | `"FLUSHALL"` |  |
| slave.extraEnvVars | list | `[]` |  |
| slave.extraEnvVarsCM | list | `[]` |  |
| slave.extraEnvVarsSecret | list | `[]` |  |
| slave.extraFlags | list | `[]` |  |
| slave.hostAliases | list | `[]` |  |
| slave.livenessProbe.enabled | bool | `true` |  |
| slave.livenessProbe.failureThreshold | int | `5` |  |
| slave.livenessProbe.initialDelaySeconds | int | `30` |  |
| slave.livenessProbe.periodSeconds | int | `10` |  |
| slave.livenessProbe.successThreshold | int | `1` |  |
| slave.livenessProbe.timeoutSeconds | int | `5` |  |
| slave.persistence.accessModes[0] | string | `"ReadWriteOnce"` |  |
| slave.persistence.enabled | bool | `true` |  |
| slave.persistence.matchExpressions | object | `{}` |  |
| slave.persistence.matchLabels | object | `{}` |  |
| slave.persistence.path | string | `"/data"` |  |
| slave.persistence.size | string | `"8Gi"` |  |
| slave.persistence.subPath | string | `""` |  |
| slave.podAnnotations | object | `{}` |  |
| slave.podLabels | object | `{}` |  |
| slave.port | int | `6379` |  |
| slave.preExecCmds | string | `""` |  |
| slave.priorityClassName | string | `nil` |  |
| slave.readinessProbe.enabled | bool | `true` |  |
| slave.readinessProbe.failureThreshold | int | `5` |  |
| slave.readinessProbe.initialDelaySeconds | int | `5` |  |
| slave.readinessProbe.periodSeconds | int | `10` |  |
| slave.readinessProbe.successThreshold | int | `1` |  |
| slave.readinessProbe.timeoutSeconds | int | `10` |  |
| slave.service.annotations | object | `{}` |  |
| slave.service.externalTrafficPolicy | string | `"Cluster"` |  |
| slave.service.labels | object | `{}` |  |
| slave.service.loadBalancerIP | string | `nil` |  |
| slave.service.port | int | `6379` |  |
| slave.service.type | string | `"ClusterIP"` |  |
| slave.shareProcessNamespace | bool | `false` |  |
| slave.spreadConstraints | object | `{}` |  |
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
| tls.authClients | bool | `true` |  |
| tls.certCAFilename | string | `nil` |  |
| tls.certFilename | string | `nil` |  |
| tls.certKeyFilename | string | `nil` |  |
| tls.certificatesSecret | string | `nil` |  |
| tls.enabled | bool | `false` |  |
| usePassword | bool | `true` |  |
| usePasswordFile | bool | `false` |  |
| volumePermissions.enabled | bool | `false` |  |
| volumePermissions.image.pullPolicy | string | `"Always"` |  |
| volumePermissions.image.registry | string | `"docker.io"` |  |
| volumePermissions.image.repository | string | `"bitnami/bitnami-shell"` |  |
| volumePermissions.image.tag | string | `"10"` |  |
| volumePermissions.resources | object | `{}` |  |
| volumePermissions.securityContext.runAsUser | int | `0` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
