# secret

![Version: 2.0.0](https://img.shields.io/badge/Version-2.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

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
| secrets | object | `{}` | Dict with key/values. Each key is the name of the secret. Each value may contain an object with `.name`, `.type`, `.stringData`, `.data`, `.labels`, `.annotations`. If `.name` is not given, the key name is used. You can pass additional templated stringData using `.stringDataTemplates`. See an example in `values.yaml`. |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
