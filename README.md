# Helm Charts for APPUiO

This repository contains Helm Charts for running applications on APPUiO.

*Please Note*: This repository is in it's early stages. More to come.

## Usage

Add the repo:

```
helm repo add appuio http://charts.appuio.ch
```

## Add / Update Charts

```
helm package MYCHART
mv *.tgz docs/
helm repo index docs --url http://charts.appuio.ch
git add -i
git commit -av
git push
```
