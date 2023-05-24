package iavl

import (
	"testing"

	"cosmossdk.io/log"
	"github.com/cosmos/iavl"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/wrapper"
)

func TestImmutableTreePanics(t *testing.T) {
	t.Parallel()
	db := wrapper.NewCosmosDB(dbm.NewMemDB())
	immTree := iavl.NewImmutableTree(db, 100, false, log.NewNopLogger())
	it := &immutableTree{immTree}
	require.Panics(t, func() { it.Set([]byte{}, []byte{}) })
	require.Panics(t, func() { it.Remove([]byte{}) })
	require.Panics(t, func() { it.SaveVersion() })              // nolint:errcheck
	require.Panics(t, func() { it.DeleteVersionsTo(int64(1)) }) // nolint:errcheck

	imm, err := it.GetImmutable(1)
	require.Error(t, err)
	require.Nil(t, imm)

	imm, err = it.GetImmutable(0)
	require.NoError(t, err)
	require.NotNil(t, imm)
	require.Equal(t, immTree, imm)
}
