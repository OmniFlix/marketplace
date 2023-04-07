#!/usr/bin/env bash

set -eo pipefail
PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts_proto"

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname |
sort | uniq)
for dir in $proto_dirs; do
  protoc \
  -I "proto" \
  -I "third_party/proto" \
  --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
  --ts_out=build  $(find "${dir}" -maxdepth 1 -name '*.proto')
done
