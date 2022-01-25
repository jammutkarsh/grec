package main

import (
	"fmt"
)

func main() {
	input := -1
	fmt.Println("Enter 1 to add data ")
	fmt.Println("Enter 2 to display all ")
	fmt.Scan(&input)
	switch input {
	case -1:
		break
	case 1:
		writeJson()
	case 2:
		display()
	}
}

//To make a program to
//read and write json file DONE
//parse through json elements and sort them DONE
// to display information in a particular manner
// make a cli based parsing command to retrieve data according to en_no.
