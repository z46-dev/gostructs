package lists

type SkipListNode[T any] struct {
	Value      T
	Next, Down *SkipListNode[T]
}

type SkipList[T any] struct {
	Head                  *SkipListNode[T]
	Size, level, maxLevel int
	comparator            func(a, b T) bool
}

func NewSkipList[T any](maxLevel int, comparator func(a, b T) bool) *SkipList[T] {
	return &SkipList[T]{
		Head:       nil,
		Size:       0,
		level:      0,
		maxLevel:   maxLevel,
		comparator: comparator,
	}
}

func (s *SkipList[T]) Add(value T) {
	s.Size++

	var node *SkipListNode[T] = &SkipListNode[T]{Value: value}
	if s.Head == nil {
		s.Head = node
		return
	}

	var current *SkipListNode[T] = s.Head
	var prev []*SkipListNode[T] = make([]*SkipListNode[T], s.maxLevel)

	for i := s.level - 1; i >= 0; i-- {
		for current != nil && s.comparator(current.Value, value) {
			prev[i] = current
			current = current.Next
		}
	}

	if current == nil || !s.comparator(current.Value, value) {
		node.Next = current
		if s.level < s.maxLevel {
			s.level++
			node.Down = nil
			s.Head = node
		} else {
			for i := s.level - 1; i >= 0; i-- {
				prev[i].Next = node
				node.Down = prev[i]
				node = prev[i]
			}
		}
	} else {
		node.Next = current.Next
		current.Next = node
	}
}

func (s *SkipList[T]) Remove(value T) bool {
	if s.Head == nil {
		return false
	}

	var current *SkipListNode[T] = s.Head
	var prev []*SkipListNode[T] = make([]*SkipListNode[T], s.maxLevel)

	for i := s.level - 1; i >= 0; i-- {
		for current != nil && s.comparator(current.Value, value) {
			prev[i] = current
			current = current.Next
		}
	}

	if current == nil || !s.comparator(current.Value, value) {
		return false
	}

	s.Size--

	for i := s.level - 1; i >= 0; i-- {
		if prev[i] != nil {
			prev[i].Next = current.Next
		}
	}

	return true
}
