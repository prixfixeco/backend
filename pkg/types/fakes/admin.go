package fakes

import (
	"github.com/prixfixeco/backend/pkg/types"

	fake "github.com/brianvoe/gofakeit/v5"
)

// BuildFakeUserAccountStatusUpdateInput builds a faked UserAccountStatusUpdateInput.
func BuildFakeUserAccountStatusUpdateInput() *types.UserAccountStatusUpdateInput {
	return &types.UserAccountStatusUpdateInput{
		TargetUserID: BuildFakeID(),
		NewStatus:    string(types.GoodStandingUserAccountStatus),
		Reason:       fake.Sentence(10),
	}
}
