package main


import (
	"fmt"
	"os"	
)

func main() {
	fmt.Println("Starting assignment - trunc.go ...")
	fmt.Print("Insert a float number: ")
    var inputNumber float32

	_, err :=fmt.Scanf("%f", &inputNumber)

	if err == nil {
		fmt.Printf("Float input: %v is converted in Int: %v \n", inputNumber, int(inputNumber))
        // Exit successfully
		os.Exit(0)
	}

	fmt.Println("Error: ", err)
	os.Exit(1)
}

