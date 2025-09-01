package option_resolvers

import (
	"loc-go/domain"
	"os"
	"strings"
)

type OptionResolver interface {
	resolve(args []string, options *domain.Options) *domain.Options
}

func readLocIgnore() (directoriesOrFilesToIgnore []string) {
	locIgnoreFilePath := os.Getenv("HOME") + "/.locignore"
	locIgnore, _ := os.ReadFile(locIgnoreFilePath)
	directoriesOrFilesToIgnore = strings.Split(string(locIgnore), "\n")
	return
}

func Resolve(cliArgs []string) domain.Options {
	options := &domain.Options{}
	path := "."

	for i := 0; i < len(cliArgs); {
		if cliArgs[i] == "-ignore" {
			IgnoreOptionResolver{}.resolve(cliArgs[i+1:i+2], options)
			i = i + 2
		} else if cliArgs[i] == "-sort" {
			SortOptionResolver{}.resolve(cliArgs[i+1:i+3], options)
			i = i + 3
		} else if i == (len(cliArgs) - 1) {
			path = cliArgs[i]
			i = i + 1
		}
	}
	PathOptionResolver{}.resolve([]string{path}, options)

	return *options
}
