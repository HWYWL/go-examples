package main

import (
	"fmt"
	"time"
	"log"
)

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"

	//自定义
	MyRFC	   = "2006-01-02 15:04:05"
	MyRFC1	   = "2006-01-02"
)

//获取当前时间
func getNow()  {
	//获取当前时间，返回time.Time对象
	//2018-05-15 17:41:09.9962013 +0800 CST m=+0.003000001
	// 其中CST可视为美国，澳大利亚，古巴或中国的标准时间
	// +0800表示比UTC时间快8个小时
	fmt.Println(time.Now())

	//获取当前时间戳
	fmt.Println(time.Now().Unix())
	//精确到纳秒，通过纳秒就可以计算出毫秒和微妙
	fmt.Println(time.Now().UnixNano())
}

//格式化时间显示
func formatUnixTime()  {
	//获取当前时间，进行格式化
	//2018-05-15 17:50:45
	fmt.Println(time.Now().Format(MyRFC))

	//指定的时间格式化
	//2018-05-15 17:49:29
	fmt.Println(time.Unix(1526377769, 0).Format(MyRFC))
}

//获取指定时间戳的年份
func getYear()  {
	//获取指定时间戳的年月日时分秒
	t := time.Unix(1526377769, 0)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

//时间字符串转换时间戳
func strToUnix()  {
	//先用time.Parse对时间字符串进行分析，如果正确会得到一个time.Time对象
	//后面就可以用time.Time对象的函数Unix进行获取
	t, err := time.Parse(MyRFC, "2018-05-15 18:01:32")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(t)
	fmt.Println(t.Unix())
}

//根据时间戳获取当日开始的时间戳
func getDayStartUnix()  {
	t := time.Unix(1526407292, 0).Format(MyRFC1)
	sts, err := time.Parse(MyRFC1, t)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sts.Unix())
}

//休眠
func sleep()  {
	//time.Millisecond  1毫秒
	// time.Microsecond 1纳秒

	//休眠一秒
	time.Sleep(1 * time.Second)
	fmt.Println("我休眠了1秒")

	//休眠100毫秒
	time.Sleep(100 * time.Millisecond)
	fmt.Println("我休眠了100毫秒")
}

func main() {
	getNow()
	fmt.Println("==========邪恶的分割线==========")

	formatUnixTime()
	fmt.Println("==========邪恶的分割线==========")

	getYear()
	fmt.Println("==========邪恶的分割线==========")

	strToUnix()
	fmt.Println("==========邪恶的分割线==========")

	getDayStartUnix()
	fmt.Println("==========邪恶的分割线==========")

	sleep()

}
