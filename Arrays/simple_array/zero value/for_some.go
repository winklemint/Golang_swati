package simple

import "fmt"

func Sum() {
	a := [...]float64{12.3, 45.78, 19, 27, 89.07}
	sum := float64(0)
	for i, v := range a { //returns both index value
		fmt.Printf("%d element of a is %.2f", i, v)
		sum += v
	}
	fmt.Println(sum)
}
