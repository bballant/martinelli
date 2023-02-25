package ll

import "fmt"

type LLVal struct {
	value interface{}
	next  *LLVal
}

type LL struct {
	head *LLVal
	size uint64
}

func (l LL) String() string {
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

func (l LL) Cons(value interface{}) LL {
	val := LLVal{value, l.head}
	return LL{&val, l.size + 1}
}

func (l LL) First() interface{} {
	head := l.head
	if head == nil {
		return nil
	}
	return (*l.head).value
}

func (l LL) Rest() LL {
	if l.head == nil {
		return LL{}
	}
	return LL{(*l.head).next, l.size - 1}
}

func (l LL) Size() uint64 {
	return l.size
}

func (l LL) Map(fn func(interface{}) interface{}) LL {
	newl := LL{}
	currv := l.head
	var prev *LLVal
	for currv != nil {
		newv := LLVal{fn(currv.value), nil}
		if prev == nil {
			newl.head = &newv
		} else {
			prev.next = &newv
		}
		prev = &newv
		currv = currv.next
	}
	return newl
}

func LLCreateWith(values []interface{}) LL {
	var l LL = LL{}
	for i := len(values) - 1; i >= 0; i-- {
		val := LLVal{values[i], l.head}
		l.head = &val
		l.size++
	}
	return l
}
