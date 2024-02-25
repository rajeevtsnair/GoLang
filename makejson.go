/*Write a program which prompts the user to first enter a name, and then enter an address. 
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. 
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
package main

import (
	"bufio"
	"encoding/json"
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

	Usermap := make(map[string]string)

	if name, err := ReadUserInput("Enter the name: "); err != nil {
		fmt.Println("Input error:", err)
	} else {
		name = strings.ReplaceAll(name, "\n", "")
		//fmt.Println("Name: ", name)
		Usermap["name"] = name
	}

	if address, err := ReadUserInput("Enter the address: "); err != nil {
		fmt.Println("Input error:", err)
	} else {
		address = strings.ReplaceAll(address, "\n", "")
		//fmt.Println("address: ", address)
		Usermap["address"] = address
	}

	if jobject, err := json.Marshal(Usermap); err != nil {
		fmt.Println("error in json conversion. ", err)
	} else {
		fmt.Println("json object is ", string(jobject))
	}

}
