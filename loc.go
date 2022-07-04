package main

import (
	"os"
	"strings"
	"sync"
)
import "loc-go/utils"

func calculateLines(data string) (lines int) {
	lines = strings.Count(data, "\n") + 1
	return
}

func calculateLinesOfAllFilesInDir(
	dirPath string,
	directoriesOrFilesToIgnore []string,
	wg *sync.WaitGroup,
) {
	dir, _ := os.ReadDir(dirPath)
	for _, value := range dir {
		fileOrDirectoryName := utils.PrefixPath(dirPath, value.Name())
		if utils.ShouldIgnore(directoriesOrFilesToIgnore, value.Name()) && value.IsDir() {
			wg.Add(1)
			go calculateLinesOfAllFilesInDir(fileOrDirectoryName, directoriesOrFilesToIgnore, wg)
		}
		if utils.ShouldIgnore(directoriesOrFilesToIgnore, value.Name()) && !value.IsDir() {
			fileName := fileOrDirectoryName
			file, _ := os.ReadFile(fileName)
			numberOfLines := calculateLines(string(file))
			utils.PrettyPrint(fileName, numberOfLines)
		}
	}
	wg.Done()
}

func readLocIgnore() (directoriesOrFilesToIgnore []string) {
	locIgnore, _ := os.ReadFile("/Users/aftabshk/.locignore")
	directoriesOrFilesToIgnore = strings.Split(string(locIgnore), "\n")
	return
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go calculateLinesOfAllFilesInDir(".", readLocIgnore(), &wg)
	wg.Wait()
}
