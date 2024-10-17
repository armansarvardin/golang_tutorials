package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value if i is: %v\n", i)

	p = &i
	*p = 1

	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", *p)

	var k int32 = 2
	i = k

	slicesTutorial()

	functionsTutorial()
}

func slicesTutorial() {

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice

	sliceCopy[2] = 4
	
	// arrays are reference types
	fmt.Println(slice) // [1, 2, 4]
	fmt.Println(sliceCopy) // [1, 2 ,4]
}

func functionsTutorial() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v\n", result)
	fmt.Printf("\nThe value of thing1 is %v", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}