#!/usr/bin/env bash

CHAIN_ID=testnet-1

if sed --version >/dev/null 2>&1; then
  sed_inplace() { sed -i "$@"; }
else
  sed_inplace() { sed -i '' "$@"; }
fi

echo "Initializing testnet..."
rm -rf ~/.sequencer ~/.attestor0 ~/.attestor1 ~/.full0 ~/.full1

# initialize nodes
echo "Initializing nodes..."
minitiad init operator --chain-id $CHAIN_ID --home ~/.sequencer
minitiad init attestor0 --chain-id $CHAIN_ID --home ~/.attestor0
minitiad init attestor1 --chain-id $CHAIN_ID --home ~/.attestor1
minitiad init full0 --chain-id $CHAIN_ID --home ~/.full0
minitiad init full1 --chain-id $CHAIN_ID --home ~/.full1

# create genesis and copy to other nodes
echo "Creating genesis..."
minitiad genesis add-genesis-account operator 1000000000000000000umin --home ~/.sequencer
minitiad genesis add-genesis-validator operator --home ~/.sequencer
cp ~/.sequencer/config/genesis.json ~/.attestor0/config/genesis.json
cp ~/.sequencer/config/genesis.json ~/.attestor1/config/genesis.json
cp ~/.sequencer/config/genesis.json ~/.full0/config/genesis.json
cp ~/.sequencer/config/genesis.json ~/.full1/config/genesis.json

# update config to register each other as persistent peers
echo "Updating persistent peers..."
SEQUENCER_P2P=$(minitiad cometbft show-node-id --home ~/.sequencer)@127.0.0.1:26656
ATTESTOR0_P2P=$(minitiad cometbft show-node-id --home ~/.attestor0)@127.0.0.1:26666
ATTESTOR1_P2P=$(minitiad cometbft show-node-id --home ~/.attestor1)@127.0.0.1:26676
FULL0_P2P=$(minitiad cometbft show-node-id --home ~/.full0)@127.0.0.1:26686
FULL1_P2P=$(minitiad cometbft show-node-id --home ~/.full1)@127.0.0.1:26696
PERSISTENT_PEERS="$SEQUENCER_P2P,$ATTESTOR0_P2P,$ATTESTOR1_P2P,$FULL0_P2P,$FULL1_P2P"
sed_inplace "s/persistent_peers = \"\"/persistent_peers = \"$PERSISTENT_PEERS\"/" ~/.sequencer/config/config.toml
sed_inplace "s/persistent_peers = \"\"/persistent_peers = \"$PERSISTENT_PEERS\"/" ~/.attestor0/config/config.toml
sed_inplace "s/persistent_peers = \"\"/persistent_peers = \"$PERSISTENT_PEERS\"/" ~/.attestor1/config/config.toml
sed_inplace "s/persistent_peers = \"\"/persistent_peers = \"$PERSISTENT_PEERS\"/" ~/.full0/config/config.toml
sed_inplace "s/persistent_peers = \"\"/persistent_peers = \"$PERSISTENT_PEERS\"/" ~/.full1/config/config.toml

# update config to change ports
echo "Updating ports..."
sed_inplace 's/26656/26666/g' ~/.attestor0/config/config.toml
sed_inplace 's/26657/26667/g' ~/.attestor0/config/config.toml
sed_inplace 's/1317/1327/g' ~/.attestor0/config/app.toml
sed_inplace 's/9090/9091/g' ~/.attestor0/config/app.toml
sed_inplace 's/8545/8555/g' ~/.attestor0/config/app.toml
sed_inplace 's/8546/8556/g' ~/.attestor0/config/app.toml
sed_inplace 's/26656/26676/g' ~/.attestor1/config/config.toml
sed_inplace 's/26657/26677/g' ~/.attestor1/config/config.toml
sed_inplace 's/1317/1337/g' ~/.attestor1/config/app.toml
sed_inplace 's/9090/9092/g' ~/.attestor1/config/app.toml
sed_inplace 's/8545/8565/g' ~/.attestor1/config/app.toml
sed_inplace 's/8546/8566/g' ~/.attestor1/config/app.toml
sed_inplace 's/26656/26686/g' ~/.full0/config/config.toml
sed_inplace 's/26657/26687/g' ~/.full0/config/config.toml
sed_inplace 's/1317/1347/g' ~/.full0/config/app.toml
sed_inplace 's/9090/9093/g' ~/.full0/config/app.toml
sed_inplace 's/8545/8575/g' ~/.full0/config/app.toml
sed_inplace 's/8546/8576/g' ~/.full0/config/app.toml
sed_inplace 's/26656/26696/g' ~/.full1/config/config.toml
sed_inplace 's/26657/26697/g' ~/.full1/config/config.toml
sed_inplace 's/1317/1357/g' ~/.full1/config/app.toml
sed_inplace 's/9090/9094/g' ~/.full1/config/app.toml
sed_inplace 's/8545/8585/g' ~/.full1/config/app.toml
sed_inplace 's/8546/8586/g' ~/.full1/config/app.toml

# allow duplicate IPs for local testnet
sed_inplace 's/allow_duplicate_ip = false/allow_duplicate_ip = true/g' ~/.sequencer/config/config.toml
sed_inplace 's/addr_book_strict = true/addr_book_strict = false/g' ~/.sequencer/config/config.toml
sed_inplace 's/allow_duplicate_ip = false/allow_duplicate_ip = true/g' ~/.attestor0/config/config.toml
sed_inplace 's/addr_book_strict = true/addr_book_strict = false/g' ~/.attestor0/config/config.toml
sed_inplace 's/allow_duplicate_ip = false/allow_duplicate_ip = true/g' ~/.attestor1/config/config.toml
sed_inplace 's/addr_book_strict = true/addr_book_strict = false/g' ~/.attestor1/config/config.toml
sed_inplace 's/allow_duplicate_ip = false/allow_duplicate_ip = true/g' ~/.full0/config/config.toml
sed_inplace 's/addr_book_strict = true/addr_book_strict = false/g' ~/.full0/config/config.toml
sed_inplace 's/allow_duplicate_ip = false/allow_duplicate_ip = true/g' ~/.full1/config/config.toml
sed_inplace 's/addr_book_strict = true/addr_book_strict = false/g' ~/.full1/config/config.toml

echo "Testnet initialized"
