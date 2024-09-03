package msg

import (
	"github.com/TomasCruz/users/internal/configuration"
	"github.com/TomasCruz/users/internal/core"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func InitMsg(config configuration.Config) (core.Msg, error) {
	var kp *kafka.Producer
	var err error
	if kp, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBroker,
	}); err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	return kafkaMsg{kp: kp}, nil
}

type kafkaMsg struct {
	kp *kafka.Producer
}
