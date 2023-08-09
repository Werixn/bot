package set

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) All() {
	return allSets
}
