package v14

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	ibchookstypes "github.com/percosis-labs/percosis/x/ibc-hooks/types"

	"github.com/percosis-labs/percosis/v16/app/upgrades"
	downtimetypes "github.com/percosis-labs/percosis/v16/x/downtime-detector/types"
)

// UpgradeName defines the on-chain upgrade name for the Percosis v14 upgrade.
const UpgradeName = "v14"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{downtimetypes.StoreKey, ibchookstypes.StoreKey},
		Deleted: []string{},
	},
}
