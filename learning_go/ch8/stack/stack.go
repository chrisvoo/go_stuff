package stack

type Stack[T comparable] struct {
	Vals []T
}

func (s *Stack[T]) Push(val T) {
	s.Vals = append(s.Vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.Vals[len(s.Vals)-1]
	s.Vals = s.Vals[:len(s.Vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.Vals {
		if v == val {
			return true
		}
	}
	return false
}