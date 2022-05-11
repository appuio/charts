```bash
kubectl apply -f https://github.com/vshn/appcat-service-postgresql/releases/download/provider-postgresql-{{ template "chart.version" . }}/crds.yaml
```

<!---
The README.md file is automatically generated with helm-docs!

Edit the README.gotmpl.md template instead.
-->

## Handling CRDs

* Always upgrade the CRDs before upgrading the Helm release.
* Watch out for breaking changes in the {{ title .Name }} release notes.

{{ template "chart.sourcesSection" . }}

{{ template "chart.requirementsSection" . }}
<!---
The values below are generated with helm-docs!

Document your changes in values.yaml and let `make chart-docs` generate this section.
-->
{{ template "chart.valuesSection" . }}
