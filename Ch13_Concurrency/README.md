# Concurrency
## 作業
* 1~100 當中偶數與奇數，幫我用evenChan,oddChan兩個Channel做分類，並使用evenList與oddList收集偶數奇數
```go
evenChan := make(chan int)
oddChan := make(chan int)
evenList, oddList := receive1(....)
......
```


