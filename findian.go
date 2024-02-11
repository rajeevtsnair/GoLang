package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {

	var userInput string
	var userInputLow string

	fmt.Println("Enter a string...")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput = scanner.Text()
		break
	}

	//fmt.Println("Entered string...", userInput)
	
	userInputLow = strings.ToLower(userInput)

	if strings.HasPrefix(userInputLow, "i") && strings.HasSuffix(userInputLow, "n") && strings.Contains(userInputLow, "a") { 
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
