package simulation

import (
	"encoding/json"
	"fmt"

	"github.com/percosis-labs/percosis/v16/x/superfluid/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// RandomizedGenState generates a random GenesisState for staking.
func RandomizedGenState(simState *module.SimulationState) {
	superfluidGenesis := &types.GenesisState{
		Params: types.Params{
			MinimumRiskFactor: sdk.NewDecWithPrec(5, 2), // 5%
		},
		SuperfluidAssets:          []types.SuperfluidAsset{},
		PercoEquivalentMultipliers: []types.PercoEquivalentMultiplierRecord{},
	}

	bz, err := json.MarshalIndent(&superfluidGenesis.Params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated superfluid parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(superfluidGenesis)
}
