#!/bin/bash

set -ex

# initialize Hermes relayer configuration
mkdir -p /root/.hermes/
touch /root/.hermes/config.toml
echo $PERCO_A_E2E_VAL_MNEMONIC > /root/.hermes/PERCO_A_MNEMONIC.txt
echo $PERCO_B_E2E_VAL_MNEMONIC > /root/.hermes/PERCO_B_MNEMONIC.txt
# setup Hermes relayer configuration
tee /root/.hermes/config.toml <<EOF
[global]
log_level = 'info'
[mode]
[mode.clients]
enabled = true
refresh = true
misbehaviour = true
[mode.connections]
enabled = false
[mode.channels]
enabled = true
[mode.packets]
enabled = true
clear_interval = 1
clear_on_start = true
tx_confirmation = true
[rest]
enabled = true
host = '0.0.0.0'
port = 3031
[telemetry]
enabled = true
host = '127.0.0.1'
port = 3001
[[chains]]
id = '$PERCO_A_E2E_CHAIN_ID'
rpc_addr = 'http://$PERCO_A_E2E_VAL_HOST:26657'
grpc_addr = 'http://$PERCO_A_E2E_VAL_HOST:9090'
websocket_addr = 'ws://$PERCO_A_E2E_VAL_HOST:26657/websocket'
rpc_timeout = '30s'
account_prefix = 'perco'
key_name = 'val01-percosis-a'
store_prefix = 'ibc'
max_gas = 9000000
gas_multiplier = 1.5
default_gas = 400000
gas_price = { price = 0.0025, denom = 'e2e-default-feetoken' }
clock_drift = '1m' # to accomdate docker containers
trusting_period = '239seconds'
trust_threshold = { numerator = '1', denominator = '3' }
[[chains]]
id = '$PERCO_B_E2E_CHAIN_ID'
rpc_addr = 'http://$PERCO_B_E2E_VAL_HOST:26657'
grpc_addr = 'http://$PERCO_B_E2E_VAL_HOST:9090'
websocket_addr = 'ws://$PERCO_B_E2E_VAL_HOST:26657/websocket'
rpc_timeout = '30s'
account_prefix = 'perco'
key_name = 'val01-percosis-b'
store_prefix = 'ibc'
max_gas = 9000000
gas_multiplier = 1.5
default_gas = 400000
gas_price = { price = 0.0025, denom = 'e2e-default-feetoken' }
clock_drift = '1m' # to accomdate docker containers
trusting_period = '239seconds'
trust_threshold = { numerator = '1', denominator = '3' }
EOF

# import keys
hermes keys add --chain ${PERCO_B_E2E_CHAIN_ID} --key-name "val01-percosis-b" --mnemonic-file /root/.hermes/PERCO_B_MNEMONIC.txt
hermes keys add --chain ${PERCO_A_E2E_CHAIN_ID} --key-name "val01-percosis-a" --mnemonic-file /root/.hermes/PERCO_A_MNEMONIC.txt

# start Hermes relayer
hermes start
