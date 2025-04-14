package logistics_test

import (
	"testing"

	keepertest "logistics/testutil/keeper"
	"logistics/testutil/nullify"
	logistics "logistics/x/logistics/module"
	"logistics/x/logistics/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LogisticsKeeper(t)
	logistics.InitGenesis(ctx, k, genesisState)
	got := logistics.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
