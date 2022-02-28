package keeper

import (
	"context"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CommitCouncilResponse(goCtx context.Context, msg *types.MsgCommitCouncilResponse) (*types.MsgCommitCouncilResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	colDep := k.GetParams(ctx).CollateralDeposit

	creator, err := k.GetUserFromString(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	council := k.GetCouncil(ctx, msg.CouncilId)
	if !stringItemInList(msg.Creator, council.Voters) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid Voter")
	}

	if council.Status != types.CouncelingStatus_councilCreated {
		return nil, sdkerrors.Wrapf(types.ErrCouncilStatus, "Have '%s', want '%s'", council.Status.String(), types.CouncelingStatus_councilCreated.String())
	}

	var allreadyVoted []string
	for _, response := range council.HashResponses {
		allreadyVoted = append(allreadyVoted, response.User)
	}

	if stringItemInList(msg.Creator, allreadyVoted) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Allready voted")
	}

	resp := types.WrapHashResponse{msg.Creator, msg.Response}
	council.HashResponses = append(council.HashResponses, &resp)

	if len(council.HashResponses) == 5 {
		council.Status = types.CouncelingStatus_commited
	}

	err = k.BurnCoinsFromAddr(ctx, creator.Addr, sdk.Coins{colDep})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Voter does not have enough coins")
	}
	council.Treasury = council.Treasury.Add(colDep)

	k.SetCouncil(ctx, msg.CouncilId, council)

	return &types.MsgCommitCouncilResponseResponse{}, nil
}
