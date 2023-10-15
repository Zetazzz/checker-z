package checkerz_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zetazzz/checker-z/testutil/keeper"
	"github.com/zetazzz/checker-z/testutil/nullify"
	"github.com/zetazzz/checker-z/x/checkerz"
	"github.com/zetazzz/checker-z/x/checkerz/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckerzKeeper(t)
	checkerz.InitGenesis(ctx, *k, genesisState)
	got := checkerz.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
