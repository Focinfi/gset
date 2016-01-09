Set data structure in Golang

#### Install

`go get github.com/Focinfi/gset`

#### Elementer interface

`gset` can receive element which has implemented Elmementer interface.

```go
  type Elementer interface {
    Element() interface{}
  }
```

And Here is an simple way to wrap a value to a Elementer


```go
  Elementerlize(1)
```

#### API

```go
  func NewSet(elements ...Elementer) *Set
  func (set *Set) Add(elements ...Elementer) *Set
  func (s *Set) Clear()
  func (set *Set) Has(element Elementer) bool
  func (s *Set) Length() int
  func (set *Set) Remove(elements ...Elementer) *Set
  func (set *Set) ToSlice() []Elementer
```
