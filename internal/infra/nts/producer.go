package nts

import (
	"time"

	"github.com/pkg/errors"
)

func (nProd natsProducer) UserCreatedRequest(data []byte, timeout time.Duration) (string, error) {
	msg, err := nProd.nc.Request(nProd.config.NatsSubjectCreateUser, data, timeout)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(msg.Data), nil
}

func (nProd natsProducer) Drain() error {
	return nProd.nc.Drain()
}
