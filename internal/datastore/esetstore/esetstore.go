package esetstore

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type EsetStore interface {
	Save(eset entity.BaseEset) (entity.Eset, error)
	Update(esid string, updates entity.EsetUpdate) error
	Delete(esid string) error
	FindEsetById(esid string) (entity.Eset, bool, error)
	FindAllEsetByEID(eid string) ([]entity.Eset, error)
}
