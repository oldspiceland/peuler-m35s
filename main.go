package main

import (
	"github.com/oldspiceland/peuler-m35s/multieuler"
	"flag"
)

func main() { //TODO: Add arguments to remove hard coded limits feeding into functions
	xPtr := flag.Int("x", 3, "an int")
	yPtr := flag.Int("y", 5, "an int")
	limitPtr := flag.Int("limit", 1000, "an int")
	
	flag.Parse()
	multieuler.Multiples(*xPtr, *yPtr, *limitPtr)
}
