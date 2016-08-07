package multieuler

import "fmt"

func Multiples(x int, y int, limit int) { //Multiples takes 3 int argsa
	xMult := false
	yMult := false
	for i := 1; i <= limit; i++ { //Loops through 'limit' arg
		if i%x == 0 { //Checks if current number is a multiple of x
			xMult = true
		} else {
			xMult = false
		}
		if i%y == 0 { //Checks if current number is a multiple of y
			yMult = true
		} else {
			yMult = false
		}
		switch {
		case xMult == true && yMult == true:
			fmt.Printf("%v is a multiple of %v and %v\n", i, x, y)
		case xMult == true && yMult == false:
			fmt.Printf("%v is a multiple of %v\n", i, x)
		case xMult == false && yMult == true:
			fmt.Printf("%v is a multiple of %v\n", i, y)
		case xMult == false && yMult == false:
		}
	}
	//This is not a very elegant solution, and does not provide an elegant output
}
