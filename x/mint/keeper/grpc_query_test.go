package keeper_test

import (
	"context"

	"github.com/percosis-labs/percosis/x/mint/types"
)

func (s *KeeperTestSuite) TestGRPCParams() {
	_, _, queryClient := s.App, s.Ctx, s.queryClient

	_, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
	s.Require().NoError(err)

	_, err = queryClient.EpochProvisions(context.Background(), &types.QueryEpochProvisionsRequest{})
	s.Require().NoError(err)
}
