package entities

import "errors"

var (
	ErrBadEmail         = errors.New("bad email")
	ErrExistingEmail    = errors.New("existing email")
	ErrEmptyName        = errors.New("empty name")
	ErrBadHashedPswd    = errors.New("bad hashed password")
	ErrCountryLength    = errors.New("country length can only be 2 or 3")
	ErrInsertUser       = errors.New("insert user failed")
	ErrUpdateUser       = errors.New("update user failed")
	ErrDeleteUser       = errors.New("delete user failed")
	ErrGetUser          = errors.New("get user failed")
	ErrListUser         = errors.New("list user failed")
	ErrNonexistingUser  = errors.New("user not found")
	ErrBadUUID          = errors.New("bad UUID")
	ErrKafkaProduce     = errors.New("kafka produce error")
	ErrPageSize         = errors.New("invalid page size")
	ErrPageNumber       = errors.New("invalid page number")
	ErrPaginationValues = errors.New("only one value for page size or number")
)
