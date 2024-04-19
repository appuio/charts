# Helm Charts for APPUiO

[![License](https://img.shields.io/github/license/appuio/charts)](https://github.com/appuio/charts/blob/master/LICENSE)
[![Downloads](https://img.shields.io/github/downloads/appuio/charts/total)](https://github.com/appuio/charts/releases)

## Usage

Add the repo:

```
helm repo add appuio https://charts.appuio.ch
```

## List of Charts

| Downloads & Changelog | Chart |
| --- | --- |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/cloud-portal-0.4.1/total)](https://github.com/appuio/charts/releases/tag/cloud-portal-0.4.1) | [cloud-portal](appuio/cloud-portal/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/data-cube-curation-0.3.1/total)](https://github.com/appuio/charts/releases/tag/data-cube-curation-0.3.1) | [data-cube-curation](appuio/data-cube-curation/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/exoip-1.0.4/total)](https://github.com/appuio/charts/releases/tag/exoip-1.0.4) | [exoip](appuio/exoip/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/generic-0.1.2/total)](https://github.com/appuio/charts/releases/tag/generic-0.1.2) | [generic](appuio/generic/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/haproxy-2.5.0/total)](https://github.com/appuio/charts/releases/tag/haproxy-2.5.0) | [haproxy](appuio/haproxy/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/mariadb-galera-1.2.4/total)](https://github.com/appuio/charts/releases/tag/mariadb-galera-1.2.4) | [mariadb-galera](appuio/mariadb-galera/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/maxscale-2.0.1/total)](https://github.com/appuio/charts/releases/tag/maxscale-2.0.1) | [maxscale](appuio/maxscale/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/metrics-server-2.12.1/total)](https://github.com/appuio/charts/releases/tag/metrics-server-2.12.1) | [metrics-server](appuio/metrics-server/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/openshift-oauth-proxy-0.2.3/total)](https://github.com/appuio/charts/releases/tag/openshift-oauth-proxy-0.2.3) | [openshift-oauth-proxy](appuio/openshift-oauth-proxy/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/openshift-route-1.1.4/total)](https://github.com/appuio/charts/releases/tag/openshift-route-1.1.4) | [openshift-route](appuio/openshift-route/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/prometheus-blackbox-exporter-0.3.1/total)](https://github.com/appuio/charts/releases/tag/prometheus-blackbox-exporter-0.3.1) | [prometheus-blackbox-exporter](appuio/prometheus-blackbox-exporter/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/redis-1.3.5/total)](https://github.com/appuio/charts/releases/tag/redis-1.3.5) | [redis](appuio/redis/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/secret-1.1.0/total)](https://github.com/appuio/charts/releases/tag/secret-1.1.0) | [secret](appuio/secret/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/signalilo-0.12.1/total)](https://github.com/appuio/charts/releases/tag/signalilo-0.12.1) | [signalilo](appuio/signalilo/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/snappass-1.0.0/total)](https://github.com/appuio/charts/releases/tag/snappass-1.0.0) | [snappass](appuio/snappass/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/stardog-0.26.0/total)](https://github.com/appuio/charts/releases/tag/stardog-0.26.0) | [stardog](appuio/stardog/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/stardog-userrole-operator-0.2.3/total)](https://github.com/appuio/charts/releases/tag/stardog-userrole-operator-0.2.3) | [stardog-userrole-operator](appuio/stardog-userrole-operator/README.md) |
| [![chart downloads](https://img.shields.io/github/downloads/appuio/charts/trifid-2.0.2/total)](https://github.com/appuio/charts/releases/tag/trifid-2.0.2) | [trifid](appuio/trifid/README.md) |

## Add / Update Charts

New charts and versions will be built and published automatically as GitHub Releases. All charts use the Semantic Versioning release strategy.

Each PR shall be labelled with the chart name using the pattern `chart/<chart-name>`. The labels help building a changelog when releaseing a new chart version. Please request a new label if the label is missing for a new chart.

The following GitHub pull request labels require certain SemVer version increments:

| Label | SemVer increment | Usage examples |
| --- | --- | --- |
| `bug` | Patch | Fix an unintended behaviour |
| `change` | Patch | Reimplementation of existing features, code improvements |
| `change` | Minor | Deprecate an existing feature |
| `dependency` | Patch | Patches or minor updates from a dependency that causes no known significant change |
| `dependency` | Major | Dependency update that causes breaking changes for the users |
| `enhancement` | Minor | New features |
| `breaking` | Major | See below |

### Breaking changes

Pull request authors and reviewers are obliged to correctly identify breaking changes.
The following changes are examples of breaking changes:

* Changes of existing parameters in default configuration (new optional or default parameters should’t be breaking).
* Changes that require manual upgrade steps when upgrading the chart version (can also be caused by dependency updates).
* Removal of existing behaviour or features.
