package fakes

import (
	fake "github.com/brianvoe/gofakeit/v5"
	"github.com/segmentio/ksuid"

	"github.com/prixfixeco/api_server/pkg/types"
)

// BuildFakeUserReputationUpdateInput builds a faked UserReputationUpdateInput.
func BuildFakeUserReputationUpdateInput() *types.UserReputationUpdateInput {
	return &types.UserReputationUpdateInput{
		TargetUserID:  ksuid.New().String(),
		NewReputation: types.GoodStandingHouseholdStatus,
		Reason:        fake.Sentence(10),
	}
}
