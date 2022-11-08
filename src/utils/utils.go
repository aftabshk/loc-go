package utils

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"loc-go/src/domain"
	"os"
	"sort"
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
