package simple

import "fmt"

type sequence struct {
	first stateInterface
	second stateInterface
}

func Sequence(first stateInterface, second stateInterface) *sequence {
	return &sequence{first: first, second: second}
}

func (s *sequence) To_s() string {
	return s.first.To_s() + "; " + s.second.To_s()
}

func (s *sequence) Reduce(env environment) stateInterface {
	switch s.first.(type) {
	case *doNothing:
		return s.second
	default:
		reduced_first := s.first.Reduce(env)
		return Sequence(reduced_first, s.second)
	}
}

func (s *sequence) IsReducible() bool {
	return true
}

func (s *sequence) Expression(env environment) {
	fmt.Printf(s.To_s() + " %v\n", env)
}
