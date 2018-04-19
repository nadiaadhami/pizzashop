package main

import (
	"time"
)

func makeDough(fromChannel  chan order, toChannel chan order) {
	for order := range fromChannel{
		wg.Add(1)
		if verbose {
			pr(order, "Make Dough and send for Sauce")
		}
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(&order, d)
		toChannel <- order
		wg.Done()
	}
}
func addSauce(fromChannel  chan order, toChannel chan order){
	for order := range fromChannel{
		wg.Add(1)
		if verbose {
			pr(order,"Add Sauce and Send for Topping")
		}
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(&order, d)
		toChannel <- order
		wg.Done()
	}
}
func addToppings(fromChannel  chan order){
	for order := range fromChannel{
		wg.Add(1)
		if verbose {
			pr(order,"Add Topping")
		}
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(&order, d)

		//todo why is cumulative cooking time more than duration time?
		cookingTime = cookingTime + order.duration // all orders
		//pr("====== cookingTime", cookingTime, order)

		wg.Done()
	}
}
