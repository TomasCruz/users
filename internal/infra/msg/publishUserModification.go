package msg

import (
	"encoding/json"

	"github.com/TomasCruz/users/internal/domain/entities"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (k kafkaMsgProducer) PublishUserModification(user entities.User, modificationType entities.UserModification) error {
	var topic string
	switch modificationType {
	case entities.CREATE_MODIFICATION:
		topic = k.config.CreateUserTopic
	case entities.UPDATE_MODIFICATION:
		topic = k.config.UpdateUserTopic
	case entities.DELETE_MODIFICATION:
		topic = k.config.DeleteUserTopic
	default:
		return errors.WithStack(entities.ErrBadMsgType)
	}

	serialized, err := json.Marshal(user)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = k.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          serialized,
	}, nil); err != nil {
		return errors.Wrap(entities.ErrPublishMsg, err.Error())
	}

	k.kp.Flush(100)

	return nil
}
