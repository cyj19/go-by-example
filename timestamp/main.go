package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millils := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(nanos)
	fmt.Println(millils)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
