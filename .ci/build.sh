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
if git status -s | grep -v '^??' -q; then
  echo '----> No new charts, exiting'
  exit 0
fi

echo '----> Reindexing Helm repository'
helm repo index --url="$REPO_URL" .

popd

#echo '----> Commiting changes'
#new_charts="$(git status --short | awk '/^\?\?/{print "*", $2}')"
#commit_message="Adding new charts
#
#New charts:
#$new_charts
#"
#git add .
#git commit -m "$commit_message"
#git push origin gh-pages
