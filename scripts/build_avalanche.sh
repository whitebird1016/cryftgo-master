#!/usr/bin/env bash

set -euo pipefail

print_usage() {
  printf "Usage: build_avalanche [OPTIONS]

  Build cryftgo

  Options:

    -r  Build with race detector
"
}

race=''
while getopts 'r' flag; do
  case "${flag}" in
    r) race='-race' ;;
    *) print_usage
      exit 1 ;;
  esac
done

# Cryftgo root folder
AVALANCHE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the constants
source "$AVALANCHE_PATH"/scripts/constants.sh

build_args="$race"
echo "Building CryftGo..."
go build $build_args -ldflags "-X github.com/cryft-labs/cryftgo/version.GitCommit=$git_commit $static_ld_flags" -o "$cryftgo_path" "$AVALANCHE_PATH/main/"*.go
