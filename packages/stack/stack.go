package stack

type Stack struct {
	items []uint8
}

func (s *Stack) Push(item uint8) int {
	if len(s.items) >= 256 {
		return -1
	}
	s.items = append(s.items, item)
	return 0
}

func (s *Stack) Pop() uint8 {
	if len(s.items) == 0 {
		return 0
	}
	lastItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lastItem
}
