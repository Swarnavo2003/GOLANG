package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning slice with hello world")

	// numbers := []int{1, 2, 3}
	// numbers = append(numbers, 3,10,12,13,14,15,16,17,18,19)
	// fmt.Println("Numbers :", numbers)
	// fmt.Printf("Number has a data type : %T\n", numbers)
	// fmt.Println("Length :",len(numbers))

	// name := []string{}
	// fmt.Println("name :",name)

	number := make([]int, 3, 5)
	number = append(number, 4)
	number = append(number, 98)
	number = append(number, 6)
	number = append(number, 4)
	number = append(number, 98)
	number = append(number, 6)
	number = append(number, 6)
	number = append(number, 6)

	fmt.Println("Slice :", number)
	fmt.Println("Length :",len(number))
	fmt.Println("Capacity :", cap(number))

	stock := make([]int, 0)
	fmt.Println("Slice :", stock)
	fmt.Println("Length :",len(stock))
	fmt.Println("Capacity :", cap(stock))
}