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

func calculateLinesOfAllFilesInDir(
	dirPath string,
	directoriesOrFilesToIgnore []string,
	wg *sync.WaitGroup,
) {
	dir, _ := os.ReadDir(dirPath)
	for _, value := range dir {
		fileOrDirectoryName := prefixPath(dirPath, value.Name())
		if shouldIgnore(directoriesOrFilesToIgnore, value.Name()) && value.IsDir() {
			wg.Add(1)
			go calculateLinesOfAllFilesInDir(fileOrDirectoryName, directoriesOrFilesToIgnore, wg)
		}
		if shouldIgnore(directoriesOrFilesToIgnore, value.Name()) && !value.IsDir() {
			fileName := fileOrDirectoryName
			file, _ := os.ReadFile(fileName)
			numberOfLines := calculateLines(string(file))
			prettyPrint(fileName, numberOfLines)
		}
	}
	wg.Done()
}

func shouldIgnore(directoriesOrFilesToIgnore []string, dirOrFileName string) bool {
	return !contains(directoriesOrFilesToIgnore, dirOrFileName)
}

func prefixPath(dirPath, name string) string {
	return dirPath + "/" + name
}

func readLocIgnore() (directoriesOrFilesToIgnore []string) {
	locIgnore, _ := os.ReadFile("/Users/aftabshk/.locignore")
	directoriesOrFilesToIgnore = strings.Split(string(locIgnore), "\n")
	return
}

func contains(arr []string, s string) (isFound bool) {
	for i := 0; i < len(arr) && !isFound; i++ {
		isFound = arr[i] == s
	}
	return
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go calculateLinesOfAllFilesInDir(".", readLocIgnore(), &wg)
	wg.Wait()
}
