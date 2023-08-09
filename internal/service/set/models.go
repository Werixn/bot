package set

var allSets = []Set{
	{ID: 1, Exercises: allExercises},
}

type Set struct {
	ID        int
	Exercises []Exercise
}
