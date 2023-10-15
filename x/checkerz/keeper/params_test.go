package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/zetazzz/checker-z/testutil/keeper"
	"github.com/zetazzz/checker-z/x/checkerz/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CheckerzKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
