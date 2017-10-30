package main

import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "encoding/json"
import "bytes"

const API_Key string = "f44bfd10a8cfed257d7dc6cb1ba18038"
const Token string = "90f6221553f46fb8bcedd905e3f74cd7e13327336c7f18ec6754620bf520627f"
const Card_ID string = "59f6650a357ae208592c00f9"
const Endpoint_format_string string = "https://api.trello.com/1/cards/%s?key=%s&token=%s" // Card_ID, key, token
const New_Name string = "SHIDD"



func main() {
    request_url := fmt.Sprintf(Endpoint_format_string, Card_ID, API_Key, Token)

    fmt.Printf("request_url: %s\n\n", request_url)

    payload_map := make(map[string]interface{})
    var payload_bytes []byte

    payload_map["name"] = New_Name

    payload_bytes, _ = json.Marshal(payload_map)

    fmt.Printf("payload_bytes: %s\n\n", payload_bytes)


    // PUT is not easily accessed! Need to use the client interface of http lib...
	resp, err := http.Put(request_url, "application/json", bytes.NewBuffer(payload_bytes))
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
}
