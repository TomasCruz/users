package msg

import (
	"os"

	"github.com/TomasCruz/users/internal/domain/core"
	"github.com/TomasCruz/users/internal/domain/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func InitConsumer(config configuration.Config, cr core.Core, logger ports.Logger) (ports.MsgConsumer, error) {
	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaURL,
		"group.id":          "group.id",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	consumer := kafkaMsgConsumer{kc: kc, config: config, cr: cr, logger: logger, shutdownReceived: make(chan os.Signal, 1)}

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
	cr               core.Core
	logger           ports.Logger
	shutdownReceived chan os.Signal
}
