# [TestFileExist](https://golang.org/pkg/os/#File.Stat)
查看檔案是否存在透過判斷Error
```go
f, _ := os.Open(`C:\MyDir\msg.txt`)
_, err := f.Stat()
if err == nil {
	fmt.Println("exist!")
} 
//也可用封裝好的 os.Stat與 os.IsNotExist 搭配
filePath := `C:\MyDir\msg.txt`
_, err2 := os.Stat(filePath)
fmt.Println(os.IsNotExist(err2))

```


