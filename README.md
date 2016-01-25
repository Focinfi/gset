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
  gset.T(1) // return a object with Element() function

```

#### Usage
----

##### New

Create a normal Set
```go
set := gset.NewSet(gset.T(1), gset.T("hello"))
```

Create a thread-safe Set
```go 
set := gset.NewSetThreadSafe(gset.T(1), gset.T("hello")) 
```

##### Manipulation

```go
  // add 1
  set.Add(gset.T(1))

  // length
  set.Length() // still 2

  // clear
  set.Clear()

  // chekc if set has 1
  set.Has(gset.T(1)) // true

  // get 1
  value, ok := set.Get(1) // value is gset.T(1), ok is true

  // remove 1 from set
  set.Remove(gset.T(1))
  set.Has(gset.T(1)) // false

  // get a array contains evey element in set
  elements := set.ToSlice() // []interface{}{"hello"}
```
