package main

import "fmt"



func main() {
	/* 
	 * for loops
	 */

	// the classic
	for j := 0; j < 3; j++ {
		fmt.Println("j: ", j)
	}

	// acting like a while
	i := 0
	for i <= 3 {
		fmt.Println("i:", i)
		i++
	}

	for {
		fmt.Println("Would normally run forever...")
		break
	}

	/*
	 * Conditionals
	 */
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    // Num is scoped to just those conditional blocks
    // fmt.Println(num) //  Error undefined
}