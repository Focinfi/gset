package gset

import (
	"testing"
)

type Value struct {
	val int
}

func (v Value) Id() interface{} {
	return v.val
}

type Count struct {
	cout string
}

func (c Count) Id() interface{} {
	return c.cout
}

var v = &Value{1}
var v1 = &Count{"1"}
var set = Set{make(map[interface{}]IdGetter)}

func TestLen(t *testing.T) {
	set.Add(v, v, v1)
	if set.Len() != 2 {
		t.Error("Can distinguish same elements")
	}
}

func TestClear(t *testing.T) {
	set.Clear()
	if set.Len() != 0 {
		t.Error("can not clear the set")
	}
}

func TestAdd(t *testing.T) {
	set.Add(v)
	if !set.Has(v) {
		t.Error("can not add a element")
	}
}

func TestRemove(t *testing.T) {
	set.Add(v)
	set.Remove(v)
	if set.Has(v) {
		t.Error("can not add a element")
	}
}

func TestNewSet(t *testing.T) {
	set := NewSet(v)
	if set == nil {
		t.Error("can not new a set")
	}
}

func TestHas(t *testing.T) {
	set := NewSet(v)
	if !set.Has(v) {
		t.Error("can not detect a element")
	}
}

func TestGet(t *testing.T) {
	set := NewSet(v)
	value, ok := set.Get(v.Id())
	if !ok || value != v {
		t.Error("can get a element")
	}
}

func TestT(t *testing.T) {
	set.Clear()
	set.Add(T(1))
	if set.Len() != 1 {
		t.Error("can not return a KeyGeter")
	}
}

func TestToSlice(t *testing.T) {
	elements := NewSet(v).ToSlice()
	if elements[0] != v {
		t.Errorf("can not to slice, the first element is: %v", elements[0])
	}
}
