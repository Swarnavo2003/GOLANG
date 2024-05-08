package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string 
	Phone string
	fax string
}

type Address struct {
	House int
	Area string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var swarna Person
	fmt.Println("Swarna Person :", swarna)
	swarna.FirstName = "Swarnavo"
	swarna.LastName = "Majumder"
	// swarna.Age = 20
	// fmt.Println("Swarna Person :", swarna)

	person1 := Person {
		FirstName: "Susrutu",
		LastName: "Banerjee",
		Age: 21,
	}
	fmt.Println("Person 1 :", person1)

	var person2 = new(Person)
	person2.FirstName = "Rishika"
	person2.LastName = "Das"
	person2.Age = 18
	// fmt.Println("Person 2 :", person2.FirstName)
	// fmt.Println("Age of prince is :", swarna.Age)

	var employee1 Employee
	employee1.Person_Details = Person {
		FirstName: "Swarnavo",
		LastName: "Majumder",
		Age: 20,
	}

	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "9876543210"

	employee1.Person_Address = Address {
		House: 12,
		Area: "Kolkata",
		State: "West Bengal",
	}
	employee1.Person_Contact.fax = "fax@4567890"
	fmt.Println("Employee 1 :", employee1)
}