package lamnv_test

import (
	"testing"

	keepertest "github.com/NguyenLam05200/Lamnv/testutil/keeper"
	"github.com/NguyenLam05200/Lamnv/testutil/nullify"
	"github.com/NguyenLam05200/Lamnv/x/lamnv"
	"github.com/NguyenLam05200/Lamnv/x/lamnv/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LamnvKeeper(t)
	lamnv.InitGenesis(ctx, *k, genesisState)
	got := lamnv.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
