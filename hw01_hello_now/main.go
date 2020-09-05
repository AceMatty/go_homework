package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntptime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		panic(err)
	}
	exect, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		panic(err)
	}
	times := time.Now().Add(exect.ClockOffset)
	fmt.Println("current time:", ntptime.Format("2006-01-02 15:04:05 -0700 MST"))
	fmt.Println("exact time:", times.Format("2006-01-02 15:04:05 -0700 MST"))
}
