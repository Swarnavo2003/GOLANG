package main

import (
	"fmt"
	"os"
)

func main() {

	/*
	file,err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file :", err)
		return 
	}
	defer file.Close()

	content := "hello world by swarnavo"
	byte , error := io.WriteString(file, content+"\n")
	fmt.Println("Byte written :", byte)
	if error != nil {
		fmt.Println("Error while writing file :", error)
		return 
	}

	fmt.Println("successfully created a file")
	*/

	/* 
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file :", err)
		return 
	}
	defer file.Close()

	// create buffer too read the file content
	buffer := make([]byte, 1024)

	// Read the file content into the buffer
	for {
		n,err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file", err)
		}
		// Process to read content
		fmt.Println(string(buffer[:n]))
	} 
	*/

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}
	fmt.Println(string(content))
}