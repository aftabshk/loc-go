package main

import (
	"loc-go/src"
	"loc-go/src/domain"
	"loc-go/src/utils"
	"os"
	"strings"
	"sync"
)

func calculateLines(data string) (lines int) {
	lines = strings.Count(data, "\n") + 1
	return
}

func calculateLinesOfAllFilesInDir(
	dirPath string,
	options src.Options,
	allFiles chan domain.FileMetadata,
	wg *sync.WaitGroup,
	safeCounter *domain.SafeCounter,
) {
	dir, _ := os.ReadDir(dirPath)
	for _, value := range dir {
		fileOrDirectoryName := utils.PrefixPath(dirPath, value.Name())
		if !utils.ShouldIgnore(options.Ignore, value.Name()) && value.IsDir() {
			wg.Add(1)
			safeCounter.Inc()
			go calculateLinesOfAllFilesInDir(fileOrDirectoryName, options, allFiles, wg, safeCounter)
		}
		if !utils.ShouldIgnore(options.Ignore, value.Name()) && !value.IsDir() {
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
	options := src.Resolve(os.Args[1:])
	go calculateLinesOfAllFilesInDir(".", options, allFiles, &wg, &safeCounter)
	wg.Add(1)
	go collectFileMetadataAndPrint(allFiles, &wg)
	wg.Wait()
}
