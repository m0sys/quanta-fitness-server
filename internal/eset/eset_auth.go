package eset

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type EsetAuth interface {
	AuthorizeExerciseAccess(uname string, eid int64) (bool, error)
	AuthorizeEsetAccess(uname string, esid int64) (bool, error)
}
