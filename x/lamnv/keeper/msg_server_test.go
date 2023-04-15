package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/NguyenLam05200/Lamnv/testutil/keeper"
	"github.com/NguyenLam05200/Lamnv/x/lamnv/keeper"
	"github.com/NguyenLam05200/Lamnv/x/lamnv/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LamnvKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
