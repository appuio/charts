{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "haproxy.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "haproxy.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "haproxy.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
HAProxy config for stats and metrics
*/}}
{{- define "haproxy.metricsConfig" -}}
{{- if .Values.metrics.enabled }}
frontend stats
  mode http
  bind *:9000
  stats enable
  stats uri /stats
  stats refresh 3s
  acl server_stopping stopping
  monitor-uri /healthz
  monitor fail if server_stopping
  option http-use-htx
  option dontlog-normal
  option httplog
  http-request use-service prometheus-exporter if { path /metrics }
{{- end }}
{{- end -}}

{{/*
HAProxy config for dns resolver
*/}}
{{- define "haproxy.resolversConfig" -}}
resolvers mydns
  parse-resolv-conf
  accepted_payload_size 8192 # allow larger DNS payloads
  hold valid 3s
  hold timeout 3s
  hold other 3s
  resolve_retries 999
{{- end -}}
