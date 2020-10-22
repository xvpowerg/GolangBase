# 常數 與 iota
不可變動的數
## 宣告方式
```go
const PI float64 = 3.1415927
	const (
        a int     = 42
    )
```
## iota
+ 可自動給流水號
+ 取代Enum
## 使用方式
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
