# signalilo

![Version: 0.9.1](https://img.shields.io/badge/Version-0.9.1-informational?style=flat-square) ![AppVersion: v0.12.0](https://img.shields.io/badge/AppVersion-v0.12.0-informational?style=flat-square)

Alertmanager to Icinga2 bridge

**Homepage:** <https://git.vshn.net/vshn/signalilo>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install signalilo appuio/signalilo
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Configuration

The following table lists the configurable parameters of the Signalilo chart and their default values.

Parameter | Description | Default
--- | --- | ---
`replicaCount` | Number of replicas to run | `1`
`image.registry` | Image registry | `docker.io`
`image.repository` | Image repository | `vshn/signalilo`
`image.tag` | Image tag | `v0.8.0`
`image.pullPolicy` | Image pull policy | `IfNotPresent`
`image.username` | Username to pull the image | `""`
`image.password` | Password to pull the image | `""`
`image.existingPullSecret` | Existing image pull secret | `""`
`config.uuid` | Signalilo UUID (required) |
`config.icinga_hostname` | Icinga Servicehost name (required) |
`config.icinga_url` | Icinga master URL (required) |
`config.icinga_username` | Icinga API user username (required) |
`config.icinga_password` | Icinga API user password (required if `config.icinga_password_secret` is unset) |
`config.icinga_password_secret` | Pre-existing secret for icinga API user password (the secret is expected to have a key `icinga_password`) |
`config.alertmanager_bearer_token` | Bearer token for incoming webhooks (required if `config.alertmanager_bearer_token_secret is unset`) |
`config.alertmanager_bearer_token_secret` | Pre-existing secret for bearer token for incoming webhooks (the secret is expected to have a key `alertmanager_bearer_token`) |
`config.alertmanager_port` | Port for incoming webhooks from Alertmanager | `8888`
`extraEnvVars` | Extra Signalilo configuration (see values.yaml for optional configuration values, and their defaults) | `[]`
`extraVolumes` | Extra Signalilo extra volumes | `[]`
`extraVolumeMounts` | Extra Signalilo extra volume mounts | `[]`
`securityContext.enabled` | Enable security context for the pod | `false`
`securityContext.runAsUser` | User to run the pod as | `999`
`securityContext.fsGroup` | fs group to use for the pod | `999`
`resources` | Pod resource requests and limits | `{}`
`nodeSelector` | Pod node selector | `{}`
`tolerations` | Pod tolerations | `[]`
`affinity` | Pod affinity rules | `{}`

## Source Code

* <https://git.vshn.net/vshn/signalilo>

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
