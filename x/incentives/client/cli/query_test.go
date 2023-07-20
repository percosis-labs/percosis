package cli_test

import (
	gocontext "context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/percosis-labs/percosis/app/apptesting"
	"github.com/percosis-labs/percosis/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type QueryTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func (s *QueryTestSuite) SetupSuite() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)

	// create a pool
	s.PrepareBalancerPool()
	// set up lock with id = 1
	s.LockTokens(s.TestAccs[0], sdk.Coins{sdk.NewCoin("gamm/pool/1", sdk.NewInt(1000000))}, time.Hour*24)

	s.Commit()
}

func (s *QueryTestSuite) TestQueriesNeverAlterState() {
	testCases := []struct {
		name   string
		query  string
		input  interface{}
		output interface{}
	}{
		{
			"Query active gauges",
			"/percosis.incentives.Query/ActiveGauges",
			&types.ActiveGaugesRequest{},
			&types.ActiveGaugesResponse{},
		},
		{
			"Query active gauges per denom",
			"/percosis.incentives.Query/ActiveGaugesPerDenom",
			&types.ActiveGaugesPerDenomRequest{Denom: "stake"},
			&types.ActiveGaugesPerDenomResponse{},
		},
		{
			"Query gauge by id",
			"/percosis.incentives.Query/GaugeByID",
			&types.GaugeByIDRequest{Id: 1},
			&types.GaugeByIDResponse{},
		},
		{
			"Query all gauges",
			"/percosis.incentives.Query/Gauges",
			&types.GaugesRequest{},
			&types.GaugesResponse{},
		},
		{
			"Query lockable durations",
			"/percosis.incentives.Query/LockableDurations",
			&types.QueryLockableDurationsRequest{},
			&types.QueryLockableDurationsResponse{},
		},
		{
			"Query module to distibute coins",
			"/percosis.incentives.Query/ModuleToDistributeCoins",
			&types.ModuleToDistributeCoinsRequest{},
			&types.ModuleToDistributeCoinsResponse{},
		},
		{
			"Query reward estimate",
			"/percosis.incentives.Query/RewardsEst",
			&types.RewardsEstRequest{Owner: s.TestAccs[0].String()},
			&types.RewardsEstResponse{},
		},
		{
			"Query upcoming gauges",
			"/percosis.incentives.Query/UpcomingGauges",
			&types.UpcomingGaugesRequest{},
			&types.UpcomingGaugesResponse{},
		},
		{
			"Query upcoming gauges",
			"/percosis.incentives.Query/UpcomingGaugesPerDenom",
			&types.UpcomingGaugesPerDenomRequest{Denom: "stake"},
			&types.UpcomingGaugesPerDenomResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.SetupSuite()
			err := s.QueryHelper.Invoke(gocontext.Background(), tc.query, tc.input, tc.output)
			s.Require().NoError(err)
			s.StateNotAltered()
		})
	}
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
