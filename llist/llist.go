// Copyright 1977 Our Mothers. All Rights Reserved.
package llist

import "fmt"

type LVal struct {
	value any
	next  *LVal
}

type LList struct {
	head *LVal
	len  int
}

func (v *LVal) Next() *LVal {
	return v.next
}

func New() *LList {
	return &LList{}
}

func NewLList() *LList {
	return &LList{}
}

func NewWithValues(values []any) *LList {
	var l LList = LList{}
	for i := len(values) - 1; i >= 0; i-- {
		val := LVal{values[i], l.head}
		l.head = &val
		l.len++
	}
	return &l
}

func (l *LList) String() string {
	val := l.head
	outStr := "("
	delim := ""
	for val != nil {
		outStr = fmt.Sprintf("%s%s%v", outStr, delim, val.value)
		delim = " "
		val = val.next
	}
	outStr = fmt.Sprintf("%s)", outStr)
	return fmt.Sprintln(outStr)
}

func (l *LList) Push(value any) {
	val := LVal{value, l.head}
	l.head = &val
	l.len++
}

func (l *LList) Pop() any {
	head := l.head
	if head == nil {
		return nil
	}
	l.head = head.next
	l.len--

	return head.value
}

func (l *LList) First() any {
	if l.head == nil {
		return nil
	}
	return l.head.value
}

func (l *LList) Rest() *LList {
	if l.head == nil {
		return nil
	}
	return &LList{l.head.next, l.len - 1}
}

func (l *LList) Len() int {
	return l.len
}

func (l *LList) Map(fn func(any) any) *LList {
	newl := LList{}
	currv := l.head
	var prev *LVal
	for currv != nil {
		newv := LVal{fn(currv.value), nil}
		if prev == nil {
			newl.head = &newv
		} else {
			prev.next = &newv
		}
		prev = &newv
		currv = currv.next
	}
	return &newl
}
