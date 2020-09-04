package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
	"time"
)
func main(){
	ntptime, err := ntp.Time("time.nist.gov")
	if err != nil {
		_, err := io.WriteString(os.Stderr, "GfG\n")
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("current time:", ntptime)
	fmt.Println("exact time:", time.Now())
}
