package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web services")
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response",err)
	}
	defer res.Body.Close()
	fmt.Printf("Type of response : %T\n",res)

	// Read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response ",err)
		return
	}

	fmt.Println("response :",string(data))
}