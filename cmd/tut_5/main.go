// strings, runes and bytes
package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString string = "sàhilâāver"
	var myString2 = []rune("sàhilâāver") // cast them into an array of runes
	// "àâā > these are not utf8 characters and hence cant be displyayed using 8bits
	// runes are just an alias for int32 characters > uses 32 bits to stores characters

	fmt.Println(myString)

	var indexed, indexed2 = myString[0], myString2[0]

	fmt.Printf("%v %T \n%v, %T\n", indexed, indexed, indexed2, indexed2) // %T prints the datatype

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Println(len(myString)) // returns number of bytes

	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	fmt.Println(len(myString2)) // we would get the correct length in this case

	var myRune = 'a'
	fmt.Printf("printing a: %v\n", myRune) // prints 97
	/*
		x := 'a'
		fmt.Printf("%c\n", x) // Output: a
		fmt.Printf("%d\n", x) // Output: 97 // ascii of 'a'
		fmt.Printf("%T\n", x) // Output: int32
		fmt.Printf("%x\n", x) // Output: 61 ??
	*/

	// strings
	// strings are immutable in go
	// we cannot reassign a new value at an index once formed

	// when we concatenate strings using + symbol we create a completely new string
	// this is ineffecient

	var mySlice = []string{"a", "b", "c", "d", "e"} // a slice of strings
	// var catString string
	// rather we use the strings.Builder
	var strBuilder strings.Builder
	for ind := range mySlice {
		// catString += mySlice[ind]	// this creates a new string every time
		strBuilder.WriteString(mySlice[ind])
		// internaly an array is created and values are appended to the array
	}
	var catString = strBuilder.String()
	fmt.Println(catString)

	allStrings := []string{
		"a",
		"b",
		"c",
		"d",
	}

	var finalString string
	// var myStringBuilder strings.Builder
	for _, s := range allStrings {
		strBuilder.WriteString(s)
	}

	finalString = strBuilder.String()
	fmt.Println(finalString) // abcdeabcd
	/*
		strings.Builder vs []string{}


					parts := []string{"Hello", " ", "World", "!"}
			    // Join at the end
			    result := strings.Join(parts, "")

					Every time you Join, Go has to allocate a new string for the concatenation.
					If you keep appending dynamically, you might end up repeatedly reallocating memory, which is less efficient.
	*/
}
