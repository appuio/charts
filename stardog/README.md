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

| Parameter                                    | Description                                                                  |
| ---                                          | ---                                                                          |
| `replicaCount`                               | Number of Stardog instances to run. If greater than 1, `zookeeper.enabled=true` must be set |
| `image.registry`                             | The Docker registry to pull from                                             |
| `image.repository`                           | The Docker image repostory to pull                                           |
| `image.tag`                                  | The Docker image tag to pull                                                 |
| `image.username`                             | The Docker registry user name                                                |
| `image.password`                             | The Docker registry user name                                                |
| `image.existingPullSecret`                   | An existing image pull secret to use |
| `stardog.existingLicenseSecret`              | Name of an existing secret with a key `stardog-license-key.bin` which contains the Stardog license |
| `stardog.config`                             | Custom configuration to append to the `stardog.properties` file (as string) |
| `stardog.adminPassword`                      | Password for the Administrator |
| `stardog.users`                              | Array of users, each element is an object with `name` and `password` |
| `stardog.roles`                              | Object Array of Roles to create, see below |
| `stardog.roles[].name`                       | Name of the Role to create |
| `stardog.roles[].grants`                     | Object Array of permissions to add to objects |
| `stardog.roles[].grants[].action`            | Name of the permission to add, see [Docs](https://www.stardog.com/docs/6.0.0/#_permissions) |
| `stardog.roles[].grants[].object`            | Object that this grant should apply to, see [Docs](https://www.stardog.com/docs/6.0.0/#_permissions) |
| `stardog.roles[].users          `            | String array of users that should have the role assigned to. The users need to exist first |
| `stardog.databases          `                | Arrays of databases that should be created at Helm install/upgrade. Existing databases are left untouched. |
| `stardog.databases[].options`                | Object Array of additional [Database Options](https://www.stardog.com/docs/6.0.0/#_configuration_options) (each with `name` and `value`) that should be applied when creating the database. |
| `stardog.javaArgs`                           | Custom java args, passed to the JVM |
| `stardog.securityContext`                    | K8s security context |
| `stardog.spatial.enabled`                    | Enables the geospatial search index for the database |
| `stardog.backup.databases[]`                 | List of DBs to be backed up. The schedule is in UTC |
| `stardog.backup.s3Region`                    | Region in which the S3 bucket is |
| `stardog.backup.s3BucketName`                | Name of the S3 bucket |
| `stardog.backup.s3PathPrefix`                | Path prefix to use. Defaults to release full name |
| `stardog.backup.s3AccessKey`                 | S3 API access key id |
| `stardog.backup.s3SecretKey`                 | S3 API access secret key |
| `stardog.backup.s3CredentialsExistingSecret` | Optionally existing secret with S3 credentials |
| `ingress.enabled`                            | If an ingress object should be created |
| `ingress.annotations`                        | Annotations to set on the ingress object |
| `ingress.host`                               | Host name which the ingress should resolve |
| `ingress.tls.enabled`                        | If TLS should be enabled on the ingress |
| `ingress.tls.secretName`                     | Name of the secret containing the TLS certificate and key |
| `persistence.enabled`                        | Enable persistence using PVC |
| `persistence.storageClass`                   | PVC storage class for Stardog data volume |
| `persistence.size`                           | PVC storage request size for Stardog data volume |
| `affinityTopologyKey`                        | Topology key on which pod anti affinity is configured |
| `zookeeper.enabled`                          | Enable ZooKeeper. Required for Stardog Cluster deployments |
| `zookeeper.replicaCount`                     | Number of ZooKeeper instances to run. Should always be an odd number > 3 |
| `zookeeper.updateStrategy`                   | [Update strategy](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#update-strategies) of the ZooKeeper StatefulSet |
| `zookeeper.allowAnonymousLogin`              | Allow anonymous logins to ZooKeeper |
| `zookeeper.auth.enabled`                     | Enable ZooKeeper authentication |
| `zookeeper.auth.clientPassword`              | Password to use for ZooKeeper clients. Will be generated if unset |
| `zookeeper.auth.serverPasswords`             | Password to use for ZooKeeper servers (delimited by `,`). Will be generated if unset |
| `zookeeper.metrics.enabled`                  | Enable ZooKeeper Proemtheus exporter |
| `zookeeper.metrics.prometheusOperator`       | Enable Proemtheus Operator integration (requires the operator to be installed) |

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
