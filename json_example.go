package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	
	type Person struct {
		name string
		address string
		phone string
	}

	p1 := Person{"Joe", "A St", "123"}
	var p2 Person
		
	//barr, err := json.Marshal(p1);

	/*if byteArray, err := json.Marshal(p1); err == nil {
		fmt.Println("Created JSON successfuly.")
		fmt.Println(byteArray, err)
		} else {
			fmt.Println("Error while creating JSON." , err)
		}*/

	byteArray, err1 := json.Marshal(&p1);
	err2 := json.Unmarshal(byteArray, &p2)
	fmt.Println(byteArray, p1, p2, err1, err2)

	//b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	//fmt.Println(b)


}