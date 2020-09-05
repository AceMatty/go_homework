package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntptime, err := ntp.Time("time.nist.gov")
	if err != nil {
		panic(err)
	}
	fmt.Println("current time:", ntptime.Format("2006-01-02 15:04:05"))
	fmt.Println("exact time:", time.Now().Format("2006-01-02 15:04:05"))
}
