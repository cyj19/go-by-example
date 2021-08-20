package main

import (
	"fmt"
	"time"
)

/*
	常用标准库-time包练习
*/

func timeDemo() {
	// 获取当前时间对象
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)
}

// 时间戳
func timestampDemo() {
	now := time.Now()
	// 时间戳
	timestamp1 := now.Unix()
	// 纳秒时间戳
	timestamp2 := now.UnixNano()
	fmt.Printf("%d %d \n", timestamp1, timestamp2)
}

func timestampDemo2(timestamp int64) {
	// 将时间戳转为时间对象
	now := time.Unix(timestamp, 0)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n", year, month, day, hour, minute, second)
}

func addDemo() {
	now := time.Now()
	// 原有时间上加一小时
	later := now.Add(time.Hour)
	fmt.Println(later)
}

func subDemo() {
	// 计算两个时间之差 t1-t2
	t1 := time.Now()
	t2 := time.Unix(1629442719, 0)
	fmt.Printf("t1-t2: %fs \n", t1.Sub(t2).Seconds())
	// 获取某个时间点 时间点-时间间隔
	t3 := t1.Add(-time.Hour)
	fmt.Printf("t1 before one hour: %d-%02d-%02d %02d:%02d:%02d \n", t3.Year(), t3.Month(), t3.Day(), t3.Hour(), t3.Minute(), t3.Second())
}

// 判断两个时间是否相等，会受时区影响
func equalDemo() {
	t1 := time.Now()
	t2 := time.Unix(1629442719, 0)
	t3 := time.Unix(1629442719, 0)
	fmt.Printf("t1 == t2: %t \n", t1.Equal(t2))
	fmt.Printf("t2 == t3: %t \n", t2.Equal(t3))
}

// 判断是否在某个时间点之前
func beforeDemo() {
	t1 := time.Now()
	t2 := time.Unix(1629442719, 0)
	fmt.Printf("t2 before t1 : %t \n", t2.Before(t1))
}

func afterDemo() {
	t1 := time.Now()
	t2 := time.Unix(1629442719, 0)
	fmt.Printf("t2 after t1 : %t \n", t2.After(t1))
}

func formatDemo() {
	now := time.Now()
	// 12小时制，需指定AM 或 PM
	fmt.Println(now.Format("2006-01-02 03:04:05 AM Mon Jon"))
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jon"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("15:04:05 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

// 解析字符串时间格式
func parseTimeDemo() {
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	// 解析字符串
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-08-20 15:04:05", loc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(time.Now()))
}

func main() {
	timeDemo()
	timestampDemo()
	timestampDemo2(time.Now().Unix())
	addDemo()
	subDemo()
	equalDemo()
	beforeDemo()
	afterDemo()
	formatDemo()
	parseTimeDemo()
}
