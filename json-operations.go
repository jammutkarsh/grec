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

const fileLocation = "./.hidden/student.json"

func readJson() (student []students) {
	fileBytes, err := ioutil.ReadFile(fileLocation)
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
	var studentData students
	fmt.Println("Enter EnNo. First Name, Last Name, Department, Age ")
	_, _ = fmt.Scan(&studentData.EnNo)
	_, _ = fmt.Scan(&studentData.FirstName)
	_, _ = fmt.Scan(&studentData.LastName)
	_, _ = fmt.Scan(&studentData.Department)
	_, _ = fmt.Scan(&studentData.Age)
	userData := append(readJson(), studentData)
	infoByte, err := json.Marshal(userData)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileLocation, infoByte, 0644)
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
