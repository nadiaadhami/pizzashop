package main

import (
	"time"
	"fmt"
	"sync"
)

var pr = fmt.Println
var pf = fmt.Printf
var cookingTime time.Duration = 0
var wg sync.WaitGroup
var verbose = false

func main() {
	a := parseCommandLineArgs()
	makePizza(a)
}
func makePizza(a args) {
	start := time.Now()

	pr("makePizza numPizzas:",a.numPizzas, "numStations:", a.numStations, "verbose:", a.verbose, start)

	// CHANNELS allow us to pass data between go routines
	doughChan 	:= make(chan order)
	sauceChan 	:= make(chan order)
	toppingChan := make(chan order)

	for i := 0; i < a.numStations; i++ {
		go makeDough(doughChan, sauceChan)
		go addSauce(sauceChan, toppingChan)
		go addToppings(toppingChan)
	}
	for i := 0; i < a.numPizzas; i++{
		order := getOrder(i)
		if a.verbose {
			pr(i, order)
		}
		doughChan <- order
	}
	wg.Wait()

	timeTrack(start, a.numPizzas, a.numStations, cookingTime)
}
