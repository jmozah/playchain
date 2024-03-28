package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "playchain/testutil/keeper"
	"playchain/x/playchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PlaychainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
