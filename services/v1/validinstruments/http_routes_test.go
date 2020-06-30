package validinstruments

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mockencoding "gitlab.com/prixfixe/prixfixe/internal/v1/encoding/mock"
	mockmetrics "gitlab.com/prixfixe/prixfixe/internal/v1/metrics/mock"
	models "gitlab.com/prixfixe/prixfixe/models/v1"
	fakemodels "gitlab.com/prixfixe/prixfixe/models/v1/fake"
	mockmodels "gitlab.com/prixfixe/prixfixe/models/v1/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocknewsman "gitlab.com/verygoodsoftwarenotvirus/newsman/mock"
)

func TestValidInstrumentsService_ListHandler(T *testing.T) {
	T.Parallel()

	T.Run("happy path", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrumentList := fakemodels.BuildFakeValidInstrumentList()

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstruments", mock.Anything, mock.AnythingOfType("*models.QueryFilter")).Return(exampleValidInstrumentList, nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrumentList")).Return(nil)
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ListHandler()(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager, ed)
	})

	T.Run("with no rows returned", func(t *testing.T) {
		s := buildTestService()

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstruments", mock.Anything, mock.AnythingOfType("*models.QueryFilter")).Return((*models.ValidInstrumentList)(nil), sql.ErrNoRows)
		s.validInstrumentDataManager = validInstrumentDataManager

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrumentList")).Return(nil)
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ListHandler()(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager, ed)
	})

	T.Run("with error fetching valid instruments from database", func(t *testing.T) {
		s := buildTestService()

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstruments", mock.Anything, mock.AnythingOfType("*models.QueryFilter")).Return((*models.ValidInstrumentList)(nil), errors.New("blah"))
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ListHandler()(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error encoding response", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrumentList := fakemodels.BuildFakeValidInstrumentList()

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstruments", mock.Anything, mock.AnythingOfType("*models.QueryFilter")).Return(exampleValidInstrumentList, nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrumentList")).Return(errors.New("blah"))
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ListHandler()(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager, ed)
	})
}

func TestValidInstrumentsService_CreateHandler(T *testing.T) {
	T.Parallel()

	T.Run("happy path", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		exampleInput := fakemodels.BuildFakeValidInstrumentCreationInputFromValidInstrument(exampleValidInstrument)

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("CreateValidInstrument", mock.Anything, mock.AnythingOfType("*models.ValidInstrumentCreationInput")).Return(exampleValidInstrument, nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		mc := &mockmetrics.UnitCounter{}
		mc.On("Increment", mock.Anything)
		s.validInstrumentCounter = mc

		r := &mocknewsman.Reporter{}
		r.On("Report", mock.AnythingOfType("newsman.Event")).Return()
		s.reporter = r

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(nil)
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), CreateMiddlewareCtxKey, exampleInput))

		s.CreateHandler()(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager, mc, r, ed)
	})

	T.Run("without input attached", func(t *testing.T) {
		s := buildTestService()

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.CreateHandler()(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	T.Run("with error creating valid instrument", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		exampleInput := fakemodels.BuildFakeValidInstrumentCreationInputFromValidInstrument(exampleValidInstrument)

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("CreateValidInstrument", mock.Anything, mock.AnythingOfType("*models.ValidInstrumentCreationInput")).Return(exampleValidInstrument, errors.New("blah"))
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), CreateMiddlewareCtxKey, exampleInput))

		s.CreateHandler()(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error encoding response", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		exampleInput := fakemodels.BuildFakeValidInstrumentCreationInputFromValidInstrument(exampleValidInstrument)

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("CreateValidInstrument", mock.Anything, mock.AnythingOfType("*models.ValidInstrumentCreationInput")).Return(exampleValidInstrument, nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		mc := &mockmetrics.UnitCounter{}
		mc.On("Increment", mock.Anything)
		s.validInstrumentCounter = mc

		r := &mocknewsman.Reporter{}
		r.On("Report", mock.AnythingOfType("newsman.Event")).Return()
		s.reporter = r

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(errors.New("blah"))
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), CreateMiddlewareCtxKey, exampleInput))

		s.CreateHandler()(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager, mc, r, ed)
	})
}

func TestValidInstrumentsService_ExistenceHandler(T *testing.T) {
	T.Parallel()

	T.Run("happy path", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("ValidInstrumentExists", mock.Anything, exampleValidInstrument.ID).Return(true, nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ExistenceHandler()(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with no such valid instrument in database", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("ValidInstrumentExists", mock.Anything, exampleValidInstrument.ID).Return(false, sql.ErrNoRows)
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ExistenceHandler()(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error fetching valid instrument from database", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("ValidInstrumentExists", mock.Anything, exampleValidInstrument.ID).Return(false, errors.New("blah"))
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ExistenceHandler()(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})
}

func TestValidInstrumentsService_ReadHandler(T *testing.T) {
	T.Parallel()

	T.Run("happy path", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return(exampleValidInstrument, nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(nil)
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ReadHandler()(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager, ed)
	})

	T.Run("with no such valid instrument in database", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return((*models.ValidInstrument)(nil), sql.ErrNoRows)
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ReadHandler()(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error fetching valid instrument from database", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return((*models.ValidInstrument)(nil), errors.New("blah"))
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ReadHandler()(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error encoding response", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return(exampleValidInstrument, nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(errors.New("blah"))
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ReadHandler()(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager, ed)
	})
}

func TestValidInstrumentsService_UpdateHandler(T *testing.T) {
	T.Parallel()

	T.Run("happy path", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		exampleInput := fakemodels.BuildFakeValidInstrumentUpdateInputFromValidInstrument(exampleValidInstrument)

		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return(exampleValidInstrument, nil)
		validInstrumentDataManager.On("UpdateValidInstrument", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		r := &mocknewsman.Reporter{}
		r.On("Report", mock.AnythingOfType("newsman.Event")).Return()
		s.reporter = r

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(nil)
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), UpdateMiddlewareCtxKey, exampleInput))

		s.UpdateHandler()(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		mock.AssertExpectationsForObjects(t, r, validInstrumentDataManager, ed)
	})

	T.Run("without update input", func(t *testing.T) {
		s := buildTestService()

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.UpdateHandler()(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	T.Run("with no rows fetching valid instrument", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		exampleInput := fakemodels.BuildFakeValidInstrumentUpdateInputFromValidInstrument(exampleValidInstrument)

		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return((*models.ValidInstrument)(nil), sql.ErrNoRows)
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), UpdateMiddlewareCtxKey, exampleInput))

		s.UpdateHandler()(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error fetching valid instrument", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		exampleInput := fakemodels.BuildFakeValidInstrumentUpdateInputFromValidInstrument(exampleValidInstrument)

		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return((*models.ValidInstrument)(nil), errors.New("blah"))
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), UpdateMiddlewareCtxKey, exampleInput))

		s.UpdateHandler()(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error updating valid instrument", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		exampleInput := fakemodels.BuildFakeValidInstrumentUpdateInputFromValidInstrument(exampleValidInstrument)

		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return(exampleValidInstrument, nil)
		validInstrumentDataManager.On("UpdateValidInstrument", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(errors.New("blah"))
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), UpdateMiddlewareCtxKey, exampleInput))

		s.UpdateHandler()(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error encoding response", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		exampleInput := fakemodels.BuildFakeValidInstrumentUpdateInputFromValidInstrument(exampleValidInstrument)

		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("GetValidInstrument", mock.Anything, exampleValidInstrument.ID).Return(exampleValidInstrument, nil)
		validInstrumentDataManager.On("UpdateValidInstrument", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		r := &mocknewsman.Reporter{}
		r.On("Report", mock.AnythingOfType("newsman.Event")).Return()
		s.reporter = r

		ed := &mockencoding.EncoderDecoder{}
		ed.On("EncodeResponse", mock.Anything, mock.AnythingOfType("*models.ValidInstrument")).Return(errors.New("blah"))
		s.encoderDecoder = ed

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		req = req.WithContext(context.WithValue(req.Context(), UpdateMiddlewareCtxKey, exampleInput))

		s.UpdateHandler()(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		mock.AssertExpectationsForObjects(t, r, validInstrumentDataManager, ed)
	})
}

func TestValidInstrumentsService_ArchiveHandler(T *testing.T) {
	T.Parallel()

	T.Run("happy path", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("ArchiveValidInstrument", mock.Anything, exampleValidInstrument.ID).Return(nil)
		s.validInstrumentDataManager = validInstrumentDataManager

		r := &mocknewsman.Reporter{}
		r.On("Report", mock.AnythingOfType("newsman.Event")).Return()
		s.reporter = r

		mc := &mockmetrics.UnitCounter{}
		mc.On("Decrement", mock.Anything).Return()
		s.validInstrumentCounter = mc

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ArchiveHandler()(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager, mc, r)
	})

	T.Run("with no valid instrument in database", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("ArchiveValidInstrument", mock.Anything, exampleValidInstrument.ID).Return(sql.ErrNoRows)
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ArchiveHandler()(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})

	T.Run("with error writing to database", func(t *testing.T) {
		s := buildTestService()

		exampleValidInstrument := fakemodels.BuildFakeValidInstrument()
		s.validInstrumentIDFetcher = func(req *http.Request) uint64 {
			return exampleValidInstrument.ID
		}

		validInstrumentDataManager := &mockmodels.ValidInstrumentDataManager{}
		validInstrumentDataManager.On("ArchiveValidInstrument", mock.Anything, exampleValidInstrument.ID).Return(errors.New("blah"))
		s.validInstrumentDataManager = validInstrumentDataManager

		res := httptest.NewRecorder()
		req, err := http.NewRequest(
			http.MethodGet,
			"http://prixfixe.app",
			nil,
		)
		require.NotNil(t, req)
		require.NoError(t, err)

		s.ArchiveHandler()(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

		mock.AssertExpectationsForObjects(t, validInstrumentDataManager)
	})
}
