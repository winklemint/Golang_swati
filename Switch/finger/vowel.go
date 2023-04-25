package fingers

import "fmt"

func Vowels() (j string) {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")

	}
	return j

}
