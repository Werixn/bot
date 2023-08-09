package set

var allSets = []Set{
	{ID: 1, Exercises: allExercises},
}

type Set struct {
	ID        int
	Exercises []Exercise
}

var allExercises = []Exercise{
	{ID: 1, Title: "Bench Press", Description: "Lie on the bench and press."},
	{ID: 2, Title: "Squats", Description: "Take something heavy and squat with it. "},
	{ID: 3, Title: "Biceps", Description: "You know whtat to do"},
	{ID: 4, Title: "Chill", Description: "Actually not and exercise but a very important thing"},
}

type Exercise struct {
	ID          int
	Title       string
	Description string
}
