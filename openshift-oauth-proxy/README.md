# OpenShift oauth-proxy

This Chart deploys the OpenShift oauth-proxy

## Configuration

The following table lists the configurable parameters of the OpenShift oauth-proxy chart. For defaults consult `values.yaml`

| Parameter | Description
| --- | ---
| `image.repository` | Image repository
| `image.tag` | Image tag
| `image.pullPolicy` | Image pull policy
| `extraArgs` | Extra arguments to pass to the oauth-proxy
| `resources` | Define resource limits
| `route.enabled` | Create route
| `route.annotations` | Add annotations to route
| `route.host` | Define route host
| `route.tls` | Enable tls for the route
