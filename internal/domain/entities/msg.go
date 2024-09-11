package entities

// Msg is an interface through which to talk with DB
type Msg interface {
	Close()
	PublishUserModification(user User, modificationType UserModification) error
}
