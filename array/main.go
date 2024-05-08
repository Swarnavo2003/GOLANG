package main

import "fmt"

func main() {
	fmt.Println("We are learning Array in Golang")

	// var name[5]string
	// name[0] = "swarnavo"
	// name[2] = "Aakash"
	// fmt.Println("Names of Person is :", name)

	// var numbers = [8]int{1,2,3,4,5}
	// fmt.Println("Number is :", numbers)

	// fmt.Println("Length of Numbers array is :", len(numbers))

	// fmt.Println("value of name at 2 index is :",name[2])

	var name[5] string
	name[2] = "Swarnavo"
	name[0] = "Aakash"
	fmt.Println("Name is :", name)
	fmt.Printf("Price Array is %q\n", name)
}