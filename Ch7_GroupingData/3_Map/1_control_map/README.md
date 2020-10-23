# Map操作

## Create Map
```go
stm := map[string]int{
		"Ken":   80,
		"Lindy": 70,
	}
map2 := make(map[string]int)
```
## keｙ是否存在　Keyexist
```go
if v, ok := stm["Lindy"]; ok {
		fmt.Println("exist",v)
	}
```
## range 輪巡map
```go
for key, value := range stm {
		fmt.Println(key, value)
	}
```



