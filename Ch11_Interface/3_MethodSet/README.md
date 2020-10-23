# MethodSet
指Struct的方法是否可相容某Interface
```text
T  可看到 T的方法
*T 可看到 T的方法與*T的方法

type T struct {}
func (i T) 方法名稱() {}
func (i *T) 方法名稱() {}

```
