package main

import (
	"math/rand"
	"encoding/json"
	"time"
)

func getRandomDuration() time.Duration {
	return time.Duration(rand.Intn(1000))
}
func timeTrack(start time.Time, numPizzas int, numStations int, cookingTime time.Duration) {
	elapsed := time.Since(start)
	cookingPercentage := cookingTime/elapsed * 100
	processingPercentage := ((elapsed - cookingTime) * 100 )/elapsed
	pf("timeTrack Stations:%d Pizzas:%d Duration:%v CookingTime:%v (CookingTime:%d%% GoRoutines:%d%%)\n",
		numStations,
		numPizzas,
		elapsed,
		cookingTime,
		cookingPercentage,
		processingPercentage)
}
func addDurationToOrder(o *order, d time.Duration) {
	o.duration = o.duration + d
	//pr("addDurationToOrder", o)
}
func prettyJsonPrint (a interface{}) {
	out, err := json.Marshal(a)
	if err != nil {
		panic (err)
	}
	pr(string(out))
}