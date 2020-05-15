# K8up - Kubernetes and OpenShift Backup Operator based on restic

[K8up](https://github.com/vshn/k8up) is a Kubernetes operator which simplifies the process of application data backup in a kubernetes cluster. It is based on [restic](https://restic.net/).

## TL;DR;

```console
helm repo add appuio https://charts.appuio.ch
helm install k8up appuio/k8up
```

## Introduction

This chart bootstraps a [K8up](https://vshn.github.io/k8up/) operator on a [Kubernetes](https://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager. To start backing up your application data just create a schedule object in the namespace you would like to backup. For more information on K8up and its capabilities please check the [documentation](https://vshn.github.io/k8up/)

## Prerequisites Details

As the K8up operator is working cluster wide it needs an appropriate cluster role and rolebinding. If your executing user or tiller (the server-side component of Helm v2) doesnt have the sufficient privileges to define roles and rolebindings, it might be that you need to define these per hand with elevated privileges.

K8up uses Wrestic as a backup runner, to learn more about it, please visit [wrestic on github](https://github.com/vshn/wrestic/tree/master).

## Installing the Chart

To install the chart with the release name `k8up`:

```console
helm install k8up appuio/k8up
```

## Uninstalling the Chart

To uninstall/delete the `k8up` deployment:

```console
helm delete k8up
```

## Installing the Chart (without cluster-admin)

```console
helm fetch appuio/k8up --untar
helm template ./k8up/ -s templates/clusterrole.yaml -s templates/clusterrolebinding.yaml | kubectl apply -f -
rm -rf ./k8up/
helm install k8up appuio/k8up --set rbac.enabled=false
```

## Uninstalling the Chart (without cluster-admin)

```console
helm delete k8up
helm fetch appuio/k8up --untar
helm template ./k8up/ -s templates/clusterrole.yaml -s templates/clusterrolebinding.yaml | kubectl delete -f -
rm -rf ./k8up/
```

## Configuration

The following table lists the configurable parameters of the k8up chart. For the usual parameters and defaults please consult `values.yaml`.

| Parameter                     | Description                                             | Default
| ---                           | ---                                                     | ---
| `image.tag`                   | The operator image tag (version)                        | see `values.yaml` for latest supported version
| `k8up.backupImage.repository` | The backup runner image repository                      | `docker.io/vshn/wrestic`
| `k8up.backupImage.tag`        | The backup runner image tag                             | see `values.yaml` for latest supported version
| `k8up.envVars`                | Allows the specification of additional environment variables for the k8up operator | `[]`
| `rbac.create`                 | Create cluster roles and rolebinding                    | `true`
| `metrics.service.type` | Service type of the metrics endpoint | `ClusterIP`
| `metrics.service.port` | Service port of the metrics endpoint | `8080`
| `metrics.service.nodePort` | Service node port of the metrics endpoint, requires `metrics.service.type=NodePort` | `0`
| `metrics.serviceMonitor.enabled` | Whether to enable ServiceMonitor manifests for [Prometheus Operator](https://github.com/coreos/prometheus-operator) | `false`
| `metrics.serviceMonitor.scrapeInterval` | Scrape interval to collect K8up metrics | `60s`
| `metrics.serviceMonitor.namespace` | If the object should be installed in a different namespace than K8up | `""`
| `metrics.serviceMonitor.additionalLabels` | Add labels to the ServiceMonitor object | `{}`
| `metrics.prometheusRule.enabled` | Whether to enable PrometheusRule manifest for [Prometheus Operator](https://github.com/coreos/prometheus-operator) | `false`
| `metrics.prometheusRule.namespace` | If the object should be installed in a different namespace than K8up | `""`
| `metrics.prometheusRule.additionalLabels` | Add labels to the PrometheusRule object | `{}`
| `metrics.prometheusRule.createDefaultRules` | Whether the K8up default rules should be installed | `false`
| `metrics.prometheusRule.additionalRules` | Provide additional alert rules in addition to the defaults | `{}`

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.
