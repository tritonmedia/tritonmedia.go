#!/usr/bin/env bash
#
# Pulls down the latest protobufs

set -e

PROTO_DIR="_proto"

rm -rf "${PROTO_DIR}/api"
mkdir -p "${PROTO_DIR}"

echo "fetching latest protobufs ..."

pushd "${PROTO_DIR}" >/dev/null || exit 1
git clone git@github.com:tritonmedia/core _core
mv _core/protos/api api
rm -rf _core
popd >/dev/null|| exit 1
echo "update protobufs"