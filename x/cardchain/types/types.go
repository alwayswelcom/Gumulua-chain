package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Here can constructors for types be placed, sice the files containing those
// types shall not be edited

// NewUser Constructor for User
func NewUser() User {
	return User{
		Alias:            "newPlayer",
		Cards:            []uint64{},
		OwnedCardSchemes: []uint64{},
		OwnedPrototypes:  []uint64{},
		VoteRights:       []*VoteRight{},
		CouncilStatus:    CouncilStatus_unavailable,
		ReportMatches:    false,
	}
}

// NewVoteRight Constructor for VoteRights
func NewVoteRight(cardId uint64, expireBlock int64) VoteRight {
	return VoteRight{
		CardId:      cardId,
		ExpireBlock: expireBlock,
	}
}

// NewCard Constructor for Card
func NewCard(owner sdk.AccAddress) Card {
	return Card{
		Owner:              owner.String(),
		Content:            []byte{},
		Image:              []byte{},
		Notes:              "",
		FullArt:            true,
		Status:             Status_scheme,
		VotePool:           sdk.NewInt64Coin("ucredits", 0),
		FairEnoughVotes:    0,
		OverpoweredVotes:   0,
		UnderpoweredVotes:  0,
		InappropriateVotes: 0,
		Nerflevel:          0,
	}
}

// NewVotingResults Constructor for VoteResults
func NewVotingResults() VotingResults {
	return VotingResults{
		TotalVotes:              0,
		TotalFairEnoughVotes:    0,
		TotalOverpoweredVotes:   0,
		TotalUnderpoweredVotes:  0,
		TotalInappropriateVotes: 0,
		CardResults:             []*VotingResult{},
		Notes:                   "none",
	}
}
