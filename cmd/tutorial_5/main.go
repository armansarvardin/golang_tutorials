package main

import (
	"fmt"
)

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh 
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{ 25, 15, owner{"Alex"} }
	var myEngine2 = struct {
		mpg uint8
		gallons uint8
	} {25, 15}
	fmt.Println(myEngine2)
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())

	var myElectricEngine = electricEngine {25, 15}
	canMakeIt(myElectricEngine, 50)
}