# Varargs 特性
```go
func sum(value ...int){

}
//Varargs 只能是最後一筆參數 
```
# Slice傳入Varargs
```go
values := []int{5, 6, 7, 8, 9}
//Slice傳入Varargs
v2 := sum(values...)
```
