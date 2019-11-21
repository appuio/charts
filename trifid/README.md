# Trifid Chart

[Trifid](https://github.com/zazuko/trifid) provides a lightweight and easy way to access Linked Data URIs via HTTP.

## TL;DR;

```console
helm install appuio/trifid
```

## Configuration

The following table lists the configurable parameters chart. For default values consult `values.yaml`.

| Parameter                                    | Description |
| ---                                          | --- |
| `replicaCount`                               | Number of trifid instances to run. |
| `image.registry`                             | The Docker registry to pull from |
| `image.repository`                           | The Docker image repostory to pull |
| `image.tag`                                  | The Docker image tag to pull |
| `trifid.sparql.endpoint.url`                 | URL of the SPARQL endpoint to use |
| `trifid.sparql.endpoint.username`            | Username used to connect to the SPARQL endpoint |
| `trifid.sparql.endpoint.password`            | Password used to connect to the SPARQL endpoint |
| `trifid.sparql.endpoint.secretName`          | Name of an existing secret containing `username` and `password` of the SPARQL endpoint |
| `ingress.enabled`                            | If an ingress object should be created |
| `ingress.annotations`                        | Annotations to set on the ingress object |
| `ingress.host`                               | Host name which the ingress should resolve |
| `ingress.tls.enabled`                        | If TLS should be enabled on the ingress |
| `ingress.tls.secretName`                     | Name of the secret containing the TLS certificate and key |
| `resources`                                  | Pod resource requests and limits |
| `nodeSelector`                               | Pod node selector |
| `affinity`                                   | Pod affinity rules |
| `tolerations`                                | Pod tolerations |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
helm install --name my-release appuio/trifid
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```console
helm install --name my-release -f values.yaml appuio/trifid
```

> **Tip**: You can use the default [values.yaml](values.yaml)


## Testing

### Template testing

This chart uses some Go code to test the generated templates against the
Kubernetes client API. With this we are more flexible than `helm lint`, as
linting would not catch errors such as: "If we are in clustered mode, we expect
an init container, otherwise not". The test code would also catch typos in a
limited manner (e.g. if `initContainer` is spelled `initcontainer`, it will
generate an error).

```console
cd trifid/test
go test -v ./...
```
