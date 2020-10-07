# IO

## os.File
跟檔案相關的操作可使用[File](https://golang.org/pkg/os/#File)
* 可透過os.Open 開啟檔案如下
```go
file, err := os.Open("file.go") // For read access.
if err != nil {
	log.Fatal(err)
}
```

