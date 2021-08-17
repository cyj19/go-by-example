package main

import "os"

func main() {

	_, err := os.Create("/temp/json.txt")
	if err != nil {
		panic(err)
	}
}
