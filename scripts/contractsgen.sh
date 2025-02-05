#!/bin/bash
set -e
BUILD_DIR=build
CONTRACTS_DIR=x/evm/contracts
VERSION="0.8.25"

echo "If you don't have solc-select installed, please install it first from https://github.com/crytic/solc-select?tab=readme-ov-file#quickstart"
solc-select use $VERSION --always-install
for CONTRACT_HOME in $CONTRACTS_DIR/*; do
    if [ -d "$CONTRACT_HOME" ]; then
        PKG_NAME=$(basename $CONTRACT_HOME)
        for CONTRACT_PATH in $CONTRACT_HOME/*; do
            if [ "${CONTRACT_PATH: -4}" == ".sol" ]; then 
                CONTRACT_NAME=$(basename $CONTRACT_PATH .sol)
                echo "compiling $CONTRACT_NAME"
                solc $CONTRACT_PATH --metadata-hash none --bin --abi -o $BUILD_DIR --overwrite
                abigen --pkg $PKG_NAME \
                    --bin=$BUILD_DIR/$CONTRACT_NAME.bin \
                    --abi=$BUILD_DIR/$CONTRACT_NAME.abi \
                    --out=$CONTRACT_HOME/$CONTRACT_NAME.go
            fi
        done
    fi
done