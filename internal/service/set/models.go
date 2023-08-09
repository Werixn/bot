package set

var allSets = []Set{
	{ID: 1, Exercises: *exercise.allExercises},
}

type Set struct {
	ID        int
	Exercises []Exercise
}
