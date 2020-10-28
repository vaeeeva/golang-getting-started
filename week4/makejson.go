package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {
	fmt.Println("Starting assignment - makejson.go ...")
	userMap := make(map[string]string)
	fmt.Print("Insert name: ")
	var name string
	_, err :=fmt.Scanf("%s",&name)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	
	fmt.Print("Insert address: ")
	var address string
	_, err =fmt.Scanf("%s",&address)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	userMap["name"] = name
	userMap["address"] = address

	jsonOutput, _ := json.Marshal(userMap)
	fmt.Println(string(jsonOutput))

}
