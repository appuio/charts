# secret

![Version: 1.1.0](https://img.shields.io/badge/Version-1.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

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

Note: Properties suffixed with `Templates` support template functions and variables.
All standard Helm functions are available (https://helm.sh/docs/chart_template_guide/function_list/).

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| secrets | object | `{}` | Dict with key/values. Each key is the name of the secret. Each value may contain an object with `.nameTemplate`, `.type`, `.stringData`, `.data`, `.labels`, `.annotations`. If `.nameTemplate` is not given, the key name is used. You can pass additional templates to `.stringDataTemplates` and `.nameTemplate`. See an example in `values.yaml`. |

## Upgrade configuration syntax

From version `1.1.0` onwards, the type of `secrets` changed from array to object to better support deep-merging behaviour of hierarchical configuration management tools.
Specifying `secrets` as an array is deprecated.
The changes are backwards compatible, altough you will note some default label changes.

Going forward, `secrets` is now an object by default.
Each key of `secrets` is the `metadata.name` of a secret.

To upgrade to the new structure, take the value of `secrets[*].name` and make it a key.
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
v1.2 also now supports additional fields, see the example in `values.yaml`.

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
