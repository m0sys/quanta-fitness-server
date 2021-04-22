package record

import a "github.com/mhd53/quanta-fitness-server/athlete"

type AthleteGateway interface {
	FindAtheleteByUname(uname string) (a.Athlete, bool)
}
