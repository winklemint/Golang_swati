package simple

import "fmt"

func Values() {
	var a [5]int // int array with len 5
	a[0] = 12    //assigning values in array thru index i.e starts from zero
	a[1] = 24
	a[2] = 36
	a[3] = 48
	a[4] = 60
	fmt.Println(a)
}
