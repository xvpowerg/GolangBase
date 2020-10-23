# 限定Channel接收或傳送 Directional Channel 
宣告時語法限定Channel只能接收或傳送
```go
	cr := make(<-chan int) //recive
	cs := make(chan<- int) //send
    func testSend(ch chan<- int) {　｝
    func testRecive(ch <-chan int) {
```


