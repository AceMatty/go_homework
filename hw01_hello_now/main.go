package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)
func main(){
	ntptime, err := ntp.Time("time.nist.gov")
	if err != nil {
		panic(err)
	}
	fmt.Println("current time:", ntptime)
	fmt.Println("exact time:", time.Now())
}
