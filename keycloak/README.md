# Keycloak

[Keycloak](http://www.keycloak.org/) is an open source identity and access management for modern applications and services.

The README for the upstream chart of [Keycloak can be found here](https://github.com/codecentric/helm-charts/tree/master/charts/keycloak)

## Adaptations for OpenShift

There are some minor adaptations necessary to run the Keycloak chart on OpenShift. To make these adaptations the APPUiO Keycloak chart defines the
codecentric chart as a dependency. The values.yaml then adds an empty securityContext and a configuration for the route to the chart. The templates-folder contains the openshift-specific definition for the route.

## Configuration

To override parameters regarding the keycloak-chart (which is helm/stable) you need to address the chart in your values.yaml like following:

```yaml
keycloak: #prepend
  keycloak:
    <key>:<value>
    ...
```
