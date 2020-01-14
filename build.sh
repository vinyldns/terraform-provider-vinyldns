#!/usr/bin/env bash

# Safety forced
set -exuo pipefail

# Ensure we leave the user's prompt whereever they were before
function cleanup {
  popd
}
trap cleanup EXIT

# Declare an anchor so relative paths aren't an issue
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
pushd $DIR

# Cleanup old artifacts
find build -not -name build -not -name '.dockerignore' -not -name '.gitignore' -print
find build -not -name build -not -name '.dockerignore' -not -name '.gitignore' -delete
find release -not -name release -not -name '.dockerignore' -not -name '.gitignore' -print
find release -not -name release -not -name '.dockerignore' -not -name '.gitignore' -delete

# Build artifacts
docker-compose build --no-cache builder
docker-compose up --force-recreate --abort-on-container-exit --exit-code-from builder builder
docker-compose down
