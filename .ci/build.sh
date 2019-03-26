#!/bin/bash
set -e

# set to one of the following values
# - don't use HelmQA: (empty)
# - use HelmQA in SaaS mode: 'http://helmqa-zhaw-prod-demos.appuioapp.ch/livecheck?repo'
# - use HelmQA locally in Docker mode: 'http://127.0.0.1:5000/livecheck?repo'
HELMQA_URL='http://helmqa-zhaw-prod-demos.appuioapp.ch/livecheck?repo'
REPO_URL="${REPO_URL:-https://charts.appuio.ch}"
echo "----> Deploying to $REPO_URL"

GIT_REPO="$(git config remote.origin.url)"

if [ -n "$HELMQA_URL" ]; then
  did=
  if [[ "$HELMQA_URL" =~ 127.0.0.1 ]]; then
    docker pull docker.io/jszhaw/helmqa
    did="$(docker run -d -p 127.0.0.1:5000:5000 docker.io/jszhaw/helmqa)"
  fi

  while true; do
    sleep 1
    tresponse="$(curl --connect-timeout 5 --max-time 60 $HELMQA_URL=$GIT_REPO)" || true
    if [ -n "$tresponse" ]; then
      break
    fi
    echo "Still spinning up HelmQA container image..."
  done

  if [[ $? != 28 ]]; then
    tstatus=$(echo "$tresponse" | jq -r .status)

    if [[ "$tstatus" == "fail" ]]; then
      echo "HelmQA test failed. Check response"
      echo "$tresponse"
      exit 1
    elif [[ "$tstatus" == "success" ]]; then
      echo "HelmQA test succeeded. No issues found"
    fi
  else
    echo "HelmQA connection failed. Timed out!"
  fi

  if [ -n "$did" ]; then
    docker kill "$did"
  fi
else
  echo "Skipping HelmQA test."
fi

tmp="$(mktemp -d)"
echo "----> Working in $tmp"

cleanup() {
  rm -rf "$tmp"
}
trap cleanup 0 1 2 3 6 15

charts=$(find ./* -maxdepth 1 -name Chart.yaml -exec dirname "{}" \;)
changed_files=$(git diff --name-only $TRAVIS_COMMIT_RANGE)

echo "Changed files: $changed_files"

for chart in $charts; do
  chart=${chart:2}
  chart_changed=false
  for file in $changed_files; do
    if [[ $file == "$chart/"* ]]; then
      chart_changed=true
      break
    fi
  done

  if [[ $chart_changed == false ]]; then
    echo "No changes for chart $chart"
    continue
  fi

  echo "----> Build dependencies for ${chart}"
  helm dependency build "$chart"

  helm lint "$chart"

  echo "----> Packaging ${chart}"
  helm package \
    --destination="$tmp" \
    --save=false \
    "$chart"
done

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
