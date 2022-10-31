package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for v := range values {
		if v%2 == 0 {
			fmt.Println("Even")
		}
		if v%2 != 0 {
			fmt.Println("Odd")
		}
	}
}
