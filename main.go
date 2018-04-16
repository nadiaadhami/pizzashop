package main

import (
	"math/rand"
	"strconv"
	"time"
	"fmt"
	"encoding/json"
)

var cnt = 0
var pr = fmt.Println
var pf = fmt.Printf
var verbose = false

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
	done 		:= make(chan bool)

	start := time.Now()
	for i := 0; i < numStations; i++ {
		go makeDough(doughChan, sauceChan)
		go addSauce(sauceChan, toppingChan)
		go addToppings(toppingChan, done)
	}

	for i := 0; i < numPizza; i++{
		cnt++
		order := getOrder()
		if verbose {
			pr(i, order)
		}
		doughChan <- "Pizza"+strconv.Itoa(i)+"|"+order+"|0"
	}
	<- done
	timeTrack(start, "makePizza", numPizza, numStations)
}
func getDuration() time.Duration {
	return time.Duration(rand.Intn(100))
}
func timeTrack(start time.Time, name string, numPizzas int, numStations int) {
	elapsed := time.Since(start)
	pf("%s numStations:%d numPizzas:%d took %s\n", name, numStations, numPizzas, elapsed)
}
func prettyJsonPrint (a interface{}) {
	out, err := json.Marshal(a)
	if err != nil {
		panic (err)
	}
	pr(string(out))
}