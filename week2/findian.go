package main


import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	fmt.Println("Starting assignment - findian.go ...")
	fmt.Print("Insert a String: ")
    var inputString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString = scanner.Text()
	err := scanner.Err()
	if err == nil {
		lenghtString := len(inputString)
		if lenghtString >= 3 {
			inputString = strings.TrimSpace(strings.ToLower(inputString))
			if strings.HasPrefix(inputString, "i") && strings.HasSuffix(inputString, "n") && strings.Contains(inputString, "a") {
				fmt.Println("Found!")
		        os.Exit(0)
			}
		}
	    fmt.Println("Not Found!")
        // Exit successfully
		os.Exit(0)
	}
	fmt.Println("Error: ", err)
	os.Exit(1)
}

