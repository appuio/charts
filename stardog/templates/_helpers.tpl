{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "stardog.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "stardog.fullname" -}}
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
{{- define "stardog.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "stardog.imagePullSecret" -}}
{{- printf "{\"auths\": {\"%s\": {\"auth\": \"%s\"}}}" .Values.image.registry (printf "%s:%s" .Values.image.username .Values.image.password | b64enc) | b64enc }}
{{- end }}

{{- define "stardog.zookeeper.fullname" -}}
{{- $zookeeperContext := dict "Values" .Values.zookeeper "Release" .Release "Chart" (dict "Name" "zookeeper") -}}
{{ template "zookeeper.fullname" $zookeeperContext }}
{{- end -}}

{{/*
Create ZooKeeper connection string.
*/}}
{{- define "stardog.zookeeperConnection" -}}
{{- $zookeeperFullName := (include "stardog.zookeeper.fullname" $) }}
{{- range $i, $e := until (int .Values.zookeeper.replicaCount) -}}
{{- printf "%v-%v.%v-headless:%d," $zookeeperFullName $i $zookeeperFullName (int $.Values.zookeeper.service.port) -}}
{{- end -}}
{{- end -}}
