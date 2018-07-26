# BaaS Operator

This Chart deploys the APPUiO BaaS Operator

## Configuration

The following table lists the configurable parameters of the BaaS Operator chart. For defaults consult `values.yaml`

| Parameter                   | Description                                             |
| ---                         | ---                                                     |
| `baas_operator.image`       | BaaS operator image                                     |
| `baas_operator.environment` | Environment variables for configuring the BaaS operator |
| `rbac.create`               | Create cluster roles and rolebinding                    |

