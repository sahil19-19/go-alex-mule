// struct, and interface
// struct is a custom date type
// properties plus methods
package main

import (
	"fmt"
)

type address struct {
	pincode uint32
}

func (a address) printPinCode() {
	fmt.Println("Pincode is", a.pincode)
}

type student struct {
	name    string
	age     uint16
	rollno  uint16
	address // no need to write name of variable
	// promotes printPinCode() to student
	// int  // can also do this
}

// these functions are called methods
func (e student) showDetails() uint16 { // e student means it belongs to struct studnt
	return e.age * e.rollno
}

type teacher struct {
	posno    uint16
	yearsExp uint16
}

func (e teacher) showDetails() uint16 {
	return uint16(e.posno) * uint16(e.yearsExp)
}

type member interface {
	showDetails() uint16 // method signature
}

// interface is a type that defines behavior

//	func idValue(e student, eId uint16) uint16 {
//		return e.showDetails() * eId;
//	}
//
// say we want to use only one function for both teacher and student
func idValue(e member, eId uint16) uint16 {
	return e.showDetails() * eId
}

// now this method can work with struct type given, that that struct has a showDetails() method with the given method signature
// this is polymorphism in Go without inheritance

type Shape interface {
	Area() float64
	Perimeter() float64
}

// any type must implement both methods

func main() {
	var student1 student = student{name: "sahil", age: 23}
	// we can also omit the property names but order matters
	var student2 student = student{"verma", 23, 24, address{171009}} // we need to pass all arguments in this case
	student1.rollno = 24
	fmt.Println(student1, student2)

	//	anonymous struct
	//	we define and initialize it at the same time
	//	not reusable
	var student3 = struct {
		name string
		age  uint8
	}{"sah", 25}

	fmt.Println(student3)
	// fmt.Println(student1.showDetails())

	var teacher1 teacher = teacher{10, 4}
	fmt.Println(idValue(teacher1, 20))
	fmt.Println(idValue(student1, 10))

}
