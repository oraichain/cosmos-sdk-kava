package keeper_test

import (
	abci "github.com/kava-labs/tendermint/abci/types"

	"github.com/kava-labs/cosmos-sdk/simapp"
	sdk "github.com/kava-labs/cosmos-sdk/types"
	"github.com/kava-labs/cosmos-sdk/x/mint/internal/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}
