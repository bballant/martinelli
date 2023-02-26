package htable

import (
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"math"
)

type HVal struct {
	key   string
	value any
}

type HTable struct {
	vals     []*HVal
	capacity int
	len      int
}

func hashIndex(capacity int, key string) int {
	hash := fnv.New32()
	hash.Write([]byte(key))
	hashVal := hash.Sum32()
	return int(hashVal % uint32(capacity-1))
}

func NewHTableWithCap(capacity int) *HTable {
	return &(HTable{make([]*HVal, capacity), capacity, 0})
}

func New() *HTable {
	return NewHTableWithCap(16)
}

func (t *HTable) String() string {
	outStr := ""
	for _, v := range t.vals {
		if v == nil {
			outStr = fmt.Sprintf("%s\nnil", outStr)
		} else {
			outStr = fmt.Sprintf("%s\n%v", outStr, *v)
		}
	}
	return outStr
}

/*
 *  Returns the index for the matching key or the first empty
 *  spot in the table.entities slice
 */
func (t *HTable) find(key string) (int, error) {
	index := hashIndex(t.capacity, key)
	//fmt.Printf("%s: %d\n%v\n", key, index, t)
	counter := index
	var idx int
	for {
		idx = counter % t.capacity
		// prevent infinite loop
		if idx == index && counter > t.capacity {
			return 0, errors.New("Table is full!")
		}
		if t.vals[idx] == nil || t.vals[idx].key == key {
			break
		}
		counter++
	}
	return idx, nil
}

func (t *HTable) setVal(val *HVal) {
	index, err := t.find(val.key)
	if err != nil {
		log.Fatalln(err)
	}
	// it's either nil or ours, so set it
	t.vals[index] = val
	t.len++
}

func (t *HTable) grow() (*HTable, error) {
	cap := t.capacity
	if cap == math.MaxInt {
		return nil, errors.New("Cannot grow table!")
	}

	var newCap int
	if cap > math.MaxInt/2 {
		newCap = math.MaxInt
	} else {
		newCap = cap * 2
	}
	newHT := NewHTableWithCap(newCap)
	for _, v := range t.vals {
		if v == nil {
			continue
		}
		newHT.setVal(v)
	}
	return newHT, nil
}

func (t *HTable) Get(key string) any {
	index, err := t.find(key)
	if err != nil {
		log.Fatalln(err)
	}
	if t.vals[index] == nil {
		return nil
	}
	return t.vals[index].value
}

func (t *HTable) Set(key string, value any) *HTable {
	index, err := t.find(key)
	if err != nil {
		log.Fatalln(err)
	}
	table := *t
	if table.vals[index] != nil {
		table.vals[index].value = value
	} else {
		e := HVal{key, value}
		table.vals[index] = &e
		table.len++
	}
	if table.len > (table.capacity/4)*3 {
		log.Println("Need to grow table")
		newTable, newErr := table.grow()
		log.Println(newTable)
		if newErr != nil {
			log.Fatalln(newErr)
		}
		table = *newTable
	}
	return &table
}

func (t *HTable) Delete(key string) {
	index, err := t.find(key)
	if err != nil {
		log.Fatalln(err)
	}
	if t.vals[index] == nil ||
		t.vals[index].key != key {
		// not found
		return
	}
	t.vals[index] = nil
	t.len--
}
