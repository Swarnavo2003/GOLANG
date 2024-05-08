package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple.orange.banana.swarna"
	parts := strings.Split(data, ".")
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("count :", count)

	str = "    Hello, Go!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed :", trimmed)	 

	str1 := "Swarnavo"
	str2 := "Majumder"
	result := strings.Join([]string{str1, "Raha" ,str2}, " ")
	fmt.Println("result :", result)
}