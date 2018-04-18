package main

import (
	"time"
)

func makeDough(fromChannel  chan order, toChannel chan order) {
	for order := range fromChannel{
		if verbose {
			pr(order, "Make Dough and send for Sauce")
		}
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(&order, d)
		toChannel <- order
	}
}
func addSauce(fromChannel  chan order, toChannel chan order){
	for order := range fromChannel{
		if verbose {
			pr(order,"Add Sauce and Send for Topping")
		}
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(&order, d)
		toChannel <- order
	}
}
func addToppings(fromChannel  chan order, toChannel chan bool){
	for order := range fromChannel{
		if verbose {
			pr(order,"Add Topping")
		}
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(&order, d)

		cookingTime = cookingTime + order.duration
		pr("====== cookingTime", cookingTime)
		cnt--
		if cnt <= 0 {
			if verbose {
				pr("done!")
			}
			toChannel <- true
		}
	}
}
