package main

import (
	"math/rand"
	"encoding/json"
	"strconv"
	"time"
)

type order struct {
	name string
	dough string
	sauce string
	topping string
	duration time.Duration
}

func (o *order) string() string {
	out, err := json.Marshal(o)
	if err != nil {
		panic (err)
	}
	return(string(out))
}

func getOrder(i int) order {
	o := order {"Pizza"+strconv.Itoa(i), getDough(), getSauce(), getTopping(), 0}
	return o
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
