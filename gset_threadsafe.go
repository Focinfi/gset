package gset

import (
	"sync"
)

// SetThreadSafe is wrapper of Set,
// ensure the thread safe via sync.RWMutex
type SetThreadSafe struct {
	sync.RWMutex
	*Set
}

func NewSetThreadSafe(elements ...Elementer) *SetThreadSafe {
	return &SetThreadSafe{Set: NewSet(elements...)}
}

func (set *SetThreadSafe) Len() int {
	set.RLock()
	defer set.RUnlock()
	return set.Set.Len()
}

func (set *SetThreadSafe) Clear() {
	set.Lock()
	set.Set.Clear()
	set.Unlock()
}

func (set *SetThreadSafe) Add(elements ...Elementer) *Set {
	set.Lock()
	defer set.Unlock()
	return set.Set.Add(elements...)
}

func (set *SetThreadSafe) Remove(elements ...Elementer) *Set {
	set.Lock()
	defer set.Unlock()
	return set.Set.Remove(elements...)
}

func (set *SetThreadSafe) Has(element Elementer) bool {
	set.RLock()
	defer set.RUnlock()
	return set.Set.Has(element)
}

func (set *SetThreadSafe) Get(key interface{}) (Elementer, bool) {
	set.RLock()
	defer set.Unlock()
	return set.Set.Get(key)
}

func (set *SetThreadSafe) ToSlice() []interface{} {
	set.RLock()
	defer set.RUnlock()
	return set.Set.ToSlice()
}
