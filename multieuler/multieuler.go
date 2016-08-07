package multieuler

import "fmt"

func Multiples(x int, y int, limit int) {
	for i := 0; i < limit; i++ {
		if i%x == 0 {
			fmt.Printf("%v is a multiple of %v\n", i, x)
		}
		if i%y == 0 {
			fmt.Printf("%v is a multiple of %v\n", i, y)
		}
	}
}
