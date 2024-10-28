package validingredientstates

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/encoding"
	mockpublishers "github.com/dinnerdonebetter/backend/internal/messagequeue/mock"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/pkg/testutils"
	mocksearch "github.com/dinnerdonebetter/backend/internal/search/mock"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"
	mocktypes "github.com/dinnerdonebetter/backend/pkg/types/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestValidIngredientStatesService_CreateValidIngredientStateHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateCreationRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"CreateValidIngredientState",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientStateDatabaseCreationInput) bool { return true }),
		).Return(helper.exampleValidIngredientState, nil)
		helper.service.validIngredientStateDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(nil)
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.CreateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusCreated, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, helper.exampleValidIngredientState)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})

	T.Run("without input attached", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(nil))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		helper.service.CreateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusBadRequest, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with invalid input attached", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := &types.ValidIngredientStateCreationRequestInput{}
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		helper.service.CreateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusBadRequest, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateCreationRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.CreateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with error writing to database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateCreationRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"CreateValidIngredientState",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientStateDatabaseCreationInput) bool { return true }),
		).Return((*types.ValidIngredientState)(nil), errors.New("blah"))
		helper.service.validIngredientStateDataManager = dbManager

		helper.service.CreateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, dbManager)
	})

	T.Run("with error publishing event", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateCreationRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"CreateValidIngredientState",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientStateDatabaseCreationInput) bool { return true }),
		).Return(helper.exampleValidIngredientState, nil)
		helper.service.validIngredientStateDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(errors.New("blah"))
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.CreateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusCreated, helper.res.Code)

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})
}

func TestValidIngredientStatesService_ReadValidIngredientStateHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(helper.exampleValidIngredientState, nil)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.ReadValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, helper.exampleValidIngredientState)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.ReadValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no such valid ingredient state in the database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return((*types.ValidIngredientState)(nil), sql.ErrNoRows)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.ReadValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusNotFound, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with error fetching from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return((*types.ValidIngredientState)(nil), errors.New("blah"))
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.ReadValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})
}

func TestValidIngredientStatesService_ListValidIngredientStatesHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		exampleValidIngredientStateList := fakes.BuildFakeValidIngredientStatesList()

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientStates",
			testutils.ContextMatcher,
			mock.IsType(&types.QueryFilter{}),
		).Return(exampleValidIngredientStateList, nil)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.ListValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[[]*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, exampleValidIngredientStateList.Data)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.ListValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no rows returned", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientStates",
			testutils.ContextMatcher,
			mock.IsType(&types.QueryFilter{}),
		).Return((*types.QueryFilteredResult[types.ValidIngredientState])(nil), sql.ErrNoRows)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.ListValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[[]*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with error retrieving valid ingredient states from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientStates",
			testutils.ContextMatcher,
			mock.IsType(&types.QueryFilter{}),
		).Return((*types.QueryFilteredResult[types.ValidIngredientState])(nil), errors.New("blah"))
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.ListValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})
}

func TestValidIngredientStatesService_SearchValidIngredientStatesHandler(T *testing.T) {
	T.Parallel()

	exampleQuery := "whatever"
	exampleLimit := uint8(123)
	exampleValidIngredientStateList := fakes.BuildFakeValidIngredientStatesList()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		helper.req.URL.RawQuery = url.Values{
			types.QueryKeySearch: []string{exampleQuery},
			types.QueryKeyLimit:  []string{strconv.Itoa(int(exampleLimit))},
		}.Encode()

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"SearchForValidIngredientStates",
			testutils.ContextMatcher,
			exampleQuery,
		).Return(exampleValidIngredientStateList.Data, nil)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.SearchValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[[]*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, exampleValidIngredientStateList.Data)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("using external service", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.cfg.UseSearchService = true

		helper.req.URL.RawQuery = url.Values{
			types.QueryKeySearch: []string{exampleQuery},
			types.QueryKeyLimit:  []string{strconv.Itoa(int(exampleLimit))},
		}.Encode()

		expectedIDs := []string{}
		validIngredientStateSearchSubsets := make([]*types.ValidIngredientStateSearchSubset, len(exampleValidIngredientStateList.Data))
		for i := range exampleValidIngredientStateList.Data {
			expectedIDs = append(expectedIDs, exampleValidIngredientStateList.Data[i].ID)
			validIngredientStateSearchSubsets[i] = converters.ConvertValidIngredientStateToValidIngredientStateSearchSubset(exampleValidIngredientStateList.Data[i])
		}

		searchIndex := &mocksearch.IndexManager[types.ValidIngredientStateSearchSubset]{}
		searchIndex.On(
			"Search",
			testutils.ContextMatcher,
			exampleQuery,
		).Return(validIngredientStateSearchSubsets, nil)
		helper.service.validIngredientStatesSearchIndex = searchIndex

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientStatesWithIDs",
			testutils.ContextMatcher,
			expectedIDs,
		).Return(exampleValidIngredientStateList.Data, nil)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.SearchValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager, searchIndex)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.SearchValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no rows returned", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		helper.req.URL.RawQuery = url.Values{
			types.QueryKeySearch: []string{exampleQuery},
			types.QueryKeyLimit:  []string{strconv.Itoa(int(exampleLimit))},
		}.Encode()

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"SearchForValidIngredientStates",
			testutils.ContextMatcher,
			exampleQuery,
		).Return([]*types.ValidIngredientState{}, sql.ErrNoRows)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.SearchValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[[]*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with error retrieving from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.req.URL.RawQuery = url.Values{
			types.QueryKeySearch: []string{exampleQuery},
			types.QueryKeyLimit:  []string{strconv.Itoa(int(exampleLimit))},
		}.Encode()

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"SearchForValidIngredientStates",
			testutils.ContextMatcher,
			exampleQuery,
		).Return([]*types.ValidIngredientState{}, errors.New("blah"))
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.SearchValidIngredientStatesHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})
}

func TestValidIngredientStatesService_UpdateValidIngredientStateHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"GetValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(helper.exampleValidIngredientState, nil)

		dbManager.ValidIngredientStateDataManagerMock.On(
			"UpdateValidIngredientState",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientState) bool { return true }),
		).Return(nil)
		helper.service.validIngredientStateDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(nil)
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.UpdateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, helper.exampleValidIngredientState)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})

	T.Run("with invalid input", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := &types.ValidIngredientStateUpdateRequestInput{}
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		helper.service.UpdateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusBadRequest, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.UpdateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("without input attached to context", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(nil))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		helper.service.UpdateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusBadRequest, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no such valid ingredient state", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return((*types.ValidIngredientState)(nil), sql.ErrNoRows)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.UpdateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusNotFound, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with error retrieving valid ingredient state from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"GetValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return((*types.ValidIngredientState)(nil), errors.New("blah"))
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.UpdateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with problem writing to database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"GetValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(helper.exampleValidIngredientState, nil)

		dbManager.ValidIngredientStateDataManagerMock.On(
			"UpdateValidIngredientState",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientState) bool { return true }),
		).Return(errors.New("blah"))
		helper.service.validIngredientStateDataManager = dbManager

		helper.service.UpdateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, dbManager)
	})

	T.Run("with error publishing to message queue", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"GetValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(helper.exampleValidIngredientState, nil)

		dbManager.ValidIngredientStateDataManagerMock.On(
			"UpdateValidIngredientState",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientState) bool { return true }),
		).Return(nil)
		helper.service.validIngredientStateDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(errors.New("blah"))
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.UpdateValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})
}

func TestValidIngredientStatesService_ArchiveValidIngredientStateHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"ValidIngredientStateExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(true, nil)

		dbManager.ValidIngredientStateDataManagerMock.On(
			"ArchiveValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(nil)
		helper.service.validIngredientStateDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(nil)
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.ArchiveValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.ArchiveValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no such valid ingredient state in the database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"ValidIngredientStateExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(false, nil)
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.ArchiveValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusNotFound, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with error checking for item in database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateDataManager := &mocktypes.ValidIngredientStateDataManagerMock{}
		validIngredientStateDataManager.On(
			"ValidIngredientStateExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(false, errors.New("blah"))
		helper.service.validIngredientStateDataManager = validIngredientStateDataManager

		helper.service.ArchiveValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateDataManager)
	})

	T.Run("with error writing to database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"ValidIngredientStateExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(true, nil)

		dbManager.ValidIngredientStateDataManagerMock.On(
			"ArchiveValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(errors.New("blah"))
		helper.service.validIngredientStateDataManager = dbManager

		helper.service.ArchiveValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientState]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, dbManager)
	})

	T.Run("with error publishing to message queue", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateDataManagerMock.On(
			"ValidIngredientStateExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(true, nil)

		dbManager.ValidIngredientStateDataManagerMock.On(
			"ArchiveValidIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
		).Return(nil)
		helper.service.validIngredientStateDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(errors.New("blah"))
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.ArchiveValidIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})
}