// FUNCTIONS AND CONTROL STRUCTURES

package main

import (
	"errors"
	"fmt"
)

func main() {
	var myString = "bye world"
	var age int16 = 20
	helper(myString, age)
	// fmt.Println(divHelper(10,2));
}

func helper(myString string, age int16) {
	// we need to specify data types?
	// fmt.Println("executing helper function")
	// fmt.Println(`hello, my age is ${age}`)
	fmt.Println(myString, age)
	var result, rem, err = divHelper(63, 4)
	// if(err != nil){
	// 	fmt.Println(err.Error())
	// }
	// using switch / case
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case rem == 0, rem == 1:
		fmt.Printf("the result of the division is %v", result)
	default:
		fmt.Printf("the result of the division is %v and the remainder is %v", result, rem)
	}
	// we dont need to write break statements in golang
	// in switch both , or || can be used but
	// fun(), gun() and fun()  || gun() are different, [ with , both functions are evaluated, with ||, concept of short circuit works]

	// conditions switch statement
	switch rem {
	case 0:
		fmt.Println("the remainder is 0")
	case 1, 2: // 1 or 2
		fmt.Println("the remainder is not 0")
	default:
		fmt.Println("the remainder is greater than 2")
	}
}

// a function that does int division and returns and int value
func divHelper(a int, b int) (int, int, error) {
	var err error // a builtin type in go
	if b == 0 {
		err = errors.New("cannot diviede by 0")
		return 0, 0, err
	}
	return a / b, a % b, err // we can return multiple values at the same time
}
