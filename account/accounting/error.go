package accounting

import "fmt"

const (
	errSlug = "accounting"
)

var (
	errInternal           = fmt.Errorf("%s: internal error", errSlug)
	ErrUnameAlreadyExists = fmt.Errorf("%s: username already exists", errSlug)
	ErrEmailAlreadyExists = fmt.Errorf("%s: email already exists", errSlug)
	ErrUnameNotFound      = fmt.Errorf("%s: username not found", errSlug)
	ErrIncorrectPassword  = fmt.Errorf("%s: incorrect password", errSlug)
)
