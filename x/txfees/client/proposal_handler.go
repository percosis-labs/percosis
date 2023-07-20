package client

import (
	"github.com/percosis-labs/percosis/x/txfees/client/cli"
	"github.com/percosis-labs/percosis/x/txfees/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	SubmitUpdateFeeTokenProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdateFeeTokenProposal, rest.ProposalUpdateFeeTokenProposalRESTHandler)
)
