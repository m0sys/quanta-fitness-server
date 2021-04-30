package planner

import "github.com/mhd53/quanta-fitness-server/athlete"

type Workout struct {
	uuid      string
	athlete   athlete.Athlete
	title     string
	exercises []Exercise
}
