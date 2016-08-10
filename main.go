package main

import (
	"github.com/oldspiceland/peuler-m35s/multieuler"
	"flag"
	"fmt"
	"os"
)

func main() { 
	xPtr := flag.Int("x", 3, "an int")
	yPtr := flag.Int("y", 5, "an int")
	limitPtr := flag.Int("limit", 1000, "an int")
	
	flag.Parse()
	validate(*xPtr, *yPtr, *limitPtr)
	multieuler.Multiples(*xPtr, *yPtr, *limitPtr)
}

func validate(x int, y int, limit int) {
	if (x <= 0 || y <= 0 || limit <= 0) {
		fmt.Print("Invalid Input, negative values not allowed\n")
		os.Exit(1)
	}
}

