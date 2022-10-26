package main

import (
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
	safeCounter *domain.SafeCounter,
) {
	dir, _ := os.ReadDir(dirPath)
	for _, value := range dir {
		fileOrDirectoryName := utils.PrefixPath(dirPath, value.Name())
		if !utils.ShouldIgnore(directoriesOrFilesToIgnore, value.Name()) && value.IsDir() {
			wg.Add(1)
			safeCounter.Inc()
			go calculateLinesOfAllFilesInDir(fileOrDirectoryName, directoriesOrFilesToIgnore, allFiles, wg, safeCounter)
		}
		if !utils.ShouldIgnore(directoriesOrFilesToIgnore, value.Name()) && !value.IsDir() {
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
	safeCounter.Dec()
	if safeCounter.Count() == 0 {
		close(allFiles)
	}
}

func readLocIgnore() (directoriesOrFilesToIgnore []string) {
	locIgnoreFilePath := os.Getenv("HOME") + "/.locignore"
	locIgnore, _ := os.ReadFile(locIgnoreFilePath)
	directoriesOrFilesToIgnore = strings.Split(string(locIgnore), "\n")
	return
}

func collectFileMetadataAndPrint(metadataOfAllFilesChan chan domain.FileMetadata, wg *sync.WaitGroup) {
	var metadataOfAllFiles []domain.FileMetadata
	for fileMetadata := range metadataOfAllFilesChan {
		metadataOfAllFiles = append(metadataOfAllFiles, fileMetadata)
	}

	sortedFiles := utils.SortDescending(metadataOfAllFiles)
	utils.PrettyPrintAll(sortedFiles)
	wg.Done()
}

func main() {
	allFiles := make(chan domain.FileMetadata)
	var wg sync.WaitGroup
	wg.Add(1)
	safeCounter := domain.SafeCounter{}
	safeCounter.Inc()
	go calculateLinesOfAllFilesInDir(".", readLocIgnore(), allFiles, &wg, &safeCounter)
	wg.Add(1)
	go collectFileMetadataAndPrint(allFiles, &wg)
	wg.Wait()
}
