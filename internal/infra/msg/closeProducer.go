package msg

func (k kafkaMsgProducer) Close() {
	k.kp.Close()
}
