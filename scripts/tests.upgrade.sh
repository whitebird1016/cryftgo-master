#!/usr/bin/env bash

set -euo pipefail

# e.g.,
# ./scripts/tests.upgrade.sh                                               # Use default version
# ./scripts/tests.upgrade.sh 1.11.0                                        # Specify a version
# CRYFTGO_PATH=./path/to/cryftgo ./scripts/tests.upgrade.sh 1.11.0 # Customization of cryftgo path
if ! [[ "$0" =~ scripts/tests.upgrade.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

# The CryftGo local network does not support long-lived
# backwards-compatible networks. When a breaking change is made to the
# local network, this flag must be updated to the last compatible
# version with the latest code.
#
# v1.11.0 activates Durango.
DEFAULT_VERSION="1.11.0"

VERSION="${1:-${DEFAULT_VERSION}}"
if [[ -z "${VERSION}" ]]; then
  echo "Missing version argument!"
  echo "Usage: ${0} [VERSION]" >>/dev/stderr
  exit 255
fi

CRYFTGO_PATH="$(realpath "${CRYFTGO_PATH:-./build/cryftgo}")"

#################################
# download cryftgo
# https://github.com/cryft-labs/cryftgo/releases
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
DOWNLOAD_URL=https://github.com/cryft-labs/cryftgo/releases/download/v${VERSION}/cryftgo-linux-${GOARCH}-v${VERSION}.tar.gz
DOWNLOAD_PATH=/tmp/cryftgo.tar.gz
if [[ ${GOOS} == "darwin" ]]; then
  DOWNLOAD_URL=https://github.com/cryft-labs/cryftgo/releases/download/v${VERSION}/cryftgo-macos-v${VERSION}.zip
  DOWNLOAD_PATH=/tmp/cryftgo.zip
fi

rm -f ${DOWNLOAD_PATH}
rm -rf "/tmp/cryftgo-v${VERSION}"
rm -rf /tmp/cryftgo-build

echo "downloading cryftgo ${VERSION} at ${DOWNLOAD_URL}"
curl -L "${DOWNLOAD_URL}" -o "${DOWNLOAD_PATH}"

echo "extracting downloaded cryftgo"
if [[ ${GOOS} == "linux" ]]; then
  tar xzvf ${DOWNLOAD_PATH} -C /tmp
elif [[ ${GOOS} == "darwin" ]]; then
  unzip ${DOWNLOAD_PATH} -d /tmp/cryftgo-build
  mv /tmp/cryftgo-build/build "/tmp/cryftgo-v${VERSION}"
fi
find "/tmp/cryftgo-v${VERSION}"

# Sourcing constants.sh ensures that the necessary CGO flags are set to
# build the portable version of BLST. Without this, ginkgo may fail to
# build the test binary if run on a host (e.g. github worker) that lacks
# the instructions to build non-portable BLST.
source ./scripts/constants.sh

#################################
echo "building upgrade.test"
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo@v2.13.1
ACK_GINKGO_RC=true ginkgo build ./tests/upgrade
./tests/upgrade/upgrade.test --help

#################################
# By default, it runs all upgrade test cases!
echo "running upgrade tests against the local cluster with ${CRYFTGO_PATH}"
./tests/upgrade/upgrade.test \
  --ginkgo.v \
  --cryftgo-path="/tmp/cryftgo-v${VERSION}/cryftgo" \
  --cryftgo-path-to-upgrade-to="${CRYFTGO_PATH}"
