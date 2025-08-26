package utils

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"loc-go/domain"
	"os"
	"sort"
	"unicode"
)

func PrettyPrintAll(allFiles []domain.FileMetadata) {
	t := table.NewWriter()
	t.SetStyle(table.StyleColoredBlackOnCyanWhite)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"File Name", "LOC"})
	for _, fileMetadata := range allFiles {
		t.AppendRow([]interface{}{text.WrapHard(fileMetadata.FileName, 100), fileMetadata.NumberOfLines})
		t.AppendSeparator()
	}
	t.Render()
}

func ShouldIgnore(directoriesOrFilesToIgnore []string, dirOrFileName string) bool {
	return contains(directoriesOrFilesToIgnore, dirOrFileName)
}

func PrefixPath(dirPath, name string) string {
	return dirPath + "/" + name
}

func SortDescending(files []domain.FileMetadata) []domain.FileMetadata {
	result := make([]domain.FileMetadata, len(files))
	copy(result, files)
	sort.Slice(result, func(i, j int) bool {
		return result[i].NumberOfLines > result[j].NumberOfLines
	})
	return result
}

func SortAscending(files []domain.FileMetadata) []domain.FileMetadata {
	result := make([]domain.FileMetadata, len(files))
	copy(result, files)
	sort.Slice(result, func(i, j int) bool {
		return result[i].NumberOfLines < result[j].NumberOfLines
	})
	return result
}

func Compare(str1 string, str2 string) bool {
	str1Runes := []rune(str1)
	str2Runes := []rune(str2)

	for idx := 0; idx < len(str1Runes) && idx < len(str2Runes); idx++ {
		ir := str1Runes[idx]
		jr := str2Runes[idx]

		lir := unicode.ToLower(ir)
		ljr := unicode.ToLower(jr)

		if lir != ljr {
			return lir < ljr
		}

		if ir != jr {
			return ir < jr
		}
	}
	return len(str1Runes) < len(str2Runes)
}

func PrintMetadata(fileMetadata []domain.FileMetadata) {
	fmt.Printf("\nTotal number of files: %d\n\n", len(fileMetadata))
}
