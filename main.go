package main

import (
	"time"
	"fmt"
	"math"
)

var cnt = 0
var pr = fmt.Println
var pf = fmt.Printf
var verbose = false
var cookingTime int64 = 0

func main() {
	a := parseCommandLineArgs()
	verbose = a.verbose
	makePizza(a.numPizzas, a.numStations)
}
func makePizza(numPizza int, numStations int) {
	pr("makePizza numPizzas:",numPizza, "numStations:", numStations)
	// CHANNELS allow us to pass data between go routines
	doughChan 	:= make(chan string)
	sauceChan 	:= make(chan string)
	toppingChan := make(chan string)
	done 		:= make(chan int64)

	start := time.Now()
	for i := 0; i < numStations; i++ {
		go makeDough(doughChan, sauceChan)
		go addSauce(sauceChan, toppingChan)
		go addToppings(toppingChan, done)
	}
	for i := 0; i < numPizza; i++{
		cnt++
		order := getOrder(i)
		if verbose && math.Mod(float64(i), 50) == 0{
			pr(i, order)
		}
		doughChan <- order
	}
	defer close (doughChan)
	defer close (sauceChan)
	defer close (toppingChan)

	c := <- done
	timeTrack(start, "makePizza", numPizza, numStations, c)
}
