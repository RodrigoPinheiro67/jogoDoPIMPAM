package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Pim")
		}
		if i%5 == 0 && i%3 != 0 {
			fmt.Println("Pam")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pim Pam")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}
	}
}
