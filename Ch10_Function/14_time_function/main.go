package main

import (
	"fmt"
	"time"
)

func main() {
	//time文件如下
	//https://golang.org/pkg/time/
	//無影片
	//時間相關函數
	//取得目前時間
	now := time.Now()
	fmt.Printf("no=%v now type=%T\n", now, now)
	//取得年月日
	fmt.Printf("年:%v 月:%v 日:%v 時:%v 分:%v 秒:%v \n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())
	fmt.Printf("顯示數字day:%v\n", int(now.Month()))
	//格式化日期時間
	//注意 2006/01/02 15:04:05 這組格式數字是固定的
	//表示要顯示的年 月 日 時 分 秒
	//但順序或中間的分隔方式可修改
	fmt.Println(now.Format("2006-01-02"))
	//但順序可改
	fmt.Println(now.Format("01-02"))
	fmt.Println(now.Format("15:04:05"))
	//unix 與　unixnano
	fmt.Printf("unix %v unixnano %v\n", time.Now().Unix(), time.Now().UnixNano())

	//time常數
	/*
		const (
		    Nanosecond  Duration = 1
		    Microsecond          = 1000 * Nanosecond
		    Millisecond          = 1000 * Microsecond
		    Second               = 1000 * Millisecond
		    Minute               = 60 * Second
		    Hour                 = 60 * Minute
		)
	*/

	//time.Second
	fmt.Println(time.Second)
	//配合Sleep
	//Sleep可暫停
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second) //表示暫停1秒
		if i == 3 {
			break
		}
	}

}
