<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

Note: Properties suffixed with `Templates` support template functions and variables.
All standard Helm functions are available (https://helm.sh/docs/chart_template_guide/function_list/).

{{ template "chart.valuesSection" . }}

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
