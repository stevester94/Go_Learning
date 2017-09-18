package main

import "fmt"

const c int = 9001

func main() {
	var s string = "Jej"
	fmt.Println(s)


	var i int // Will be initialized to its zero value
	fmt.Println("Uninitialzied value: ", i) 

	i = 0
	fmt.Println("Initialized: ", i)

	auto_guessed_var := 420
	fmt.Println(auto_guessed_var)

	fmt.Println(auto_guessed_var + c)
}