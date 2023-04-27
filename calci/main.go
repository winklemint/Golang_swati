package main

import (
	calci "calci/struct_method"
	"fmt"
)

func main() {
	a := calci.Calc{
		Num1: 2,
		Num2: 2,
	}
	fmt.Printf("Addition of 2 numbers is %d", a.Add())

	s := calci.Calc{
		Num1: 4,
		Num2: 10,
	}
	fmt.Printf("\nSubtraction of 2 numbers is %d", s.Sub())

	m := calci.Calc{
		Num1: 5,
		Num2: 5,
	}
	fmt.Printf("\n Multiplication of 2 numbers is %d", m.Multi())
	d := calci.Calc{
		Num1: 4,
		Num2: 2,
	}
	fmt.Printf("\n Division of 2 numbers is %d", d.Div())

}
