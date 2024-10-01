package ports

type NatsConsumer interface {
	SubUserCreated() error
	Unsubscribe() error
}
