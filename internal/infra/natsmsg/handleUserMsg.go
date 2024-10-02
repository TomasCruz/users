package natsmsg

import "github.com/nats-io/nats.go"

func (nCons *natsConsumer) userCreatedHandler(msg *nats.Msg) {
	msg.Respond([]byte("hello, " + string(msg.Data)))
}
