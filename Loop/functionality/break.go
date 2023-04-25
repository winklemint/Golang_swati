package funny

import "fmt"

func Breaking() int {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%d\n", i)
	}
	return 1
}
