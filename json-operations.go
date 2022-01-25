package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type students struct {
	EnNo       string `json:"enrollment"`
	FirstName  string `json:"fName"`
	LastName   string `json:"lName"`
	Department string `json:"department"`
	Age        int    `json:"age"`
}

func readJson() (student []students) {
	fileBytes, err := ioutil.ReadFile("./student.json")
	//var student []students
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &student)
	if err != nil {
		panic(err)
	}
	return student
}

func writeJson() {
	var input students
	fmt.Println("Enter EnNo. First Name, Last Name, Department, Age ")
	_, _ = fmt.Scan(&input.EnNo)
	_, _ = fmt.Scan(&input.FirstName)
	_, _ = fmt.Scan(&input.LastName)
	_, _ = fmt.Scan(&input.Department)
	_, _ = fmt.Scan(&input.Age)
	infoByte, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./student-updated.json", infoByte, 0644)
	if err != nil {
		panic(err)
	}
}

func display() {
	metadata := readJson()
	for _, data := range metadata {
		//goland:noinspection GoPrintFunctions
		fmt.Printf(
			"| Enrollment No: %s | Name: %s %s | Department: %s | Age: %d |\n",
			data.EnNo,
			data.FirstName,
			data.LastName,
			data.Department,
			data.Age,
		)
	}
}
