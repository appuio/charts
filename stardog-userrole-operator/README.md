# Stardog-UserRole-Operator - Kubernetes and OpenShift Operator that manages Stardog Users and Roles

[stardog-userrole-operator](https://github.com/vshn/stardog-userrole-operator) is a Kubernetes operator which helps to maintain users and roles of [Stardog](https://www.stardog.com/) with plain Kubernetes yaml files.

## TL;DR;

```console
helm repo add appuio https://charts.appuio.ch
helm install stardog-userrole-operator appuio/stardog-userrole-operator
```

## Introduction

This chart bootstraps a [stardog-userrole-operator](https://vshn.github.io/stardog-userrole-operator/) operator on a [Kubernetes](https://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager. To start using the operator create a StardogInstance CRD that connects to Stardog, then define users and roles CRD that should be synced to Stardog. 

## Prerequisites Details

As the Stardog-UserRole-Operator is working cluster wide it needs an appropriate cluster role and rolebinding. If your executing user or tiller (the server-side component of Helm v2) doesn't have the sufficient privileges to define roles and rolebindings, it might be that you need to define them per hand with elevated privileges.

Stardog-UserRole-Operator uses Autorest client to connect to the Stardog API, to learn more about Stardog REST API, please visit [stardog api](https://stardog.docs.apiary.io/).

## Installing the Chart

To install the chart with the release name `stardog-userrole-operator`:

```console
helm install stardog-userrole-operator appuio/stardog-userrole-operator
```

## Uninstalling the Chart

To uninstall/delete the `stardog-userrole-operator` deployment:

```console
helm delete stardog-userrole-operator
```

## Configuration

The following table lists the configurable parameters of the stardog-userrole-operator chart. For the usual parameters and defaults please consult `values.yaml`.

| Parameter                     | Description                                             | Default
| ---                           | ---                                                     | ---
| `image.registry`              | The operator image registry                             | docker.io
| `image.repository`            | The operator image repository                           | docker.io
| `image.tag`                   | The operator image tag (version)                        | vshn/stardog-userrole-operator
| `image.pullPolicy`            | The operator image pull policy from repository          | IfNotPresent
| `imagePullSecrets`            | The operator image pull secrets for registry            | []
| `replicaCount`                | The operator number of replicas                         | 1
| `leaderElection.enabled`      | Management of multiple operator instances via master-slave relationship | true
| `serviceAccount.create`       | Create a custom service account for the operator        | true
| `serviceAccount.name`         | Set a specific name for the service account             | disabled by default
| `resources.limits`            | Define resource limit that are assigned to the operator | cpu 100m, memory 100Mi
| `resources.requests`          | Define resource request that are assigned to the operator | cpu 30m, memory 20Mi
| `metrics.service.enabled`     | Enable metrics via service | true
| `metrics.service.port`        | Define the port where the metrics can be accessed       | 80
| `metrics.proxy.enabled`       | Enable metrics via service behind a authenticated proxy | false
| `metrics.proxy.port`          | Define the proxy port where the metrics can be accessed | 8443

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.