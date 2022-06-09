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
prepare: ## Build dependencies for this chart - pinned to index.yaml before upstream cleanup (see: https://github.com/bitnami/charts/pull/10530)
	helm repo add bitnami https://raw.githubusercontent.com/bitnami/charts/defb094c658024e4aa8245622dab202874880cbc/bitnami/
	helm dep build
