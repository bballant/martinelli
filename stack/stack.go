// Copyright 1977 Our Mothers. All Rights Reserved.
package stack

import (
	"github.com/bballant/llist"
)

type Stack struct {
	l *llist.LList
}

func New() *Stack {
	return &Stack{llist.New()}
}

func (s *Stack) String() string {
	if s.l == nil {
		return "nil"
	}
	return s.l.String()
}

func (s *Stack) Push(value any) {
	s.l.Push(value)
}

func (s *Stack) Pop() any {
	return s.l.Pop()
}

func (s *Stack) Peek() any {
	return s.l.First()
}

func (s *Stack) Len() int {
	return s.l.Len()
}
