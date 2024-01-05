package main

import	(
	"fmt"
	"strings"
)

func main(){
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n",indexed, indexed)
	for i, v := range myString{
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of 'myString' is %v",len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s"}
	// var catStr = ""
	var strBuilder strings.Builder
	for i := range strSlice{
		// catStr += strSlice[i]
		strBuilder.WriteString(strSlice[i])	// So builder appends every value
	}
	var catStr = strBuilder.String		// and new string created at the end
	fmt.Printf("\n%v",catStr)
	// Strings are immutable in go
	// cannot modify them once created
	// catStr [0] = "a"
}