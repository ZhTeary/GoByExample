package builtin

import (
	"fmt"
	"time"
)

func ITime() {
	p := fmt.Println

	//get current time
	now := time.Now()
	p(now)

	//build time, time通常会呆着地址，那是市区
	then := time.Date(2023, 10, 13, 18, 14, 45, 123456789, time.Local /*世界时time.UTC*/)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond()) //纳秒 十亿分之一 10^-9
	p(then.Location())

	//duration 返回两个时间的时间段
	diff := now.Sub(then)
	fmt.Println(diff)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())

	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))

	//Epoch
	fmt.Println("============")
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	//time.Unix(秒，纳秒)
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))

	fmt.Println("============")
	// time format
	tt := time.Now()
	fmt.Println(tt.Format(time.RFC3339))

	tp, _ := time.Parse(
		time.RFC3339,
		"2023-10-17T11:28:45+00:00")
	fmt.Println(tp)

}
