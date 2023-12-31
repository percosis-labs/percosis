#!/bin/sh

CHAIN_ID=localpercosis
PERCOSIS_HOME=$HOME/.percosisd
CONFIG_FOLDER=$PERCOSIS_HOME/config
MONIKER=val
STATE='false'

MNEMONIC="bottom loan skill merry east cradle onion journey palm apology verb edit desert impose absurd oil bubble sweet glove shallow size build burst effort"
POOLSMNEMONIC="traffic cool olive pottery elegant innocent aisle dial genuine install shy uncle ride federal soon shift flight program cave famous provide cute pole struggle"

while getopts s flag
do
    case "${flag}" in
        s) STATE='true';;
    esac
done

install_prerequisites () {
    apk add dasel
}

edit_genesis () {

    GENESIS=$CONFIG_FOLDER/genesis.json

    # Update staking module
    dasel put string -f $GENESIS '.app_state.staking.params.bond_denom' 'ufury'
    dasel put string -f $GENESIS '.app_state.staking.params.unbonding_time' '240s'

    # Update bank module
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[0].description' 'Registered denom uion for localpercosis testing'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[0].denom_units.[0].denom' 'uion'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[0].denom_units.[0].exponent' 0
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[0].base' 'uion'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[0].display' 'uion'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[0].name' 'uion'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[0].symbol' 'uion'

    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[1].description' 'Registered denom ufury for localpercosis testing'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[1].denom_units.[0].denom' 'ufury'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[1].denom_units.[0].exponent' 0
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[1].base' 'ufury'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[1].display' 'ufury'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[1].name' 'ufury'
    dasel put string -f $GENESIS '.app_state.bank.denom_metadata.[1].symbol' 'ufury'

    # Update crisis module
    dasel put string -f $GENESIS '.app_state.crisis.constant_fee.denom' 'ufury'

    # Udpate gov module
    dasel put string -f $GENESIS '.app_state.gov.voting_params.voting_period' '60s'
    dasel put string -f $GENESIS '.app_state.gov.deposit_params.min_deposit.[0].denom' 'ufury'

    # Update epochs module
    dasel put string -f $GENESIS '.app_state.epochs.epochs.[1].duration' "60s"

    # Update poolincentives module
    dasel put string -f $GENESIS '.app_state.poolincentives.lockable_durations.[0]' "120s"
    dasel put string -f $GENESIS '.app_state.poolincentives.lockable_durations.[1]' "180s"
    dasel put string -f $GENESIS '.app_state.poolincentives.lockable_durations.[2]' "240s"
    dasel put string -f $GENESIS '.app_state.poolincentives.params.minted_denom' "ufury"

    # Update incentives module
    dasel put string -f $GENESIS '.app_state.incentives.lockable_durations.[0]' "1s"
    dasel put string -f $GENESIS '.app_state.incentives.lockable_durations.[1]' "120s"
    dasel put string -f $GENESIS '.app_state.incentives.lockable_durations.[2]' "180s"
    dasel put string -f $GENESIS '.app_state.incentives.lockable_durations.[3]' "240s"
    dasel put string -f $GENESIS '.app_state.incentives.params.distr_epoch_identifier' "hour"

    # Update mint module
    dasel put string -f $GENESIS '.app_state.mint.params.mint_denom' "ufury"
    dasel put string -f $GENESIS '.app_state.mint.params.epoch_identifier' "hour"

    # Update poolmanager module
    dasel put string -f $GENESIS '.app_state.poolmanager.params.pool_creation_fee.[0].denom' "ufury"

    # Update txfee basedenom
    dasel put string -f $GENESIS '.app_state.txfees.basedenom' "ufury"

    # Update wasm permission (Nobody or Everybody)
    dasel put string -f $GENESIS '.app_state.wasm.params.code_upload_access.permission' "Everybody"

    # Update concentrated-liquidity (enable pool creation)
    dasel put bool -f $GENESIS '.app_state.concentratedliquidity.params.is_permissionless_pool_creation_enabled' true
}

add_genesis_accounts () {

    percosisd add-genesis-account perco12smx2wdlyttvyzvzg54y2vnqwq2qjateuf7thj 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    # note such large amounts are set for e2e tests on FE 
    percosisd add-genesis-account perco1cyyzpxplxdzkeea7kwsydadg87357qnahakaks 9999999999999999999999999999999999999999999999999ufury,9999999999999999999999999999999999999999999999999uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco18s5lynnmx37hq4wlrw9gdn68sg2uxp5rgk26vv 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco1qwexv7c6sm95lwhzn9027vyu2ccneaqad4w8ka 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco14hcxlnwlqtq75ttaxf674vk6mafspg8xwgnn53 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco12rr534cer5c0vj53eq4y32lcwguyy7nndt0u2t 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco1nt33cjd5auzh36syym6azgc8tve0jlvklnq7jq 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco10qfrpash5g2vk3hppvu45x0g860czur8ff5yx0 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco1f4tvsdukfwh6s9swrc24gkuz23tp8pd3e9r5fa 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco1myv43sqgnj5sm4zl98ftl45af9cfzk7nhjxjqh 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco14gs9zqh8m49yy9kscjqu9h72exyf295afg6kgk 100000000000ufury,100000000000uion,100000000000stake,100000000000uusdc,100000000000uweth --home $PERCOSIS_HOME
    percosisd add-genesis-account perco1jllfytsz4dryxhz5tl7u73v29exsf80vz52ucc 1000000000000ufury,1000000000000uion,1000000000000stake,1000000000000uusdc,1000000000000uweth --home $PERCOSIS_HOME

    echo $MNEMONIC | percosisd keys add $MONIKER --recover --keyring-backend=test --home $PERCOSIS_HOME
    echo $POOLSMNEMONIC | percosisd keys add pools --recover --keyring-backend=test --home $PERCOSIS_HOME
    percosisd gentx $MONIKER 500000000ufury --keyring-backend=test --chain-id=$CHAIN_ID --home $PERCOSIS_HOME

    percosisd collect-gentxs --home $PERCOSIS_HOME
}

edit_config () {

    # Remove seeds
    dasel put string -f $CONFIG_FOLDER/config.toml '.p2p.seeds' ''

    # Expose the rpc
    dasel put string -f $CONFIG_FOLDER/config.toml '.rpc.laddr' "tcp://0.0.0.0:26657"
    
    # Expose pprof for debugging
    # To make the change enabled locally, make sure to add 'EXPOSE 6060' to the root Dockerfile
    # and rebuild the image.
    dasel put string -f $CONFIG_FOLDER/config.toml '.rpc.pprof_laddr' "0.0.0.0:6060"
}

enable_cors () {

    # Enable cors on RPC
    dasel put string -f $CONFIG_FOLDER/config.toml -v "*" '.rpc.cors_allowed_origins.[]'
    dasel put string -f $CONFIG_FOLDER/config.toml -v "Accept-Encoding" '.rpc.cors_allowed_headers.[]'
    dasel put string -f $CONFIG_FOLDER/config.toml -v "DELETE" '.rpc.cors_allowed_methods.[]'
    dasel put string -f $CONFIG_FOLDER/config.toml -v "OPTIONS" '.rpc.cors_allowed_methods.[]'
    dasel put string -f $CONFIG_FOLDER/config.toml -v "PATCH" '.rpc.cors_allowed_methods.[]'
    dasel put string -f $CONFIG_FOLDER/config.toml -v "PUT" '.rpc.cors_allowed_methods.[]'

    # Enable unsafe cors and swagger on the api
    dasel put bool -f $CONFIG_FOLDER/app.toml -v "true" '.api.swagger'
    dasel put bool -f $CONFIG_FOLDER/app.toml -v "true" '.api.enabled-unsafe-cors'

    # Enable cors on gRPC Web
    dasel put bool -f $CONFIG_FOLDER/app.toml -v "true" '.grpc-web.enable-unsafe-cors'
}

run_with_retries() {
  cmd=$1
  success_msg=$2

  substring='code: 0'
  COUNTER=0

  while [ $COUNTER -lt 15 ]; do
    string=$(eval $cmd 2>&1)
    echo $string

    if [ "$string" != "${string%"$substring"*}" ]; then
      echo "$success_msg"
      break
    else
      COUNTER=$((COUNTER+1))
      sleep 0.5
    fi
  done
}

# Define the functions using the new function
create_two_asset_pool() {
  run_with_retries "percosisd tx gamm create-pool --pool-file=$1 --from pools --chain-id=$CHAIN_ID --home $PERCOSIS_HOME --keyring-backend=test -b block --fees 5000ufury --yes" "create two asset pool: successful"
}

create_three_asset_pool() {
  run_with_retries "percosisd tx gamm create-pool --pool-file=nativeDenomThreeAssetPool.json --from pools --chain-id=$CHAIN_ID --home $PERCOSIS_HOME --keyring-backend=test -b block --fees 5000ufury --gas 900000 --yes" "create three asset pool: successful"
}

if [[ ! -d $CONFIG_FOLDER ]]
then
    echo $MNEMONIC | percosisd init -o --chain-id=$CHAIN_ID --home $PERCOSIS_HOME --recover $MONIKER
    install_prerequisites
    edit_genesis
    add_genesis_accounts
    edit_config
    enable_cors
fi

if [[ $STATE == 'true' ]]
then
    # Enter the script folder
    cd cl-genesis-positions

    # Build script with dependencies
    apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers
    go mod download
    WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | cut -d ' ' -f 2) && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$(uname -m).a \
        -O /lib/libwasmvm_muslc.a && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm_muslc.a | grep $(cat /tmp/checksums.txt | grep $(uname -m) | cut -d ' ' -f 1)
    go mod tidy
    go build \
        -mod=readonly \
        -tags "netgo,ledger,muslc" \
        -ldflags \
            "-X github.com/cosmos/cosmos-sdk/version.Name="percosis" \
            -X github.com/cosmos/cosmos-sdk/version.AppName="percosisd" \
            -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
            -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
            -X github.com/cosmos/cosmos-sdk/version.BuildTags=netgo,ledger,muslc \
            -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
        -trimpath \
        -o script \
        .

    # Get relevant data is not present on the mounted volume.
    if [[ ! -f "genesis.json" ]]; then
        if [[ ! -f "subgraph_positions.json" ]]; then
            echo "Getting concentrated liquidity data from Uniswap subgraph"
            ./script --operation 0 --localpercosis
        fi

        echo "Generating Percosis genesis for the concentrated liquidity module from Uniswap data"
        ./script --operation 1 --localpercosis
    fi

    # Run genesis merge script
    ./script --operation 2 --localpercosis
    cd ..
fi

percosisd start --home $PERCOSIS_HOME &

if [[ $STATE == 'true' ]]
then
    create_two_asset_pool "nativeDenomPoolA.json"
    create_two_asset_pool "nativeDenomPoolB.json"
    create_three_asset_pool
fi
wait
