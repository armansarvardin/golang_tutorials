package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "resume"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i,v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\n The length of 'myString' is %v\n", len(myString))

	var myRunes = []rune("resume")
	var indexedRune = myRunes[0]
	fmt.Printf("%v, %T\n", indexedRune, indexedRune)

	for i,v := range myRunes {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	var strSlice = []string {"a", "r", "m", "a", "n"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v",catStr)
}