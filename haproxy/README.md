# Customizable HAProxy

This Helm chart uses the [HAProxy](http://www.haproxy.org/) image from [haproxytech](https://github.com/haproxytech/haproxy-docker-debian) based on the Debian Linux Distribution. Its configurable by passing values to a configmap which contains haproxy.cfg - the configuration file for HAProxy.

## TL;DR

```console
helm repo add appuio https://charts.appuio.ch
helm install appuio/haproxy
```

## Introduction

The Debian based image was choosen because it already contains the ca-certificates which makes it
easy to use HAProxy as a reverse TLS Proxy.

## Installing the Chart

To install the chart with the release name `haproxy`:

```console
helm install --name haproxy appuio/haproxy
```

## Uninstalling the Chart

To uninstall/delete the `haproxy` deployment:

```console
helm delete haproxy
```

## Configuration

The configuration of HAProxy is derived through a configmap which in turn can be configured by adjusting the corresponding values in values.yaml.  
The present example is a configmap to configure HAProxy as a TLS reverse proxy to communicate with an LDAP backend. For different purposes you can  
create your own configmap and values.  

| Parameter              | Description            | Default
|---                     | ---                    | ---
| `replicaCount`         | Number of replicas for the proxy. | 1
| `image.repository`     | Repository for the base image | docker.io/haproxytech/haproxy-debian
| `image.tag`            | Tag of the base image | 2.1
| `image.pullPolicy`     | Determines when the image is pulled | IfNotPresent
| `ingress.enabled`      | If an ingress object should be created | false
| `ingress.annotations`  | Annotations to set on the ingress object |
| `ingress.host`         | Host name which the ingress should resolve |
| `ingress.tls.enabled`  | If TLS should be enabled on the ingress |
| `ingress.tls.secretName`| Name of the secret containing the TLS certificate and key |
| `haproxy.frontendPort` | Port for the Proxy to listen on the frontend | 30636
| `haproxy.config`       | Suffix of the used config-map | ldap-tls
| `haproxy.ldapTls.backend` | Service the proxy should connect to | none
| `haproxy.ldapTls.certificatePath` | Path to the corresponding root-certificate | /etc/ssl/certs/ca-certificates.crt   
| `nodeSelector` | Pod node selector | `{}`
| `tolerations` | Pod tolerations | `[]`
| `affinity` | Pod affinity rules | `{}`
