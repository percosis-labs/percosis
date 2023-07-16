# Download a genesis.json for testing. The node that you this on will be your "validator"
# It should be on version v4.x

percosisd init --chain-id=testing testing --home=$HOME/.percosisd
percosisd keys add validator --keyring-backend=test --home=$HOME/.percosisd
percosisd add-genesis-account $(percosisd keys show validator -a --keyring-backend=test --home=$HOME/.percosisd) 1000000000ufury,1000000000valtoken --home=$HOME/.percosisd
sed -i -e "s/stake/ufury/g" $HOME/.percosisd/config/genesis.json
percosisd gentx validator 500000000ufury --commission-rate="0.0" --keyring-backend=test --home=$HOME/.percosisd --chain-id=testing
percosisd collect-gentxs --home=$HOME/.percosisd

cat $HOME/.percosisd/config/genesis.json | jq '.initial_height="711800"' > $HOME/.percosisd/config/tmp_genesis.json && mv $HOME/.percosisd/config/tmp_genesis.json $HOME/.percosisd/config/genesis.json
cat $HOME/.percosisd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]["denom"]="valtoken"' > $HOME/.percosisd/config/tmp_genesis.json && mv $HOME/.percosisd/config/tmp_genesis.json $HOME/.percosisd/config/genesis.json
cat $HOME/.percosisd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]["amount"]="100"' > $HOME/.percosisd/config/tmp_genesis.json && mv $HOME/.percosisd/config/tmp_genesis.json $HOME/.percosisd/config/genesis.json
cat $HOME/.percosisd/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > $HOME/.percosisd/config/tmp_genesis.json && mv $HOME/.percosisd/config/tmp_genesis.json $HOME/.percosisd/config/genesis.json
cat $HOME/.percosisd/config/genesis.json | jq '.app_state["staking"]["params"]["min_commission_rate"]="0.050000000000000000"' > $HOME/.percosisd/config/tmp_genesis.json && mv $HOME/.percosisd/config/tmp_genesis.json $HOME/.percosisd/config/genesis.json

# Now setup a second full node, and peer it with this v3.0.0-rc0 node.

# start the chain on both machines
percosisd start
# Create proposals

percosisd tx gov submit-proposal --title="existing passing prop" --description="passing prop"  --from=validator --deposit=1000valtoken --chain-id=testing --keyring-backend=test --broadcast-mode=block  --type="Text"
percosisd tx gov vote 1 yes --from=validator --keyring-backend=test --chain-id=testing --yes
percosisd tx gov submit-proposal --title="prop with enough perco deposit" --description="prop w/ enough deposit"  --from=validator --deposit=500000000ufury --chain-id=testing --keyring-backend=test --broadcast-mode=block  --type="Text"
# Check that we have proposal 1 passed, and proposal 2 in deposit period
percosisd q gov proposals
# CHeck that validator commission is under min_commission_rate
percosisd q staking validators
# Wait for upgrade block.
# Upgrade happened
# your full node should have crashed with consensus failure

# Now we test post-upgrade behavior is as intended

# Everything in deposit stayed in deposit
percosisd q gov proposals
# Check that commissions was bumped to min_commission_rate
percosisd q staking validators
# pushes 2 into voting period
percosisd tx gov deposit 2 1valtoken --from=validator --keyring-backend=test --chain-id=testing --yes