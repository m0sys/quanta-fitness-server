package managerrepo

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/manager"
)

type repo struct {
	aths       map[string]inRepoAthlete
	records    map[int]inRepoRecord
	lastRecord int
}

// NewMangerRepo create a new in memory Manager Repository.
func NewMangerRepo() manager.Repository {
	return &repo{
		aths:       make(map[string]inRepoAthlete),
		records:    make(map[int]inRepoRecord),
		lastRecord: 0,
	}
}
func (r *repo) FindAllWeightRecords(aid string) ([]manager.WeightRecordRes, error) {
	var records []manager.WeightRecordRes
	for _, val := range r.records {
		if val.AthleteID == aid {
			record := manager.WeightRecordRes{
				Amount: val.Weight,
				Date:   val.Date,
			}
			records = append(records, record)
		}
	}

	return records, nil
}
func (r *repo) StoreWeightRecord(aid string, record athlete.WeightRecord) error {
	r.records[r.lastRecord] = inRepoRecord{
		AthleteID: aid,
		Weight:    record.Amount(),
		Date:      record.Date(),
	}
	r.lastRecord++
	return nil
}

type inRepoAthlete struct {
	AthleteID string
	Height    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type inRepoRecord struct {
	AthleteID string
	Weight    float64
	Date      time.Time
}
