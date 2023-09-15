package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.ReadFile("./normalFile.txt")
	if err != nil {
		log.Fatal(err)
	}
}
