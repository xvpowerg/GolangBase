# Struct Method
與Struct關聯的函數
## 語法
```go
type person struct {
	id   int
	name string
}

//Typte模式
func (p person) info() string {
	return fmt.Sprintf("id:%d name:%s\n", p.id, p.name)
}

//Point模式
func (p *person) infoP() string {
	return fmt.Sprintf("id:%d name:%s\n", p.id, p.name)
}
```
