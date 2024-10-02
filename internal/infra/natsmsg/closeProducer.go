package natsmsg

func (nProd natsProducer) Close() error {
	return nProd.nc.Drain()
}
