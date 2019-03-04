# Keycloak

[Keycloak](http://www.keycloak.org/) is an open source identity and access management for modern applications and services.

The readme for the upstream chart of [Keycloak can be found here](https://github.com/helm/charts/blob/master/stable/keycloak/README.md)  

## Adaptations for Openshift

After helm updated its dependencies there is a new folder called /charts. In this folder you find the upstream chart from helm/stable.
There are some minor adaptations necessary to run the [helm/stable Keycloak-Chart] on openshift. To make these adaptations the appuio Keycloak-Chart takes the
helm/stable Chart just as a dependency. The values.yaml then adds an empty securityContext and a configuration for the route to the chart. The templates-folder contains the openshift-specific definition for the route. Everything else is plain Keycloak.

## Configuration

To override parameters regarding the keycloak-chart (which is helm/stable) you need to address the chart in your values.yaml like following:

```yaml
keycloak: #prepend
  keycloak:
    <key>:<value>
    ...
```
