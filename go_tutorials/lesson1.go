package main 	//special
import "fmt"
import "unicode/utf8"

func main(){
	fmt.Println("Hello World!")
	var intNum int16 = 32767
	// var intNum int16 = 32767 + 1
	intNum = intNum + 1 	//-32768
	// uint8: (0, 255)
	// int8: (-128, 127)
	fmt.Println(intNum)

	// No float type either float32 or float64
	var floatNum float32 = 12345678.9
	var floatNum1 float64 = 12345678.9

	fmt.Println(floatNum)
	fmt.Println(floatNum1)
	
	fmt.Printf("%f\n",floatNum1)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString1 string = "Hello World"
	var myString2 string = `Hello
	World`
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)
	fmt.Println(myString1)
	fmt.Println(myString2)

	fmt.Println(len("A")) // 1
	fmt.Println(len("€")) // 3

	fmt.Println(utf8.RuneCountInString("€"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3)

	// assignment

	var myVar string = "text"
	myVar1:= "text"
	// using shorthand is not recommendable

	fmt.Println(myVar)
	fmt.Println(myVar1)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"  // cant change its value
	// cant do this
	// const myConst string
	// have to initialize with a value
	const pi float32 = 3.1415
}