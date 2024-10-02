package natsmsg

import "github.com/pkg/errors"

func (nCons *natsConsumer) subUserCreated() error {
	var err error
	nCons.sub, err = nCons.nc.Subscribe(nCons.config.NatsSubjectCreateUser, nCons.userCreatedHandler)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
