package fingers

import (
	"fmt"
)

func Wedding() (f string) {
	finger := 3
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")

	}
	return f
}
