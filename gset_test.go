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

var v = &Value{1}
var set = Set{make(map[Elementer]bool)}

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
