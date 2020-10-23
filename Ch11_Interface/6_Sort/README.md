# Sort
如何排序Slice可使用
* [sort.Interface](https://golang.org/pkg/sort/#Interface)
```go
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```
* [SliceStable](https://golang.org/pkg/sort/#Interface)
```go
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
```

