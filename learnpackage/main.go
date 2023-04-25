package main

import (
	"fmt"
	"learnpackage/SimpleInterest"
	"log"
)

var p, r, t = 500.09, 10.0, 2.0

func init() {
	println("Main package initialised")
	if p < 0 {
		log.Fatal("Principal amount is less than zero")
	}
	if r < 0 {
		log.Fatal("rate is less than zero")
	}
	if t < 0 {
		log.Fatal("time is less than zero")
	}
}
func main() {
	fmt.Println("Simple interest calculation")
	si := SimpleInterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
