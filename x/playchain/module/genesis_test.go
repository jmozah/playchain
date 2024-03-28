package playchain_test

import (
	"testing"

	keepertest "playchain/testutil/keeper"
	"playchain/testutil/nullify"
	playchain "playchain/x/playchain/module"
	"playchain/x/playchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PlaychainKeeper(t)
	playchain.InitGenesis(ctx, k, genesisState)
	got := playchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
