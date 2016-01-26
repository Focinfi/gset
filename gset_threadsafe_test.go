package gset

import (
	"testing"
)

var setThreadSafe = SetThreadSafe{Set: &Set{make(map[interface{}]Elementer)}}

func TestThreadSafeLen(t *testing.T) {
	setThreadSafe.Add(v, v, v1)
	if setThreadSafe.Len() != 2 {
		t.Error("Can distinguish same elements")
	}
}

func TestThreadSafeClear(t *testing.T) {
	setThreadSafe.Clear()
	if setThreadSafe.Len() != 0 {
		t.Error("can not clear the setThreadSafe")
	}
}

func TestThreadSafeAdd(t *testing.T) {
	setThreadSafe.Add(v)
	if !setThreadSafe.Has(v) {
		t.Error("can not add a element")
	}
}

func TestThreadSafeRemove(t *testing.T) {
	setThreadSafe.Add(v)
	setThreadSafe.Remove(v)
	if setThreadSafe.Has(v) {
		t.Error("can not add a element")
	}
}

func TestThreadSafeNewSet(t *testing.T) {
	setThreadSafe := NewSetThreadSafe(v)
	if setThreadSafe == nil {
		t.Error("can not new a setThreadSafe")
	}
}

func TestThreadSafeHas(t *testing.T) {
	setThreadSafe := NewSetThreadSafe(v)
	if !setThreadSafe.Has(v) {
		t.Error("can not detect a element")
	}
}

func TestThreadSafeGet(t *testing.T) {
	set := NewSetThreadSafe(v)
	value, ok := set.Get(v.Element())
	if !ok || value != v {
		t.Error("can get a element")
	}
}

func TestThreadSafeToSlice(t *testing.T) {
	elements := NewSetThreadSafe(v).ToSlice()
	if elements[0] != v {
		t.Errorf("can not to slice, the first element is: %v", elements[0])
	}
}
