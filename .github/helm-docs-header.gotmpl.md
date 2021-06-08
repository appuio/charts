{{ template "chart.header" . }}
{{ template "chart.deprecationWarning" . }}

{{ template "chart.badgesSection" . }}

{{ template "chart.description" . }}

{{ template "chart.homepageLine" . }}

## Installation

```bash
helm repo add appuio https://charts.appuio.ch
helm install {{ template "chart.name" . }} appuio/{{ template "chart.name" . }}
```
