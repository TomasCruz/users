package msg

func (k *kafkaMsgConsumer) Close() {
	k.logger.Debug(nil, "worker shutdownReceived")
	k.shutdownReceived = true
	<-k.shutdownComplete
}
