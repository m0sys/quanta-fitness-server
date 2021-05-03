package training

import "fmt"

const (
	errSlur = "translator"
)

var (
	errInternal = fmt.Errorf("%s: internal error", errSlur)
)
