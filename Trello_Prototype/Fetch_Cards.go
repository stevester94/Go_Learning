package main

import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "encoding/json"

const API_Key string = "f44bfd10a8cfed257d7dc6cb1ba18038"
const Token string = "90f6221553f46fb8bcedd905e3f74cd7e13327336c7f18ec6754620bf520627f"
const Board_ID string = "db6BDqSJ"
const Endpoint_format_string string = "https://api.trello.com/1/boards/%s/cards?key=%s&token=%s" // Board_ID, key, token


func deJsonifyCardArray(s []byte) {
	var dat = make([]interface{}, 50)

	if err := json.Unmarshal(s, &dat); err != nil {
        panic(err)
    }

    fmt.Println("All fields:")
    for  key, _ := range dat[0].(map[string]interface{}) {
        fmt.Printf("%s ", key)
    }

    fmt.Println()

    fmt.Println("Index, name, dateLastActivity")
    for index, element := range dat {
        card_map := element.(map[string]interface{})
        fmt.Printf("%d: %s, %s\n", index, card_map["name"], card_map["dateLastActivity"])
    }
}


func main() {
    request_url := fmt.Sprintf(Endpoint_format_string, Board_ID, API_Key, Token)

    fmt.Printf("request_url: %s\n\n", request_url)

	resp, err := http.Get(request_url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close() // Believe defer is some kinf of synchronization thing


    fmt.Printf("resp:%s\n\n", resp)

	contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
	    fmt.Printf("%s", err)
	    os.Exit(1)
    }

    fmt.Printf("contents:%s\n\n", string(contents))

    deJsonifyCardArray(contents)

}
