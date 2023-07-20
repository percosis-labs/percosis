#!/bin/bash
# microtick and bitcanna contributed significantly here.
# rocksdb doesn't work yet
# sage prediction: it will state sync fine with v12.2.1 and it won't work with v12.3.0 and the issue will be a blockheader.apphash error, which is a p2p issue, NOT an issue with the commit state of the store, as those would halt the local percosis instance.




# PRINT EVERY COMMAND
set -ux

# uncomment the three lines below to build percosis

go mod edit -replace github.com/tendermint/tm-db=github.com/baabeetaa/tm-db@pebble
go mod tidy
go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb -X github.com/tendermint/tm-db.ForceSync=1' -tags pebbledb ./...


# MAKE HOME FOLDER AND GET GENESIS
percosisd init test
wget -O ~/.percosisd/config/genesis.json https://github.com/percosis-labs/percosis/raw/main/networks/percosis-1/genesis.json


INTERVAL=1500

# GET TRUST HASH AND TRUST HEIGHT

LATEST_HEIGHT=$(curl -s https://rpc.percosis.zone/block | jq -r .result.block.header.height);
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL))
TRUST_HASH=$(curl -s "https://rpc.percosis.zone/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)


# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"


# export state sync vars
export PERCOSISD_P2P_MAX_NUM_OUTBOUND_PEERS=200
export PERCOSISD_STATESYNC_ENABLE=true
export PERCOSISD_STATESYNC_RPC_SERVERS="https://rpc.percosis.zone:443,https://rpc.percosis.zone:443"
export PERCOSISD_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export PERCOSISD_STATESYNC_TRUST_HASH=$TRUST_HASH



# THERE, NOW IT'S SYNCED AND YOU CAN PLAY
percosisd start
