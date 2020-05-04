package client

import (
	govclient "github.com/kava-labs/cosmos-sdk/x/gov/client"
	"github.com/kava-labs/cosmos-sdk/x/upgrade/client/cli"
	"github.com/kava-labs/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
