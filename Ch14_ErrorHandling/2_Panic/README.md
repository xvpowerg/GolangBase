# Panic
會中斷程式碼，通常嚴重性高的

```go
func testSize(size int) {
	if size <= 0 {
		err := errors.New("size is Zero")
		panic(err)
	}
	fmt.Println("Pass!")
}
```
# recover()
Panic 可搭配recover取的error，注意必須使用defer
```go
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
```

