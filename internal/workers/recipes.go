package workers

import (
	"context"

	"github.com/prixfixeco/api_server/internal/observability"
	"github.com/prixfixeco/api_server/pkg/types"
)

func (w *WritesWorker) createRecipe(ctx context.Context, msg *types.PreWriteMessage) error {
	ctx, span := w.tracer.StartSpan(ctx)
	defer span.End()

	logger := w.logger.WithValue("data_type", msg.DataType)

	recipe, err := w.dataManager.CreateRecipe(ctx, msg.Recipe)
	if err != nil {
		return observability.PrepareError(err, logger, span, "creating recipe")
	}

	if err = w.recipesIndexManager.Index(ctx, recipe.ID, recipe); err != nil {
		return observability.PrepareError(err, logger, span, "indexing the recipe")
	}

	if w.postWritesPublisher != nil {
		dcm := &types.DataChangeMessage{
			DataType:                  msg.DataType,
			MessageType:               "recipeCreated",
			Recipe:                    recipe,
			AttributableToUserID:      msg.AttributableToUserID,
			AttributableToHouseholdID: msg.AttributableToHouseholdID,
		}

		if err = w.postWritesPublisher.Publish(ctx, dcm); err != nil {
			return observability.PrepareError(err, logger, span, "publishing to post-writes topic")
		}
	}

	return nil
}

func (w *UpdatesWorker) updateRecipe(ctx context.Context, msg *types.PreUpdateMessage) error {
	ctx, span := w.tracer.StartSpan(ctx)
	defer span.End()

	logger := w.logger.WithValue("data_type", msg.DataType)

	if err := w.dataManager.UpdateRecipe(ctx, msg.Recipe); err != nil {
		return observability.PrepareError(err, logger, span, "creating recipe")
	}

	if err := w.recipesIndexManager.Index(ctx, msg.Recipe.ID, msg.Recipe); err != nil {
		return observability.PrepareError(err, logger, span, "indexing the recipe")
	}

	if w.postUpdatesPublisher != nil {
		dcm := &types.DataChangeMessage{
			DataType:                  msg.DataType,
			MessageType:               "recipeUpdated",
			Recipe:                    msg.Recipe,
			AttributableToUserID:      msg.AttributableToUserID,
			AttributableToHouseholdID: msg.AttributableToHouseholdID,
		}

		if err := w.postUpdatesPublisher.Publish(ctx, dcm); err != nil {
			return observability.PrepareError(err, logger, span, "publishing data change message")
		}
	}

	return nil
}

func (w *ArchivesWorker) archiveRecipe(ctx context.Context, msg *types.PreArchiveMessage) error {
	ctx, span := w.tracer.StartSpan(ctx)
	defer span.End()

	logger := w.logger.WithValue("data_type", msg.DataType)

	if err := w.dataManager.ArchiveRecipe(ctx, msg.RecipeID, msg.AttributableToHouseholdID); err != nil {
		return observability.PrepareError(err, w.logger, span, "archiving recipe")
	}

	if err := w.recipesIndexManager.Delete(ctx, msg.RecipeID); err != nil {
		return observability.PrepareError(err, w.logger, span, "removing recipe from index")
	}

	if w.postArchivesPublisher != nil {
		dcm := &types.DataChangeMessage{
			DataType:                  msg.DataType,
			MessageType:               "recipeArchived",
			AttributableToUserID:      msg.AttributableToUserID,
			AttributableToHouseholdID: msg.AttributableToHouseholdID,
		}

		if err := w.postArchivesPublisher.Publish(ctx, dcm); err != nil {
			return observability.PrepareError(err, logger, span, "publishing data change message")
		}
	}

	return nil
}
