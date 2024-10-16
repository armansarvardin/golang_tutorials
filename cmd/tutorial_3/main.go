package main

import "fmt"

func main() {

	var intArr [3]int32

	intArr[1] = 123

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr1 [3]int32 = [3]int32{1, 2 ,3}
	fmt.Println(intArr1)

	intArr2 := [...]int32{1, 2, 3}
	fmt.Println(intArr2)

	var intSlice []int = []int{1, 2 ,3}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice2 []int = []int{8,9}

	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int = make([]int, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{
		"Adam": 4, 
		"Sarah": 2,
	}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Jason"]

	delete(myMap2, "Adam")

	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	var i int = 0

	for i < 10 {
		fmt.Println(i)
		i += 1
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}