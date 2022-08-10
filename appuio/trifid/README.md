# trifid

![Version: 1.2.3](https://img.shields.io/badge/Version-1.2.3-informational?style=flat-square) ![AppVersion: 2.3.7](https://img.shields.io/badge/AppVersion-2.3.7-informational?style=flat-square)

Trifid provides a lightweight and easy way to access Linked Data URIs via HTTP.

**Homepage:** <https://github.com/zazuko/trifid>

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install trifid appuio/trifid
```
<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Configuration

The following table lists the configurable parameters chart. For default values consult `values.yaml`.

| Parameter                                    | Description |
| ---                                          | --- |
| `replicaCount`                               | Number of trifid instances to run |
| `serviceAccount.create`                      | Specifies whether a service account should be created |
| `serviceAccount.name`                        | The name of the service account to use |
| `imagePullSecrets`                           | The image pull secrets to use |
| `image.registry`                             | The Docker registry to pull from |
| `image.repository`                           | The Docker image repostory to pull |
| `image.tag`                                  | The Docker image tag to pull |
| `podSecurityContext`                         | The security context which should be applied to the pod |
| `securityContext`                            | The security context which should be applied to the container |
| `pdb.enabled`                                | Create a pod disruption budget |
| `pdb.minAvailable`                           | How many pods need to be available (absolute number or in percentage) |
| `trifid.sparql.endpoint.url`                 | URL of the SPARQL endpoint to use |
| `trifid.sparql.endpoint.username`            | Username used to connect to the SPARQL endpoint |
| `trifid.sparql.endpoint.password`            | Password used to connect to the SPARQL endpoint |
| `trifid.sparql.endpoint.secretName`          | Name of an existing secret containing `username` and `password` of the SPARQL endpoint |
| `ingress.enabled`                            | If an ingress object should be created |
| `ingress.annotations`                        | Annotations to set on the ingress object |
| `ingress.hosts.host`                         | Host name which the ingress should resolve |
| `ingress.hosts.paths`                        | Paths which the ingress should resolve |
| `ingress.tls.enabled`                        | If TLS should be enabled on the ingress |
| `ingress.tls.secretName`                     | Name of the secret containing the TLS certificate and key |
| `ingress.tls.hosts`                          | Hosts to use the certificate secret for |
| `resources`                                  | Pod resource requests and limits |
| `nodeSelector`                               | Pod node selector |
| `tolerations`                                | Pod tolerations |

## Source Code

* <https://github.com/zazuko/trifid>

<!---
Common/Useful Link references from values.yaml
-->
[resource-units]: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes
[prometheus-operator]: https://github.com/coreos/prometheus-operator
