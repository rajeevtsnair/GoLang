/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line. 
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. 
Each field will be a string of size 20 (characters).Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadUserInput(message string) (string, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	text, err := reader.ReadString('\n')

	return text, err
}

func main() {

	type Person struct {
		fname string
		lname string
	}
	var p1 Person
		
	filename, err := ReadUserInput("Enter the file name: ")

	if err != nil {
		fmt.Println("Input error:", err)
	}else{
		filename = strings.ReplaceAll(filename, "\n", "")
		fmt.Println("Entered file name:", filename)
	}
	
	slice:= make([]Person,0) 

	if fh, err := os.Open(filename); err != nil {
		fmt.Println("Error while opening the file." , filename,  err)
		} else {
		scanner := bufio.NewScanner(fh)
        for scanner.Scan() {
			name := strings.Split(scanner.Text(), " ")
			p1.fname = name[0]
			p1.lname = name[1]
			slice = append(slice, p1)
   		 }

	}

	for _, person := range slice{
		fmt.Println("fname: ", person.fname, "lname: ", person.lname)
	}
	
}