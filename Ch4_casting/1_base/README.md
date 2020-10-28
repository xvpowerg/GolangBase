# 基本型態轉型
基本型態可轉型為其他基本型態
## 案例
轉變為的類型(被轉換類型)
```go
var v1 int32 = 20
var v2 float32 = 25.67
//整數轉浮點數
v2 = float32(v1)
fmt.Println(v2)
//浮點數轉整數
v1 = int32(v2)
fmt.Println(v1)
```
