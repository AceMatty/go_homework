package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/beevik/ntp"
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
