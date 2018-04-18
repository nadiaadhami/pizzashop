package main

import (
	"testing"
)

func TestPizzaShop(t *testing.T) {
	p := 5
	s := 1
	a := args{p, s, true}
	makePizza(a)
}