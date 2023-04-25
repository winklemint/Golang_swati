package funny

import "fmt"

func Contra() (k int) {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d\n", i)
	}
	return k
}
