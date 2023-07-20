package cli

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
	"github.com/percosis-labs/percosis/v16/x/incentives/types"
)


func TestGetCmdGauges(t *testing.T) {
	desc, _ := GetCmdGauges()
	tcs := map[string]percocli.QueryCliTestCase[*types.GaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.GaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
	}
	percocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdToDistributeCoins(t *testing.T) {
	desc, _ := GetCmdToDistributeCoins()
	tcs := map[string]percocli.QueryCliTestCase[*types.ModuleToDistributeCoinsRequest]{
		"basic test": {
			Cmd: "", ExpectedQuery: &types.ModuleToDistributeCoinsRequest{},
		},
	}
	percocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdGaugeByID(t *testing.T) {
	desc, _ := GetCmdGaugeByID()
	tcs := map[string]percocli.QueryCliTestCase[*types.GaugeByIDRequest]{
		"basic test": {
			Cmd: "1", ExpectedQuery: &types.GaugeByIDRequest{Id: 1},
		},
	}
	percocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdActiveGauges(t *testing.T) {
	desc, _ := GetCmdActiveGauges()
	tcs := map[string]percocli.QueryCliTestCase[*types.ActiveGaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.ActiveGaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
	}
	percocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdActiveGaugesPerDenom(t *testing.T) {
	desc, _ := GetCmdActiveGaugesPerDenom()
	tcs := map[string]percocli.QueryCliTestCase[*types.ActiveGaugesPerDenomRequest]{
		"basic test": {
			Cmd: "ufury --offset=2",
			ExpectedQuery: &types.ActiveGaugesPerDenomRequest{
				Denom:      "ufury",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
	}
	percocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdUpcomingGauges(t *testing.T) {
	desc, _ := GetCmdUpcomingGauges()
	tcs := map[string]percocli.QueryCliTestCase[*types.UpcomingGaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.UpcomingGaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
	}
	percocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdUpcomingGaugesPerDenom(t *testing.T) {
	desc, _ := GetCmdUpcomingGaugesPerDenom()
	tcs := map[string]percocli.QueryCliTestCase[*types.UpcomingGaugesPerDenomRequest]{
		"basic test": {
			Cmd: "ufury --offset=2",
			ExpectedQuery: &types.UpcomingGaugesPerDenomRequest{
				Denom:      "ufury",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
	}
	percocli.RunQueryTestCases(t, desc, tcs)
}
