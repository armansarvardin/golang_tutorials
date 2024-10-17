package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contactInfo struct {
	Name string
	Email string
}

type purchaseInfo struct {
	Name string
	Price float32
	Amount int
}

func main() {
	var intSlice []int32 = []int32{1, 2, 3, 4}
	fmt.Println(intSlice)

	var float32Slice []float32 = []float32{1, 2, 3}
	fmt.Println(float32Slice)

	fmt.Printf("\n The sum of ints is: %v", sumSlice(float32Slice))
	fmt.Printf("\n The sum of floats is: %v", sumSlice(float32Slice))
}

func loadJSON[T contactInfo | purchaseInfo] (filepath string) []T {
	data, _ := ioutil.ReadFile(filepath)
	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}