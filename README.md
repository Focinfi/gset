Set data structure in Golang

#### Install

`go get github.com/Focinfi/gset`

#### Elementer interface

`gset` can receive elements implemented Elementer interface.

```go
  type Elementer interface {
    Element() interface{}
  }
```

And Here is an simple way to wrap a value to a Elementer


```go
  Elementerlize(1) // return a object with Element() function

  Elementerslize(1, 2, 3) // return a array of Elementer objects
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
