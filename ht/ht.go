package ds

import (
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"math"
)

type HTEntry struct {
	key   string
	value interface{}
}

type HT struct {
	entries  []*HTEntry
	capacity uint64
	size   uint64
}

func hashKey(key []byte) uint64 {
	hash := fnv.New64()
	hash.Write(key)
	return hash.Sum64()
}

func keyIndex(table *HT, key string) uint64 {
	hash := hashKey([]byte(key))
	return hash % ((*table).capacity - 1)
}

// func-away some boilerplate to prevent long lines
func found(table *HT, index uint64, key string) bool {
	return (*(*table).entries[index]).key == key
}

/*
 *  Returns the index for the matching key or the first empty
 *  spot in the table.entities slice
 */
func findIndex(table *HT, key string) (uint64, error) {
	index := keyIndex(table, key)
	counter := index
	var idx uint64
	for {
		idx = counter % (*table).capacity;
		// prevent infinite loop
		if idx == index && counter > (*table).capacity {
			return 0, errors.New("Table is full!")
		}
		if (*table).entries[idx] == nil || found(table, idx, key) {
			break
		}
		counter++
	}
	return idx, nil
}

func setEntry(table *HT, entry *HTEntry) {
	index, err := findIndex(table, (*entry).key)
	if err != nil {
		log.Fatalln(err)
	}
	// it's either nil or ours, so set it
	(*table).entries[index] = entry
	(*table).size++
}

func grow(table *HT) error {
	cap := (*table).capacity
	if cap > math.MaxInt64 / 2 {
		return errors.New("Cannot grow table!")
	}
	newHT := HT{make([]*HTEntry, cap * 2), cap * 2, 0}
	for _, e := range (*table).entries {
		if (e == nil) {
			continue
		}
		setEntry(&newHT, e)
	}
	*table = newHT
	return nil
}


func HTCreate() *HT {
	return &(HT{make([]*HTEntry, 16), 16, 0})
}

func HTPrint(table *HT) {
	fmt.Printf("Table w/ capacity %d, size %d\n", (*table).capacity, (*table).size)
	for _, v := range (*table).entries {
		if v == nil {
			fmt.Println("nil")
		} else {
			fmt.Println(*v)
		}
	}
}

func HTGet(table *HT, key string) interface{} {
	index, err := findIndex(table, key)
	if err != nil {
		log.Fatalln(err)
	}
	if (*table).entries[index] == nil {
		return nil
	}
	return (*(*table).entries[index]).value
}

func HTSet(table *HT, key string, value interface{}) {
	index, err := findIndex(table, key)
	if err != nil {
		log.Fatalln(err)
	}
	if (*table).entries[index] != nil {
		(*(*table).entries[index]).value = value
	} else  {
		e := HTEntry{key, value}
		(*table).entries[index] = &e
		(*table).size++
	}
	if (*table).size > ((*table).capacity / 4) * 3 {
		log.Println("Need to grow table")
		grow(table)
	}
}

func HTDelete(table *HT, key string) {
	index, err := findIndex(table, key)
	if err != nil {
		log.Fatalln(err)
	}
	if (*table).entries[index] == nil ||  !found(table, index, key) {
		// not found
		return
	}
	(*table).entries[index] = nil
	(*table).size--
}

func Run() {
	table := HTCreate()
	HTSet(table, "cool", "dude")
	HTSet(table, "rad", "man")
	HTDelete(table, "rad")
	HTSet(table, "rad", "dude")
	HTSet(table, "rad", "skilla")
	HTPrint(table)
	fmt.Println(hashKey([]byte("Cool")))
	fmt.Println(HTGet(table, "cool"))
}
