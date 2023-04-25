package simple

import "fmt"

func Change(num [5]int) {
	num[0] = 55
	fmt.Println("inside function", num)
}
func main() {
	num := [...]int{1, 5, 8, 8, 7}
	fmt.Println("before", num)
	Change(num)
	fmt.Println("after", num)
}
