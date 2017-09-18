package main

import "fmt"
import "time"


func main() {
	/*
	 * Switches don't bleed through like C
	 * Can be used just as an alternative to if-else
	 */

	// Looks like functions are first class citizens
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
	        case bool:
	            fmt.Println("I'm a bool")
	        case int:
	            fmt.Println("I'm an int")
	        default:
	            fmt.Printf("Don't know type %T\n", t)
        }
    }

    whatAmI(true)
    whatAmI(1337)
    whatAmI(time.Now())

    switch time.Now() {
	    case time.Saturday, time.Sunday:
	    	fmt.Println("It's the weekend")
	    default:
	    	fmt.Println("It's not the weekend...")
    }

    // Using as an if else
    t := time.Now()
    switch {
    	case t.Hour() > 12:
    		fmt.Println("It's past noon")
		default:
			fmt.Println("It's before noon")
    }



}