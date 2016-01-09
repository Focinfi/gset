package gset

import (
	"sync"
)

var mux sync.RWMutex

type Elementer interface {
	Element() interface{}
}

type ElementerFunc func() interface{}

func (e ElementerFunc) Element() interface{} {
	return e()
}

func Elementerlize(val interface{}) ElementerFunc {
	return ElementerFunc(func() interface{} { return val })
}

func Elementerslize(elementers ...interface{}) []Elementer {
	elementersArr := make([]Elementer, len(elementers))
	for i, elementer := range elementers {
		elementersArr[i] = Elementerlize(elementer)
	}
	return elementersArr
}

type Set struct {
	elements map[interface{}]bool
}

func NewSet(elements ...Elementer) *Set {
	set := &Set{make(map[interface{}]bool)}
	set.Add(elements...)
	return set
}

func (set *Set) Length() int {
	mux.RLock()
	defer mux.RUnlock()
	return len(set.elements)
}

func (set *Set) Clear() {
	mux.Lock()
	set.elements = make(map[interface{}]bool)
	mux.Unlock()
}

func (set *Set) Add(elements ...Elementer) *Set {
	mux.Lock()
	for _, element := range elements {
		set.elements[element.Element()] = true
	}
	mux.Unlock()
	return set
}

func (set *Set) Remove(elements ...Elementer) *Set {
	mux.Lock()
	for _, element := range elements {
		delete(set.elements, element.Element())
	}
	mux.Unlock()
	return set
}

func (set *Set) Has(element Elementer) bool {
	mux.RLock()
	_, has := set.elements[element.Element()]
	mux.RUnlock()
	return has
}

func (set *Set) ToSlice() []interface{} {
	elements := []interface{}{}
	mux.RLock()
	for element := range set.elements {
		elements = append(elements, element)
	}
	mux.RUnlock()
	return elements
}
