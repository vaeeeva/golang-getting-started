package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)
type Person struct {
	firstName string
	lastName string
}

func printError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
func main(){
	fmt.Println("Starting assignment - read.go")
	fmt.Print("Insert the file containings names input: ")
	var fileName string
	var lines []string
	var person Person
	var persons []Person
	_, err := fmt.Scanf("%s",&fileName)

	if err == nil {
		textFile,err := ioutil.ReadFile(fileName)
		printError(err)
		lines = strings.Split(string(textFile), "\n")
		for _,line := range lines{
			fields := strings.Fields(string(line))
			if len(fields)<2 {
					break
				}
			person.firstName = fields[0]
			person.lastName = fields[1]

			persons = append(persons,person)						
		}
		fmt.Println("name list from file: ", fileName)
		 
		for _,person := range persons{
			        fmt.Println(person.firstName, person.lastName)
					    
		}
	}
	printError(err)
}
