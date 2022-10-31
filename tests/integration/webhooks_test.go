package integration

import (
	"github.com/prixfixeco/api_server/pkg/types/converters"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/prixfixeco/api_server/internal/observability/tracing"
	"github.com/prixfixeco/api_server/pkg/types"
	"github.com/prixfixeco/api_server/pkg/types/fakes"
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
	s.runForCookieClient("should be creatable and readable and deletable", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			// Create webhook.
			exampleWebhook := fakes.BuildFakeWebhook()
			exampleWebhookInput := converters.ConvertWebhookToWebhookCreationRequestInput(exampleWebhook)

			createdWebhook, err := testClients.user.CreateWebhook(ctx, exampleWebhookInput)
			require.NoError(t, err)
			t.Logf("created webhook %s", createdWebhook.ID)

			createdWebhook, err = testClients.user.GetWebhook(ctx, createdWebhook.ID)
			requireNotNilAndNoProblems(t, createdWebhook, err)

			// assert webhook equality
			checkWebhookEquality(t, exampleWebhook, createdWebhook)

			actual, err := testClients.user.GetWebhook(ctx, createdWebhook.ID)
			requireNotNilAndNoProblems(t, actual, err)
			checkWebhookEquality(t, exampleWebhook, actual)

			// Clean up.
			assert.NoError(t, testClients.user.ArchiveWebhook(ctx, createdWebhook.ID))
		}
	})

	s.runForPASETOClient("should be creatable and readable and deletable", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			// Create webhook.
			exampleWebhook := fakes.BuildFakeWebhook()
			exampleWebhookInput := converters.ConvertWebhookToWebhookCreationRequestInput(exampleWebhook)
			createdWebhook, err := testClients.user.CreateWebhook(ctx, exampleWebhookInput)
			require.NoError(t, err)

			// assert webhook equality
			checkWebhookEquality(t, exampleWebhook, createdWebhook)

			actual, err := testClients.user.GetWebhook(ctx, createdWebhook.ID)
			requireNotNilAndNoProblems(t, actual, err)
			checkWebhookEquality(t, exampleWebhook, actual)

			// Clean up.
			assert.NoError(t, testClients.user.ArchiveWebhook(ctx, createdWebhook.ID))
		}
	})
}

func (s *TestSuite) TestWebhooks_Reading_Returns404ForNonexistentWebhook() {
	s.runForEachClientExcept("should fail to read non-existent webhook", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			// Fetch webhook.
			_, err := testClients.user.GetWebhook(ctx, nonexistentID)
			assert.Error(t, err)
		}
	})
}

func (s *TestSuite) TestWebhooks_Listing() {
	s.runForCookieClient("should be able to be read in a list", func(testClients *testClientWrapper) func() {
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
				createdWebhook, webhookCreationErr := testClients.user.CreateWebhook(ctx, exampleWebhookInput)
				requireNotNilAndNoProblems(t, createdWebhook, webhookCreationErr)

				expected = append(expected, createdWebhook)
			}

			// Assert webhook list equality.
			actual, err := testClients.user.GetWebhooks(ctx, nil)
			requireNotNilAndNoProblems(t, actual, err)

			assert.GreaterOrEqual(t, len(actual.Webhooks), len(expected))

			// Clean up.
			//for _, webhook := range actual.Webhooks {
			//	assert.NoError(t, testClients.user.ArchiveWebhook(ctx, webhook.ID))
			//}
		}
	})

	s.runForPASETOClient("should be able to be read in a list", func(testClients *testClientWrapper) func() {
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
				createdWebhook, err := testClients.user.CreateWebhook(ctx, exampleWebhookInput)
				require.NoError(t, err)

				requireNotNilAndNoProblems(t, createdWebhook, err)

				expected = append(expected, createdWebhook)
			}

			// Assert webhook list equality.
			actual, err := testClients.user.GetWebhooks(ctx, nil)
			requireNotNilAndNoProblems(t, actual, err)
			assert.True(t, len(expected) <= len(actual.Webhooks))

			// Clean up.
			for _, webhook := range actual.Webhooks {
				assert.NoError(t, testClients.user.ArchiveWebhook(ctx, webhook.ID))
			}
		}
	})
}

func (s *TestSuite) TestWebhooks_Archiving_Returns404ForNonexistentWebhook() {
	s.runForEachClientExcept("should fail to archive a non-existent webhook", func(testClients *testClientWrapper) func() {
		return func() {
			t := s.T()

			ctx, span := tracing.StartCustomSpan(s.ctx, t.Name())
			defer span.End()

			assert.Error(t, testClients.user.ArchiveWebhook(ctx, nonexistentID))
		}
	})
}
