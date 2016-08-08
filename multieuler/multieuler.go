package multieuler

import "fmt"

func Multiples(x int, y int, limit int) { //Multiples takes 3 int argsa
	xySlice := make([]int, limit)

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
			xySlice = append(xySlice, i)
		case xMult == true && yMult == false:
			fmt.Printf("%v is a multiple of %v\n", i, x)
			xySlice = append(xySlice, i)
		case xMult == false && yMult == true:
			fmt.Printf("%v is a multiple of %v\n", i, y)
			xySlice = append(xySlice, i)
		case xMult == false && yMult == false:
		}
	}

	sum := 0
	for n := 0; n <= len(xySlice)-1; n++ {
		sum += xySlice[n]
	}
	fmt.Printf("%v is the sum of all multiples of %v and %v\n", sum, x, y)

	//This is not a very elegant solution, and does not provide an elegant output
}
