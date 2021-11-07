package fakes

import (
	fake "github.com/brianvoe/gofakeit/v5"
	"github.com/segmentio/ksuid"

	"github.com/prixfixeco/api_server/internal/authorization"
	"github.com/prixfixeco/api_server/pkg/types"
)

// BuildFakeHouseholdUserMembership builds a faked HouseholdUserMembership.
func BuildFakeHouseholdUserMembership() *types.HouseholdUserMembership {
	return &types.HouseholdUserMembership{
		ID:                 ksuid.New().String(),
		BelongsToUser:      fake.UUID(),
		BelongsToHousehold: fake.UUID(),
		HouseholdRoles:     []string{authorization.HouseholdMemberRole.String()},
		CreatedOn:          0,
		ArchivedOn:         nil,
	}
}

// BuildFakeHouseholdUserMembershipList builds a faked HouseholdUserMembershipList.
func BuildFakeHouseholdUserMembershipList() *types.HouseholdUserMembershipList {
	var examples []*types.HouseholdUserMembership
	for i := 0; i < exampleQuantity; i++ {
		examples = append(examples, BuildFakeHouseholdUserMembership())
	}

	return &types.HouseholdUserMembershipList{
		Pagination: types.Pagination{
			Page:          1,
			Limit:         20,
			FilteredCount: exampleQuantity / 2,
			TotalCount:    exampleQuantity,
		},
		HouseholdUserMemberships: examples,
	}
}

// BuildFakeHouseholdUserMembershipUpdateInputFromHouseholdUserMembership builds a faked HouseholdUserMembershipUpdateInput from a household user membership.
func BuildFakeHouseholdUserMembershipUpdateInputFromHouseholdUserMembership(householdUserMembership *types.HouseholdUserMembership) *types.HouseholdUserMembershipUpdateInput {
	return &types.HouseholdUserMembershipUpdateInput{
		BelongsToUser:      householdUserMembership.BelongsToUser,
		BelongsToHousehold: householdUserMembership.BelongsToHousehold,
	}
}

// BuildFakeHouseholdUserMembershipCreationInput builds a faked HouseholdUserMembershipCreationInput.
func BuildFakeHouseholdUserMembershipCreationInput() *types.HouseholdUserMembershipCreationInput {
	return BuildFakeHouseholdUserMembershipCreationInputFromHouseholdUserMembership(BuildFakeHouseholdUserMembership())
}

// BuildFakeHouseholdUserMembershipCreationInputFromHouseholdUserMembership builds a faked HouseholdUserMembershipCreationInput from a household user membership.
func BuildFakeHouseholdUserMembershipCreationInputFromHouseholdUserMembership(householdUserMembership *types.HouseholdUserMembership) *types.HouseholdUserMembershipCreationInput {
	return &types.HouseholdUserMembershipCreationInput{
		ID:                 householdUserMembership.ID,
		BelongsToUser:      householdUserMembership.BelongsToUser,
		BelongsToHousehold: householdUserMembership.BelongsToHousehold,
	}
}
