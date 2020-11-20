#!/bin/bash
set -e

REPO_URL="${REPO_URL:-https://charts.appuio.ch}"
echo "----> Deploying to $REPO_URL"

GIT_REPO="$(git config remote.origin.url)"

tmp="$(mktemp -d)"
echo "----> Working in $tmp"

cleanup() {
  rm -rf "$tmp"
}
trap cleanup 0 1 2 3 6 15

# Dependency repos (we still have to download them manually, even if they are added in requirements.yaml)
# (https://github.com/helm/helm/issues/6005)
helm repo add stable https://charts.helm.sh/stable
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add codecentric https://codecentric.github.io/helm-charts
helm repo add jetstack https://charts.jetstack.io
helm repo add flux https://fluxcd.github.io/flux

charts=$(find ./* -maxdepth 1 -name Chart.yaml -exec dirname "{}" \;)
changed_files=$(git diff --name-only $TRAVIS_COMMIT_RANGE)
any_chart_changed=false

echo "Changed files: $changed_files"

for chart in $charts; do
  chart=${chart:2}
  chart_changed=false

  for file in $changed_files; do
    if [[ $file != "$chart/test/"* && $file == "$chart/"* ]]; then
      chart_changed=true
      break
    fi
  done

  if [[ $chart_changed == false ]]; then
    echo "No changes for chart $chart"
    continue
  fi

  any_chart_changed=true

  echo "----> Build dependencies for ${chart}"
  helm dependency build "$chart"

  helm lint "$chart"

  # Run go test if there are chart unit tests
  if [[ -f "$chart/test/go.mod" ]]; then
    pushd "$chart/test"
    echo "----> Running unit tests for ${chart}"
    go test ./...
    popd
  fi

  echo "----> Packaging ${chart}"
  helm package \
    --destination="$tmp" \
    "$chart"
done

if [[ $any_chart_changed == false ]]; then
  echo "No charts changes at all. Nothing more to do."
  exit 0
fi

echo "----> Check out gh-pages branch"
git clone --depth=1 "$GIT_REPO" --branch=gh-pages "$OUT_DIR"

echo "----> Moving new charts into repo"
mv -v "${tmp}"/*.tgz "${OUT_DIR}/"

pushd "$OUT_DIR"
new_charts="$(git status --short | grep '^??' --count ||:)"
modified_charts="$(git status --short | grep '^.M' --count ||:)"

if [ "$modified_charts" -ne 0 ]; then
  echo 'ERROR: Chart version number was not updated'
  git status --short | grep '^.M'
  exit 1
fi

if [ "$new_charts" -eq 0 ]; then
  echo '----> No new charts, exiting'
  exit 0
fi

echo '----> Reindexing Helm repository'
helm repo index --url="$REPO_URL" .

git diff
