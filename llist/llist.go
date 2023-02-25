package llist

import "fmt"

type LVal struct {
	value interface{}
	next  *LVal
}

type LList struct {
	head *LVal
	size uint64
}

func (l *LList) String() string {
	val := l.head
	outStr := "("
	delim := ""
	for val != nil {
		outStr = fmt.Sprintf("%s%s%v", outStr, delim, (*val).value)
		delim = " "
		val = (*val).next
	}
	outStr = fmt.Sprintf("%s)", outStr)
	return fmt.Sprintln(outStr)
}

func (l *LList) Cons(value interface{}) *LList {
	val := LVal{value, l.head}
	return &LList{&val, l.size + 1}
}

func (l *LList) First() interface{} {
	head := l.head
	if head == nil {
		return nil
	}
	return (*l.head).value
}

func (l *LList) Rest() *LList {
	if l.head == nil {
		return nil
	}
	return &LList{(*l.head).next, l.size - 1}
}

func (l *LList) Size() uint64 {
	return l.size
}

func (l *LList) Map(fn func(interface{}) interface{}) *LList {
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

func LListCreateWith(values []interface{}) *LList {
	var l LList = LList{}
	for i := len(values) - 1; i >= 0; i-- {
		val := LVal{values[i], l.head}
		l.head = &val
		l.size++
	}
	return &l
}
