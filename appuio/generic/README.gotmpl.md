<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Configuration

The following table lists the configurable parameters chart. For default values and examples, consult `values.yaml`.

| Parameter      | Description                                                   |
| ---            | ---                                                           |
| `resources`    | List of resources in string format. Allows to use templating. |

> **Tip**: You can use the default [values.yaml](values.yaml)

## Templating

The resource strings allow to use Go templating like regular Helm charts.
Make sure you use the root context `$` form instead of `.` (see [docs](https://helm.sh/docs/chart_template_guide/variables/)).

The following named templates are provided as helpers:
* `generic.name` - Expand the name of the chart (truncated to 63 characters)
* `generic.fullname` - Default fully qualified app name
* `generic.chart` - Chart name and version
* `generic.labels` - Common labels
* `generic.selectorLabels` - Selector (name and instance) labels
