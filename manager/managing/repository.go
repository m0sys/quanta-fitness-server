package manager

import "github.com/mhd53/quanta-fitness-server/manager/athlete"

// Repository repo for persisting all Athlete related data.
type Repository interface {
	FindAllWeightRecords(aid string) ([]WeightRecordRes, error)
	StoreWeightRecord(aid string, record athlete.WeightRecord) error
}
