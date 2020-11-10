# Generic Chart

A Helm chart to create generic Kubernetes resources.

## TL;DR;

```console
$ helm install appuio/generic
```

## Configuration

The following table lists the configurable parameters chart. For default values and examples, consult `values.yaml`.

| Parameter      | Description                                                   |
| ---            | ---                                                           |
| `resources`    | List of resources in string format. Allows to use templating. |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
$ helm install --name my-release appuio/generic
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```console
$ helm install --name my-release -f values.yaml appuio/generic
```

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

## Testing

### Template testing

This chart uses some Go code to test the generated templates against the
Kubernetes client API.

```console
cd generic/test
go test -v ./...
```
