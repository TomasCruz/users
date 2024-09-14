package msg

import (
	"github.com/TomasCruz/users/internal/domain/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func InitProducer(config configuration.Config, logger ports.Logger) (ports.MsgProducer, error) {
	kp, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaURL,
	})
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	return kafkaMsgProducer{kp: kp, config: config, logger: logger}, nil
}

type kafkaMsgProducer struct {
	kp     *kafka.Producer
	config configuration.Config
	logger ports.Logger
}
