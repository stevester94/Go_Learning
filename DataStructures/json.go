/*
 * This demonstrates marshalling and demarshalling JSON objects, as well as more examples of advanced map usage
 */

package main

import "fmt"
import "encoding/json"


func main() {
	// Setup our struct
	var attributes map[string]interface{} // interface{} is an empty interface, lets us use anything

	attributes = make(map[string]interface{})

	attributes["Header"] = "fugg"
	attributes["Wildcard"] = 1337

    body := make(map[string]interface{})
    body["0"] = 1337
    body["1"] = "This is body nested field"

	attributes["Body"] = body


	//Encode that shit
	payload, _ := json.Marshal(attributes)

	fmt.Println("JSON encoding:", string(payload))


    // Let's decode that shit
    decoded := make(map[string]interface{})

    json.Unmarshal([]byte(payload), &decoded)

    fmt.Println("decoded json map:", decoded)

    header, ok := decoded["Header"].(string) 
    if !ok {
        panic("Header cast to string is not ok!")
    }

    in_body, ok := decoded["Body"].(map[string]interface{})
    if !ok {
        panic("Body cast to map of interfaces not ok!")
    }

    fmt.Println("Header:", header)
    fmt.Println("Body:", in_body)
}
