package main

import (
	"First_test/maths"
	"fmt"
)

func main() {
	a := maths.Add(6, 7)
	fmt.Println(a)
	b := maths.Sub(5, 9)
	fmt.Println(b)
	c := maths.Multiply(2, 2)
	fmt.Println(c)
	d := maths.Div(99.12, 11.12)
	fmt.Printf("%f", d)

}
