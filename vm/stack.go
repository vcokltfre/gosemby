package vm

type Stack struct {
	data []int
	size int
	indx int
}

func (s *Stack) Init(size int) {
	s.data = make([]int, size)
	s.size = size
	s.indx = 0
}

func (s *Stack) Pop() int {
	if s.indx == 0 {
		panic("stack underflow")
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	s.indx -= 1

	return value
}

func (s *Stack) Push(value int) {
	if s.indx == s.size {
		panic("stack overflow")
	}

	s.data = append(s.data, value)
	s.indx += 1
}

func (s *Stack) Dup() {
	s.Push(s.data[len(s.data)-1])
}
