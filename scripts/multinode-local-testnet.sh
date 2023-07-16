#!/bin/bash
set -e

# always returns true so set -e doesn't exit if it is not running.
killall percosisd || true
rm -rf $HOME/.percosisd/

# make four percosis directories
mkdir $HOME/.percosisd
mkdir $HOME/.percosisd/validator1
mkdir $HOME/.percosisd/validator2
mkdir $HOME/.percosisd/validator3

# init all three validators
percosisd init --chain-id=testing validator1 --home=$HOME/.percosisd/validator1
percosisd init --chain-id=testing validator2 --home=$HOME/.percosisd/validator2
percosisd init --chain-id=testing validator3 --home=$HOME/.percosisd/validator3
# create keys for all three validators
percosisd keys add validator1 --keyring-backend=test --home=$HOME/.percosisd/validator1
percosisd keys add validator2 --keyring-backend=test --home=$HOME/.percosisd/validator2
percosisd keys add validator3 --keyring-backend=test --home=$HOME/.percosisd/validator3

update_genesis () {    
    cat $HOME/.percosisd/validator1/config/genesis.json | jq "$1" > $HOME/.percosisd/validator1/config/tmp_genesis.json && mv $HOME/.percosisd/validator1/config/tmp_genesis.json $HOME/.percosisd/validator1/config/genesis.json
}

# change staking denom to ufury
update_genesis '.app_state["staking"]["params"]["bond_denom"]="ufury"'

# create validator node with tokens to transfer to the three other nodes
percosisd add-genesis-account $(percosisd keys show validator1 -a --keyring-backend=test --home=$HOME/.percosisd/validator1) 100000000000ufury,100000000000stake --home=$HOME/.percosisd/validator1
percosisd gentx validator1 500000000ufury --keyring-backend=test --home=$HOME/.percosisd/validator1 --chain-id=testing
percosisd collect-gentxs --home=$HOME/.percosisd/validator1


# update staking genesis
update_genesis '.app_state["staking"]["params"]["unbonding_time"]="240s"'

# update crisis variable to ufury
update_genesis '.app_state["crisis"]["constant_fee"]["denom"]="ufury"'

# udpate gov genesis
update_genesis '.app_state["gov"]["voting_params"]["voting_period"]="60s"'
update_genesis '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="ufury"'

# update epochs genesis
update_genesis '.app_state["epochs"]["epochs"][1]["duration"]="60s"'

# update poolincentives genesis
update_genesis '.app_state["poolincentives"]["lockable_durations"][0]="120s"'
update_genesis '.app_state["poolincentives"]["lockable_durations"][1]="180s"'
update_genesis '.app_state["poolincentives"]["lockable_durations"][2]="240s"'
update_genesis '.app_state["poolincentives"]["params"]["minted_denom"]="ufury"'

# update incentives genesis
update_genesis '.app_state["incentives"]["lockable_durations"][0]="1s"'
update_genesis '.app_state["incentives"]["lockable_durations"][1]="120s"'
update_genesis '.app_state["incentives"]["lockable_durations"][2]="180s"'
update_genesis '.app_state["incentives"]["lockable_durations"][3]="240s"'
update_genesis '.app_state["incentives"]["params"]["distr_epoch_identifier"]="day"'

# update mint genesis
update_genesis '.app_state["mint"]["params"]["mint_denom"]="ufury"'
update_genesis '.app_state["mint"]["params"]["epoch_identifier"]="day"'

# update gamm genesis
update_genesis '.app_state["gamm"]["params"]["pool_creation_fee"][0]["denom"]="ufury"'


# port key (validator1 uses default ports)
# validator1 1317, 9090, 9091, 26658, 26657, 26656, 6060
# validator2 1316, 9088, 9089, 26655, 26654, 26653, 6061
# validator3 1315, 9086, 9087, 26652, 26651, 26650, 6062


# change app.toml values
VALIDATOR2_APP_TOML=$HOME/.percosisd/validator2/config/app.toml
VALIDATOR3_APP_TOML=$HOME/.percosisd/validator3/config/app.toml

# validator2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1316|g' $VALIDATOR2_APP_TOML
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $VALIDATOR2_APP_TOML
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $VALIDATOR2_APP_TOML

# validator3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1315|g' $VALIDATOR3_APP_TOML
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $VALIDATOR3_APP_TOML
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $VALIDATOR3_APP_TOML


# change config.toml values
VALIDATOR1_CONFIG=$HOME/.percosisd/validator1/config/config.toml
VALIDATOR2_CONFIG=$HOME/.percosisd/validator2/config/config.toml
VALIDATOR3_CONFIG=$HOME/.percosisd/validator3/config/config.toml

# validator1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $VALIDATOR1_CONFIG
# validator2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26655|g' $VALIDATOR2_CONFIG
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26654|g' $VALIDATOR2_CONFIG
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26653|g' $VALIDATOR2_CONFIG
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $VALIDATOR2_CONFIG
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $VALIDATOR2_CONFIG
# validator3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26652|g' $VALIDATOR3_CONFIG
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26651|g' $VALIDATOR3_CONFIG
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $VALIDATOR3_CONFIG
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $VALIDATOR3_CONFIG
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $VALIDATOR3_CONFIG


# copy validator1 genesis file to validator2-3
cp $HOME/.percosisd/validator1/config/genesis.json $HOME/.percosisd/validator2/config/genesis.json
cp $HOME/.percosisd/validator1/config/genesis.json $HOME/.percosisd/validator3/config/genesis.json


# copy tendermint node id of validator1 to persistent peers of validator2-3
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(percosisd tendermint show-node-id --home=$HOME/.percosisd/validator1)@localhost:26656\"|g" $HOME/.percosisd/validator2/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(percosisd tendermint show-node-id --home=$HOME/.percosisd/validator1)@localhost:26656\"|g" $HOME/.percosisd/validator3/config/config.toml


# start all three validators
tmux new -s validator1 -d percosisd start --home=$HOME/.percosisd/validator1
tmux new -s validator2 -d percosisd start --home=$HOME/.percosisd/validator2
tmux new -s validator3 -d percosisd start --home=$HOME/.percosisd/validator3


# send ufury from first validator to second validator
echo "Waiting 7 seconds to send funds to validators 2 and 3..."
sleep 7
percosisd tx bank send validator1 $(percosisd keys show validator2 -a --keyring-backend=test --home=$HOME/.percosisd/validator2) 500000000ufury --keyring-backend=test --home=$HOME/.percosisd/validator1 --chain-id=testing --broadcast-mode block --node http://localhost:26657 --yes
percosisd tx bank send validator1 $(percosisd keys show validator3 -a --keyring-backend=test --home=$HOME/.percosisd/validator3) 400000000ufury --keyring-backend=test --home=$HOME/.percosisd/validator1 --chain-id=testing --broadcast-mode block --node http://localhost:26657 --yes

# create second & third validator
percosisd tx staking create-validator --amount=500000000ufury --from=validator2 --pubkey=$(percosisd tendermint show-validator --home=$HOME/.percosisd/validator2) --moniker="validator2" --chain-id="testing" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="500000000" --keyring-backend=test --home=$HOME/.percosisd/validator2 --broadcast-mode block --node http://localhost:26657 --yes
percosisd tx staking create-validator --amount=400000000ufury --from=validator3 --pubkey=$(percosisd tendermint show-validator --home=$HOME/.percosisd/validator3) --moniker="validator3" --chain-id="testing" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="400000000" --keyring-backend=test --home=$HOME/.percosisd/validator3 --broadcast-mode block --node http://localhost:26657 --yes

echo "All 3 Validators are up and running!"