package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	timeDemo()
	timestampDemo()
	now := time.Now()
	timestampDemo2(now.Unix())
	tickerDemo()
}

func timeDemo() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

/*时间戳*/
func timestampDemo() {
	now := time.Now()
	timestamp := now.Unix()         //时间戳
	timestampNano := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp)
	fmt.Printf("current timestampNano:%v\n", timestampNano)
}

/*时间戳转为时间*/
func timestampDemo2(timestamp int64) {
	now := time.Unix(timestamp, 0)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

/*时间间隔*/
const (
	Nanosecond  time.Duration = 1 //1纳秒
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

/*定时器*/
func tickerDemo() {
	//定时器的本质是channel
	tiker := time.Tick(time.Second)
	for t := range tiker {
		log.Printf("定时器执行了 %v 秒 \n", t)
	}
}
