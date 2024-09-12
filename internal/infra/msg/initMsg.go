package msg

import (
	"github.com/TomasCruz/users/internal/domain/ports"
	"github.com/TomasCruz/users/internal/infra/configuration"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func InitMsg(config configuration.Config, logger ports.Logger) (ports.Msg, error) {
	var kp *kafka.Producer
	var err error
	if kp, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaURL,
	}); err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	return kafkaMsg{kp: kp, config: config, logger: logger}, nil
}

type kafkaMsg struct {
	kp     *kafka.Producer
	config configuration.Config
	logger ports.Logger
}
