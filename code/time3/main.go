package main

import (
	"fmt"
	"time"
)

func main() {
	// timestamp := time.Now().Unix()
	// t := time.Unix(timestamp, 0)
	// nt := t.Format("2006-01-02 15:04")
	// fmt.Println(nt)

	t := time.Now().Add(-120 * time.Minute)
	s := t.Format("2006-01-02 15:04")
	timestampInt, err := time.Parse("2006-01-02 15:04:05", s+":00")
	fmt.Println(timestampInt)
	fmt.Println(err)
	rt := timestampInt.Add(-8 * 60 * 60 * time.Second).Unix()
	fmt.Println(rt)

	fmt.Println(time.Since(t) >= 2*time.Hour)
}
