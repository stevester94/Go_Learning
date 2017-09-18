/*
 * -variadic input
   -errors
   -multiple return
   -range
   -slices
 */

package main

import "fmt"
import "errors"

func makeArray(input ...int) ([]int, error) {
	if len(input) == 0 {
		return []int{}, errors.New("Need at least one input")
	}

	retAr := make([]int, 0)
	for _, value := range input {
		retAr = append(retAr, value)
	}

	return retAr, nil
}


func main() {
	slice, err := makeArray()
	if err != nil {
		fmt.Println("Error when making array: ", err)
	}

	slice, err = makeArray(1, 2)
	if err != nil {
		fmt.Println("Error when making array: ", err)
	}

	fmt.Println(slice)
}