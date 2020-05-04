package client

import (
	"github.com/kava-labs/cosmos-sdk/x/distribution/client/cli"
	"github.com/kava-labs/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/kava-labs/cosmos-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
