# K8up - Kubernetes and OpenShift Backup Operator based on restic

![Version: 1.0.1](https://img.shields.io/badge/Version-1.0.1-informational?style=flat-square) ![AppVersion: v1.0.1](https://img.shields.io/badge/AppVersion-v1.0.1-informational?style=flat-square)

<!---
This README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## TL;DR;

**K8up Documentation**: https://k8up.io/

```bash
# Install CRDs for K8s >= 1.16:
kubectl apply -f https://github.com/vshn/k8up/releases/download/v1.0.1/k8up-crd.yaml
# Install CRDs for K8s <= 1.15 (e.g. OpenShift 3.11):
kubectl apply -f https://github.com/vshn/k8up/releases/download/v1.0.1/k8up-crd-legacy.yaml

helm repo add appuio https://charts.appuio.ch
helm install k8up appuio/k8up
```

## Handling CRDs

* Always upgrade the CRDs before upgrading the Helm release.
* Watch out for breaking changes in the K8up release notes.

## Source Code

* <https://github.com/vshn/k8up>
* <https://github.com/vshn/wrestic>

<!---
The values below are generated with helm-docs!

Document your changes in values.yaml and let `make helm-docs` generate this section.
-->
## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| image.pullPolicy | string | `"IfNotPresent"` | Operator image pull policy |
| image.registry | string | `"quay.io"` | Operator image registry |
| image.repository | string | `"vshn/k8up"` | Operator image repository |
| image.tag | string | `"v1.0.1"` | Operator image tag (version) |
| imagePullSecrets | list | `[]` |  |
| k8up.backupImage.repository | string | `"quay.io/vshn/wrestic"` | The backup runner image repository |
| k8up.backupImage.tag | string | `"v0.1.9"` | The backup runner image tag |
| k8up.enableLeaderElection | bool | `true` | Specifies whether leader election should be enabled. Disable this for K8s versions < 1.16 |
| k8up.envVars | list | `[]` | envVars allows the specification of additional environment variables. See [values.yaml](values.yaml) how to specify See documentation which variables are supported. |
| k8up.globalResources | object | empty values, [see supported units][supported-units] | Specify the resource requests and limits that the Pods should have when they are scheduled by K8up. You are still able to override those via K8up resources, but this gives cluster administrators custom defaults. |
| k8up.globalResources.limits.cpu | string | `""` | Global CPU resource limit |
| k8up.globalResources.limits.memory | string | `""` | Global Memory resource limit |
| k8up.globalResources.requests.cpu | string | `""` | Global CPU resource requests |
| k8up.globalResources.requests.memory | string | `""` | Global Memory resource requests |
| k8up.operatorNamespace | string | `""` | Specifies the namespace in which K8up's `EffectiveSchedules` are stored. Defaults to release namespace if left empty. |
| k8up.timezone | string | `""` | Specifies the timezone K8up is using for scheduling. Empty value defaults to the timezone in which Kubernetes is deployed. Accepts `tz database` compatible entries, e.g. `Europe/Zurich` |
| metrics.prometheusRule.additionalLabels | object | `{}` | Add labels to the PrometheusRule object |
| metrics.prometheusRule.additionalRules | list | `[]` | Provide additional alert rules in addition to the defaults |
| metrics.prometheusRule.createDefaultRules | bool | `true` | Whether the default rules should be installed |
| metrics.prometheusRule.enabled | bool | `false` | Whether to enable PrometheusRule manifest for [Prometheus Operator][prometheus-operator] |
| metrics.prometheusRule.namespace | string | `""` | If the object should be installed in a different namespace than operator |
| metrics.service.nodePort | int | `0` | Service node port of the metrics endpoint, requires `metrics.service.type=NodePort` |
| metrics.service.port | int | `8080` |  |
| metrics.service.type | string | `"ClusterIP"` |  |
| metrics.serviceMonitor.additionalLabels | object | `{}` | Add labels to the ServiceMonitor object |
| metrics.serviceMonitor.enabled | bool | `false` | Whether to enable ServiceMonitor manifests for [Prometheus Operator][prometheus-operator] |
| metrics.serviceMonitor.namespace | string | `""` | If the object should be installed in a different namespace than operator |
| metrics.serviceMonitor.scrapeInterval | string | `"60s"` | Scrape interval to collect metrics |
| nodeSelector | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| rbac.create | bool | `true` | Create cluster roles and rolebinding. May need elevated permissions to create cluster roles and -bindings. |
| replicaCount | int | `1` | How many operator pods should run. Note: Operator features leader election for K8s 1.16 and later, so that only 1 pod is reconciling/scheduling jobs. Follower pods reduce interruption time as they're on hot standby when leader is unresponsive. |
| resources.limits.memory | string | `"256Mi"` |  |
| resources.requests.cpu | string | `"20m"` |  |
| resources.requests.memory | string | `"128Mi"` |  |
| securityContext | object | `{}` |  |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| serviceAccount.name | string | `""` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |
| tolerations | list | `[]` |  |

## Upgrading from Charts 0.x to 1.x

* In `image.repository` the registry domain was moved into its own parameter `image.registry`.
* K8up 1.x features leader election, this enables rolling updates and multiple replicas.
  `k8up.enableLeaderElection` defaults to `true`. Disable this for older Kubernetes versions (<= 1.15)
* `replicaCount` is now configurable, defaults to `1`.
* Note: Deployment strategy type has changed from `Recreate` to `RollingUpdate`.
* CRDs need to be installed separately, they are no longer included in this chart.

<!---
Link references from values.yaml
-->
[supported-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
