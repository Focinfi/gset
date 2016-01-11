package gset

import (
	"testing"
)

type Value struct {
	val int
}

func (v Value) Element() interface{} {
	return v.val
}

type Count struct {
	cout string
}

func (c Count) Element() interface{} {
	return c.cout
}

var v = &Value{1}
var v1 = &Count{"1"}
var set = Set{make(map[interface{}]bool)}

func TestLength(t *testing.T) {
	set.Add(v, v, v1)
	if set.Length() != 2 {
		t.Error("Can distinguish same elements")
	}
}

func TestClear(t *testing.T) {
	set.Clear()
	if set.Length() != 0 {
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

func TestT(t *testing.T) {
	set.Clear()
	set.Add(T(1))
	if set.Length() != 1 {
		t.Error("can not return a elementer")
	}
}
