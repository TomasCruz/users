package kafkaque

import (
	"encoding/json"

	"github.com/TomasCruz/users/internal/core/entities"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (k kafkaProducer) publishUserCreatedEvent(user entities.User) error {
	serialized, err := json.Marshal(user)
	if err != nil {
		return errors.WithStack(err)
	}

	topic := k.config.CreateUserTopic
	if err = k.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          serialized,
	}, nil); err != nil {
		return errors.Wrap(entities.ErrPublishMsg, err.Error())
	}

	k.kp.Flush(100)
	return nil
}
