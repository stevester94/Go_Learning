package main

import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "encoding/json"


func deJsonify(s []byte) {
	var dat map[string]interface{} // Don't understand this at all

 	if err := json.Unmarshal(s, &dat); err != nil {
        panic(err)
    }

    fmt.Println(dat)
}


func main() {
	resp, err := http.Get("http://httpbin.org/html")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close() // Believe defer is some kinf of synchronization thing


    fmt.Println("resp:", resp)

	contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
	    fmt.Printf("%s", err)
	    os.Exit(1)
    }

    fmt.Println("contents:", string(contents))

    deJsonify(contents)

}
