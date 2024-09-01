package msg

func (m kafkaMsg) Close() {
	m.kp.Close()
}
