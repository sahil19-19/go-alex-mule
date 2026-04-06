package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_PRICE float32 = 5
const MAX_TOFU_PRICE float32 = 4

func main() {
	websites := []string{"walmart", "costco", "wholefoods"}
	// var wg = sync.WaitGroup{}

	chickenChannel := make(chan string)
	tofuChannel := make(chan string)

	for i := range websites {
		// wg.Add(1)
		go checkPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}
	// wg.Wait()
	sendMessage(chickenChannel, tofuChannel)
}

func checkPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20
		if chickenPrice <= MAX_TOFU_PRICE {
			chickenChannel <- website // the code blocks until we read from it
			break
		}
	}
}

func checkTofuPrice(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		tofuPrice := rand.Float32() * 20
		if tofuPrice <= MAX_PRICE {
			tofuChannel <- website // the code blocks until we read from it
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// fmt.Println("Deal found:", <-chickenChannel)
	select {
	case website := <-chickenChannel:
		fmt.Println("text sent:", website)
	case website := <-tofuChannel:
		fmt.Println("email sent:", website)
	}

}

/*
	we dont need a waitGroup here as we just need to know the first website with prices lower than max_price, and then we can exit the code


	select statement
	select statement is used to wait on multiple channel operations
> it is like a switch statement for channels
> it will block until one of the cases can execute


no need for close(channel) cause we only read once and not over a range

In this code, we spawn 6 go routines, but read from only one
this is not an issue here cause the main function exits and kills are routines, but
if the code was such that main function did not die, it would cause memory consumption
*/
