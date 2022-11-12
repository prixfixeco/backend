package integration

import (
	"fmt"
	"github.com/prixfixeco/backend/pkg/types/converters"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/prixfixeco/backend/internal/observability/tracing"
	"github.com/prixfixeco/backend/pkg/types/fakes"

	"github.com/prixfixeco/backend/pkg/types"
)

func checkValidInstrumentEquality(t *testing.T, expected, actual *types.ValidInstrument) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.Name, actual.Name, "expected Name for valid instrument %s to be %v, but it was %v", expected.ID, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description, "expected Description for valid instrument %s to be %v, but it was %v", expected.ID, expected.Description, actual.Description)
	assert.Equal(t, expected.IconPath, actual.IconPath, "expected IconPath for valid instrument %s to be %v, but it was %v", expected.ID, expected.IconPath, actual.IconPath)
	assert.Equal(t, expected.UsableForStorage, actual.UsableForStorage, "expected UsableForStorage for valid instrument %s to be %v, but it was %v", expected.ID, expected.UsableForStorage, actual.UsableForStorage)
	assert.Equal(t, expected.DisplayInSummaryLists, actual.DisplayInSummaryLists, "expected DisplayInSummaryLists for valid instrument %s to be %v, but it was %v", expected.ID, expected.DisplayInSummaryLists, actual.DisplayInSummaryLists)
	assert.Equal(t, expected.Slug, actual.Slug, "expected UsableForStorage for valid instrument %s to be %v, but it was %v", expected.ID, expected.UsableForStorage, actual.UsableForStorage)
	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestValidInstruments_CompleteLifecycle() {
	s.runForEachClient("should be creatable and readable and updatable and deletable", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			t.Log("creating valid instrument")
			exampleValidInstrument := fakes.BuildFakeValidInstrument()
			exampleValidInstrumentInput := converters.ConvertValidInstrumentToValidInstrumentCreationRequestInput(exampleValidInstrument)
			createdValidInstrument, err := testClients.admin.CreateValidInstrument(ctx, exampleValidInstrumentInput)
			require.NoError(t, err)
			t.Logf("valid instrument %q created", createdValidInstrument.ID)
			checkValidInstrumentEquality(t, exampleValidInstrument, createdValidInstrument)

			t.Log("changing valid instrument")
			newValidInstrument := fakes.BuildFakeValidInstrument()
			createdValidInstrument.Update(converters.ConvertValidInstrumentToValidInstrumentUpdateRequestInput(newValidInstrument))
			assert.NoError(t, testClients.admin.UpdateValidInstrument(ctx, createdValidInstrument))

			t.Log("fetching changed valid instrument")
			actual, err := testClients.admin.GetValidInstrument(ctx, createdValidInstrument.ID)
			requireNotNilAndNoProblems(t, actual, err)

			// assert valid instrument equality
			checkValidInstrumentEquality(t, newValidInstrument, actual)
			assert.NotNil(t, actual.LastUpdatedAt)

			t.Log("cleaning up valid instrument")
			assert.NoError(t, testClients.admin.ArchiveValidInstrument(ctx, createdValidInstrument.ID))
		}
	})
}

func (s *TestSuite) TestValidInstruments_GetRandom() {
	s.runForEachClient("should be able to get a random valid instrument", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			t.Log("creating valid instrument")
			exampleValidInstrument := fakes.BuildFakeValidInstrument()
			exampleValidInstrumentInput := converters.ConvertValidInstrumentToValidInstrumentCreationRequestInput(exampleValidInstrument)
			createdValidInstrument, err := testClients.admin.CreateValidInstrument(ctx, exampleValidInstrumentInput)
			require.NoError(t, err)
			t.Logf("valid instrument %q created", createdValidInstrument.ID)
			checkValidInstrumentEquality(t, exampleValidInstrument, createdValidInstrument)

			t.Log("fetching changed valid instrument")
			actual, err := testClients.admin.GetRandomValidInstrument(ctx)
			requireNotNilAndNoProblems(t, actual, err)

			t.Log("cleaning up valid instrument")
			assert.NoError(t, testClients.admin.ArchiveValidInstrument(ctx, createdValidInstrument.ID))
		}
	})
}

func (s *TestSuite) TestValidInstruments_Listing() {
	s.runForEachClient("should be readable in paginated form", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			t.Log("creating valid instruments")
			var expected []*types.ValidInstrument
			for i := 0; i < 5; i++ {
				exampleValidInstrument := fakes.BuildFakeValidInstrument()
				exampleValidInstrumentInput := converters.ConvertValidInstrumentToValidInstrumentCreationRequestInput(exampleValidInstrument)
				createdValidInstrument, createdValidInstrumentErr := testClients.admin.CreateValidInstrument(ctx, exampleValidInstrumentInput)
				require.NoError(t, createdValidInstrumentErr)
				t.Logf("valid instrument %q created", createdValidInstrument.ID)

				checkValidInstrumentEquality(t, exampleValidInstrument, createdValidInstrument)

				expected = append(expected, createdValidInstrument)
			}

			// assert valid instrument list equality
			actual, err := testClients.admin.GetValidInstruments(ctx, nil)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual.ValidInstruments),
				"expected %d to be <= %d",
				len(expected),
				len(actual.ValidInstruments),
			)

			t.Log("cleaning up")
			for _, createdValidInstrument := range expected {
				assert.NoError(t, testClients.admin.ArchiveValidInstrument(ctx, createdValidInstrument.ID))
			}
		}
	})
}

func (s *TestSuite) TestValidInstruments_Searching() {
	s.runForEachClient("should be able to be search for valid instruments", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			t.Log("creating valid instruments")
			var expected []*types.ValidInstrument
			exampleValidInstrument := fakes.BuildFakeValidInstrument()
			exampleValidInstrument.Name = fmt.Sprintf("example_%s", testClients.authType)
			searchQuery := exampleValidInstrument.Name
			for i := 0; i < 5; i++ {
				exampleValidInstrument.Name = fmt.Sprintf("%s %d", searchQuery, i)
				exampleValidInstrumentInput := converters.ConvertValidInstrumentToValidInstrumentCreationRequestInput(exampleValidInstrument)
				createdValidInstrument, createdValidInstrumentErr := testClients.admin.CreateValidInstrument(ctx, exampleValidInstrumentInput)
				require.NoError(t, createdValidInstrumentErr)
				t.Logf("valid instrument %q created", createdValidInstrument.ID)
				checkValidInstrumentEquality(t, exampleValidInstrument, createdValidInstrument)

				expected = append(expected, createdValidInstrument)
			}

			exampleLimit := uint8(20)

			// assert valid instrument list equality
			actual, err := testClients.admin.SearchValidInstruments(ctx, searchQuery, exampleLimit)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(
				t,
				len(expected) <= len(actual),
				"expected %d to be <= %d",
				len(expected),
				len(actual),
			)

			t.Log("cleaning up")
			for _, createdValidInstrument := range expected {
				assert.NoError(t, testClients.admin.ArchiveValidInstrument(ctx, createdValidInstrument.ID))
			}
		}
	})
}
