# Helm Charts for APPUiO

[![License](https://img.shields.io/github/license/appuio/charts)](https://github.com/appuio/charts/blob/master/LICENSE)
[![Downloads](https://img.shields.io/github/downloads/appuio/charts/total)](https://github.com/appuio/charts/releases)

## Usage

Add the repo:

```
helm repo add appuio https://charts.appuio.ch
```

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
