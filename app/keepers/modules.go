package keepers

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	transfer "github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v4/modules/core/02-client/client"
	"github.com/strangelove-ventures/packet-forward-middleware/v4/router"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v4"
	ica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"

	_ "github.com/percosis-labs/percosis/client/docs/statik"
	clclient "github.com/percosis-labs/percosis/x/concentrated-liquidity/client"
	concentratedliquidity "github.com/percosis-labs/percosis/x/concentrated-liquidity/clmodule"
	cwpoolclient "github.com/percosis-labs/percosis/x/cosmwasmpool/client"
	cosmwasmpoolmodule "github.com/percosis-labs/percosis/x/cosmwasmpool/module"
	downtimemodule "github.com/percosis-labs/percosis/x/downtime-detector/module"
	"github.com/percosis-labs/percosis/x/gamm"
	gammclient "github.com/percosis-labs/percosis/x/gamm/client"
	"github.com/percosis-labs/percosis/x/ibc-rate-limit/ibcratelimitmodule"
	"github.com/percosis-labs/percosis/x/incentives"
	"github.com/percosis-labs/percosis/x/lockup"
	"github.com/percosis-labs/percosis/x/mint"
	poolincentives "github.com/percosis-labs/percosis/x/pool-incentives"
	poolincentivesclient "github.com/percosis-labs/percosis/x/pool-incentives/client"
	poolmanager "github.com/percosis-labs/percosis/x/poolmanager/module"
	"github.com/percosis-labs/percosis/x/protorev"
	superfluid "github.com/percosis-labs/percosis/x/superfluid"
	superfluidclient "github.com/percosis-labs/percosis/x/superfluid/client"
	"github.com/percosis-labs/percosis/x/tokenfactory"
	"github.com/percosis-labs/percosis/x/twap/twapmodule"
	"github.com/percosis-labs/percosis/x/txfees"
	txfeesclient "github.com/percosis-labs/percosis/x/txfees/client"
	valsetprefmodule "github.com/percosis-labs/percosis/x/valset-pref/valpref-module"
	"github.com/osmosis-labs/osmosis/x/epochs"
	ibc_hooks "github.com/osmosis-labs/osmosis/x/ibc-hooks"
)

// AppModuleBasics returns ModuleBasics for the module BasicManager.
var AppModuleBasics = []module.AppModuleBasic{
	auth.AppModuleBasic{},
	genutil.AppModuleBasic{},
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	downtimemodule.AppModuleBasic{},
	distr.AppModuleBasic{},
	gov.NewAppModuleBasic(
		append(
			wasmclient.ProposalHandlers,
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
			upgradeclient.CancelProposalHandler,
			poolincentivesclient.UpdatePoolIncentivesHandler,
			poolincentivesclient.ReplacePoolIncentivesHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
			superfluidclient.SetSuperfluidAssetsProposalHandler,
			superfluidclient.RemoveSuperfluidAssetsProposalHandler,
			superfluidclient.UpdateUnpoolWhitelistProposalHandler,
			gammclient.ReplaceMigrationRecordsProposalHandler,
			gammclient.UpdateMigrationRecordsProposalHandler,
			clclient.CreateConcentratedLiquidityPoolProposalHandler,
			clclient.TickSpacingDecreaseProposalHandler,
			cwpoolclient.UploadCodeIdAndWhitelistProposalHandler,
			cwpoolclient.MigratePoolContractsProposalHandler,
			txfeesclient.SubmitUpdateFeeTokenProposalHandler,
		)...,
	),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	ibc.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	vesting.AppModuleBasic{},
	gamm.AppModuleBasic{},
	poolmanager.AppModuleBasic{},
	twapmodule.AppModuleBasic{},
	concentratedliquidity.AppModuleBasic{},
	protorev.AppModuleBasic{},
	txfees.AppModuleBasic{},
	incentives.AppModuleBasic{},
	lockup.AppModuleBasic{},
	poolincentives.AppModuleBasic{},
	epochs.AppModuleBasic{},
	superfluid.AppModuleBasic{},
	tokenfactory.AppModuleBasic{},
	valsetprefmodule.AppModuleBasic{},
	wasm.AppModuleBasic{},
	icq.AppModuleBasic{},
	ica.AppModuleBasic{},
	ibc_hooks.AppModuleBasic{},
	ibcratelimitmodule.AppModuleBasic{},
	router.AppModuleBasic{},
	cosmwasmpoolmodule.AppModuleBasic{},
}
