package ports

type MsgConsumer interface {
	Close() error
}
