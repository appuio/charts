# Stardog Chart

[Stardog](https://www.stardog.com/) is the world’s leading knowledge graph platform for the enterprise. Stardog makes it fast and easy to turn enterprise data into knowledge.

## TL;DR;

```console
# Create a secret with your Stardog license
$ kubectl create secret generic stardog-license --from-file stardog-license-key.bin=/path/to/your/license.bin

$ helm install appuio/stardog
```

## Configuration

The following table lists the configurable parameters chart. For default values consult `values.yaml`.

| Parameter                            | Description                                                                  |
| ---                                  | ---                                                                          |
| `replicaCount`                       | Number of Stardog instances to run. If greater than 1, `zookeeper.enabled=true` must be set |
| `image.registry`                     | The Docker registry to pull from                                             |
| `image.repository`                   | The Docker image repostory to pull                                           |
| `image.tag`                          | The Docker image tag to pull                                                 |
| `image.username`                     | The Docker registry user name                                                |
| `image.password`                     | The Docker registry user name                                                |
| `image.existingPullSecret`           | An existing image pull secret to use |
| `stardog.existingLicenseSecret`      | Name of an existing secret with a key `stardog-license-key.bin` which contains the Stardog license |
| `stardog.config`                     | Custom configuration to append to the `stardog.properties` file (as string) |
| `stardog.securityContext`            | K8s security context |
| `ingress.enabled`                    | If an ingress object should be created |
| `ingress.annotations`                | Annotations to set on the ingress object |
| `ingress.host`                       | Host name which the ingress should resolve |
| `ingress.tls.enabled`                | If TLS should be enabled on the ingress |
| `ingress.tls.secretName`             | Name of the secret containing the TLS certificate and key |
| `persistence.enabled`                | Enable persistence using PVC |
| `persistence.storageClass`           | PVC storage class for Stardog data volume |
| `persistence.size`                   | PVC storage request size for Stardog data volume |
| `affinityTopologyKey`                | Topology key on which pod anti affinity is configured |
| `zookeeper.enabled`                  | Enable ZooKeeper. Required for Stardog Cluster deployments |
| `zookeeper.allowAnonymousLogin`      | Allow anonymous logins to ZooKeeper |
| `zookeeper.auth.enabled`             | Enable ZooKeeper authentication |
| `zookeeper.auth.clientPassword`      | Password to use for ZooKeeper clients. Will be generated if unset |
| `zookeeper.auth.serverPasswords`     | Password to use for ZooKeeper servers (delimited by `,`). Will be generated if unset |
| `zookeeper.metrics.enabled`          | Enable ZooKeeper Proemtheus exporter |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
$ helm install --name my-release appuio/stardog
```

The above command enables the ZooKeeper installation.

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```console
$ helm install --name my-release -f values.yaml appuio/stardog
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
cd stardog/test
go test -v ./...
```
