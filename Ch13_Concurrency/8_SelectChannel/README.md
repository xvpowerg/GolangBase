# select Channel
可判定是哪個Channel做接收，注意一定要給一個中斷點不然運行會出錯
```go
		select {
		case v := <-c1:
			fmt.Printf("c1 value:%d\n", v)
		case v := <-c2:
			fmt.Printf("c2 value:%d\n", v)
		case <-quit:
			return
		}
```


