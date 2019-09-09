# K8up - Kubernetes and OpenShift Backup Operator based on restic

[K8up](https://github.com/vshn/k8up) is a Kubernetes operator which simplifies the process of application data backup in a kubernetes cluster. It is based on [restic](https://restic.net/).

## TL;DR;

```console
helm repo add appuio https://charts.appuio.ch
helm install appuio/k8up
```

## Introduction

This chart bootstraps a [K8up](https://vshn.github.io/k8up/) operator on a [Kubernetes](https://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager. To start backing up your application data just create a schedule object in the namespace you would like to backup. For more information on K8up and its capabilities please check the [documentation](https://vshn.github.io/k8up/)

## Prerequisites Details

As the K8up operator is working cluster wide it needs an appropriate cluster role and rolebinding. If your tiller (the server-side component of Helm) doesnt have the right privileges to define roles and rolebindings, it might be that you need to define these per hand.

K8up uses Wrestic as a backup runner, to learn more about it, please visit [wrestic on github](https://github.com/vshn/wrestic/tree/master). 

## Installing the Chart

To install the chart with the release name `k8up`:

```console
helm install --name k8up appuio/k8up
```

## Uninstalling the Chart

To uninstall/delete the `k8up` deployment:

```console
helm delete k8up
```

## Installing the Chart (without cluster-admin for tiller)

```console
helm fetch appuio/k8up --untar
helm template ./k8up/ -x templates/clusterrole.yaml -x templates/clusterrolebinding.yaml --set rbac.enabled=true | kubectl apply -f -
rm -rf ./k8up/
helm install --name k8up appuio/k8up --set rbac.enabled=false
```

## Uninstalling the Chart (without cluster-admin for tiller)

```console
helm delete k8up
helm fetch appuio/k8up --untar
helm template ./k8up/ -x templates/clusterrole.yaml -x templates/clusterrolebinding.yaml --set rbac.enabled=true | kubectl delete -f -
rm -rf ./k8up/
```

## Configuration

The following table lists the configurable parameters of the k8up chart. For defaults please consult `values.yaml`

| Parameter                   | Description                                             | Default
| ---                         | ---                                                     | ---
| `k8up_operator.image`       | The K8up operator image                                     | docker.io/vshn/k8up:v0.1.4
| `k8up_operator.envVars`     | Allows the specification of additional environment variables for the k8up operator. | BACKUP_IMAGE:docker.io/vshn/wrestic:v0.0.10
| `rbac.create`               | Create cluster roles and rolebinding                    | true

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

### Usage of the `tpl` Function

The `tpl` function allows us to pass string values from `values.yaml` through the templating engine. It is used for the following values:

* `k8up_operator.envVars`

It is important that these values be configured as strings. Otherwise, installation will fail. See example for Google Cloud Proxy or default affinity configuration in `values.yaml`.
