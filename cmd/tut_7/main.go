// pointers

package main

import "fmt"

// here we create a copy of the array
func square(arr [5]int) [5]int {
	//	func square(arr *[5]int) [5]int {
	//	if we wanted to pass the array's pointer
	fmt.Printf("%p \n", &arr)
	for i := range arr {
		arr[i] *= arr[i]
	}
	return arr
}

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}
}

// this will modify the array that was passed
func modifyArray(arr *[5]int) {
	// no need to do (*arr)[0] = 100, go handles that itself
	if len(arr) > 0 {
		arr[0] = 100
	}
}

func main() {
	var p *int // takes 32 or 64 bit depending upon the archi
	// intially this stores the value nil
	var i int = 10
	p = &i
	fmt.Println(p)

	var q *int = new(int) // in this case q will store a random free address
	// the value at that memory location will be 0, "", false depending upon the type of pointer

	fmt.Println(q)
	fmt.Println(*q) // deferencing the pointer

	myArr := [5]int{1, 2, 3, 4, 5}
	var newArr = myArr
	// fmt.Println(&newArr, &myArr)
	fmt.Printf("printing address of arrays >>> %p %p \n", &myArr, &newArr) // to print address
	// these 2 point to the different memory address
	// this is different for slices, shown below

	newArray := square(myArr)
	fmt.Println(myArr) // array is still the same
	fmt.Println(newArray)
	// values of myArr are not changed cause we created a copy of the array

	modifyArray(&myArr)
	fmt.Println(myArr) // this now modifies the array that we passed

	// slices on the other hand are passed by refernce
	mySlice := []int{1, 3, 5, 7, 9}

	copySlice := mySlice
	fmt.Printf("printing address of slices >>> %p %p \n", mySlice, copySlice) // this does not print the address to the the slice variable, but rather
	// it prints the address to which the pointer to underlying array points (address of first element of underlying array)
	// AND &mySlice would print the address of the slider headers
	// Slices in go are a struct

	modifySlice(mySlice) // this does modify the slice
	fmt.Println(mySlice)
	copySlice[1] = 90;
	fmt.Println(mySlice)
	/*
		   Important Slice Gotcha

		   If you append, things can change:

		   func modifySlice(s []int) {
		       s = append(s, 100)
		   }

		   This may:

		   Create a new underlying array
		   Not affect the original slice

			 We handle this >> we return the slice

			func modifySlice(s []int) []int {
				s = append(s, 100)
				return s
			}
	*/

}
