package v15

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v4/types"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"

	"github.com/percosis-labs/percosis/v16/app/upgrades"
	poolmanagertypes "github.com/percosis-labs/percosis/v16/x/poolmanager/types"
	protorevtypes "github.com/percosis-labs/percosis/v16/x/protorev/types"
	valsetpreftypes "github.com/percosis-labs/percosis/v16/x/valset-pref/types"
)

// UpgradeName defines the on-chain upgrade name for the Percosis v15 upgrade.
const UpgradeName = "v15"

// pool ids to migrate
const (
	stPERCO_PERCOPoolId   = 833
	stJUNO_JUNOPoolId   = 817
	stSTARS_STARSPoolId = 810
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{poolmanagertypes.StoreKey, valsetpreftypes.StoreKey, protorevtypes.StoreKey, icqtypes.StoreKey, packetforwardtypes.StoreKey},
		Deleted: []string{},
	},
}
