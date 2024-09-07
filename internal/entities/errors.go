package entities

import "errors"

var (
	ErrBadEmail           = errors.New("bad email")
	ErrExistingEmail      = errors.New("existing email")
	ErrEmptyName          = errors.New("empty name")
	ErrBadHashedPswd      = errors.New("bad hashed password")
	ErrCountryLength      = errors.New("country length can only be 2 or 3")
	ErrDatabaseError      = errors.New("database error")
	ErrInsertUser         = errors.New("insert user failed")
	ErrUpdateUser         = errors.New("update user failed")
	ErrDeleteUser         = errors.New("delete user failed")
	ErrGetUser            = errors.New("get user failed")
	ErrListUser           = errors.New("list user failed")
	ErrNonexistingUser    = errors.New("user not found")
	ErrBadUUID            = errors.New("bad UUID")
	ErrKafkaProduce       = errors.New("kafka produce error")
	ErrCountFilteredQuery = errors.New("count filtered query")
)
