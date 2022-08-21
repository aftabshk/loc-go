package utils

import (
	"fmt"
	"loc-go/domain"
	"sort"
)

func PrettyPrint(fileName string, numberOfLines int) {
	fmt.Printf("%v	%v\n", fileName, numberOfLines)
}

func PrettyPrintAll(allFiles []domain.FileMetadata) {
	for _, fileMetadata := range allFiles {
		PrettyPrint(fileMetadata.FileName, fileMetadata.NumberOfLines)
	}
}

func ShouldIgnore(directoriesOrFilesToIgnore []string, dirOrFileName string) bool {
	return !contains(directoriesOrFilesToIgnore, dirOrFileName)
}

func PrefixPath(dirPath, name string) string {
	return dirPath + "/" + name
}

func SortDescending(files []domain.FileMetadata) []domain.FileMetadata {
	result := make([]domain.FileMetadata, len(files))
	copy(result, files)
	sort.Slice(result, func(i, j int) bool {
		return files[i].NumberOfLines > files[j].NumberOfLines
	})
	return result
}
