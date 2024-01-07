// -------- CHANNELS ----------
// channels - way to enable go routines to pass around information

// channels i	- hold data
// channels ii 	- thread safe i.e avoid data races when reading and writing from memory
// channels iii - listen for data i.e we can listen when data is added or removed from a channel during code execution until one of these happens

package main

import "fmt"

func main(){
	var c = make(chan int)
	// c <- 1
	// var i = <- c
	// fmt.Println(i)		//results in deadlock

	go process(c)
	for i:= range c{		// results in deadlock if close is not used
		fmt.Println(i)
	}
	// for i:=0; i<5; i++{
	// 	fmt.Println(<-c)
	// }
	// fmt.Println(<-c)
}

func process(c chan int){
	defer close(c)			//do this right before the function exits
	for i:=0; i<5; i++{
		c <- i
	}
	// c <- 123
}