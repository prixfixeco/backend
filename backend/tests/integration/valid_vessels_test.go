package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/pkg/pointer"
	"github.com/dinnerdonebetter/backend/pkg/apiclient"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkValidVesselEquality(t *testing.T, expected, actual *types.ValidVessel) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.PluralName, actual.PluralName)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.IconPath, actual.IconPath)
	assert.Equal(t, expected.Slug, actual.Slug)
	assert.Equal(t, expected.Shape, actual.Shape)
	assert.Equal(t, expected.CapacityUnit.ID, actual.CapacityUnit.ID)
	assert.Equal(t, expected.WidthInMillimeters, actual.WidthInMillimeters)
	assert.Equal(t, expected.LengthInMillimeters, actual.LengthInMillimeters)
	assert.Equal(t, expected.HeightInMillimeters, actual.HeightInMillimeters)
	assert.Equal(t, expected.Capacity, actual.Capacity)
	assert.Equal(t, expected.IncludeInGeneratedInstructions, actual.IncludeInGeneratedInstructions)
	assert.Equal(t, expected.DisplayInSummaryLists, actual.DisplayInSummaryLists)
	assert.Equal(t, expected.UsableForStorage, actual.UsableForStorage)
	assert.NotZero(t, actual.CreatedAt)
}

func createValidVesselForTest(t *testing.T, ctx context.Context, vessel *types.ValidVessel, adminClient *apiclient.Client) *types.ValidVessel {
	t.Helper()

	createdValidMeasurementUnit := createValidMeasurementUnitForTest(t, ctx, adminClient)

	exampleValidVessel := vessel
	if exampleValidVessel == nil {
		exampleValidVessel = fakes.BuildFakeValidVessel()
	}

	exampleValidVessel.CapacityUnit = &types.ValidMeasurementUnit{ID: createdValidMeasurementUnit.ID}
	exampleValidVesselInput := converters.ConvertValidVesselToValidVesselCreationRequestInput(exampleValidVessel)
	createdValidVessel, err := adminClient.CreateValidVessel(ctx, exampleValidVesselInput)
	require.NoError(t, err)
	checkValidVesselEquality(t, exampleValidVessel, createdValidVessel)

	createdValidVessel, err = adminClient.GetValidVessel(ctx, createdValidVessel.ID)
	requireNotNilAndNoProblems(t, createdValidVessel, err)
	checkValidVesselEquality(t, exampleValidVessel, createdValidVessel)

	return createdValidVessel
}

// TODO: handle creating and reading a vessel that doesn't have a CapacityUnit

func (s *TestSuite) TestValidVessels_CompleteLifecycle() {
	s.runTest("should CRUD", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			createdValidVessel := createValidVesselForTest(t, ctx, nil, testClients.adminClient)

			newValidVessel := fakes.BuildFakeValidVessel()
			newValidVessel.CapacityUnit = createdValidVessel.CapacityUnit
			updateInput := converters.ConvertValidVesselToValidVesselUpdateRequestInput(newValidVessel)
			createdValidVessel.Update(updateInput)
			assert.NoError(t, testClients.adminClient.UpdateValidVessel(ctx, createdValidVessel.ID, updateInput))

			actual, err := testClients.adminClient.GetValidVessel(ctx, createdValidVessel.ID)
			requireNotNilAndNoProblems(t, actual, err)

			// assert valid vessel equality
			checkValidVesselEquality(t, newValidVessel, actual)
			assert.NotNil(t, actual.LastUpdatedAt)

			assert.NoError(t, testClients.adminClient.ArchiveValidVessel(ctx, createdValidVessel.ID))
		}
	})
}

func (s *TestSuite) TestValidVessels_GetRandom() {
	s.runTest("should be able to get a random valid vessel", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			createdValidVessel := createValidVesselForTest(t, ctx, nil, testClients.adminClient)

			actual, err := testClients.adminClient.GetRandomValidVessel(ctx)
			requireNotNilAndNoProblems(t, actual, err)

			assert.NoError(t, testClients.adminClient.ArchiveValidVessel(ctx, createdValidVessel.ID))
		}
	})
}

func (s *TestSuite) TestValidVessels_Listing() {
	s.runTest("should be readable in paginated form", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			var expected []*types.ValidVessel
			for i := 0; i < 5; i++ {
				createdValidVessel := createValidVesselForTest(t, ctx, nil, testClients.adminClient)
				expected = append(expected, createdValidVessel)
			}

			// assert valid vessel list equality
			actual, err := testClients.adminClient.GetValidVessels(ctx, nil)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual.Data),
				"expected %d to be <= %d",
				len(expected),
				len(actual.Data),
			)

			for _, createdValidVessel := range expected {
				assert.NoError(t, testClients.adminClient.ArchiveValidVessel(ctx, createdValidVessel.ID))
			}
		}
	})
}

func (s *TestSuite) TestValidVessels_Searching() {
	s.runTest("should be able to be search for valid vessels", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			var expected []*types.ValidVessel
			exampleValidVessel := fakes.BuildFakeValidVessel()
			exampleValidVessel.Name = fmt.Sprintf("example_%s", testClients.authType)
			searchQuery := exampleValidVessel.Name
			for i := 0; i < 5; i++ {
				exampleValidVessel.Name = fmt.Sprintf("%s %d", searchQuery, i)
				createdValidVessel := createValidVesselForTest(t, ctx, exampleValidVessel, testClients.adminClient)
				expected = append(expected, createdValidVessel)
			}

			filter := types.DefaultQueryFilter()
			filter.Limit = pointer.To(uint8(20))

			// assert valid vessel list equality
			actual, err := testClients.adminClient.SearchForValidVessels(ctx, searchQuery, filter)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual.Data),
				"expected %d to be <= %d",
				len(expected),
				len(actual.Data),
			)

			for _, createdValidVessel := range expected {
				assert.NoError(t, testClients.adminClient.ArchiveValidVessel(ctx, createdValidVessel.ID))
			}
		}
	})
}
