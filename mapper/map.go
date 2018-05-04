package mapper

import (
	"errors"
)

const loadFactor = 0.75

type Func func(k string) int

type entry struct {
	key   string
	value string
}

// Map is basic hash based data structure
type Map struct {
	size         int
	cap          int
	hashFunction Func
	array        [][]entry
}

type option func(*Map)

func defaultHashFunction(key string) int {
	return len(key)
}

// Size returns the size of the map
func (m *Map) Size() int {
	return m.size
}

// WithCapacity allows developers to write custom capacity
func WithCapacity(cap int) option {
	return func(hashmap *Map) {
		hashmap.cap = cap
	}
}

// WithHashFunction allows developers to write custom hashFunction
func WithHashFunction(hashFunction Func) option {
	return func(hashmap *Map) {
		hashmap.hashFunction = hashFunction
	}
}

// Mapper defines a basic interface for a map
type Mapper interface {
	Put(string, string)
	Get(string) string
}

// Get gets an element to the map
func (m *Map) Get(key string) (string, error) {
	hash := m.hashFunction(key)
	slot := hash % m.cap
	for _, entry := range m.array[slot] {
		if entry.key == key {
			return entry.value, nil
		}
	}
	return "", errors.New("key not found")
}

func (m *Map) rebuild() {

	tempArray := make([][]entry, m.cap*2)
	for i := 0; i < m.cap; i++ {
		tempArray[i] = make([]entry, 0)
	}
	for _, list := range m.array {
		for _, entry := range list {
			hash := m.hashFunction(entry.key)
			slot := hash % (m.cap * 2)
			tempArray[slot] = append(tempArray[slot], entry)
		}
	}
	m.array = tempArray
	m.cap = m.cap * 2
}

// Put adds an element to the map
func (m *Map) Put(key, value string) {
	hash := m.hashFunction(key)
	slot := hash % m.cap
	// check if the key already exists
	for _, entry := range m.array[slot] {
		if entry.key == key {
			return
		}
	}
	m.array[slot] = append(m.array[slot], entry{key, value})
	m.size++
	if float64(m.size)/float64(m.cap) > loadFactor {
		m.rebuild()
	}
}

// NewMap is the constructor for map instance
func NewMap(opts ...option) *Map {
	m := &Map{
		cap:          1,
		hashFunction: defaultHashFunction,
	}

	for _, opt := range opts {
		opt(m)
	}

	m.array = make([][]entry, m.cap)
	for i := 0; i < m.cap; i++ {
		m.array[i] = make([]entry, 0)
	}

	return m
}
