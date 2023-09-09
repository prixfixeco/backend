package mocktypes

import (
	"context"

	"github.com/dinnerdonebetter/backend/pkg/types"

	"github.com/stretchr/testify/mock"
)

var _ types.WebhookDataManager = (*WebhookDataManagerMock)(nil)

// WebhookDataManagerMock is a mocked types.WebhookDataManager for testing.
type WebhookDataManagerMock struct {
	mock.Mock
}

// WebhookExists satisfies our WebhookDataManagerMock interface.
func (m *WebhookDataManagerMock) WebhookExists(ctx context.Context, webhookID, householdID string) (bool, error) {
	args := m.Called(ctx, webhookID, householdID)
	return args.Bool(0), args.Error(1)
}

// GetWebhook satisfies our WebhookDataManagerMock interface.
func (m *WebhookDataManagerMock) GetWebhook(ctx context.Context, webhookID, householdID string) (*types.Webhook, error) {
	args := m.Called(ctx, webhookID, householdID)
	return args.Get(0).(*types.Webhook), args.Error(1)
}

// GetWebhooks satisfies our WebhookDataManagerMock interface.
func (m *WebhookDataManagerMock) GetWebhooks(ctx context.Context, householdID string, filter *types.QueryFilter) (*types.QueryFilteredResult[types.Webhook], error) {
	args := m.Called(ctx, householdID, filter)
	return args.Get(0).(*types.QueryFilteredResult[types.Webhook]), args.Error(1)
}

// GetWebhooksForHouseholdAndEvent satisfies our WebhookDataManagerMock interface.
func (m *WebhookDataManagerMock) GetWebhooksForHouseholdAndEvent(ctx context.Context, householdID string, eventType types.CustomerEventType) ([]*types.Webhook, error) {
	args := m.Called(ctx, householdID, eventType)
	return args.Get(0).([]*types.Webhook), args.Error(1)
}

// GetAllWebhooks satisfies our WebhookDataManagerMock interface.
func (m *WebhookDataManagerMock) GetAllWebhooks(ctx context.Context, results chan []*types.Webhook, bucketSize uint16) error {
	return m.Called(ctx, results, bucketSize).Error(0)
}

// CreateWebhook satisfies our WebhookDataManagerMock interface.
func (m *WebhookDataManagerMock) CreateWebhook(ctx context.Context, input *types.WebhookDatabaseCreationInput) (*types.Webhook, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*types.Webhook), args.Error(1)
}

// UpdateWebhook satisfies our WebhookDataManagerMock interface.
func (m *WebhookDataManagerMock) UpdateWebhook(ctx context.Context, updated *types.Webhook) error {
	return m.Called(ctx, updated).Error(0)
}

// ArchiveWebhook satisfies our WebhookDataManagerMock interface.
func (m *WebhookDataManagerMock) ArchiveWebhook(ctx context.Context, webhookID, householdID string) error {
	return m.Called(ctx, webhookID, householdID).Error(0)
}
