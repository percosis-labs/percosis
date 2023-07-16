#!/bin/bash

rm -rf $HOME/.percosisd/

cd $HOME

percosisd init --chain-id=testing testing --home=$HOME/.percosisd
percosisd keys add validator --keyring-backend=test --home=$HOME/.percosisd
percosisd add-genesis-account $(percosisd keys show validator -a --keyring-backend=test --home=$HOME/.percosisd) 100000000000stake,100000000000valtoken --home=$HOME/.percosisd
percosisd gentx validator 500000000stake --keyring-backend=test --home=$HOME/.percosisd --chain-id=testing
percosisd collect-gentxs --home=$HOME/.percosisd

update_genesis () {    
    cat $HOME/.percosisd/validator1/config/genesis.json | jq "$1" > $HOME/.percosisd/validator1/config/tmp_genesis.json && mv $HOME/.percosisd/validator1/config/tmp_genesis.json $HOME/.percosisd/validator1/config/genesis.json
}

# update staking genesis
update_genesis '.app_state["staking"]["params"]["unbonding_time"]="120s"'

# update governance genesis
update_genesis '.app_state["gov"]["voting_params"]["voting_period"]="10s"'

# update epochs genesis
update_genesis '.app_state["epochs"]["epochs"][0]["identifier"]="min"'
update_genesis '.app_state["epochs"]["epochs"][0]["duration"]="60s"'

# update poolincentives genesis
update_genesis '.app_state["poolincentives"]["lockable_durations"][0]="120s"'
update_genesis '.app_state["poolincentives"]["lockable_durations"][1]="180s"'
update_genesis '.app_state["poolincentives"]["lockable_durations"][2]="240s"'

# update incentives genesis
update_genesis '.app_state["incentives"]["params"]["distr_epoch_identifier"]="min"'
update_genesis '.app_state["incentives"]["lockable_durations"][0]="1s"'
update_genesis '.app_state["incentives"]["lockable_durations"][1]="120s"'
update_genesis '.app_state["incentives"]["lockable_durations"][2]="180s"'
update_genesis '.app_state["incentives"]["lockable_durations"][3]="240s"'

# update mint genesis
update_genesis '.app_state["mint"]["params"]["epoch_identifier"]="min"'

# update gamm genesis
update_genesis '.app_state["gamm"]["params"]["pool_creation_fee"][0]["denom"]="stake"'

# update superfluid genesis
update_genesis '.app_state["superfluid"]["params"]["minimum_risk_factor"]="0.500000000000000000"'

percosisd start --home=$HOME/.percosisd
