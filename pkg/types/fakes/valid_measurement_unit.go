package fakes

import (
	fake "github.com/brianvoe/gofakeit/v5"

	"github.com/prixfixeco/backend/pkg/types"
	"github.com/prixfixeco/backend/pkg/types/converters"
)

// BuildFakeValidMeasurementUnit builds a faked valid ingredient.
func BuildFakeValidMeasurementUnit() *types.ValidMeasurementUnit {
	return &types.ValidMeasurementUnit{
		ID:          BuildFakeID(),
		Name:        buildUniqueString(),
		Description: buildUniqueString(),
		Volumetric:  fake.Bool(),
		IconPath:    buildUniqueString(),
		Universal:   fake.Bool(),
		Metric:      fake.Bool(),
		Imperial:    fake.Bool(),
		PluralName:  buildUniqueString(),
		CreatedAt:   BuildFakeTime(),
	}
}

// BuildFakeValidMeasurementUnitList builds a faked ValidMeasurementUnitList.
func BuildFakeValidMeasurementUnitList() *types.ValidMeasurementUnitList {
	var examples []*types.ValidMeasurementUnit
	for i := 0; i < exampleQuantity; i++ {
		examples = append(examples, BuildFakeValidMeasurementUnit())
	}

	return &types.ValidMeasurementUnitList{
		Pagination: types.Pagination{
			Page:          1,
			Limit:         20,
			FilteredCount: exampleQuantity / 2,
			TotalCount:    exampleQuantity,
		},
		ValidMeasurementUnits: examples,
	}
}

// BuildFakeValidMeasurementUnitUpdateRequestInput builds a faked ValidMeasurementUnitUpdateRequestInput from a valid ingredient.
func BuildFakeValidMeasurementUnitUpdateRequestInput() *types.ValidMeasurementUnitUpdateRequestInput {
	validMeasurementUnit := BuildFakeValidMeasurementUnit()
	return &types.ValidMeasurementUnitUpdateRequestInput{
		Name:        &validMeasurementUnit.Name,
		Description: &validMeasurementUnit.Description,
		Volumetric:  &validMeasurementUnit.Volumetric,
		IconPath:    &validMeasurementUnit.IconPath,
		Universal:   &validMeasurementUnit.Universal,
		Metric:      &validMeasurementUnit.Metric,
		Imperial:    &validMeasurementUnit.Imperial,
		PluralName:  &validMeasurementUnit.PluralName,
	}
}

// BuildFakeValidMeasurementUnitCreationRequestInput builds a faked ValidMeasurementUnitCreationRequestInput.
func BuildFakeValidMeasurementUnitCreationRequestInput() *types.ValidMeasurementUnitCreationRequestInput {
	validMeasurementUnit := BuildFakeValidMeasurementUnit()
	return converters.ConvertValidMeasurementUnitToValidMeasurementUnitCreationRequestInput(validMeasurementUnit)
}
