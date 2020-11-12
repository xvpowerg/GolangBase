# [TestFileExist](https://golang.org/pkg/os/#File.Stat)
查看檔案是否存在透過判斷Error
```go

filePath := `C:\MyDir\msg.txt`
	_, err2 := os.Stat(filePath)
	if os.IsNotExist(err2) {
		fmt.Println("檔案不存在!")
	} else {
		fmt.Println("檔案存在!")
	}
```


