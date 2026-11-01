package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Password Manager")
	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")
	fmt.Println("1. add a new password")
	fmt.Println("2. get a password")
	fmt.Println("3. delete a password") 
}

// function for creating a file. with ur like name
func createFile(filename string) {
	err := os.WriteFile(filename, []byte("test"), 0755) // 0755 is the permission to the file
	if err != nil {
		fmt.Println("unable to create file: %w", err)
	}
}