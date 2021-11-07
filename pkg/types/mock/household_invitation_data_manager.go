package mocktypes

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/prixfixeco/api_server/pkg/types"
)

var _ types.HouseholdInvitationDataManager = (*HouseholdInvitationDataManager)(nil)

// HouseholdInvitationDataManager is a mocked types.HouseholdInvitationDataManager for testing.
type HouseholdInvitationDataManager struct {
	mock.Mock
}

// HouseholdInvitationExists is a mock function.
func (m *HouseholdInvitationDataManager) HouseholdInvitationExists(ctx context.Context, householdInvitationID string) (bool, error) {
	args := m.Called(ctx, householdInvitationID)
	return args.Bool(0), args.Error(1)
}

// GetHouseholdInvitationByHouseholdAndID is a mock function.
func (m *HouseholdInvitationDataManager) GetHouseholdInvitationByHouseholdAndID(ctx context.Context, householdID, userID string) (*types.HouseholdInvitation, error) {
	args := m.Called(ctx, householdID, userID)
	return args.Get(0).(*types.HouseholdInvitation), args.Error(1)
}

// GetHouseholdInvitationByEmailAndToken is a mock function.
func (m *HouseholdInvitationDataManager) GetHouseholdInvitationByEmailAndToken(ctx context.Context, emailAddress, token string) (*types.HouseholdInvitation, error) {
	args := m.Called(ctx, emailAddress, token)
	return args.Get(0).(*types.HouseholdInvitation), args.Error(1)
}

// GetAllHouseholdInvitationsCount is a mock function.
func (m *HouseholdInvitationDataManager) GetAllHouseholdInvitationsCount(ctx context.Context) (uint64, error) {
	args := m.Called(ctx)
	return args.Get(0).(uint64), args.Error(1)
}

// GetPendingHouseholdInvitationsFromUser is a mock function.
func (m *HouseholdInvitationDataManager) GetPendingHouseholdInvitationsFromUser(ctx context.Context, userID string, filter *types.QueryFilter) (*types.HouseholdInvitationList, error) {
	args := m.Called(ctx, userID, filter)
	return args.Get(0).(*types.HouseholdInvitationList), args.Error(1)
}

// GetPendingHouseholdInvitationsForUser is a mock function.
func (m *HouseholdInvitationDataManager) GetPendingHouseholdInvitationsForUser(ctx context.Context, userID string, filter *types.QueryFilter) (*types.HouseholdInvitationList, error) {
	args := m.Called(ctx, userID, filter)
	return args.Get(0).(*types.HouseholdInvitationList), args.Error(1)
}

// CreateHouseholdInvitation is a mock function.
func (m *HouseholdInvitationDataManager) CreateHouseholdInvitation(ctx context.Context, input *types.HouseholdInvitationDatabaseCreationInput) (*types.HouseholdInvitation, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*types.HouseholdInvitation), args.Error(1)
}

// CancelHouseholdInvitation is a mock function.
func (m *HouseholdInvitationDataManager) CancelHouseholdInvitation(ctx context.Context, householdID, householdInvitationID, note string) error {
	args := m.Called(ctx, householdID, householdInvitationID, note)
	return args.Error(0)
}

// AcceptHouseholdInvitation is a mock function.
func (m *HouseholdInvitationDataManager) AcceptHouseholdInvitation(ctx context.Context, householdID, householdInvitationID, note string) error {
	args := m.Called(ctx, householdID, householdInvitationID, note)
	return args.Error(0)
}

// RejectHouseholdInvitation is a mock function.
func (m *HouseholdInvitationDataManager) RejectHouseholdInvitation(ctx context.Context, householdID, householdInvitationID, note string) error {
	args := m.Called(ctx, householdID, householdInvitationID, note)
	return args.Error(0)
}

// ArchiveHouseholdInvitation is a mock function.
func (m *HouseholdInvitationDataManager) ArchiveHouseholdInvitation(ctx context.Context, householdID, userID string) error {
	return m.Called(ctx, householdID, userID).Error(0)
}
