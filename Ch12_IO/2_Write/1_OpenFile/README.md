# OpenFile
可指定檔案讀寫操作模式

```txt
O_RDONLY  // open the file read-only.只能讀
O_WRONLY // open the file write-only.只能寫
O_RDWR    // open the file read-write.能讀寫
O_APPEND // append data to the file when writing.附加文字
O_CREATE // create a new file if none exists. 如果檔案不存在建立
O_EXCL    // used with O_CREATE, file must not exist. 跟O_CREATE一起使用 建立的檔案不可存在
O_SYNC    // open for synchronous I/O.同步IO
O_TRUNC   // truncate regular writable file when opened.會清空開啟的檔案
```