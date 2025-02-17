#!/usr/bin/env bash

set -eo pipefail

# clone dependency proto files
COSMOS_URL=github.com/cosmos/cosmos-sdk
IBC_URL=github.com/cosmos/ibc-go
IBC_V=v8
INITIA_URL=github.com/initia-labs/initia
OPINIT_URL=github.com/initia-labs/OPinit
INDEXER_URL=github.com/initia-labs/kvindexer
CONNECT_URL=github.com/skip-mev/connect
CONNECT_V=v2

COSMOS_SDK_VERSION=$(grep -o "$COSMOS_URL v[^\ ]*" ./go.mod)
IBC_VERSION=$(grep -o "$IBC_URL/$IBC_V v[^\ ]*" ./go.mod)
INITIA_VERSION=$(grep -o "$INITIA_URL v[^\ ]*" ./go.mod)
OPINIT_VERSION=$(grep -o "$OPINIT_URL v[^\ ]*" ./go.mod)
INDEXER_VERSION=$(grep -o "$INDEXER_URL v[^\ ]*" ./go.mod)
CONNECT_VERSION=$(grep -o "$CONNECT_URL/$CONNECT_V v[^\ ]*" ./go.mod)

mkdir -p ./third_party
cd third_party
git clone -b $INITIA_VERSION https://$INITIA_URL
git clone -b $OPINIT_VERSION https://$OPINIT_URL
git clone -b $COSMOS_SDK_VERSION https://$COSMOS_URL
git clone -b $IBC_VERSION https://$IBC_URL
git clone -b $INDEXER_VERSION https://$INDEXER_URL
git clone -b $CONNECT_VERSION https://$CONNECT_URL
cd ..

# start generating
mkdir -p ./tmp-swagger-gen
cd proto
proto_dirs=$(find \
  ./minievm \
  ../third_party/cosmos-sdk/proto/cosmos \
  ../third_party/ibc-go/proto/ibc \
  ../third_party/initia/proto \
  ../third_party/opinit/proto \
  ../third_party/kvindexer/proto \
  ../third_party/connect/proto \
  -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done
cd ..

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen

# clean third party files
rm -rf ./third_party
