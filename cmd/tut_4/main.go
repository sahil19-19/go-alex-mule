// arrays, maps, slices and loops

package main

import "fmt"

func main() {
	// arrays
	// same type datatype
	// continuous memory locations
	//
	var intArr [5]int32 // initializes an array of length 5
	intArr[1] = 100
	fmt.Println(intArr[1:3]) // prints from 1st index to 2nd index

	var intArr1 = [3]int32{1, 2, 3}

	fmt.Println(intArr1)

	intArr2 := [4]int32{1, 2, 3} // last index values stays default
	// intArr2 := [...]int32 {1,2,3} // this too is valid and creates an array of size 3
	fmt.Println(intArr2)

	//	slices
	//	wrappers around the arrys
	//	arrays with more functions
	//	we just need to omit the length value to have slice

	intSlice := []int32{1, 2, 3} // we have a slice
	fmt.Printf("the length of the slice is %v and the capacity if %v", len(intSlice), cap(intSlice))
	// len() prints number of elements for slices and array rather than number of bytes

	intSlice = append(intSlice, 7) //
	//append returns a new array
	fmt.Printf("\nthe length of the slice is %v and the capacity if %v \n", len(intSlice), cap(intSlice))
	// Printf doesnt go to new line by default
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // appending multilple values using ... operator (spread operator)
	fmt.Println(intSlice)

	// another way to make slice
	var intSlice3 []int32 = make([]int32, 3, 8) // datatype, length, capacity(optional)
	//	if capactiy == lenght is true, then on next append the capacity will double as all values are copied into new slice
	//	to increase effeciecny of function, we
	fmt.Println(intSlice3)

	// maps
	var myMap map[string]uint8 = make(map[string]uint8)
	// means that both keys are of type string and values are of type uint8
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{
		"sahil": 23,
		"verma": 24, // trailing comma required
	}

	fmt.Println(myMap2)
	fmt.Println(myMap2["xyz"]) // returns 0, 
	// map also return an optional second value , which is a bool, its true if the key exists in the map
	var age, ok = myMap2["rahul"]
	if ok {
		println(age)
	} else {
		println("doesnt exist in map")
	}
	myMap["rahul"] = 10;
	// this adds another key-value pair into the map

	delete(myMap, "sahil") // to delete a key-value pair

	// loops
	for name := range myMap2 {
		// range based for loop
		fmt.Printf("the name is : %v \n", name)
	}

	for name, age := range myMap2 {
		// fmt.Printf("name : %v, age : %v \n", name, age)
		fmt.Println("name is", name , "and age is", age) // prints with a space
	}
	// when iterating over a map, the order of the keys is not preserved

	for i, v := range intSlice {
		fmt.Printf("index: %v, value: %v \n", i, v)
	}

	// there are no while loops in go
	// var i int = 0
	// for i<5 { >> for ( condition ) {}
	// 	fmt.Println("hello")
	// 	i++
	// }
	// or
	for i := 0; i < 5; i++ {
		fmt.Println("hello", i)
	}
	// for {
	// 	fmt.Println("Looping forever") // infinite loop
	// }
}
