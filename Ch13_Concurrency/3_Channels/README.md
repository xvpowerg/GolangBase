# Channel
Goroutine之間傳遞訊息
```go
    //建立chan 預設buffer為0
     c := make(chan int)
     //當buffer時，channel傳送資料會block
     c<-10 //block
```


