package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	now, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Error from ntp: %v", err)
		os.Exit(0)
	}
	fmt.Println(now)
}
