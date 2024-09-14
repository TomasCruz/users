package msg

import (
	"os"
	"os/signal"
)

func (k *kafkaMsgConsumer) Close() {
	k.logger.Info(nil, "before kafkaMsgConsumer.Close()")
	signal.Notify(k.shutdownReceived, os.Interrupt)
	k.shutdownReceived <- os.Interrupt
	k.logger.Info(nil, "after kafkaMsgConsumer.Close()")
}
