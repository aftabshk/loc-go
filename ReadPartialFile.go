package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readPartialFile() {
	file, err := os.Open("./normalFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for i := 1; scanner.Scan() && i <= 3000; i++ {
	}

	fmt.Println("Counted 2000 lines ")
	//fmt.Println(str)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
