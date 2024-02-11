/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. 
The program should be written as a loop. Before entering the loop, 
the program should create an empty integer slice of size (length) 3. 
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. 
The program adds the integer to the slice, sorts the slice, 
and prints the contents of the slice in sorted order. 
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"bufio"   
	"fmt"     
	"os"     
	"strconv" 
	"strings" 
)

// ReadUserInput function.
// The functions return a string as read from the console userInput.
// It eventually returns an error, trough an error code.
func ReadUserInput(message string) (string, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	text, err := reader.ReadString('\n')

	return text, err
}

//Sorting the slice
func SliceSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		temp := numbers[i] // define a temporary variable
		j := i - 1        
		for (j >= 0) && (numbers[j] >= temp) {
			numbers[j+1] = numbers[j] 
			j = j - 1                 
		}
		numbers[j+1] = temp // assign tmp to the next element
	}
}


// Script which prompts the user to enter integers and stores
// the integers in a sorted slice.
// The program prints on screen the sorted slice.
// It keeps going until the user enters X to exit.
func main() {
	slice := make([]int, 3) // create an empty integer slice of size (length) 3
	var userInput string
	var userInputLow string
	var counter int

	for userInput != "x" {
		if userInput, err := ReadUserInput("Insert an integer or X to exit: "); err != nil {
			fmt.Println("Input error:", err, ".Please try again.")
			continue
		} else {
			fmt.Println("Input string:", userInput)
			userInput = strings.ReplaceAll(userInput, "\n", "") 

			userInputLow = strings.ToLower(userInput)

			if userInputLow == "x" {
				fmt.Println("X entered, exit")
				break
			}

			if number, err := strconv.Atoi(userInput); err != nil {
				fmt.Println("Input string is not an integer number. Please try again.")
				continue
			} else {
				fmt.Println("Good! Input string is an integer number: ", number)

				if counter < 3 {
					slice[counter] = number
					fmt.Println("Slice not sorted yet:", slice)
					counter++ // increasing the counter
					continue
				}
				slice = append(slice, number)
				SliceSort(slice)
				fmt.Println("Slice in the (increasing) order:", slice)

				counter++ // increasing the counter
			}
		}
	}
}
