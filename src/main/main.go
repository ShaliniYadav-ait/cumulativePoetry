package main

import (
	"flag"
	"fmt"
)

func main() {
	reciteFlag := flag.Bool("recite", false, "")
	dayFlag := flag.Int("reveal-for-day", 0, "help message for flag ")
	flag.Parse()
	if *reciteFlag {
		fmt.Println(recitePoem())
	} else {
		fmt.Println(revealforDay(*dayFlag))
	}
}
