package simple

import "fmt"

func FOR() {
	a := [...]float64{68.9, 72.3, 45.7, 22, 19}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%dth element of a is %.2f", i, a[i])
	}
}
