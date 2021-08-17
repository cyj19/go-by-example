package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating...")
	file, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(file *os.File) {
	fmt.Println("closing...")
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}

func writeFile(file *os.File) {
	fmt.Println("writing...")
	fmt.Fprintf(file, "json json josn josn")
}

func main() {
	// file, err := os.Create("json.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// fmt.Fprintf(file, "write into json....")
	f := createFile("json.txt")
	defer closeFile(f)
	writeFile(f)
}
