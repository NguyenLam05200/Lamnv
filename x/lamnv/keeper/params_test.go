package keeper_test

import (
	"testing"

	testkeeper "github.com/NguyenLam05200/Lamnv/testutil/keeper"
	"github.com/NguyenLam05200/Lamnv/x/lamnv/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LamnvKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
