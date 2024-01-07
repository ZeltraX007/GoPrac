// ----------- Go routines ------------
package main

import (
	"fmt"
	// "math/rand"
	"time"
	"sync"		//wait groups
)

var m = sync.RWMutex{}	// mutual exclusion
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		// dbCall(i) // takes 5 seconds to call sequentially
		wg.Add(1) 	//	Add a task
		go dbCall(i)
	}
	wg.Wait()		//	wait for the counter to go back down to zero
					// ie all the tasks have been completed
	fmt.Printf("\nTotal execution time: %v",time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int){
	var delay float32 = 2000	
	// if we remove this we can cause corrupt memory
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ",dbData[i])
	save(dbData[i])
	log()
	
	wg.Done()		//decrements the counter, done with the task
}

func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}