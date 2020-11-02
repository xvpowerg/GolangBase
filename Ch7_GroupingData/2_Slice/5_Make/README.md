# Slice Make
* 如果知道需要的容量最好指定容量，可提升效能
* append方法會在容量不足時make一份容量為目前`2倍的Slice`，在將目前的內容Copy到新的Slice，這樣的動作很消耗效能