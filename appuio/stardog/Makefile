MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help
.DELETE_ON_ERROR:
.SUFFIXES:

.PHONY: help
help: ## Show this help
	@grep -E -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(: ).*?## "}; {gsub(/\\:/,":",$$1)}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

#
# "Interface" for parent Makefile
#
prepare:
	helm repo add bitnami-b661ad0 https://raw.githubusercontent.com/bitnami/charts/b661ad0e6e3277fc5b5b2db197cba47174693712/bitnami/
	helm dep build
