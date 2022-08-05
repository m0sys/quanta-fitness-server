package esetstore

import (
	"github.com/m0sys/quanta-fitness-server/internal/entity"
)

type EsetStore interface {
	Save(eset entity.BaseEset) (entity.Eset, error)
	Update(esid int64, updates entity.EsetUpdate) error
	Delete(esid int64) error
	FindEsetById(esid int64) (entity.Eset, bool, error)
	FindAllEsetByEID(eid int64) ([]entity.Eset, error)
}
