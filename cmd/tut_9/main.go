// channels
//  used for communication between goroutines
//	way to enable goroutines to pass around data without needing locks
//	> hold data
//	> thread safe -> we avoid data races when reading and writing from memory
//	> we can listen when data is added or removed from a channel

// think of a channel as an array
package main

import "fmt"

func main() {
	var c = make(chan int) // we declare an int channel , it can hold only int values
	// this is callled an unbuffer chanel, and can only store one value

	// unbuffer chanel > room for only 1
	// c <- 1	//	when we write to an unbuffer channel,
	// the code will block until something else reads from it, which we are doing on the next line but control never reaches next line
	// var i = <- c // the channel is now empty and i holds the value of 1 now
	// or i:= <-c
	var cc = make(chan int, 5) // makes a buffer channel of size 5
	// now this channel can store more than one value without the main function needing to pop out any value

	// we use chanels along with go routines
	go process(c) // the control stops here until we write to chanel
	// fmt.Println(<-c)

	for i := range c { // as c takes values from 0 to 4
		// for loop can be used along with channel
		fmt.Println(i)
	}
}

func process(c chan int) {
	// c <- 123 // we write to channel
	// defer close(c) // means close the channel right before function exits
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c) // this notifies any other process using this channel that we are done and our main function will break out of loop
	// if we dont use close(c), we will be stuck in a deadlock at line 28 after loop is over
}

/*
	select statement
	select statement is used to wait on multiple channel operations
> it is like a switch statement for channels
> it will block until one of the cases can execute

*/
