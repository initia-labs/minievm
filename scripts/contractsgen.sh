BUILD_DIR=build
CONTRACTS_DIR=x/evm/contracts
for CONTRACT_HOME in $CONTRACTS_DIR/*; do
    if [ -d "$CONTRACT_HOME" ]; then
        PKG_NAME=$(basename $CONTRACT_HOME)
        CONTRACT_PATH=$(find $CONTRACT_HOME -type f -name "*.sol")
        CONTRACT_NAME=$(basename $CONTRACT_PATH .sol)
        echo $CONTRACT_HOME $PKG_NAME $CONTRACT_PATH $CONTRACT_NAME

        solc $CONTRACT_PATH --bin --abi -o $BUILD_DIR --overwrite
        abigen --pkg $PKG_NAME \
            --bin=$BUILD_DIR/$CONTRACT_NAME.bin \
            --abi=$BUILD_DIR/$CONTRACT_NAME.abi \
            --out=$CONTRACT_HOME/$CONTRACT_NAME.go
    fi

    #solc $(ls $${file}/*.sol) --bin --abi -o build
done