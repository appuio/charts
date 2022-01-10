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
| secrets | object | `{}` | Dict with key/values. Each key is the name of the secret. Each value may contain an object with `.nameTemplate`, `.type`, `.stringData`, `.data`, `.labels`, `.annotations`. If `.nameTemplate` is not given, the key name is used. You can pass additional templates to `.stringDataTemplates` and `.nameTemplate`. See an example in `values.yaml`. |

## Upgrade chart v1 to v2

Type of `.Values.secrets` changed from List to Dict to better support deep-merging behaviour of hierarchical configuration management tools.
Each key of `.secrets` is now the default `metadata.name` of a secret.

To upgrade, take the value of `.secrets[*].name` and make it a key.
Given the example below, restructure the `name` and rename `values` to `stringData`:
```yaml
secrets:
- name: db-creds
  values:
    db-host: mysql
```
to
```yaml
secrets:
  db-creds:
    stringData:
      db-host: mysql
```
The additional fields `.type` and `.labels` also need to be moved beneath the new name key (given example: `db-creds`).
v2 also now supports additional fields, see the example in `values.yaml`.

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
