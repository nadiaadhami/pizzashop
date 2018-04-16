package main

import "time"

func makeDough(fromChannel  chan string, toChannel chan string) {
	for order := range fromChannel{
		o := parseOrder(order)
		if verbose {
			pr(o.name,"Make",o.dough,"Dough and Send for Sauce")
		};
		time.Sleep(time.Millisecond * getDuration())
		toChannel <- order
	}
}
func addSauce(fromChannel  chan string, toChannel chan string){
	for order := range fromChannel{
		o := parseOrder(order)
		if verbose {
			pr(o.name,"Add",o.sauce,"Sauce and Send for Topping")
		};
		s := time.Millisecond * getDuration()
		time.Sleep(s)
		toChannel <- order
	}
}
func addToppings(fromChannel  chan string, toChannel chan bool){
	for order := range fromChannel{
		o := parseOrder(order)
		if verbose {
			pr(o.name,"Add ",o.topping,"Topping",)
		};
		time.Sleep(time.Millisecond * getDuration())
		cnt--
		if (cnt <= 0) {
			if verbose {
				pr("done!")
			}
			toChannel <- true
		}
	}
}
