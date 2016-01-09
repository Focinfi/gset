package gset

import (
	"sync"
)

var mux sync.RWMutex

type Elementer interface {
	Element() interface{}
}

type Set struct {
	elements map[Elementer]bool
}

func NewSet(elements ...Elementer) *Set {
	set := &Set{make(map[Elementer]bool)}
	set.Add(elements...)
	return set
}

func (s *Set) Length() int {
	mux.RLock()
	defer mux.RUnlock()
	return len(s.elements)
}

func (s *Set) Clear() {
	mux.Lock()
	s.elements = make(map[Elementer]bool)
	mux.Unlock()
}

func (set *Set) Add(elements ...Elementer) *Set {
	mux.Lock()
	for _, element := range elements {
		set.elements[element] = true
	}
	mux.Unlock()
	return set
}

func (set *Set) Remove(elements ...Elementer) *Set {
	mux.Lock()
	for _, element := range elements {
		delete(set.elements, element)
	}
	mux.Unlock()
	return set
}

func (set *Set) Has(element Elementer) bool {
	mux.RLock()
	_, has := set.elements[element]
	mux.RUnlock()
	return has
}

func (set *Set) ToSlice() []Elementer {
	elements := []Elementer{}
	mux.RLock()
	for element := range set.elements {
		elements = append(elements, element)
	}
	mux.RUnlock()
	return elements
}
