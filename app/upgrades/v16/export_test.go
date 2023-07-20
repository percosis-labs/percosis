package v16

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"github.com/percosis-labs/percosis/app/keepers"
	gammkeeper "github.com/percosis-labs/percosis/x/gamm/keeper"
	"github.com/percosis-labs/percosis/x/poolmanager"
	poolmanagertypes "github.com/percosis-labs/percosis/x/poolmanager/types"
	tokenfactorykeeper "github.com/percosis-labs/percosis/x/tokenfactory/keeper"
)

var (
	AuthorizedQuoteDenoms = authorizedQuoteDenoms
	AuthorizedUptimes     = authorizedUptimes
)

func CreateConcentratedPoolFromCFMM(ctx sdk.Context, cfmmPoolIdToLinkWith uint64, desiredDenom0 string, accountKeeper authkeeper.AccountKeeper, gammKeeper gammkeeper.Keeper, poolmanagerKeeper poolmanager.Keeper) (poolmanagertypes.PoolI, error) {
	return createConcentratedPoolFromCFMM(ctx, cfmmPoolIdToLinkWith, desiredDenom0, accountKeeper, gammKeeper, poolmanagerKeeper)
}

func CreateCanonicalConcentratedLiquidityPoolAndMigrationLink(ctx sdk.Context, cfmmPoolId uint64, desiredDenom0 string, keepers *keepers.AppKeepers) (poolmanagertypes.PoolI, error) {
	return createCanonicalConcentratedLiquidityPoolAndMigrationLink(ctx, cfmmPoolId, desiredDenom0, keepers)
}

func UpdateTokenFactoryParams(ctx sdk.Context, tokenFactoryKeeper *tokenfactorykeeper.Keeper) {
	updateTokenFactoryParams(ctx, tokenFactoryKeeper)
}
