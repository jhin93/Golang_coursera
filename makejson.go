package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Create map, name, address
	nameAddr := make(map[string]string)
	var inputName string
	var inputAddr string

	// Scan input from terminal
	fmt.Println("Type your name.")
	fmt.Scan(&inputName)
	fmt.Println("Type your address.")
	fmt.Scan(&inputAddr)

	// Put input value into map
	nameAddr["name"] = inputName
	nameAddr["address"] = inputAddr

	jsonData, err := json.Marshal(nameAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Use Marshal() to convert map to json
	fmt.Println(string(jsonData))
}
