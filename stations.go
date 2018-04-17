package main

import (
	"time"
	"math"
	"strconv"
)

func makeDough(fromChannel  chan string, toChannel chan string) {
	for order := range fromChannel{
		if verbose {
			pr(order, "Make Dough and send for Sauce")
		};
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		n := addDurationToOrder(order, d)
		toChannel <- n
	}
}
func addSauce(fromChannel  chan string, toChannel chan string){
	for order := range fromChannel{
		if verbose {
			pr(order,"Add Sauce and Send for Topping")
		};
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		n := addDurationToOrder(order, d)
		toChannel <- n
	}
}
func addToppings(fromChannel  chan string, toChannel chan int64){
	for order := range fromChannel{
		if verbose {
			pr(order,"Add Topping",)
		};
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		n := addDurationToOrder(order, d)
		o := convertStringToOrder(n)
		i, _ := strconv.ParseInt(o.duration, 10, 64)
		cookingTime += i
		pr("cookingTime", cookingTime, order)
		cnt--
		if (cnt <= 0) {
			if verbose {
				pr("done!")
			}
			toChannel <- cookingTime
			return
		} else if (math.Mod(float64(cnt), 100) == 0) {
			pr(order)
		}
	}
}
