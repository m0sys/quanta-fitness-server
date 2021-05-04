package user

import "fmt"

const (
	errSlug = "account"
)

var (
	errHashFailed      = fmt.Errorf("%s: failed to hash password", errSlug)
	errInvalidUname    = fmt.Errorf("%s: username must be more than 2 characters long", errSlug)
	errInvalidPassword = fmt.Errorf("%s: password must be between 12 and 128 characters long", errSlug)
	errInvalidEmail    = fmt.Errorf("%s: email is invalid", errSlug)
)
