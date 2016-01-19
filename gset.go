package gset

// Elementer it assumes that any objects who implented Element() interface {}
// could be a member of gset
type Elementer interface {
	Element() interface{}
}

// ElementerFunc is function implemented Elementer interface()
// using ElementerFunc(f), if f's defination is the same as ElementerFunc, ElementerFunc(f)
// will be a object implemented Elementer interface
type ElementerFunc func() interface{}

func (e ElementerFunc) Element() interface{} {
	return e()
}

// T wrap any type of object to be a ElementerFunc
func T(val interface{}) ElementerFunc {
	return ElementerFunc(func() interface{} { return val })
}

type Set struct {
	elements map[interface{}]bool
}

// NewSet create new Set with uncertain number of Elementers and return its pointer
func NewSet(elements ...Elementer) *Set {
	set := &Set{make(map[interface{}]bool)}
	set.Add(elements...)
	return set
}

// Length return this set's length
func (set *Set) Length() int {
	return len(set.elements)
}

// Clear reset the set to be empty container
func (set *Set) Clear() {
	set.elements = make(map[interface{}]bool)
}

// Add add uncertain number of Elementers and return this set's pointer
func (set *Set) Add(elements ...Elementer) *Set {
	for _, element := range elements {
		set.elements[element.Element()] = true
	}
	return set
}

// Remove remove uncertain number of Elementers and return this set's pointer
func (set *Set) Remove(elements ...Elementer) *Set {
	for _, element := range elements {
		delete(set.elements, element.Element())
	}
	return set
}

// Has ckeck if this set has this element
func (set *Set) Has(element Elementer) bool {
	_, has := set.elements[element.Element()]
	return has
}

// ToSlice return a new []interface{} containing the elements in this set
func (set *Set) ToSlice() []interface{} {
	elements := []interface{}{}
	for element := range set.elements {
		elements = append(elements, element)
	}
	return elements
}
