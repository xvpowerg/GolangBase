# Select Channel TimeOut
Select Channel 可搭配TimeOut時間


```go
select {
		case v := <-c1:
			fmt.Printf("c1 value:%d\n", v)
		case v := <-c2:
			fmt.Printf("c2 value:%d\n", v)
		case <-time.After(1 * time.Second): //一秒後沒收到訊息自動結束
			fmt.Println("After Stop ")
			return
		}
```

