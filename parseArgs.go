package main

import (
	"os"
	"strconv"
)

type args struct {
	numPizzas   int
	numStations int
	verbose bool
}

func parseCommandLineArgs() args {
	numStations := 1 //defaults
	numPizzas := 1
	verbose := false
	argsWithoutProg := os.Args[1:]
	pr(argsWithoutProg)
	for i := range argsWithoutProg {
		switch argsWithoutProg[i] {
		case "-station":
			numStations, _ = strconv.Atoi(argsWithoutProg[i+1])
		case "-s":
			numStations, _ = strconv.Atoi(argsWithoutProg[i+1])
		case "-pizza":
			numPizzas, _ = strconv.Atoi(argsWithoutProg[i+1])
		case "-p":
			numPizzas, _ = strconv.Atoi(argsWithoutProg[i+1])
		case "-v":
			v := argsWithoutProg[i+1]
			if v == "Y" || v == "YES" {
				verbose = true
			}
		}
	}
	pr("numStations:", numStations, "numPizzas:", numPizzas, "verbose:",verbose )
	a := args{numPizzas, numStations, verbose }
	return a
}
