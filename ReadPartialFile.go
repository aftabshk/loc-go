package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func readPartialFileTryout(fileName string, maxLineRead int) string {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	i := 0
	for ; i < maxLineRead && scanner.Scan(); i++ {
	}

	if scanner.Scan() {
		return strconv.Itoa(i) + "+"
	}

	return strconv.Itoa(i)
}

func main1() {
	start := time.Now()
	loc := readPartialFileTryout("./loc-go", 10000)
	end := time.Now()
	fmt.Println("Loc: ", loc, " Within: ", end.Sub(start).Seconds(), " secs")

	start = time.Now()
	actualLoc := calculateFullLoc("./Tasks.md")
	end = time.Now()
	fmt.Println("Loc: ", actualLoc, " Within: ", end.Sub(start).Seconds(), " secs")
}
