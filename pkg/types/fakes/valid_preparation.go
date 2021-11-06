package fakes

import (
	fake "github.com/brianvoe/gofakeit/v5"
	"github.com/segmentio/ksuid"

	"github.com/prixfixeco/api_server/pkg/types"
)

// BuildFakeValidPreparation builds a faked valid preparation.
func BuildFakeValidPreparation() *types.ValidPreparation {
	return &types.ValidPreparation{
		ID:          ksuid.New().String(),
		Name:        fake.LoremIpsumSentence(exampleQuantity),
		Description: fake.LoremIpsumSentence(exampleQuantity),
		IconPath:    fake.LoremIpsumSentence(exampleQuantity),
		CreatedOn:   uint64(uint32(fake.Date().Unix())),
	}
}

// BuildFakeValidPreparationList builds a faked ValidPreparationList.
func BuildFakeValidPreparationList() *types.ValidPreparationList {
	var examples []*types.ValidPreparation
	for i := 0; i < exampleQuantity; i++ {
		examples = append(examples, BuildFakeValidPreparation())
	}

	return &types.ValidPreparationList{
		Pagination: types.Pagination{
			Page:          1,
			Limit:         20,
			FilteredCount: exampleQuantity,
			TotalCount:    exampleQuantity * 2,
		},
		ValidPreparations: examples,
	}
}

// BuildFakeValidPreparationUpdateRequestInput builds a faked ValidPreparationUpdateRequestInput from a valid preparation.
func BuildFakeValidPreparationUpdateRequestInput() *types.ValidPreparationUpdateRequestInput {
	validPreparation := BuildFakeValidPreparation()
	return &types.ValidPreparationUpdateRequestInput{
		Name:        validPreparation.Name,
		Description: validPreparation.Description,
		IconPath:    validPreparation.IconPath,
	}
}

// BuildFakeValidPreparationUpdateRequestInputFromValidPreparation builds a faked ValidPreparationUpdateRequestInput from a valid preparation.
func BuildFakeValidPreparationUpdateRequestInputFromValidPreparation(validPreparation *types.ValidPreparation) *types.ValidPreparationUpdateRequestInput {
	return &types.ValidPreparationUpdateRequestInput{
		Name:        validPreparation.Name,
		Description: validPreparation.Description,
		IconPath:    validPreparation.IconPath,
	}
}

// BuildFakeValidPreparationCreationRequestInput builds a faked ValidPreparationCreationRequestInput.
func BuildFakeValidPreparationCreationRequestInput() *types.ValidPreparationCreationRequestInput {
	validPreparation := BuildFakeValidPreparation()
	return BuildFakeValidPreparationCreationRequestInputFromValidPreparation(validPreparation)
}

// BuildFakeValidPreparationCreationRequestInputFromValidPreparation builds a faked ValidPreparationCreationRequestInput from a valid preparation.
func BuildFakeValidPreparationCreationRequestInputFromValidPreparation(validPreparation *types.ValidPreparation) *types.ValidPreparationCreationRequestInput {
	return &types.ValidPreparationCreationRequestInput{
		ID:          validPreparation.ID,
		Name:        validPreparation.Name,
		Description: validPreparation.Description,
		IconPath:    validPreparation.IconPath,
	}
}

// BuildFakeValidPreparationDatabaseCreationInput builds a faked ValidPreparationDatabaseCreationInput.
func BuildFakeValidPreparationDatabaseCreationInput() *types.ValidPreparationDatabaseCreationInput {
	validPreparation := BuildFakeValidPreparation()
	return BuildFakeValidPreparationDatabaseCreationInputFromValidPreparation(validPreparation)
}

// BuildFakeValidPreparationDatabaseCreationInputFromValidPreparation builds a faked ValidPreparationDatabaseCreationInput from a valid preparation.
func BuildFakeValidPreparationDatabaseCreationInputFromValidPreparation(validPreparation *types.ValidPreparation) *types.ValidPreparationDatabaseCreationInput {
	return &types.ValidPreparationDatabaseCreationInput{
		ID:          validPreparation.ID,
		Name:        validPreparation.Name,
		Description: validPreparation.Description,
		IconPath:    validPreparation.IconPath,
	}
}
