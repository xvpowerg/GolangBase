# Init function特性
* 在main運行時會執行
* 只會執行一次
* 可有多個init函數 執行順序由上至下

```go
func init() {
	fmt.Println("init() 只會運行一次 可做初始化用")
}
```