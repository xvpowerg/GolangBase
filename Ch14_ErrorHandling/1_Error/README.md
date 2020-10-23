# Error
常見的處裡錯誤方式,不會中斷程式，通常嚴重性不高的
```go
	v1, err := division(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
```
# 可自訂Error只要符合以下格式介面
```go
type error interface {
    Error() string
}

```


