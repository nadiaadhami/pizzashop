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
var numPizzas = 0

func main() {
	a := parseCommandLineArgs()
	verbose = a.verbose
	numPizzas = a.numPizzas
	makePizza(a)
}
func makePizza(a args) {
	pr("makePizza numPizzas:",a.numPizzas, "numStations:", a.numStations)
	// CHANNELS allow us to pass data between go routines
	doughChan 	:= make(chan order)
	sauceChan 	:= make(chan order)
	toppingChan := make(chan order)
	done 		:= make(chan bool)

	start := time.Now()
	for i := 0; i < a.numStations; i++ {
		go makeDough(doughChan, sauceChan)
		go addSauce(sauceChan, toppingChan)
		go addToppings(toppingChan, done)
	}
	for i := 0; i < a.numPizzas; i++{
		order := getOrder(i)
		if verbose {
			pr(i, order)
		}
		doughChan <- order
	}
	//todo
	//defer close (doughChan)
	//defer close (sauceChan)
	//defer close (toppingChan)
	 <- done

	timeTrack(start, a.numPizzas, a.numStations, cookingTime)
}
