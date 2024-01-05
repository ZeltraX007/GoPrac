// ------------- Function ------------

package main

import (
	"errors"
	"fmt")

func main(){
	var printValue string = "Hello World!"
	printMe(printValue)
}

func printMe(printValue string){
	fmt.Println(printValue)
	var result, remainder, err = intDivision(0,0)
	// fmt.Println(result, remainder)
	switch{
	case err!=nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("\nThe result of the integer division is %v",result)
	default:
		fmt.Printf("\nThe result of the integer division is %v with remainder %v", result, remainder)
	}
	switch remainder{
	case 0:
		fmt.Printf("\nThe division was exact")
	case 1,2:
		fmt.Printf("\nThe division was close")
	default:
		fmt.Printf("\nThe division was not close")
	}
}

func intDivision(numerator int, denominator int) (int, int, error) { //int is return type
	var err error
	if denominator == 0{
		err = errors.New("Cannot divide by Zero")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}