# IFELSE
## 二分判斷條件
```go
	if age >= 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
```
## 可加上initialization statement
```go
	//可以使用初始化與宣告指令
	if y := 2 + 5; y >= 7 {
		fmt.Println("y>=7")
		fmt.Println("y:", y)
    }
    //可以使用初始化指令一次初始化多個
    if y, z := 2+5, 3*2.5; y >= 7 && z > 3 {
		fmt.Println("y>=7 z > 3")
	}
```