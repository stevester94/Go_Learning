// Structs and Methods

package main

import "fmt"

type being interface {
	getAge()
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
func (p person) getAge() {
	fmt.Println(p.name, "Age: ", p.age)
}

func (p *person) appendMackey() {
	(*p).name = (*p).name + " Mackey"
}

func (p cat) getAge() {
	fmt.Println(p.name, "Age: ", p.age)
}

func (p *cat) appendMackey() {
	(*p).name = (*p).name + " Mackey"
}

func getAge(b being) {
	b.getAge()
}


func main() {
	bob := person{"Bob", 18}
	fmt.Println(bob)

	chedd :=cat{"Chedd", 3, "Orange"}

	slice := make([]being, 2)
	slice[0] = chedd
	slice[1] = bob

	for i := 0; i < len(slice); i++ {
		slice[i].getAge()
	}
}