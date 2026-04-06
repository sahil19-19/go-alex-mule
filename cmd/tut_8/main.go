// go routines > lightweight threads
// used to achieve concurreny
// different from parrallel execution
// parallelism required multiple cores
// some level of parallelism is achieved by golang even though it is meant to work on single core
package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

// var m = sync.Mutex{} // has lock() and Unlock()
var m = sync.RWMutex{}    // read write mutex > has RLock() and RUnlock() as well
var wg = sync.WaitGroup{} // they are just counters
var dbdata = []string{"db1", "db2", "db3", "db4"}
var results = []string{} //

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbdata); i++ {
		// dbCall(i);
		wg.Add(1) // increments the counter
		go dbCall(i)
		// we write go before the function call in order to execute them concurrently
	}
	wg.Wait()
	// the control waits here until the counter(wg) becomes 0

	fmt.Println("total execuction time >>> ", time.Since(t0))
	fmt.Println(results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Printf("the result from the database is %v \n", dbdata[i])
	// m.Lock()
	// results = append(results,dbdata[i])
	// // writing to the same memory location at the smame time could lead to corrupt memory
	// m.Unlock()
	//	so we use mutex
	//	when a go routine reaches a lock, a check is perfomed whether a lock has already been set by another routine
	//	if yes then the control waits there until the lock is released
	//	it matters a lot where we place the lock and unlock statement
	//	one drawback of mutex is that it completely blocks all routines
	// wg.Done() // this decrements the counter

	// using rwmutex
	save(dbdata[i])
	log()
	wg.Done() // decrements the wg counter
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("the current results are: %v\n", results)
	m.RUnlock()
}

/*
 many goroutines can read at the same time
but only one can write at a time
 if we use RLock() then we can read the data but not write to it, if any go routines has Rlock() it will stop any other routine to take Lock()
 >> more than one routines can hold Read lock at the same time
 if we use Lock() then we can write to the data but not read it
many goroutines can hold a read lock at the same time
read cannot be done while a write lock is held

Parallelism in go
runtime.GOMAXPROCS(n)
If n > 1, Go can run goroutines in parallel
*/
