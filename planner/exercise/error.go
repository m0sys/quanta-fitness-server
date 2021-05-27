package exercise

import "fmt"

const (
	errSlug = "exercise"
)

var (
	ErrInvalidName      = fmt.Errorf("%s: name must be less than 76 characters", errSlug)
	ErrInvalidTargetRep = fmt.Errorf("%s: target rep must be a positive number", errSlug)
	ErrInvalidNumSets   = fmt.Errorf("%s: num sets must be a positive number", errSlug)
	ErrInvalidWeight    = fmt.Errorf("%s: weight must be a positive number", errSlug)
	ErrInvalidRestDur   = fmt.Errorf("%s: rest duration must be a positive number", errSlug)
	ErrInvalidPos       = fmt.Errorf("%s: position must be a positive number", errSlug)
)
