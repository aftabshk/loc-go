package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)
import "loc-go/utils"
import "loc-go/domain"

func calculateLines(data string) (lines int) {
	lines = strings.Count(data, "\n") + 1
	return
}

func calculateLinesOfAllFilesInDir(
	dirPath string,
	directoriesOrFilesToIgnore []string,
	allFiles chan domain.FileMetadata,
	wg *sync.WaitGroup,
	wgCount *int,
) {
	dir, _ := os.ReadDir(dirPath)
	for _, value := range dir {
		fileOrDirectoryName := utils.PrefixPath(dirPath, value.Name())
		if utils.ShouldIgnore(directoriesOrFilesToIgnore, value.Name()) && value.IsDir() {
			wg.Add(1)
			*wgCount = *wgCount + 1
			go calculateLinesOfAllFilesInDir(fileOrDirectoryName, directoriesOrFilesToIgnore, allFiles, wg, wgCount)
		}
		if utils.ShouldIgnore(directoriesOrFilesToIgnore, value.Name()) && !value.IsDir() {
			fileName := fileOrDirectoryName
			file, _ := os.ReadFile(fileName)
			numberOfLines := calculateLines(string(file))
			fileMetadata := domain.FileMetadata{
				FileName:      fileName,
				NumberOfLines: numberOfLines,
			}
			allFiles <- fileMetadata
		}
	}
	wg.Done()
	*wgCount = *wgCount - 1
	if *wgCount == 0 {
		close(allFiles)
	}
}

func readLocIgnore() (directoriesOrFilesToIgnore []string) {
	locIgnoreFilePath := os.Getenv("HOME") + "/.locignore"
	locIgnore, _ := os.ReadFile(locIgnoreFilePath)
	directoriesOrFilesToIgnore = strings.Split(string(locIgnore), "\n")
	return
}

func prettyPrint(allFiles chan domain.FileMetadata) {
	for fileMetadata := range allFiles {
		fmt.Println(fileMetadata)
	}
}

func main() {
	allFiles := make(chan domain.FileMetadata)
	var wg sync.WaitGroup
	wg.Add(1)
	wgCount := 1
	go calculateLinesOfAllFilesInDir(".", readLocIgnore(), allFiles, &wg, &wgCount)
	go prettyPrint(allFiles)
	wg.Wait()
}
