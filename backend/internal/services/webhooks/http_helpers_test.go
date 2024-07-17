package webhooks

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dinnerdonebetter/backend/internal/authorization"
	"github.com/dinnerdonebetter/backend/internal/encoding"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"
	testutils "github.com/dinnerdonebetter/backend/tests/utils"
)

type webhooksServiceHTTPRoutesTestHelper struct {
	ctx                              context.Context
	req                              *http.Request
	res                              *httptest.ResponseRecorder
	service                          *service
	exampleUser                      *types.User
	exampleHousehold                 *types.Household
	exampleWebhook                   *types.Webhook
	exampleWebhookTriggerEvent       *types.WebhookTriggerEvent
	exampleCreationInput             *types.WebhookCreationRequestInput
	exampleTriggerEventCreationInput *types.WebhookTriggerEventCreationRequestInput
}

func newTestHelper(t *testing.T) *webhooksServiceHTTPRoutesTestHelper {
	t.Helper()

	helper := &webhooksServiceHTTPRoutesTestHelper{}

	helper.ctx = context.Background()
	helper.service = buildTestService()
	helper.exampleUser = fakes.BuildFakeUser()
	helper.exampleHousehold = fakes.BuildFakeHousehold()
	helper.exampleHousehold.BelongsToUser = helper.exampleUser.ID
	helper.exampleWebhook = fakes.BuildFakeWebhook()
	helper.exampleWebhook.BelongsToHousehold = helper.exampleHousehold.ID
	helper.exampleWebhookTriggerEvent = fakes.BuildFakeWebhookTriggerEvent()
	helper.exampleWebhookTriggerEvent.BelongsToWebhook = helper.exampleWebhook.ID
	helper.exampleCreationInput = converters.ConvertWebhookToWebhookCreationRequestInput(helper.exampleWebhook)
	helper.exampleTriggerEventCreationInput = converters.ConvertWebhookTriggerEventToWebhookTriggerEventCreationRequestInput(fakes.BuildFakeWebhookTriggerEvent())

	helper.service.webhookIDFetcher = func(*http.Request) string {
		return helper.exampleWebhook.ID
	}

	helper.service.webhookTriggerEventIDFetcher = func(*http.Request) string {
		return helper.exampleWebhookTriggerEvent.ID
	}

	sessionCtxData := &types.SessionContextData{
		Requester: types.RequesterInfo{
			UserID:                   helper.exampleUser.ID,
			AccountStatus:            helper.exampleUser.AccountStatus,
			AccountStatusExplanation: helper.exampleUser.AccountStatusExplanation,
			ServicePermissions:       authorization.NewServiceRolePermissionChecker(helper.exampleUser.ServiceRole),
		},
		ActiveHouseholdID: helper.exampleHousehold.ID,
		HouseholdPermissions: map[string]authorization.HouseholdRolePermissionsChecker{
			helper.exampleHousehold.ID: authorization.NewHouseholdRolePermissionChecker(authorization.HouseholdMemberRole.String()),
		},
	}

	helper.service.encoderDecoder = encoding.ProvideServerEncoderDecoder(logging.NewNoopLogger(), tracing.NewNoopTracerProvider(), encoding.ContentTypeJSON)
	helper.service.sessionContextDataFetcher = func(*http.Request) (*types.SessionContextData, error) {
		return sessionCtxData, nil
	}

	req := testutils.BuildTestRequest(t)

	helper.req = req.WithContext(context.WithValue(req.Context(), types.SessionContextDataKey, sessionCtxData))
	helper.res = httptest.NewRecorder()

	return helper
}