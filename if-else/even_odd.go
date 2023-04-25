package main

import "fmt"

func main() {
	if num := 8; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
