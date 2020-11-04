# Module 常用指令
* go mod init  建立Module
* go list -m all  列出Module的相依性
* go get 建立相依性
* go mod tidy  移除不必要的相依性

# 建立Module完成後的import 格式
* import 格式 module+路徑名稱
* 路徑名稱 可不跟package一樣名稱，只是用起來怪怪，所以建議一樣