# stardog

![Version: 0.18.0](https://img.shields.io/badge/Version-0.18.0-informational?style=flat-square) ![AppVersion: 9.0.0](https://img.shields.io/badge/AppVersion-9.0.0-informational?style=flat-square)

Stardog is the world’s leading knowledge graph platform for the enterprise. Stardog makes it fast and easy to turn enterprise data into knowledge.

**Homepage:** <https://www.stardog.com/>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install stardog appuio/stardog
```
Don't forget to create a secret with your Stardog license first:

```console
kubectl create secret generic stardog-license --from-file stardog-license-key.bin=/path/to/your/license.bin
```

<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Configuration

When using Stardog in clustered mode with `zookeeper.enabled=true`, it's best to also enable PodDisruptionBudgets:

```yaml
zookeeper:
  enabled: true
  pdb:
    create: true
    maxUnavailable: 1
```

Otherwise events such as draining a node can leave the cluster in a state where no cluster coordinator is left in charge, which often makes it impossible for the cluster to reassamble itself once pods become ready again.

## Values

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
| `stardog.backup.databases[]`                 | List of DBs to be backed up. The schedule is in UTC |
| `stardog.backup.s3Region`                    | Region in which the S3 bucket is |
| `stardog.backup.s3BucketName`                | Name of the S3 bucket |
| `stardog.backup.s3PathPrefix`                | Path prefix to use. Defaults to release full name |
| `stardog.backup.s3AccessKey`                 | S3 API access key id |
| `stardog.backup.s3SecretKey`                 | S3 API access secret key |
| `stardog.backup.s3CredentialsExistingSecret` | Optionally existing secret with S3 credentials |
| `stardog.memory.spilling_max_file_length`    | Configure the maximum file length when spilling data to disk |
| `ingress.enabled`                            | If an ingress object should be created |
| `ingress.annotations`                        | Annotations to set on the ingress object |
| `ingress.host`                               | Host name which the ingress should resolve |
| `ingress.tls.enabled`                        | If TLS should be enabled on the ingress |
| `ingress.tls.secretName`                     | Name of the secret containing the TLS certificate and key |
| `metrics.stardogRules.extraLabels`        | A hash of labels to add to Stardog's PrometheusRules |
| `metrics.zookeeperRules.extraLabels`        | A hash of labels to add to ZooKeeper's PrometheusRules |
| `persistence.enabled`                        | Enable persistence using PVC |
| `persistence.storageClass`                   | PVC storage class for Stardog data volume |
| `persistence.size`                           | PVC storage request size for Stardog data volume |
| `affinityTopologyKey`                        | Topology key on which pod anti affinity is configured |
| `zookeeper.enabled`                          | Enable ZooKeeper. Required for Stardog Cluster deployments |
| `zookeeper.replicaCount`                     | Number of ZooKeeper instances to run. Should always be an odd number > 3 |
| `zookeeper.updateStrategy`                   | [Update strategy](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#update-strategies) of the ZooKeeper StatefulSet |
| `zookeeper.auth.enabled`                     | Enable ZooKeeper authentication |
| `zookeeper.auth.clientPassword`              | Password to use for ZooKeeper clients. Will be generated if unset |
| `zookeeper.auth.serverPasswords`             | Password to use for ZooKeeper servers (delimited by `,`). Will be generated if unset |
| `zookeeper.metrics.enabled`                  | Enable ZooKeeper Prometheus exporter |
| `zookeeper.metrics.prometheusOperator`       | Enable Prometheus Operator integration (requires the operator to be installed) |
| `zookeeper.sessionTimeout`                   | Set the ZooKeeper [session timeout](https://docs.stardog.com/cluster/installation-and-setup/#connectionsession-timeouts) |
| `zookeeper.autopurge.snapRetainCount`        | How many snapshots ZooKeeper should keep when [autopurging](https://zookeeper.apache.org/doc/r3.4.5/zookeeperAdmin.html#sc_strengthsAndLimitations) |
| `zookeeper.autopurge.purgeInterval`        | Interval (in hours) for ZooKeeper autopurge |

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://raw.githubusercontent.com/bitnami/charts/b661ad0e6e3277fc5b5b2db197cba47174693712/bitnami/ | zookeeper | 8.1.2 |

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
