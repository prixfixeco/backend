package mock

import (
	"context"

	models "gitlab.com/prixfixe/prixfixe/models/v1"

	"github.com/stretchr/testify/mock"
)

var _ models.ValidIngredientPreparationDataManager = (*ValidIngredientPreparationDataManager)(nil)

// ValidIngredientPreparationDataManager is a mocked models.ValidIngredientPreparationDataManager for testing.
type ValidIngredientPreparationDataManager struct {
	mock.Mock
}

// ValidIngredientPreparationExists is a mock function.
func (m *ValidIngredientPreparationDataManager) ValidIngredientPreparationExists(ctx context.Context, validIngredientPreparationID uint64) (bool, error) {
	args := m.Called(ctx, validIngredientPreparationID)
	return args.Bool(0), args.Error(1)
}

// GetValidIngredientPreparation is a mock function.
func (m *ValidIngredientPreparationDataManager) GetValidIngredientPreparation(ctx context.Context, validIngredientPreparationID uint64) (*models.ValidIngredientPreparation, error) {
	args := m.Called(ctx, validIngredientPreparationID)
	return args.Get(0).(*models.ValidIngredientPreparation), args.Error(1)
}

// GetAllValidIngredientPreparationsCount is a mock function.
func (m *ValidIngredientPreparationDataManager) GetAllValidIngredientPreparationsCount(ctx context.Context) (uint64, error) {
	args := m.Called(ctx)
	return args.Get(0).(uint64), args.Error(1)
}

// GetAllValidIngredientPreparations is a mock function.
func (m *ValidIngredientPreparationDataManager) GetAllValidIngredientPreparations(ctx context.Context, results chan []models.ValidIngredientPreparation) error {
	args := m.Called(ctx, results)
	return args.Error(0)
}

// GetValidIngredientPreparations is a mock function.
func (m *ValidIngredientPreparationDataManager) GetValidIngredientPreparations(ctx context.Context, filter *models.QueryFilter) (*models.ValidIngredientPreparationList, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*models.ValidIngredientPreparationList), args.Error(1)
}

// GetValidIngredientPreparationsWithIDs is a mock function.
func (m *ValidIngredientPreparationDataManager) GetValidIngredientPreparationsWithIDs(ctx context.Context, limit uint8, ids []uint64) ([]models.ValidIngredientPreparation, error) {
	args := m.Called(ctx, limit, ids)
	return args.Get(0).([]models.ValidIngredientPreparation), args.Error(1)
}

// CreateValidIngredientPreparation is a mock function.
func (m *ValidIngredientPreparationDataManager) CreateValidIngredientPreparation(ctx context.Context, input *models.ValidIngredientPreparationCreationInput) (*models.ValidIngredientPreparation, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*models.ValidIngredientPreparation), args.Error(1)
}

// UpdateValidIngredientPreparation is a mock function.
func (m *ValidIngredientPreparationDataManager) UpdateValidIngredientPreparation(ctx context.Context, updated *models.ValidIngredientPreparation) error {
	return m.Called(ctx, updated).Error(0)
}

// ArchiveValidIngredientPreparation is a mock function.
func (m *ValidIngredientPreparationDataManager) ArchiveValidIngredientPreparation(ctx context.Context, validIngredientPreparationID uint64) error {
	return m.Called(ctx, validIngredientPreparationID).Error(0)
}
