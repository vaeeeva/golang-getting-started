package main

import (
	"fmt"
	"os"
	"strconv"
	"sort"				
)

const exitValue ="X"

func main() {
	fmt.Println("starting assignment - slice.go")
	slice := make([]int,3)
	var input string
	for {
		fmt.Println("Insert an integer number")
		_, err := fmt.Scanf("%s",&input)
		if err == nil {
			if input == exitValue {
				break						
			}
			i, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("incorrect input")					
				continue
			}
			slice = append(slice, i)
			sort.Ints(slice)
			fmt.Printf("Slice is composed of: %v \n", slice)
			continue
		}
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	os.Exit(0)
}
