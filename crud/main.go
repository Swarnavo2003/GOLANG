package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"user_id"`
	Id 		    int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting :",err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response :", res.Status)
		return
	}

	// data,err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading :",err)
	// 	return
	// }
	// fmt.Println("Data :",string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding :",err)
		return
	}
	fmt.Println("Todo :",todo)
}

func performPostRequest() {
	todo := Todo {
		UserID: 23,
		Title: "Swarnavo Majumder",
		Completed: true,
	}

	// Convert the Todo struct to json
	jsonData,err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling :", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// covert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"
	
	// send POST request
	res,err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request :",err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
}

func performUpdateRequest() {
	todo := Todo {
		UserID: 23789,
		Title: "Swarmavo Majumder Golang hello World",
		Completed: false,
	}

	jsonData,err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling :", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// covert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	// create PUT Request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request :", err)
		return
	}

	req.Header.Set("Content-Type","application/json")

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Response status :", res.Status)
}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating Delete Request :", err)
		return
	}

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request :", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response status :", res.Status)
}

func main() {
	fmt.Println("Learning CRUD...")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}