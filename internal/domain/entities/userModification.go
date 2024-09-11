package entities

type UserModification int

const (
	CREATE_MODIFICATION UserModification = iota + 1
	UPDATE_MODIFICATION
	DELETE_MODIFICATION
)
