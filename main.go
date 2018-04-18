package main

import (
	"time"
	"fmt"
)

var cnt = 0
var pr = fmt.Println
var pf = fmt.Printf
var verbose = true
var cookingTime time.Duration = 0

func main() {
	a := parseCommandLineArgs()
	verbose = a.verbose
	makePizza(a.numPizzas, a.numStations)
}
func makePizza(numPizza int, numStations int) {
	pr("makePizza numPizzas:",numPizza, "numStations:", numStations)
	// CHANNELS allow us to pass data between go routines
	doughChan 	:= make(chan order)
	sauceChan 	:= make(chan order)
	toppingChan := make(chan order)
	done 		:= make(chan bool)

	start := time.Now()
	for i := 0; i < numStations; i++ {
		go makeDough(doughChan, sauceChan)
		go addSauce(sauceChan, toppingChan)
		go addToppings(toppingChan, done)
	}
	for i := 0; i < numPizza; i++{
		cnt++
		order := getOrder(i)
		if verbose {
			pr(i, order)
		}
		doughChan <- order
	}
	//defer close (doughChan)
	//defer close (sauceChan)
	//defer close (toppingChan)

	timeTrack(start, numPizza, numStations, cookingTime)
}
