package gset

// IdGetter it assumes that any objects who implented Id() interface {}
// could be a member of gset
type IdGetter interface {
	Id() interface{}
}

// IdGetterFunc is function implemented IdGetter interface()
// using IdGetterFunc(f), if f's defination is the same as IdGetterFunc, IdGetterFunc(f)
// will be a object implemented IdGetter interface
type IdGetterFunc func() interface{}

func (e IdGetterFunc) Id() interface{} {
	return e()
}

// T wrap any type of object to be a IdGetterFunc
func T(val interface{}) IdGetterFunc {
	return IdGetterFunc(func() interface{} { return val })
}

type Set struct {
	elements map[interface{}]IdGetter
}

// NewSet create new Set with uncertain number of IdGetters and return its pointer
func NewSet(elements ...IdGetter) *Set {
	set := &Set{make(map[interface{}]IdGetter)}
	set.Add(elements...)
	return set
}

// Length return this set's length
func (set *Set) Len() int {
	return len(set.elements)
}

// Clear reset the set to be empty container
func (set *Set) Clear() {
	set.elements = make(map[interface{}]IdGetter)
}

// Add add uncertain number of IdGetters and return this set's pointer
func (set *Set) Add(elements ...IdGetter) *Set {
	for _, element := range elements {
		set.elements[element.Id()] = element
	}
	return set
}

// Remove remove uncertain number of IdGetters and return this set's pointer
func (set *Set) Remove(elements ...IdGetter) *Set {
	for _, element := range elements {
		delete(set.elements, element.Id())
	}
	return set
}

// Has ckeck if this set has this element
func (set *Set) Has(element IdGetter) bool {
	_, has := set.elements[element.Id()]
	return has
}

// Get
func (s *Set) Get(key interface{}) (interface{}, bool) {
	IdGetter, ok := s.elements[key]
	return IdGetter, ok
}

// ToSlice return a new []interface{} containing the elements in this set
func (set *Set) ToSlice() []interface{} {
	elements := []interface{}{}
	for elementKey := range set.elements {
		elements = append(elements, set.elements[elementKey])
	}
	return elements
}

func (s *Set) ToMap() map[interface{}]IdGetter {
	return s.elements
}
