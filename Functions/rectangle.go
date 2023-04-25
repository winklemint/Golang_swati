package main

import "fmt"

func rectProp(length, width int) (int, int) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
func main() {
	area, perimeter := rectProp(5, 20)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)
}
