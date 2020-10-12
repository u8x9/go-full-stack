package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// 时间戳
	ts := t.Unix()
	tsn := t.UnixNano() // 纳秒
	fmt.Println(ts, tsn)

	// 将时间戳转为时间
	t = time.Unix(ts, 0)
	fmt.Println(t)

	// 时间间隔：以纳秒为单位，可表示的最长时间段约为290年
	fmt.Println(time.Second * 10)
	fmt.Println(t.Add(time.Hour * 24))

	// 定时器
	//timer :=time.Tick(time.Second)
	//for tt:= range timer {
	//fmt.Println(tt)
	//}

	// 时间格式化 2006年1月2日 3点04分5秒
	// 2006 1 2 3 4 5
	fmt.Println(t.Format("2006年01月02日 03:04:05 PM"))
	fmt.Println(t.Format("2006年01月02日 15:04:05"))
	fmt.Println(time.Now().Format("2006年01月02日 15:04:05.000"))

	// 将字符串形式的转换成时间戳
	st, err := time.Parse("2006-01/02 15:04;05", "2020-10/13 02:53;21")
	if err != nil {
		fmt.Println(err)
	}
	var sts int64 = st.Unix()
	fmt.Println(st, sts)
}
