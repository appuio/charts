# BaaS Operator

This Chart deploys the APPUiO BaaS Operator

## Configuration

The following table lists the configurable parameters of the BaaS Operator chart. For defaults consult `values.yaml`

| Parameter                   | Description                                             |
| ---                         | ---                                                     |
| `baas_operator.image`       | BaaS operator image                                     |
| `baas_operator.envVars`     | Allows the specification of additional environment variables for the BaaS operator. Passed through the `tpl` funtion and thus to be configured a string |
| `rbac.create`               | Create cluster roles and rolebinding                    |
