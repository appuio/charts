MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help
.DELETE_ON_ERROR:
.SUFFIXES:

# git tag of the version to fetch
# defaults to tag specified in values.yaml, as this gets maintained by renovate
APP_VERSION ?= $(shell grep -E "^  tag: " values.yaml | cut -d " " -f 4 | sed s/\"//g)

SED_BIN := sed

.PHONY: help
help: ## Show this help
	@grep -E -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(: ).*?## "}; {gsub(/\\:/,":",$$1)}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: appVersion
appVersion: ## Replaces "appVersion" in Chart.yaml with 'image.tag' from values.yaml
	$(SED_BIN) -i -e 's/^appVersion.*/appVersion: $(APP_VERSION)/' Chart.yaml

#
# "Interface" for parent Makefile
#
prepare: appVersion ## Prepare helm chart
