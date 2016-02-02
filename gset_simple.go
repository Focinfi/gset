package gset

type SetSimple struct {
	elements map[interface{}]bool
}

func NewSetSimple(elements ...interface{}) *SetSimple {
	set := &SetSimple{map[interface{}]bool{}}
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

func (set *SetSimple) Len() int {
	return len(set.elements)
}

func (set *SetSimple) Has(element interface{}) bool {
	_, ok := set.elements[element]
	return ok
}

func (set *SetSimple) Add(element interface{}) {
	if !set.Has(element) {
		set.elements[element] = true
	}
}

func (set *SetSimple) Remove(element interface{}) {
	if set.Has(element) {
		delete(set.elements, element)
	}
}

func (set *SetSimple) ToSlice() []interface{} {
	slice := []interface{}{}
	for element := range set.elements {
		slice = append(slice, element)
	}
	return slice
}
