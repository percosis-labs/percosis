# Concentrated Liquidity Realistic Position Generation Script

This script adds realistic positions to localpercosis by querying Ethereum's Uniswap WETH / USDC pool and converting it
into Percosis genesis.

This script has 3 commands set via flags:

- `getData` retrieves the data from the Uniswap subgraph and writes it to the disk
- `convertPositions` converts the data from the Uniswap subgraph into Percosis genesis by reading the file from the previous step
and spinning up a mock Percosis app to create positions on and, then, exporting genesis
- `mergeSubgraphAndLocalPercosisGenesis` merges the genesis created from the subgraph data with the localpercosis genesis.
It is meant to only be called inside a localpercosis container

There are makefile commands defined that should help to interact with the script on the user's host machine.

```makefile
# This script retrieves Uniswap v3 Ethereum position data
# from subgraph. It uses WETH / USDC pool. This is helpful
# for setting up somewhat realistic positions for testing
# in localpercosis. It writes the file under
# tests/cl-genesis-positions/subgraph_positions.json
cl-refresh-subgraph-positions:
	go run ./tests/cl-genesis-positions --operation 0

# This script converts the positions data created by the
# cl-refresh-subgraph-positions makefile step into an Percosis
# genesis. It writes the file under tests/cl-genesis-positions/genesis.json
cl-refresh-subgraph-genesis:
	go run ./tests/cl-genesis-positions --operation 1

# This script converts the positions data created by the
# cl-refresh-subgraph-positions makefile step into a Big Bang
# configuration file for spinning up testnets.
# It writes the file under tests/cl-genesis-positions/bigbang_positions.json
cl-create-bigbang-config:
	go run ./tests/cl-genesis-positions --operation 1 --big-bang
```

These commands are also executed in setup.sh when starting up a localpercosis instance with make localnet-start-with-state
