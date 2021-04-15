package record

/*
Here's how you will succesfull model the critical rules of the system.
Model the whole workout experience without any concern about this being a
Software System!

This whole part of the system should be able to enable you to go through a
whole workout session.

NOTE: The key distinction between `Entity` and `UseCase` is that the former
	  is purely business driven while the latter is both business driven
	  as well as application driven. Thus, this layer doesn't reflect that
	  this is a software system. It only reflects the core values of the
	  business - and the core functionality of the product that it wishes
	  to produce for the benefit of the Customer. Hence, these said
	  functionalities must be the same as if someone were to do them manually.
*/

type ExerciseManger interface {
	Start(name string, targetRep int, weight, distance, restTime float64)
	AddNewSet(set Set)
	MoveToNextSet() bool // use closure to keep track of an internal counter?
	Complete() Exercise
}
