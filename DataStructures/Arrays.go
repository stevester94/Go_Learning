package main

import "fmt"

func main() {
	var ar1 [5]int

	for i := 0; i < 5; i++ {
		ar1[i] = i
	}

	fmt.Println("Ar1: ", ar1)

	ar2 := [5]int {0, 1, 2, 3, 4}
	fmt.Println("Ar2: ", ar2)

	fmt.Println("Ar1 == Ar2: ", ar1 == ar2)

}