package validingredientstateingredients

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/database"
	"github.com/dinnerdonebetter/backend/internal/encoding"
	mockpublishers "github.com/dinnerdonebetter/backend/internal/messagequeue/mock"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/internal/pkg/testutils"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"
	mocktypes "github.com/dinnerdonebetter/backend/pkg/types/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestValidIngredientStateIngredientsService_CreateValidIngredientStateIngredientHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientCreationRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"CreateValidIngredientStateIngredient",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientStateIngredientDatabaseCreationInput) bool { return true }),
		).Return(helper.exampleValidIngredientStateIngredient, nil)
		helper.service.validIngredientStateIngredientDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(nil)
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.CreateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusCreated, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, helper.exampleValidIngredientStateIngredient)
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

		helper.service.CreateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusBadRequest, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with invalid input attached", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := &types.ValidIngredientStateIngredientCreationRequestInput{}
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		helper.service.CreateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusBadRequest, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientCreationRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.CreateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with error writing to database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientCreationRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"CreateValidIngredientStateIngredient",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientStateIngredientDatabaseCreationInput) bool { return true }),
		).Return((*types.ValidIngredientStateIngredient)(nil), errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = dbManager

		helper.service.CreateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, dbManager)
	})

	T.Run("with error publishing event", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientCreationRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"CreateValidIngredientStateIngredient",
			testutils.ContextMatcher,
			mock.MatchedBy(func(*types.ValidIngredientStateIngredientDatabaseCreationInput) bool { return true }),
		).Return(helper.exampleValidIngredientStateIngredient, nil)
		helper.service.validIngredientStateIngredientDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(errors.New("blah"))
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.CreateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusCreated, helper.res.Code)

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})
}

func TestValidIngredientStateIngredientsService_ReadValidIngredientStateIngredientHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(helper.exampleValidIngredientStateIngredient, nil)
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.ReadValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, helper.exampleValidIngredientStateIngredient)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.ReadValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no such valid ingredient preparation in the database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return((*types.ValidIngredientStateIngredient)(nil), sql.ErrNoRows)
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.ReadValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusNotFound, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error fetching from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return((*types.ValidIngredientStateIngredient)(nil), errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.ReadValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})
}

func TestValidIngredientStateIngredientsService_ListValidIngredientStateIngredientsHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		exampleValidIngredientStateIngredientList := fakes.BuildFakeValidIngredientStateIngredientsList()

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredients",
			testutils.ContextMatcher,
			mock.IsType(&types.QueryFilter{}),
		).Return(exampleValidIngredientStateIngredientList, nil)
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.ListValidIngredientStateIngredientsHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[[]*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, exampleValidIngredientStateIngredientList.Data)
		assert.Equal(t, *actual.Pagination, exampleValidIngredientStateIngredientList.Pagination)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.ListValidIngredientStateIngredientsHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no rows returned", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredients",
			testutils.ContextMatcher,
			mock.IsType(&types.QueryFilter{}),
		).Return((*types.QueryFilteredResult[types.ValidIngredientStateIngredient])(nil), sql.ErrNoRows)
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.ListValidIngredientStateIngredientsHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[[]*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error retrieving valid ingredient preparations from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredients",
			testutils.ContextMatcher,
			mock.IsType(&types.QueryFilter{}),
		).Return((*types.QueryFilteredResult[types.ValidIngredientStateIngredient])(nil), errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.ListValidIngredientStateIngredientsHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})
}

func TestValidIngredientStateIngredientsService_UpdateValidIngredientStateIngredientHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"GetValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(helper.exampleValidIngredientStateIngredient, nil)

		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"UpdateValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient,
		).Return(nil)
		helper.service.validIngredientStateIngredientDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(nil)
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.UpdateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, helper.exampleValidIngredientStateIngredient)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})

	T.Run("with invalid input", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := &types.ValidIngredientStateIngredientUpdateRequestInput{}
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		helper.service.UpdateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusBadRequest, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.UpdateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
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

		helper.service.UpdateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusBadRequest, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no such valid ingredient preparation", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return((*types.ValidIngredientStateIngredient)(nil), sql.ErrNoRows)
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.UpdateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusNotFound, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error retrieving valid ingredient preparation from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return((*types.ValidIngredientStateIngredient)(nil), errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.UpdateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error writing to database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"GetValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(helper.exampleValidIngredientStateIngredient, nil)

		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"UpdateValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient,
		).Return(errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = dbManager

		helper.service.UpdateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, dbManager)
	})

	T.Run("with error publishing to message queue", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)

		exampleCreationInput := fakes.BuildFakeValidIngredientStateIngredientUpdateRequestInput()
		jsonBytes := helper.service.encoderDecoder.MustEncode(helper.ctx, exampleCreationInput)

		var err error
		helper.req, err = http.NewRequestWithContext(helper.ctx, http.MethodPost, "https://whatever.whocares.gov", bytes.NewReader(jsonBytes))
		require.NoError(t, err)
		require.NotNil(t, helper.req)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"GetValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(helper.exampleValidIngredientStateIngredient, nil)

		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"UpdateValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient,
		).Return(nil)
		helper.service.validIngredientStateIngredientDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(errors.New("blah"))
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.UpdateValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})
}

func TestValidIngredientStateIngredientsService_ArchiveValidIngredientStateIngredientHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"ValidIngredientStateIngredientExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(true, nil)

		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"ArchiveValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(nil)
		helper.service.validIngredientStateIngredientDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(nil)
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.ArchiveValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})

	T.Run("with error retrieving session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)
		helper.service.sessionContextDataFetcher = testutils.BrokenSessionContextDataFetcher

		helper.service.ArchiveValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with no such valid ingredient preparation in the database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"ValidIngredientStateIngredientExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(false, nil)
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.ArchiveValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusNotFound, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error checking for item in database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"ValidIngredientStateIngredientExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(false, errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.ArchiveValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error writing to database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"ValidIngredientStateIngredientExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(true, nil)

		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"ArchiveValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = dbManager

		helper.service.ArchiveValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, dbManager)
	})

	T.Run("with error publishing to message queue", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		dbManager := database.NewMockDatabase()
		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"ValidIngredientStateIngredientExists",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(true, nil)

		dbManager.ValidIngredientStateIngredientDataManagerMock.On(
			"ArchiveValidIngredientStateIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredientStateIngredient.ID,
		).Return(nil)
		helper.service.validIngredientStateIngredientDataManager = dbManager

		dataChangesPublisher := &mockpublishers.Publisher{}
		dataChangesPublisher.On(
			"Publish",
			testutils.ContextMatcher,
			testutils.DataChangeMessageMatcher,
		).Return(errors.New("blah"))
		helper.service.dataChangesPublisher = dataChangesPublisher

		helper.service.ArchiveValidIngredientStateIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)

		mock.AssertExpectationsForObjects(t, dbManager, dataChangesPublisher)
	})
}

func TestValidIngredientStateIngredientsService_SearchValidIngredientStateIngredientsByIngredientHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		exampleValidIngredientStateIngredientList := fakes.BuildFakeValidIngredientStateIngredientsList()

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredientsForIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredient.ID,
			testutils.QueryFilterMatcher,
		).Return(exampleValidIngredientStateIngredientList, nil)
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.SearchValidIngredientStateIngredientsByIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[[]*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, exampleValidIngredientStateIngredientList.Data)
		assert.Equal(t, *actual.Pagination, exampleValidIngredientStateIngredientList.Pagination)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error fetching session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		helper.service.sessionContextDataFetcher = func(request *http.Request) (*types.SessionContextData, error) {
			return nil, errors.New("blah")
		}

		helper.service.SearchValidIngredientStateIngredientsByIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with error fetching data from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredientsForIngredient",
			testutils.ContextMatcher,
			helper.exampleValidIngredient.ID,
			testutils.QueryFilterMatcher,
		).Return((*types.QueryFilteredResult[types.ValidIngredientStateIngredient])(nil), errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.SearchValidIngredientStateIngredientsByIngredientHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})
}

func TestValidIngredientStateIngredientsService_SearchValidIngredientStateIngredientsByPreparationHandler(T *testing.T) {
	T.Parallel()

	T.Run("standard", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		exampleValidIngredientStateIngredientList := fakes.BuildFakeValidIngredientStateIngredientsList()

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredientsForIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
			testutils.QueryFilterMatcher,
		).Return(exampleValidIngredientStateIngredientList, nil)
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.SearchValidIngredientStateIngredientsByIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusOK, helper.res.Code)
		var actual *types.APIResponse[[]*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Equal(t, actual.Data, exampleValidIngredientStateIngredientList.Data)
		assert.Equal(t, *actual.Pagination, exampleValidIngredientStateIngredientList.Pagination)
		assert.NoError(t, actual.Error.AsError())

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})

	T.Run("with error fetching session context data", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		helper.service.sessionContextDataFetcher = func(request *http.Request) (*types.SessionContextData, error) {
			return nil, errors.New("blah")
		}

		helper.service.SearchValidIngredientStateIngredientsByIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusUnauthorized, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)
	})

	T.Run("with error fetching data from database", func(t *testing.T) {
		t.Parallel()

		helper := buildTestHelper(t)

		validIngredientStateIngredientDataManager := &mocktypes.ValidIngredientStateIngredientDataManagerMock{}
		validIngredientStateIngredientDataManager.On(
			"GetValidIngredientStateIngredientsForIngredientState",
			testutils.ContextMatcher,
			helper.exampleValidIngredientState.ID,
			testutils.QueryFilterMatcher,
		).Return((*types.QueryFilteredResult[types.ValidIngredientStateIngredient])(nil), errors.New("blah"))
		helper.service.validIngredientStateIngredientDataManager = validIngredientStateIngredientDataManager

		helper.service.SearchValidIngredientStateIngredientsByIngredientStateHandler(helper.res, helper.req)

		assert.Equal(t, http.StatusInternalServerError, helper.res.Code)
		var actual *types.APIResponse[*types.ValidIngredientStateIngredient]
		require.NoError(t, helper.service.encoderDecoder.DecodeBytes(helper.ctx, helper.res.Body.Bytes(), &actual))
		assert.Empty(t, actual.Data)
		assert.Error(t, actual.Error)

		mock.AssertExpectationsForObjects(t, validIngredientStateIngredientDataManager)
	})
}