package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("*")
	reciteFlag := flag.Bool("recite", false, "")
	dayFlag := flag.Int("reveal-for-day", 0, "help message for flag ")
	fmt.Println("**")
	flag.Parse()
	fmt.Println("***", *reciteFlag)
	if *reciteFlag {

	fmt.Println("******")
		fmt.Println(recitePoem())
	} else {

	fmt.Println("*!")
		fmt.Println(revealforDay(*dayFlag))
	}
}
