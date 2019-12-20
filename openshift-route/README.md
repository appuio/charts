# Openshift Route

Deploys a single OpenShift route

## TL;DR;

```console
helm repo add appuio https://charts.appuio.ch
helm install appuio/openshift-route --set service.name=my-backend-service
```

## Configuration

The following table lists the configurable parameters of the chart. For defaults please consult `values.yaml`

| Parameter | Description | Default
| ---       | ---         | ---
| `host` | The hostname that should be used. If left empty, OpenShift will generate one for you with defaults. | `""`
| `nameOverride` | Overrides the name after the Chart name. Affects route name and hostname generation. | `""`
| `fullnameOverride` | If you need to give a specific name to the route, override this, as it affects the hostname generated if `host` is left empty. | `""`
| `path` | Subpath of the route. | `""`
| `annotations` | Annotations on the route object. | `{}`
| `labels` | Additional labels on the route object. | `{}`
| `service.name` | The backend service name of the route. Required. | `""`
| `service.targetPort` | The port of the backend service. The port is aware of the named ports of the service, so it can be a name too. | `http`
| `service.weight` | Weight of the service endpoint. | `100`
| `wildcardPolicy` | Wildcard Policy of the route. | `None`
| `tls.enabled` | Whether to secure the Route with TLS. | `false`
| `tls.termination` | TLS termination of the route. | `edge`
| `tls.insecureEdgeTerminationPolicy` | What to do with non-HTTPS traffic (usually port 80) | `Redirect`
| `tls.key` | Certificate private key in PEM format. | `""`
| `tls.certificate` | Certificate in PEM format. | `""`
| `tls.caCertificate` | CA (chain) of the certificate in PEM format. | `""`
| `tls.destinationCACertificate` | CA of the backend in PEM format. Only relevant when `termination=reencrypt`. | `""`
| `alternateBackends` | Up to 4 additional backends can be enabled to support the route. Consult the documentation. | `[]`

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

When you specify certificates or keys in helmfile, please use the YAML multiline strings, e.g.
```yaml
tls:
  certificate: |
    -----BEGIN CERTIFICATE-----
    ...
    -----END CERTIFICATE-----
```

You can find more information on Routes in the [Documentation](https://docs.openshift.com/container-platform/3.9/architecture/networking/routes.html).
