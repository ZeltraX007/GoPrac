//Arrays

package main
import "fmt"

func main(){
	var intArr[3] int32 	// default [0] [0] [0]
	// var int [3]int32 = [3]int32{1,2,3}
	// intArr := [3]int32{1,2,3}
	// intArr := [...]int32{1,2,3}
	// int32 hold 4 bytes so 
	// go allocates 12 bytes of contigous memory
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	// fmt.Println(&intArr[3])


	//---------- Slices -----------
	// slices are arrays with additional functionalities
	// by ommiting the length value we have a slice
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// unlike arrays we can add values to the slice by append()
	intSlice = append(intSlice, 7)
	// creates a new array
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	// spread operator for multiple 
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("\n%d",intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
	// can optionally specify the length if we have a rough idea

	// ----------- Map -----------
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])	//default value 0
	var age, ok = myMap2["Jason"]	// return a boolean true/ false found or not
	if ok{
		fmt.Print("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}
	delete(myMap2, "Adam")

	// -------------- LOOPS ------------

	for name, age:= range myMap2{
		fmt.Printf("Name: %v, Age:%v \n", name,age)		//No ordering
	}

	for i, v := range intArr{
		fmt.Printf("Index: %v, Value: %v \n",i,v)
	}

	// while loops
	var i int = 0
	for i < 10{
		fmt.Println(i)
		i = i + 1
	}

	// for {
	//	if i < 10{
	//		break
	//	}
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
}