# Channel Buffer
Goroutine之間傳遞訊息
```go
    建立chan buffer為20
     c := make(chan int,20)
     可以連續呼叫channel傳送資料20筆不會block
     
```


