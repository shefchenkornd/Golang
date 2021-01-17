package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

const host = "0.beevik-ntp.pool.ntp.org"

func main() {
	time, err := ntp.Time(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error of ntp: %s", err)
		os.Exit(1)
	}

	fmt.Println(time)
}
