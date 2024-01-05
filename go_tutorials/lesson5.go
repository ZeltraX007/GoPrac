// ---------- Structures ----------

package main

import "fmt"

type gasEngine struct {
	mpg     uint16
	gallons uint16
	ownerInfo owner
}

type electricEngine struct{
	mpkwh uint16
	kwh uint16
}

type owner struct{
	name string
}

// ---------- Interface -----------

type engine interface{
	milesLeft() uint16	//signature statement
	// takes anything only requirement is that the object has milesleft()
}

func (e gasEngine) milesLeft() uint16{
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint16{
	return e.kwh*e.mpkwh
}

func canMakeIt(e engine, miles uint16){
	if miles<=e.milesLeft(){
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{"Alex"}}
	var myEngine3 electricEngine = electricEngine{30,40}
	// var myEngine gasEngine = gasEngine{25,15}
	var myEngine2 = struct{
		mpg uint8
		gallons uint8
	}{25,15}	// not reusable
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())
	canMakeIt(myEngine3,50)
	canMakeIt(myEngine, 1000)
}