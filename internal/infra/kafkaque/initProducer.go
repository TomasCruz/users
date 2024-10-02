package kafkaque

import (
	"github.com/TomasCruz/users/internal/core/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func InitProducer(config configuration.Config, logger ports.Logger) (ports.QueueProducer, error) {
	kp, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaURL,
	})
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	return kafkaProducer{kp: kp, config: config, logger: logger}, nil
}

type kafkaProducer struct {
	kp     *kafka.Producer
	config configuration.Config
	logger ports.Logger
}
