#!/bin/bash

# Tag format from source: refs/tags/chart-1.2.3

# Remove "refs/tags/"
tag="${GITHUB_REF##*/}"

# Remove SemVer at the end
chart=$(echo ${tag} | grep --perl-regexp --only-matching '^([a-zA-Z0-9-]+)(?![0-9.]+)')

tagPattern="${chart}-(.+)"
chartLabel="chart/${chart}"

echo ::group::Configuring changelog generator
jq '.tag_resolver.filter.pattern="'$tagPattern'" | .tag_resolver.transformer.pattern="'$tagPattern'" | .categories[].labels += ["'$chartLabel'"]' \
  .github/changelog-configuration.json | tee .github/configuration.json
echo ::endgroup::
