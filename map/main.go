package main

import "fmt"

func main() {

	// name <-> grade
	studentGrades := make(map[string]int)
	studentGrades["Swarna"] = 100
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	fmt.Println("Marks of Bob :", studentGrades["Bob"])
	studentGrades["Bob"] = 100
	fmt.Println("Marks of Bob :", studentGrades["Bob"])

	delete(studentGrades, "Bob")
	fmt.Println("Marks of Bob :", studentGrades["Bob"])

	// Checking if a key exists
	grades,exists := studentGrades["David"]
	fmt.Println("Grades of David :", grades)
	fmt.Println("David exits :", exists)
	
	// Checking if a key exists
	Grades,Exists := studentGrades["Swarna"]
	fmt.Println("Grades of David :", Grades)
	fmt.Println("David exits :", Exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"Alice": 90,
		"Bob": 85,
		"Charlie": 95,
	}

	for index, value := range person {
		fmt.Printf("---------Key is %s and marks is %d-----------\n",index,value)
	}
}