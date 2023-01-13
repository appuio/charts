# Contributing Guidelines

The APPUiO Charts project accepts contributions via GitHub pull requests. This document outlines the process to help get your contribution accepted.

## How to Contribute a Chart

1. Fork this repository, develop and test your Chart.
1. Ensure your Chart follows the [technical](#technical-requirements) and [documentation](#documentation-requirements) guidelines, described below.
1. Submit a pull request.

***NOTE***: In order to make testing and merging of PRs easier, please submit changes to multiple charts in separate PRs.

### Technical requirements

* All Chart dependencies should also be submitted independently
* Must pass the linter (`helm lint`)
* Must include source GitHub repositories for images used in the Chart
* Images should not have any major security vulnerabilities
* Must be up-to-date with the latest stable Helm/Kubernetes features
  * Use Deployments in favor of ReplicationControllers
* Should follow Kubernetes best practices
  * Include Health Checks wherever practical
  * Allow configurable [resource requests and limits](http://kubernetes.io/docs/user-guide/compute-resources/#resource-requests-and-limits-of-pod-and-container)
* Provide a method for data persistence (if applicable)
* Support application upgrades
* Allow customization of the application configuration
* Provide a secure default configuration
* Do not leverage alpha features of Kubernetes
* Includes a [NOTES.txt](https://helm.sh/docs/topics/charts/#chart-license-readme-and-notes) explaining how to use the application after install
* Follows [best practices](https://helm.sh/docs/chart_best_practices/)
  (especially for [labels](https://helm.sh/docs/chart_best_practices/labels/)
  and [values](https://helm.sh/docs/chart_best_practices/values/))

### Documentation requirements

* Must include an in-depth `README.md`, including:
  * Short description of the Chart
  * Any prerequisites or requirements
  * Customization: explaining options in `values.yaml` and their defaults
* Must include a short `NOTES.txt`, including:
  * Any relevant post-installation information for the Chart
  * Instructions on how to access the application or service provided by the Chart
