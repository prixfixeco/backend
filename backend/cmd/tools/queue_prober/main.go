package main

import (
	"context"
	"log"
	"os"

	"github.com/dinnerdonebetter/backend/internal/config"
	msgconfig "github.com/dinnerdonebetter/backend/internal/messagequeue/config"
	"github.com/dinnerdonebetter/backend/internal/messagequeue/pubsub"
	"github.com/dinnerdonebetter/backend/internal/messagequeue/redis"
	"github.com/dinnerdonebetter/backend/internal/messagequeue/sqs"
	"github.com/dinnerdonebetter/backend/internal/observability"
	"github.com/dinnerdonebetter/backend/internal/observability/logging"
	loggingcfg "github.com/dinnerdonebetter/backend/internal/observability/logging/config"
	"github.com/dinnerdonebetter/backend/internal/observability/tracing"

	"github.com/spf13/pflag"
)

var (
	queueNameFlag = pflag.StringP("queue", "q", "", "the queue to write to")
	dataFlag      = pflag.StringP("data", "d", "", "the data to write to")
)

func main() {
	ctx := context.Background()
	pflag.Parse()

	if *queueNameFlag == "" || *dataFlag == "" {
		pflag.Usage()
		os.Exit(1)
	}

	logger := loggingcfg.ProvideLogger(&loggingcfg.Config{
		Level:    logging.DebugLevel,
		Provider: loggingcfg.ProviderSlog,
	})
	tracerProvider := tracing.NewNoopTracerProvider()

	eventConfig := &msgconfig.Config{
		Publisher: msgconfig.MessageQueueConfig{
			Provider: msgconfig.ProviderPubSub,
			PubSub: pubsub.Config{
				ProjectID: "dinner-done-better-dev",
			},
			SQS:   sqs.Config{},
			Redis: redis.Config{},
		},
	}

	if err := config.ApplyEnvironmentVariables(eventConfig); err != nil {
		log.Fatalln(err)
	}

	// setup baseline messaging providers

	if err := doTheThing(ctx, logger, tracerProvider, eventConfig); err != nil {
		observability.AcknowledgeError(err, logger, nil, "doing the thing")
	}
}

func doTheThing(
	ctx context.Context,
	logger logging.Logger,
	tracerProvider tracing.TracerProvider,
	eventConfig *msgconfig.Config,
) error {
	publisherProvider, err := msgconfig.ProvidePublisherProvider(ctx, logger, tracerProvider, eventConfig)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, nil, "configuring queue manager")
	}
	defer publisherProvider.Close()

	publisher, err := publisherProvider.ProvidePublisher(*queueNameFlag)
	if err != nil {
		return observability.PrepareAndLogError(err, logger, nil, "configuring publisher")
	}
	defer publisher.Stop()

	if err = publisher.Publish(ctx, *dataFlag); err != nil {
		return observability.PrepareAndLogError(err, logger, nil, "publishing data")
	}

	return nil
}
