package main

import (
	"flag"
	"fmt"
)

func main() {
	reciteFlag := flag.Bool("recite", false, "")
	dayFlag := flag.Int("reveal-for-day", 0, "help message for flag ")
	echo := flag.Bool("echo",false,"")
	flag.Parse()
	if *reciteFlag {
		fmt.Println(recitePoem(*echo))
	} else {
		fmt.Println(revealforDay(*dayFlag,*echo))
	}
}
