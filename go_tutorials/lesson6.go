// ---------- Pointers ------------
package main

import "fmt"

func main(){
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v\n", i)
	p = &i
	*p = 10
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("\nThe value if i is: %v\n", i)

	// --------- Slice copy behaiviour---------
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// --------- pointers and functions --------------
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
	var result2 [5]float64 = square2(&thing1)
	fmt.Printf("\nThe result is: %v", result2)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

func square(things2 [5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &things2)
	for i := range things2{
		things2[i] = things2[i]*things2[i]
	}
	return things2
}

func square2(things2 *[5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing2 array is: %p",things2)
	for i := range things2{
		things2[i] = things2[i]*things2[i]
	}
	return *things2
}