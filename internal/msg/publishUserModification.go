package msg

import (
	"encoding/json"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (k kafkaMsg) PublishUserModification(resp entities.UserResp, topic string) error {
	serialized, err := json.Marshal(resp)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = k.kp.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          serialized,
	}, nil); err != nil {
		return errors.Wrap(entities.ErrKafkaProduce, err.Error())
	}

	k.kp.Flush(100)

	return nil
}
