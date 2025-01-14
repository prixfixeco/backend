package integration

import (
	"testing"

	"github.com/dinnerdonebetter/backend/internal/observability/tracing"
	"github.com/dinnerdonebetter/backend/pkg/types"
	"github.com/dinnerdonebetter/backend/pkg/types/converters"
	"github.com/dinnerdonebetter/backend/pkg/types/fakes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkWebhookEquality(t *testing.T, expected, actual *types.Webhook) {
	t.Helper()

	assert.NotZero(t, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ContentType, actual.ContentType)
	assert.Equal(t, expected.URL, actual.URL)
	assert.Equal(t, expected.Method, actual.Method)
	assert.NotZero(t, actual.CreatedAt)
}

func (s *TestSuite) TestWebhooks_Creating() {
	s.runTest("should be creatable and readable and deletable", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			// Create webhook.
			exampleWebhook := fakes.BuildFakeWebhook()
			exampleWebhookInput := converters.ConvertWebhookToWebhookCreationRequestInput(exampleWebhook)

			createdWebhook, err := testClients.userClient.CreateWebhook(ctx, exampleWebhookInput)
			require.NoError(t, err)

			createdWebhook, err = testClients.userClient.GetWebhook(ctx, createdWebhook.ID)
			requireNotNilAndNoProblems(t, createdWebhook, err)

			// assert webhook equality
			checkWebhookEquality(t, exampleWebhook, createdWebhook)

			actual, err := testClients.userClient.GetWebhook(ctx, createdWebhook.ID)
			requireNotNilAndNoProblems(t, actual, err)
			checkWebhookEquality(t, exampleWebhook, actual)

			webhookTriggerEvent := fakes.BuildFakeWebhookTriggerEvent()
			webhookTriggerEvent.BelongsToWebhook = createdWebhook.ID
			webhookTriggerEvent.TriggerEvent = string(types.WebhookArchivedServiceEventType)
			eventInput := converters.ConvertWebhookTriggerEventToWebhookTriggerEventCreationRequestInput(webhookTriggerEvent)

			event, err := testClients.userClient.CreateWebhookTriggerEvent(ctx, createdWebhook.ID, eventInput)
			requireNotNilAndNoProblems(t, actual, err)

			// Archive trigger event
			assert.NoError(t, testClients.userClient.ArchiveWebhookTriggerEvent(ctx, createdWebhook.ID, event.ID))

			// Archive trigger event
			assert.NoError(t, testClients.userClient.ArchiveWebhookTriggerEvent(ctx, createdWebhook.ID, actual.Events[0].ID))

			// Clean up.
			assert.NoError(t, testClients.userClient.ArchiveWebhook(ctx, createdWebhook.ID))
		}
	})
}

func (s *TestSuite) TestWebhooks_Reading_Returns404ForNonexistentWebhook() {
	s.runTest("should error when reading non-existent webhook", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			// Fetch webhook.
			_, err := testClients.userClient.GetWebhook(ctx, nonexistentID)
			assert.Error(t, err)
		}
	})
}

func (s *TestSuite) TestWebhooks_Listing() {
	s.runTest("should be able to be read in a list", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			// Create webhooks.
			var expected []*types.Webhook
			for i := 0; i < 5; i++ {
				// Create webhook.
				exampleWebhook := fakes.BuildFakeWebhook()
				exampleWebhookInput := converters.ConvertWebhookToWebhookCreationRequestInput(exampleWebhook)
				createdWebhook, webhookCreationErr := testClients.userClient.CreateWebhook(ctx, exampleWebhookInput)
				requireNotNilAndNoProblems(t, createdWebhook, webhookCreationErr)

				expected = append(expected, createdWebhook)
			}

			// Assert webhook list equality.
			actual, err := testClients.userClient.GetWebhooks(ctx, nil)
			requireNotNilAndNoProblems(t, actual, err)

			assert.GreaterOrEqual(t, len(actual.Data), len(expected))

			// Clean up.
			for _, webhook := range actual.Data {
				assert.NoError(t, testClients.userClient.ArchiveWebhook(ctx, webhook.ID))
			}
		}
	})
}

func (s *TestSuite) TestWebhooks_Archiving_Returns404ForNonexistentWebhook() {
	s.runTest("should error when archiving a non-existent webhook", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			assert.Error(t, testClients.userClient.ArchiveWebhook(ctx, nonexistentID))
		}
	})
}

func (s *TestSuite) TestWebhookTriggerEvents_Archiving_Returns404ForNonexistentWebhook() {
	s.runTest("should error when archiving a non-existent webhook", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			assert.Error(t, testClients.userClient.ArchiveWebhookTriggerEvent(ctx, nonexistentID, nonexistentID))
		}
	})
}
