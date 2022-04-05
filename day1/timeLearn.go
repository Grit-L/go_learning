package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%02d-%02d-%4d\n", t.Day(), t.Month(), t.Year())
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	fmt.Println(t.Format("15:04 02/Jan/2006"))

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	// calculating times:
	week := 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011

	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC))  // Wed Dec 21 08:56:34 2011

	time.Sleep(5 * time.Second)
	t1 := time.Now()
	s := t1.Format("20060103")
	fmt.Println(t1, "=>", s)

}
