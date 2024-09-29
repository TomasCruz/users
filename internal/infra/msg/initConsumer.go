package msg

import (
	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/core/service/worker"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func InitConsumer(config configuration.Config, svc worker.WorkerUserService, logger ports.Logger) (ports.MsgConsumer, error) {
	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaURL,
		"group.id":          "group.id",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	consumer := kafkaMsgConsumer{kc: kc, config: config, svc: svc, logger: logger, shutdownReceived: false, shutdownComplete: make(chan struct{}, 1)}

	// Subscribe to the Kafka topic
	err = consumer.kc.SubscribeTopics([]string{config.CreateUserTopic}, nil)
	if err != nil {
		err = errors.Wrap(errors.New("Failed to subscribe to topic"), err.Error())
		return nil, err
	}

	// start Kafka consumer
	go consumer.consume()

	return &consumer, nil
}

type kafkaMsgConsumer struct {
	kc               *kafka.Consumer
	config           configuration.Config
	svc              worker.WorkerUserService
	logger           ports.Logger
	shutdownReceived bool
	shutdownComplete chan struct{}
}
