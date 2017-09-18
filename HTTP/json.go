package main

import "fmt"
import "encoding/json"

type sourceStruct struct {
	
}

func main() {
	// Setup our struct
	var attributes map[string]interface{} // interface{} is an empty interface, lets us use anything

	attributes = make(map[string]interface{})

	attributes["A"] = "fugg"
	attributes["B"] = "kek"
	attributes["wildcard"] = 1337
	attributes["nest"] = make(map[string]int)
	attributes["nest"]["1"] = 420

	fmt.Println(attributes)

	//Encode that shit
	str, _ := json.Marshal(attributes)

	fmt.Println(string(str))
}