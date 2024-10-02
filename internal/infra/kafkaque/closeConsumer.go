package kafkaque

func (k *kafkaConsumer) Close() {
	k.logger.Debug(nil, "worker shutdownReceived")
	k.shutdownReceived = true
	<-k.shutdownComplete
}
