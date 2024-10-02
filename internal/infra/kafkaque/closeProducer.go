package kafkaque

func (k kafkaProducer) Close() {
	k.kp.Close()
}
