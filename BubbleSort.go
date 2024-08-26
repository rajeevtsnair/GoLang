//The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line,
// in sorted order, from least to greatest.

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
// It eventually returns an error, through an error code.

func ReadUserInput(message string) (string, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	text, err := reader.ReadString('\n')

	return text, err
}

//Fuction to swap two elemets of slice.

func  Swap(NumSlice []int, index int) {
	temp := NumSlice[index]
	NumSlice[index] = NumSlice[index + 1] 
	NumSlice[index + 1 ] = temp
}


//Sorting the slice using Bubble Sort.

func BubbleSort(NumSlice []int) {
	for i := 1; i < len(NumSlice); i++ {
		temp := NumSlice[i] // define a temporary variable
		j := i - 1        
		for (j >= 0) && (NumSlice[j] >= temp) {
			Swap(NumSlice, j)
			j = j - 1                 
		}
		NumSlice[j+1] = temp // assign tmp to the next element
	}
}


// Script which prompts the user to enter integers and stores the integers in a slice until user enters 10 intergers.
// The program prints on screen the sorted slice.

func main() {
	slice := make([]int,10) // create an empty integer slice of size (length) 10
	var counter int = 0 

	for counter <  10 {
		if userInput, err := ReadUserInput("Insert an integer: "); err != nil {
			fmt.Println("Input error:", err, ".Please try again.")
			continue
		} else {
			//fmt.Println("Input string:", userInput)
			userInput = strings.ReplaceAll(userInput, "\n", "")
		
			if number, err := strconv.Atoi(userInput); err != nil {
				fmt.Println("Input string is not an integer number. Please try again.")
				continue
			} else {
				fmt.Println("Good! Input string is an integer number: ", number)
				slice[counter] = number
				counter++ // increasing the counter
			}
		}
	}
	fmt.Println("Entered 10 integers, Sorting... ", slice)
	BubbleSort(slice)
	fmt.Println("Sorted intergers ", slice)
}
