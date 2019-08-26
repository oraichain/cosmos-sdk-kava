package operations

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govsimops "github.com/cosmos/cosmos-sdk/x/gov/simulation/operations"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// SimulateMsgSetWithdrawAddress generates a MsgSetWithdrawAddress with random values.
func SimulateMsgSetWithdrawAddress(k distribution.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account) (opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {

		accountOrigin := simulation.RandomAcc(r, accs)
		accountDestination := simulation.RandomAcc(r, accs)
		msg := distribution.NewMsgSetWithdrawAddress(accountOrigin.Address, accountDestination.Address)

		res := app.Deliver(tx)
		if !res.IsOK() {
			return simulation.NoOpMsg(distribution.ModuleName), nil, errors.New(res.Log)
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgWithdrawDelegatorReward generates a MsgWithdrawDelegatorReward with random values.
func SimulateMsgWithdrawDelegatorReward(k distribution.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account) (opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {

		delegatorAccount := simulation.RandomAcc(r, accs)
		validatorAccount := simulation.RandomAcc(r, accs)
		msg := distribution.NewMsgWithdrawDelegatorReward(delegatorAccount.Address, sdk.ValAddress(validatorAccount.Address))

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(distribution.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		res := app.Deliver(tx)
		if !res.IsOK() {
			return simulation.NoOpMsg(distribution.ModuleName), nil, errors.New(res.Log)
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgWithdrawValidatorCommission generates a MsgWithdrawValidatorCommission with random values.
func SimulateMsgWithdrawValidatorCommission(k distribution.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account) (opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {

		account := simulation.RandomAcc(r, accs)
		msg := distribution.NewMsgWithdrawValidatorCommission(sdk.ValAddress(account.Address))

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(distribution.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		res := app.Deliver(tx)
		if !res.IsOK() {
			return simulation.NoOpMsg(distribution.ModuleName), nil, errors.New(res.Log)
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateCommunityPoolSpendProposalContent generates random community-pool-spend proposal content
func SimulateCommunityPoolSpendProposalContent(k distribution.Keeper) govsimops.ContentSimulator {
	return func(r *rand.Rand, _ *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account) gov.Content {
		var coins sdk.Coins

		recipientAcc := simulation.RandomAcc(r, accs)
		balance := k.GetFeePool(ctx).CommunityPool

		if len(balance) > 0 {
			denomIndex := r.Intn(len(balance))

			amount, err := simulation.RandPositiveInt(r, balance[denomIndex].Amount.TruncateInt())
			if err == nil {
				denom := balance[denomIndex].Denom
				coins = sdk.NewCoins(sdk.NewCoin(denom, amount.Mul(sdk.NewInt(2))))
			}
		}

		return distribution.NewCommunityPoolSpendProposal(
			simulation.RandStringOfLength(r, 10),
			simulation.RandStringOfLength(r, 100),
			recipientAcc.Address,
			coins,
		)
	}
}
