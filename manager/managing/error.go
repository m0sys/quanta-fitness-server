package managing

import "fmt"

const (
	errSlug = "managing"
)

var (
	errInternal        = fmt.Errorf("%s: internal error", errSlug)
	ErrAthleteNotFound = fmt.Errorf("%s: Athlete not found", errSlug)
)
