# [基本型態轉字串](https://golang.org/pkg/strconv/#pkg-overview)
可使用strconv將基本型態轉字串，使用Fotmat(基本形態類型)方法如下
```go
	//基本類型轉字串
var v1 int32 = 20
var x2 float32 = 123.456789	
	//參數 1 int64 位元的整數 2 10進位
str1 := strconv.FormatInt(int64(v1), 10)
fmt.Println(str1)

//參數 1 轉換為字串的數值 2 轉換後顯示的模式 3 取道小數點第幾位 4 轉換的變數是幾位元
	s32 := strconv.FormatFloat(float64(x2), 'f', 3, 32)
	fmt.Println(s32)
```
# [字串轉基本型態](https://golang.org/pkg/strconv/#pkg-overview)
可使用 strconv 將字串轉換為基本型態

```go
b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
//(字串 進位模式 轉換後的大小)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
如果轉型錯誤err不是nil
```
# 字串轉[]byte
```go
str4 := "Howard"
	b := []byte(str4)
```
