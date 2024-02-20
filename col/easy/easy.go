package easy

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

type Queue []int

func (s *Queue) Enq(x int) {
	*s = append(*s, x)
}

func (s *Queue) Deq() int {
	result := (*s)[0]
	*s = (*s)[1:len(*s)]
	return result
}

func (s *Queue) Peek() int {
	return (*s)[0]
}
