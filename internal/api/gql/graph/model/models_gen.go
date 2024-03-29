// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EditWorkoutPlanTitle struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type Exercise struct {
	ID        string  `json:"id"`
	Wpid      string  `json:"wpid"`
	Name      string  `json:"name"`
	TargetRep int     `json:"targetRep"`
	NumSets   int     `json:"numSets"`
	Weight    float64 `json:"weight"`
	RestDur   float64 `json:"restDur"`
}

type NewExercise struct {
	Wpid      string  `json:"wpid"`
	Name      string  `json:"name"`
	TargetRep int     `json:"targetRep"`
	NumSets   int     `json:"numSets"`
	Weight    float64 `json:"weight"`
	RestDur   float64 `json:"restDur"`
}

type NewWorkoutPlan struct {
	Title string `json:"title"`
}

type WorkoutPlan struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
