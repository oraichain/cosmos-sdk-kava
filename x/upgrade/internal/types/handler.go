package types

import (
	sdk "github.com/kava-labs/cosmos-sdk/types"
)

// UpgradeHandler specifies the type of function that is called when an upgrade is applied
type UpgradeHandler func(ctx sdk.Context, plan Plan)
