package gset

import (
	"testing"
)

var setThreadSafe = SetThreadSafe{Set: &Set{make(map[interface{}]bool)}}

func TestThreadSafeLength(t *testing.T) {
	setThreadSafe.Add(v, v, v1)
	if setThreadSafe.Length() != 2 {
		t.Error("Can distinguish same elements")
	}
}

func TestThreadSafeClear(t *testing.T) {
	setThreadSafe.Clear()
	if setThreadSafe.Length() != 0 {
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
	setThreadSafe := NewSet(v)
	if setThreadSafe == nil {
		t.Error("can not new a setThreadSafe")
	}
}

func TestThreadSafeHas(t *testing.T) {
	setThreadSafe := NewSet(v)
	if !setThreadSafe.Has(v) {
		t.Error("can not detect a element")
	}
}
