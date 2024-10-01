package nts

import "github.com/nats-io/nats.go"

func (nCons *natsConsumer) SubUserCreated() error {
	var err error
	nCons.sub, err = nCons.nc.Subscribe(nCons.config.NatsSubjectCreateUser, nCons.userCreatedHandler)
	if err != nil {
		return err
	}

	return nil
}

func (nCons *natsConsumer) Unsubscribe() error {
	return nCons.sub.Unsubscribe()
}

func (nCons *natsConsumer) userCreatedHandler(msg *nats.Msg) {
	msg.Respond([]byte("hello, " + string(msg.Data)))
}
