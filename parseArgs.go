package main

import (
	"os"
	"strconv"
)

type args struct {
	numPizzas   int
	numStations int
	incr int
	max int
	verbose bool
}

func parseCommandLineArgs() args {
	numChefs := 3   //defaults
	numPizzas := 10
	incr := 10
	max := 1000
	verbose := false
	argsWithoutProg := os.Args[1:]
	pr(argsWithoutProg)
	for i := range argsWithoutProg {
		switch argsWithoutProg[i] {
		case "-chef":
			numChefs, _ = strconv.Atoi(argsWithoutProg[i+1])
		case "-pizza":
			numPizzas, _ = strconv.Atoi(argsWithoutProg[i+1])
		case "-incr":
			incr, _ = strconv.Atoi(argsWithoutProg[i+1])
		case "-max":
			max, _ = strconv.Atoi(argsWithoutProg[i+1])
		case "-v":
			v := argsWithoutProg[i+1]
			if v == "Y" || v == "YES" {
				verbose = true
			}
		}
	}
	pr("numStations:", numChefs, "numPizzas:", numPizzas, "incr:", incr, "max:", max, "verbose:",verbose )
	a := args{numPizzas, numChefs, incr, max, verbose }
	return a
}
