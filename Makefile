SHELL := /usr/bin/env bash

MASTER_BRANCH=master
SOURCE_README=README.gotmpl
TARGET_README=README.md
CHARTS_DIR=charts/

HELM_DOCS_VERSION=v1.5.0

.PHONY: help
help: ## Show this help
	@grep -E -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(: ).*?## "}; {gsub(/\\:/,":",$$1)}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: docs
docs: docs\:helm

.PHONY: docs\:helm
docs\:helm: ## Creates the Chart READMEs from template and values.yaml files
	@echo --- Generating Chart READMEs
	@docker run --rm -v $$(pwd):/helm-docs -u $$(id -u) jnorwood/helm-docs:$(HELM_DOCS_VERSION) \
		--template-files ./.github/helm-docs-header.gotmpl.md \
		--template-files README.gotmpl.md \
		--template-files ./.github/helm-docs-footer.gotmpl.md

.PHONY: test
test: ## Run Chart unit tests
	@echo --- Executing unit tests
	@go test ./...

.PHONY: lint\:fmt
lint\:fmt: ## Run go fmt against code
	go fmt ./...

.PHONY: lint\:vet
lint\:vet: ## Run go vet against code
	go vet ./...

.PHONY: lint
lint: lint\:fmt lint\:vet ## All-in-one linting and checks for uncommitted changes
	@echo 'Check for uncommitted changes ...'
	git diff --exit-code
