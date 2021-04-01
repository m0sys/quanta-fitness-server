package esetstore

import (
	"log"
	"strconv"
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
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
		ID:        string(s.lastID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.esets[s.lastID] = created
	s.lastID += 1

	return created, nil
}

// Precondition: `esid` is a valid (i.e. exists).
func (s *store) Update(esid string, updates entity.EsetUpdate) error {
	prev := s.esets[parseInt64(esid)]
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

	s.esets[parseInt64(esid)] = updated
	return nil
}

// Precondition: `esid` is a valid (i.e. exists).
func (s *store) Delete(esid string) error {
	delete(s.esets, parseInt64(esid))
	return nil
}

func (s *store) FindEsetById(esid string) (entity.Eset, bool, error) {
	found := s.esets[parseInt64(esid)]
	if isEmpty(found) {
		return entity.Eset{}, false, nil
	}

	return found, true, nil
}

func isEmpty(found entity.Eset) bool {
	return found == (entity.Eset{})
}

func (s *store) FindAllEsetByEID(eid string) ([]entity.Eset, error) {
	var entries []entity.Eset

	for k := range s.esets {
		entry := s.esets[k]
		if entry.EID == eid {
			entries = append(entries, entry)
		}
	}
	return entries, nil
}

func parseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
