package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello world")
	var intNum int16 = 32767 // max number
	// var uintNum uint64 // stores only positive numbers, twice the size possible
	// int will default to 32 or64 bit depending on system type
	// int , int8,int16, int32, int64

	var floatNum float32 = 2131.52 // for float64

	fmt.Println(intNum, floatNum)

	// arithmatic ops
	// we cant perform arithmatic operation of mixed types
	// we need to typecast
	var temp = intNum * int16(floatNum)
	fmt.Println(temp)

	// strings
	var myString string = "hello \nworld"
	var myString1 = `hello 
world`
	var hellostring = "hello" + " " + "world"
	fmt.Println(myString)
	fmt.Println(myString1, hellostring)

	fmt.Println(utf8.RuneCountInString(hellostring))
	// to print length of string
	// rune is datatype
	fmt.Println(len(hellostring))
	// this returns the number of bytes
	// for non ascii characters, the number of bytes is more than the number of characters
	// for ascii characters, the number of bytes is equal to the number of characters

	var myRune rune = 'a'
	// single
	// alias for int32 (4bytes), used to represent a single unicode code point
	// go uses utf-8 encoding for strings
	// go's way of handling individual characters, especially for international text that goes beyond ASCII.

	fmt.Println(myRune) // this will print the unicode code point of 'a' which is 97

	// default value of all int is 0, string is empty string, bool is false

	// we can also drop the var keyword and use a shorthand
	myshorthandvar := "hello"

	fmt.Println(myshorthandvar)
	
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst int64 = 1000
	// const are variables for which we cant reassign values
	// we cant just declare them and not initialize
}

/*
   printf -> print formatter -> this function allows you to format numbers,
           variables and strings into the first string parameter you give it
   print -> prints line , doesnt format anything
   println -> similar to print, add '\n' at the end of the line
*/
