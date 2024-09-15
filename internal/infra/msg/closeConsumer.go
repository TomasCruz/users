package msg

func (k *kafkaMsgConsumer) Close() {
	k.shutdownReceived = true
	<-k.shutdownComplete
}
