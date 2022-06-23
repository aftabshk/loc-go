package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func calculateLines(data string) (lines int) {
	lines = strings.Count(data, "\n") + 1
	return
}

func prettyPrint(fileName string, numberOfLines int) {
	fmt.Printf("%v	%v\n", fileName, numberOfLines)
}

func calculateLinesOfAllFilesInDir(dirPath string, wg *sync.WaitGroup) {
	dir, _ := os.ReadDir(dirPath)
	for _, value := range dir {
		if value.IsDir() {
			wg.Add(1)
			go calculateLinesOfAllFilesInDir(prefixPath(dirPath, value.Name()), wg)
		}
		if !value.IsDir() {
			fileName := prefixPath(dirPath, value.Name())
			file, _ := os.ReadFile(fileName)
			numberOfLines := calculateLines(string(file))
			prettyPrint(fileName, numberOfLines)
		}
	}
	wg.Done()
}

func prefixPath(dirPath, name string) string {
	return dirPath + "/" + name
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go calculateLinesOfAllFilesInDir(".", &wg)
	wg.Wait()
}
