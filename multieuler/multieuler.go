package multieuler

import "fmt"

func Multiples(x int, y int, limit int) { //Multiples takes 3 int args
	for i := 0; i < limit; i++ { //Loops until 'limit' arg reached TODO: Currently missing last hit
		if i%x == 0 { //Checks if current number is a multiple of x
			fmt.Printf("%v is a multiple of %v\n", i, x)
		}
		if i%y == 0 { //Checks if current number is a multiple of y
			fmt.Printf("%v is a multiple of %v\n", i, y)
		}
	}
	//This is not a very elegant solution, and does not provide an elegant output
}
