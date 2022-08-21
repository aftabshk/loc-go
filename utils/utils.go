package utils

import (
	"fmt"
	"loc-go/domain"
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
