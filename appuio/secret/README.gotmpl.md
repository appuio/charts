<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

Note: Properties suffixed with `Templates` support template functions and variables.
All standard Helm functions are available (https://helm.sh/docs/chart_template_guide/function_list/).

{{ template "chart.valuesSection" . }}

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
