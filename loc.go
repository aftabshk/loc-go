package main

import (
	"fmt"
	"os"
	"strings"
)

func calculateLines(data string) (lines int) {
	lines = strings.Count(data, "\n") + 1
	return
}

func prettyPrint(fileName string, numberOfLines int) {
	fmt.Printf("%v	%v\n", fileName, numberOfLines)
}

func calculateLinesOfAllFilesInDir(dirPath string) {
	dir, _ := os.ReadDir(dirPath)
	for _, value := range dir {
		if value.IsDir() {
			calculateLinesOfAllFilesInDir(prefixPath(dirPath, value.Name()))
		}
		if !value.IsDir() {
			fileName := prefixPath(dirPath, value.Name())
			file, _ := os.ReadFile(fileName)
			numberOfLines := calculateLines(string(file))
			prettyPrint(fileName, numberOfLines)
		}
	}
}

func prefixPath(dirPath, name string) string {
	return dirPath + "/" + name
}

func main() {
	calculateLinesOfAllFilesInDir(".")
}
