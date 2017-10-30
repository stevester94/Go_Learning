package main

import "fmt"

func main() {
    //To create an empty map, use the builtin make: make(map[key-type]val-type).

    m := make(map[string]int)

    m["fugg"] = 1337

    fmt.Println(m)

    delete(m, "fugg")

    _, err := m["fugg"]
    fmt.Println(err)

    m["nested_value"] = 420

    m_nest := make(map[string]interface{})
    m_nest["map"] = m

    m2 := make(map[string]interface{})

    m2["int"] = 1337
    m2["string"] = "lel"
    m2["map"] = m_nest

    fmt.Println(m2)

}
