package main

import (
	"math/rand"
	"encoding/json"
	"time"
	"strings"
	"strconv"
)

func getRandomDuration() time.Duration {
	return time.Duration(rand.Intn(100))
}
func timeTrack(start time.Time, name string, numPizzas int, numStations int, cookingTime int64) {
	elapsed := time.Since(start)
	elapsedMs := int64(elapsed)
	cookingTimeMs := cookingTime/1000000
	cookingPercentage := (cookingTime * 100)/elapsedMs
	processingPercentage := ((elapsedMs - cookingTime) * 100)/elapsedMs
	pf("%s Stations:%d Pizzas:%d Duration:%s CookingTime:%d (CookingTime:%.2f%% GoRoutines:%.2f%%)\n",
		name, numStations, numPizzas, elapsed, cookingTimeMs,
		float64(cookingPercentage),
		float64(processingPercentage))
}
func addDurationToOrder(s string, d time.Duration) string {
	pr("addDurationToOrder", s, d)
	o := convertStringToOrder(s)
	var ms1 = int64(d)
	ms2, _ := strconv.ParseInt(o.duration, 10, 64)
	t := ms1+ms2
	o.duration = strconv.FormatInt(t, 10)
	s2 := convertOrderToString(o)
	pr("addDurationToOrder", s2)
	return s2
}
func convertStringToOrder(s string) order {
	result := strings.Split(s, "|")
	o := order {result[0], result[1], result[2], result[3], result[4]}
	return o
}
func convertOrderToString(o order) string {
	s := o.name+"|"+o.dough+"|"+o.sauce+"|"+o.topping+"|"+o.duration
	return s
}

func prettyJsonPrint (a interface{}) {
	out, err := json.Marshal(a)
	if err != nil {
		panic (err)
	}
	pr(string(out))
}