# 常數 與 iota
不可變動的數
## 宣告方式
```go
const PI float64 = 3.1415927
	const (
        a int32     = 42
    )
```
## iota
+ 自動流水號
+ 取代Enum
## 自動給流水號
```go
	//iota 會自動編號
	//宣告星期幾
	const (
		Mon = iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)
```
## 取代Enum
```go
type Animal int

const (
		Tiger Animal = iota
		Cat
		Dog
)
```