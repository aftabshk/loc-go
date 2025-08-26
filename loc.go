package main

import (
	"loc-go/domain"
	"loc-go/option-resolvers"
	"loc-go/utils"
	"os"
	"sort"
	"strings"
	"sync"
)

func calculateLOCForFile(fileName string) domain.FileMetadata {
	file, _ := os.ReadFile(fileName)
	numberOfLines := calculateLines(string(file))
	return domain.FileMetadata{
		FileName:      fileName,
		NumberOfLines: numberOfLines,
	}
}

func calculateLines(data string) int {
	return strings.Count(data, "\n") + 1
}

func calculateLinesOfAllFilesInDir(
	dirPath string,
	options domain.Options,
	allFiles chan domain.FileMetadata,
	wg *sync.WaitGroup,
	safeCounter *domain.SafeCounter,
) {
	defer wg.Done()
	dir, _ := os.ReadDir(dirPath)
	for _, value := range dir {
		fileOrDirectoryName := utils.PrefixPath(dirPath, value.Name())
		if !utils.ShouldIgnore(options.Ignore, value.Name()) && value.IsDir() {
			wg.Add(1)
			safeCounter.Inc()
			go calculateLinesOfAllFilesInDir(fileOrDirectoryName, options, allFiles, wg, safeCounter)
		}
		if !utils.ShouldIgnore(options.Ignore, value.Name()) && !value.IsDir() {
			allFiles <- calculateLOCForFile(fileOrDirectoryName)
		}
	}
	safeCounter.Dec()
	if safeCounter.Count() == 0 {
		close(allFiles)
	}
}

func sorted(files []domain.FileMetadata, options domain.Options) []domain.FileMetadata {
	if options.Sort.Key == "" || options.Sort.Direction == "" {
		return utils.SortDescending(files)
	}
	if options.Sort.Key == "name" {
		if options.Sort.Direction == "ASC" {
			sort.Slice(files, func(i, j int) bool {
				return utils.Compare(files[i].FileName, files[j].FileName)
			})
		} else if options.Sort.Direction == "DESC" {
			sort.Slice(files, func(i, j int) bool {
				return !utils.Compare(files[i].FileName, files[j].FileName)
			})
		}
	}
	if options.Sort.Key == "loc" {
		if options.Sort.Direction == "ASC" {
			return utils.SortAscending(files)
		} else if options.Sort.Direction == "DESC" {
			return utils.SortDescending(files)
		}
	}

	return files
}

func collectFileMetadataAndPrint(metadataOfAllFilesChan chan domain.FileMetadata, options domain.Options, wg *sync.WaitGroup) {
	var metadataOfAllFiles []domain.FileMetadata
	for fileMetadata := range metadataOfAllFilesChan {
		metadataOfAllFiles = append(metadataOfAllFiles, fileMetadata)
	}

	utils.PrintMetadata(metadataOfAllFiles)
	sortedFiles := sorted(metadataOfAllFiles, options)
	utils.PrettyPrintAll(sortedFiles)
	wg.Done()
}

func main() {
	allFiles := make(chan domain.FileMetadata)
	var wg sync.WaitGroup
	wg.Add(1)
	safeCounter := domain.SafeCounter{}
	safeCounter.Inc()
	options := option_resolvers.Resolve(os.Args[1:])
	go calculateLinesOfAllFilesInDir(".", options, allFiles, &wg, &safeCounter)
	wg.Add(1)
	go collectFileMetadataAndPrint(allFiles, options, &wg)
	wg.Wait()
}
