# SnapPass Chart

[SnapPass](https://github.com/pinterest/snappass) - A web app that lets you share passwords securely.

## TL;DR;

```console
helm install appuio/snappass
```

## Configuration

The following table lists the configurable chart specific parameters. For default values consult `values.yaml`.

| Parameter                                    | Description |
| ---                                          | --- |
| `redis.enabled`                              | Enables the Redis Helm chart dependency |
| `snappass.debug.enabled`                     | If the debug mode should be enabled |
| `snappass.ssl.enabled`                       | Enables SSL on application level |
| `snappass.staticUrl`                         | Location of static assets |
| `snappass.urlPrefix`                         | URL Prefix (e.g. when running SnapPass behind a reverse proxy) |
| `snappass.secretKey`                         | Secret key used to sign cookies |
| `snappass.existingSecret`                    | Existing secret with a `secretKey` field, used to sign cookies |
| `snappass.redis.host`                        | Redis host |
| `snappass.redis.port`                        | Redis port |
| `snappass.redis.password`                    | Redis password |
| `snappass.redis.existingSecret`              | Existing secret with a `redis-password` field, used to authenticate against Redis |
| `snappass.redis.db`                          | Redis DB |
| `snappass.redis.prefix`                      | Redis prefix |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
helm install --name my-release appuio/snappass
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```console
helm install --name my-release -f values.yaml appuio/snappass
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
cd snappass
helm dep build
cd test
go test -v ./...
```
