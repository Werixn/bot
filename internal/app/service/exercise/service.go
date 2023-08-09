package exercise

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Exercise {
	return allExercises
}

func (s *Service) Get(idx int) (*Exercise, error) {
	return &allExercises[idx], nil
}
