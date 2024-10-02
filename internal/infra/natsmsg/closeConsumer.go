package natsmsg

func (nCons *natsConsumer) Close() error {
	return nCons.sub.Unsubscribe()
}
