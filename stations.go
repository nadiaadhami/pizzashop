package main

import (
	"time"
)

func makeDough(fromChannel  chan string, toChannel chan string) {
	for order := range fromChannel{
		o := convertStringToOrder(order)
		if verbose {
			pr(o.name,"Make",o.dough,"Dough and Send for Sauce")
		};
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(order, d)
		toChannel <- order
	}
}
func addSauce(fromChannel  chan string, toChannel chan string){
	for order := range fromChannel{
		o := convertStringToOrder(order)
		if verbose {
			pr(o.name,"Add",o.sauce,"Sauce and Send for Topping")
		};
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(order, d)
		toChannel <- order
	}
}
func addToppings(fromChannel  chan string, toChannel chan int64){
	for order := range fromChannel{
		o := convertStringToOrder(order)
		if verbose {
			pr(o.name,"Add ",o.topping,"Topping",)
		};
		d := time.Millisecond * getRandomDuration()
		time.Sleep(d)
		addDurationToOrder(order, d)
		var ms = int64(d)
		cookingTime += ms
		pr("cookingTime", cookingTime)
		cnt--
		if (cnt <= 0) {
			if verbose {
				pr("done!")
			}
			toChannel <- cookingTime
		}
	}
}
