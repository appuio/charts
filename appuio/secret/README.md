# secret

![Version: 1.0.2](https://img.shields.io/badge/Version-1.0.2-informational?style=flat-square)

Deploy arbitrary secrets

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install secret appuio/secret
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| secrets | list | `[]` |  |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
