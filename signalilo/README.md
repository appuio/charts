# Signalilo

Signalilo is a bridge between the Prometheus Alertmanager and Icinga2 API.

## TL;DR;

```console
$ helm install appuio/signalilo
```

## Introduction

This chart deploys Signalilo on a [Kubernetes](https://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager. It provisions a fully featured Signalilo installation.

## Installing the Chart

To install the chart with the release name `signalilo`:

```console
$ helm install --name signalilo appuio/signalilo
```

## Uninstalling the Chart

To uninstall/delete the `signalilo` deployment:

```console
$ helm delete signalilo
```

## Configuration

The following table lists the configurable parameters of the Signalilo chart and their default values.

Parameter | Description | Default
--- | --- | ---
`replicaCount` | Number of replicas to run | `2`
`image.registry` | Image registry | `registry.vshn.net`
`image.repository` | Image repository | `vshn/signalilo`
`image.tag` | Image tag | `0.0.1`
`image.pullPolicy` | Image pull policy | `IfNotPresent`
`image.username` | Username to pull the image | `""`
`image.password` | Password to pull the image | `""`
`image.existingPullSecret` | Existing image pull secret | `""`
`config` | Configuration for Signalilo (in YAML format) | `{}`
`securityContext.enabled` | Enable security context for the pod | `false`
`securityContext.runAsUser` | User to run the pod as | `999`
`securityContext.fsGroup` | fs group to use for the pod | `999`
`resources` | Pod resource requests and limits | `{}`
`nodeSelector` | Pod node selector | `{}`
`tolerations` | Pod tolerations | `[]`
`affinity` | Pod affinity rules | `{}`

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```bash
$ helm install --name signalilo -f values.yaml appuio/signalilo
```
