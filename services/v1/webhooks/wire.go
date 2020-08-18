package webhooks

import (
	database "gitlab.com/prixfixe/prixfixe/database/v1"
	models "gitlab.com/prixfixe/prixfixe/models/v1"

	"github.com/google/wire"
)

var (
	// Providers is our collection of what we provide to other services.
	Providers = wire.NewSet(
		ProvideWebhooksService,
		ProvideWebhookDataManager,
		ProvideWebhookDataServer,
	)
)

// ProvideWebhookDataManager is an arbitrary function for dependency injection's sake.
func ProvideWebhookDataManager(db database.DataManager) models.WebhookDataManager {
	return db
}

// ProvideWebhookDataServer is an arbitrary function for dependency injection's sake.
func ProvideWebhookDataServer(s *Service) models.WebhookDataServer {
	return s
}
