package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

func main(){

	message := []byte(`Hello, World`) //Create a byte array message


	//creates a file test_file.txt with 755 permission and writes contents of message
	if werr := ioutil.WriteFile("out_file1.txt", message, 0755); werr == nil {
		fmt.Println("Created the file successfully")
		}else{
			fmt.Println(werr)
	}
		
	
	var fruits interface{} //Create a generic data type

	//Reads the file fruits.json as byteArray and stores in interface
	if data, rerr := ioutil.ReadFile("fruits.json"); rerr == nil {
		fmt.Println("Read the file successfully")
		fmt.Println(data)
		err := json.Unmarshal(data, &fruits)
		fmt.Println(fruits,err)
		} else {
			fmt.Println("Error while reading file." , rerr)
	}

	if fh1, err1 := os.Open("fruits.json"); err1 == nil {
		byteArray := make([]byte, 20)
		nb, err := fh1.Read(byteArray)
		fh1.Close()
		fmt.Println(nb, byteArray, err)
		} else {
			fmt.Println("Error while reading the file." , err1)
	}
		
	if fh2, err2 := os.Create("out_file2.txt"); err2 == nil {
		byteArray := []byte{1, 2, 3, 4, 5}
		nb1, err1 := fh2.Write(byteArray)
		nb2, err2 := fh2.WriteString("End")
		fh2.Close()
		fmt.Println(nb1, byteArray, err1)
		fmt.Println(nb2, err2)
		} else {
			fmt.Println("Error while wirting the file." , err2)
	}
}
