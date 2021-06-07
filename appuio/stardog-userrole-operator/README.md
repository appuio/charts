# Stardog User and Role Operator - Kubernetes Operator that manages Stardog Users and Roles

[Stardog User and Role Operator](https://github.com/vshn/stardog-userrole-operator) is a Kubernetes operator which helps to maintain users and roles of [Stardog](https://www.stardog.com/) using Kubernetes CRDs.

## TL;DR;

```console
helm repo add appuio https://charts.appuio.ch
helm install stardog-userrole-operator appuio/stardog-userrole-operator
```

## Introduction

This chart bootstraps a [Stardog User and Role Operator](https://vshn.github.io/stardog-userrole-operator/) operator on a [Kubernetes](https://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

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

| Parameter                     | Description                                                               | Default
| ---                           | ---                                                                       | ---
| `image.registry`              | Stardog User and Role Operator image registry                             | docker.io
| `image.repository`            | Stardog User and Role Operator image repository                           | vshn/stardog-userrole-operator
| `image.tag`                   | Stardog User and Role Operator image tag (version)                        | 0.0.1
| `image.pullPolicy`            | Stardog User and Role Operator image pull policy from repository          | IfNotPresent
| `imagePullSecrets`            | Stardog User and Role Operator image pull secrets for registry            | []
| `replicaCount`                | Stardog User and Role Operator number of replicas                         | 1
| `leaderElection.enabled`      | Enable leader election for multiple replicas                              | false
| `config.clusterAccess.enabled`| Enable cluster access to operator                                         | true
| `serviceAccount.create`       | Create a custom service account for the operator                          | true
| `serviceAccount.name`         | Name of service account                                                   | not set
| `resources.limits`            | Resource limits                                                           | cpu 100m, memory 100Mi
| `resources.requests`          | Resource requests                                                         | cpu 30m, memory 20Mi
| `metrics.service.enabled`     | Metrics service                                                           | true
| `metrics.service.port`        | Metrics service port                                                      | 8080
| `metrics.proxy.image  `       | Proxy image                                                               | false
| `metrics.proxy.enabled`       | Enable metrics via service behind a authenticated proxy                   | false
| `metrics.proxy.port`          | Metrics Port                                                              | 8443

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.
