// Structs and Methods

package main

import "fmt"

type being interface {
	getAge() (int)
}


type person struct {
	name string
	age int
}

type cat struct {
	name string
	age int
	color string
}

// Super weird! This is defined as the receiver
func (p person) getAge() (int) {
	return p.age
}

func (c cat) getAge() (int) {
    return c.age
}

// you can't use an interface as a receiver type, it just doesn't work like that
//func (b being) getBeingAge() (int) {
//    return b.getAge()
//}


// This however is completely fine
func getBeingAge(b being) int {
    return b.getAge()
}

func main() {
    var p person
    p.age = 1337

    fmt.Println("Person's age: ", p.getAge())

    // Implicit type conversion
    fmt.Println("A being's age(implicit):", getBeingAge(p))
    fmt.Println("A being's age(explicit):", getBeingAge(being(p)))


    var b being = person{"Steven", 23}

    // b is just an interface, but the underlying type is an actual person, we just need to use a type assertion to treat it as such
    fmt.Println("Just another being's age: ", b.getAge())
    fmt.Println("   However this being happens to be: ", b.(person).name)

    var beingSlice []being = make([]being, 0)
    beingSlice = append(beingSlice, person{"Steven", 23})
    beingSlice = append(beingSlice, cat{"Chedd", 3, "Orange"})

    // Wow, we use an empty interface kinda like a place holder, cool
    var WhatAmI interface{} = person{"DoppleGangarSteve!", -1}
    beingSlice = append(beingSlice, WhatAmI.(person))

    for _, b = range beingSlice {
        fmt.Println("Slice of being's age:", b.getAge())
    }


}














