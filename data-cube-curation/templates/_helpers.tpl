{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "data-cube-curation.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "data-cube-curation.fullname" -}}
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
{{- define "data-cube-curation.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "data-cube-curation.labels" -}}
helm.sh/chart: {{ include "data-cube-curation.chart" . }}
{{ include "data-cube-curation.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "data-cube-curation.selectorLabels" -}}
app.kubernetes.io/name: {{ include "data-cube-curation.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Secret names
*/}}
{{- define "data-cube-curation.sparql.secretName" -}}
{{- if .Values.dataCubeCuration.sparql.existingSecret -}}
{{- .Values.dataCubeCuration.sparql.existingSecret -}}
{{- else -}}
{{- include "data-cube-curation.fullname" . -}}-sparql
{{- end -}}
{{- end -}}

{{- define "data-cube-curation.gitlab.secretName" -}}
{{- if .Values.dataCubeCuration.gitlab.existingSecret -}}
{{- .Values.dataCubeCuration.gitlab.existingSecret -}}
{{- else -}}
{{- include "data-cube-curation.fullname" . -}}-gitlab
{{- end -}}
{{- end -}}

{{- define "data-cube-curation.s3.secretName" -}}
{{- if .Values.dataCubeCuration.s3.existingSecret -}}
{{- .Values.dataCubeCuration.s3.existingSecret -}}
{{- else -}}
{{- include "data-cube-curation.fullname" . -}}-s3
{{- end -}}
{{- end -}}

{{- define "data-cube-curation.baseURI" -}}
{{- if .Values.dataCubeCuration.baseURI -}}
{{- .Values.dataCubeCuration.baseURI -}}
{{- else if .Values.ingress.enabled -}}
http{{ if .Values.ingress.tls }}s{{ end }}://{{- (.Values.ingress.hosts | first).host }}/
{{- else -}}
{{ fail "Please either set the dataCubeCuration.baseURI value or enable the ingress" }}
{{- end -}}
{{- end -}}
