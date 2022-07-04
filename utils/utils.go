package utils

import "fmt"

func PrettyPrint(fileName string, numberOfLines int) {
	fmt.Printf("%v	%v\n", fileName, numberOfLines)
}

func ShouldIgnore(directoriesOrFilesToIgnore []string, dirOrFileName string) bool {
	return !contains(directoriesOrFilesToIgnore, dirOrFileName)
}

func PrefixPath(dirPath, name string) string {
	return dirPath + "/" + name
}
