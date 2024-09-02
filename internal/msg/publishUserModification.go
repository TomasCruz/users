package msg

// func (k kafkaMsg) PublishUserModification(resp entities.UserResp, topic string) error {
// 	serialized, err := json.Marshal(resp)
// 	if err != nil {
// 		return err
// 	}

// 	if err = m.kp.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Value:          serialized,
// 	}, nil); err != nil {
// 		return errors.Wrap(entities.ErrKafkaProduce, err.Error())
// 	}

// 	m.kp.Flush(100)

// 	return nil
// }
