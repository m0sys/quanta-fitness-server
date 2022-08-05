package esetstore

import (
	"time"

	"github.com/m0sys/quanta-fitness-server/internal/entity"
)

type store struct {
	esets  map[int64]entity.Eset
	lastID int64
}

func NewMockEsetStore() EsetStore {
	return &store{esets: make(map[int64]entity.Eset)}
}

func (s *store) Save(eset entity.BaseEset) (entity.Eset, error) {
	created := entity.Eset{
		BaseEset:  eset,
		ID:        s.lastID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.esets[s.lastID] = created
	s.lastID += 1

	return created, nil
}

// Precondition: `esid` is a valid (i.e. exists).
func (s *store) Update(esid int64, updates entity.EsetUpdate) error {
	prev := s.esets[esid]
	updated := entity.Eset{
		BaseEset: entity.BaseEset{
			SMetric:  updates.SMetric,
			Username: prev.BaseEset.Username,
			EID:      prev.BaseEset.EID,
		},
		ID:        prev.ID,
		CreatedAt: prev.CreatedAt,
		UpdatedAt: time.Now(),
	}

	s.esets[esid] = updated
	return nil
}

// Precondition: `esid` is a valid (i.e. exists).
func (s *store) Delete(esid int64) error {
	delete(s.esets, esid)
	return nil
}

func (s *store) FindEsetById(esid int64) (entity.Eset, bool, error) {
	found := s.esets[esid]
	if isEmpty(found) {
		return entity.Eset{}, false, nil
	}

	return found, true, nil
}

func isEmpty(found entity.Eset) bool {
	return found == (entity.Eset{})
}

func (s *store) FindAllEsetByEID(eid int64) ([]entity.Eset, error) {
	var entries []entity.Eset

	for k := range s.esets {
		entry := s.esets[k]
		if entry.EID == eid {
			entries = append(entries, entry)
		}
	}
	return entries, nil
}
