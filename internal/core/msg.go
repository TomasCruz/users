package core

// Msg is an interface through which to talk with DB
type Msg interface {
	Close()
	// PublishUserModification(resp entities.UserResp, topic string) error
}
