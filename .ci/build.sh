#!/bin/bash
set -e

HELMQA_URL=https://helmqa-helmqa.a3c1.starter-us-west-1.openshiftapps.com/livecheck?repo
REPO_URL="${REPO_URL:-https://charts.appuio.ch}"
echo "----> Deploying to $REPO_URL"

GIT_REPO="$(git config remote.origin.url)"

HELM_RESPONSE=$(curl --max-time 10 "$HELMQA_URL=$GIT_REPO")

HELM_STATUS=$("$HELM_RESPONSE" | jq -r .status)
HELM_RESPONSE_CODE=$("$HELM_RESPONSE" | jq -r .code)

echo "-------------------> $HELM_STATUS"
echo "-------------------> $HELM_RESPONSE_CODE"


tmp="$(mktemp -d)"
echo "----> Working in $tmp"

cleanup() {
  rm -rf "$tmp"
}
trap cleanup 0 1 2 3 6 15

charts=$(find ./* -maxdepth 1 -name Chart.yaml -exec dirname "{}" \;)
for chart in $charts; do
  echo "----> Packaging ${chart}"
  helm package \
    --destination="$tmp" \
    --save=false \
    "$chart"
done

echo "----> Check out gh-pages branch"
git clone --depth=1 "$GIT_REPO" --branch=gh-pages "$OUT_DIR"

echo "----> Moving new charts into repo"
mv -nv "${tmp}"/*.tgz "${OUT_DIR}/"

pushd "$OUT_DIR"
new_charts="$(git status --short | grep '^??' --count ||:)"
if [ "$new_charts" -eq 0 ]; then
  echo '----> No new charts, exiting'
  exit 0
fi

echo '----> Reindexing Helm repository'
helm repo index --url="$REPO_URL" .
