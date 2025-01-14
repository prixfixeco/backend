package mocktypes

import (
	"context"

	"github.com/dinnerdonebetter/backend/pkg/types"

	"github.com/stretchr/testify/mock"
)

var _ types.OAuth2ClientTokenDataManager = (*OAuth2ClientTokenDataManagerMock)(nil)

// OAuth2ClientTokenDataManagerMock is a mocked types.OAuth2ClientTokenDataManager for testing.
type OAuth2ClientTokenDataManagerMock struct {
	mock.Mock
}

func (m *OAuth2ClientTokenDataManagerMock) CreateOAuth2ClientToken(ctx context.Context, input *types.OAuth2ClientTokenDatabaseCreationInput) (*types.OAuth2ClientToken, error) {
	returnValues := m.Called(ctx, input)
	return returnValues.Get(0).(*types.OAuth2ClientToken), returnValues.Error(1)
}

func (m *OAuth2ClientTokenDataManagerMock) GetOAuth2ClientTokenByCode(ctx context.Context, code string) (*types.OAuth2ClientToken, error) {
	returnValues := m.Called(ctx, code)
	return returnValues.Get(0).(*types.OAuth2ClientToken), returnValues.Error(1)
}

func (m *OAuth2ClientTokenDataManagerMock) GetOAuth2ClientTokenByAccess(ctx context.Context, access string) (*types.OAuth2ClientToken, error) {
	returnValues := m.Called(ctx, access)
	return returnValues.Get(0).(*types.OAuth2ClientToken), returnValues.Error(1)
}

func (m *OAuth2ClientTokenDataManagerMock) GetOAuth2ClientTokenByRefresh(ctx context.Context, refresh string) (*types.OAuth2ClientToken, error) {
	returnValues := m.Called(ctx, refresh)
	return returnValues.Get(0).(*types.OAuth2ClientToken), returnValues.Error(1)
}

func (m *OAuth2ClientTokenDataManagerMock) DeleteOAuth2ClientTokenByAccess(ctx context.Context, access string) error {
	return m.Called(ctx, access).Error(0)
}

func (m *OAuth2ClientTokenDataManagerMock) DeleteOAuth2ClientTokenByCode(ctx context.Context, code string) error {
	return m.Called(ctx, code).Error(0)
}

func (m *OAuth2ClientTokenDataManagerMock) DeleteOAuth2ClientTokenByRefresh(ctx context.Context, refresh string) error {
	return m.Called(ctx, refresh).Error(0)
}
