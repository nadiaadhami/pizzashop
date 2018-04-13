package main

import (
	"math/rand"
	"strings"
	"encoding/json"
)

type order struct {
	name string
	dough string
	sauce string
	topping string
}

func (o *order) string() string {
	out, err := json.Marshal(o)
	if err != nil {
		panic (err)
	}
	return(string(out))
}

// create pizza orders with random dough, sauce, and topping
func getOrder() string {
	return getDough() + "|" + getSauce() + "|" + getTopping()
}
func getSauce() string {
	r := rand.Intn(3)
	switch r {
	case 0: return "Marinara"
	case 1: return "Tapenade"
	case 2:return "Pesto"
	default: panic(r)
	}
}
func getTopping() string {
	r := rand.Intn(3)
	switch r {
	case 0: return "Mushrooms"
	case 1: return "Olives"
	case 2: return "Tomatoes"
	default: panic(r)
	}
}
func getDough() string {
	r := rand.Intn(3)
	switch r {
	case 0: return "Focaccia"
	case 1: return "Flatbread"
	case 2: return "Neapolitan"
	default: panic(r)
	}
}
func parseOrder(s string) order {
	result := strings.Split(s, "|")
	o := order {result[0], result[1], result[2], result[3]}
	return o
}
