package main

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func main() {
	//Create a slice of size 3
	slice := make([]int,3) 
	//Initialize other variables.
	var counter int = 0
	var userInput string
	var userInputLow string

	for userInput != "X" {
		fmt.Println("Enter an integer. X to quit..")
		fmt.Scan(&userInput)
		//Convert the input to lower case
		userInputLow = strings.ToLower(userInput)
		
		//If user entered x exit the program
		if userInputLow == "x" {
			fmt.Println("Entered X. Exiting..")
			break
		} else {
			fmt.Println("Entered number: ", userInputLow)
			//Prompt for a new number if the user input is not integer
			if number, err := strconv.Atoi(userInputLow); err != nil {
				fmt.Println("Entered number is not an integer.")
				continue
			} else {
				//If the counter is less than 3. Add number to the slice
				if counter < 3 {
					slice[counter] = number
					counter++
					fmt.Println("Unsorted slice",slice)
					continue
				} 
				slice = append(slice, number)
				//sort the slice
				sort.Ints(slice)
				fmt.Println("Sorted slice",slice)			
				
			}
		}

	}
}
