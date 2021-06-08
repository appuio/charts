# metrics-server

![Version: 2.12.1](https://img.shields.io/badge/Version-2.12.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.3.7](https://img.shields.io/badge/AppVersion-0.3.7-informational?style=flat-square)

Metrics Server is a cluster-wide aggregator of resource usage data.

**Homepage:** <https://github.com/kubernetes-incubator/metrics-server>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install metrics-server appuio/metrics-server
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Configuration

Parameter | Description | Default
--- | --- | ---
`rbac.create` | Enable Role-based authentication | `true`
`rbac.pspEnabled` | Enable pod security policy support | `false`
`serviceAccount.create` | If `true`, create a new service account | `true`
`serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | ``
`apiService.create` | Create the v1beta1.metrics.k8s.io API service | `true`
`hostNetwork.enabled` | Enable hostNetwork mode | `false`
`image.repository` | Image repository | `k8s.gcr.io/metrics-server/metrics-server`
`image.tag` | Image tag | `v0.3.7`
`image.pullPolicy` | Image pull policy | `IfNotPresent`
`imagePullSecrets` | Image pull secrets | `[]`
`args` | Command line arguments | `[]`
`resources` | CPU/Memory resource requests/limits. | `{}`
`tolerations` | List of node taints to tolerate (requires Kubernetes >=1.6) | `[]`
`nodeSelector` | Node labels for pod assignment | `{}`
`affinity` | Node affinity | `{}`
`replicas` | Number of replicas | `1`
`extraVolumeMounts` | Ability to provide volume mounts to the pod | `[]`
`extraVolumes` | Ability to provide volumes to the pod | `[]`
`livenessProbe` | Container liveness probe | See values.yaml
`podLabels` | Labels to be added to pods | `{}`
`podAnnotations` | Annotations to be added to pods | `{}`
`priorityClassName` | Pod priority class | `""`
`readinessProbe` | Container readiness probe | See values.yaml
`service.annotations` | Annotations to add to the service | `{}`
`service.labels` | Labels to be added to the metrics-server service | `{}`
`service.port` | Service port to expose | `443`
`service.type` | Type of service to create | `ClusterIP`
`podDisruptionBudget.enabled` | Create a PodDisruptionBudget | `false`
`podDisruptionBudget.minAvailable` | Minimum available instances; ignored if there is no PodDisruptionBudget |
`podDisruptionBudget.maxUnavailable` | Maximum unavailable instances; ignored if there is no PodDisruptionBudget |
`extraContainers`   | Add additional containers  | `[]`
`testImage.repository` | Image repository and name for test pod.  See also `imagePullSecrets` | `busybox`
`testImage.tag` | Image tag for test pod | `latest`
`testImage.pullPolicy` | Image pull policy for test pod | `IfNotPresent`

## Source Code

* <https://github.com/kubernetes-incubator/metrics-server>

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
