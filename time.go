package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Seconds since epoch
	epochSeconds := now.Unix()

	// Milliseconds since epoch
	epochMillis := now.UnixMilli()

	fmt.Println("Epoch Seconds :", epochSeconds)
	fmt.Println("Epoch Millis  :", epochMillis)

	// 00:00:00 UTC on Jan 1, 1970

	
	unixTime := now.Unix()
	fmt.Println("Current Unix Time:", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Time:", t.Format("2006-01-02"))

}
