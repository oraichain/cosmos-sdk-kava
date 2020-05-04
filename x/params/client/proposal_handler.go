package client

import (
	govclient "github.com/kava-labs/cosmos-sdk/x/gov/client"
	"github.com/kava-labs/cosmos-sdk/x/params/client/cli"
	"github.com/kava-labs/cosmos-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
