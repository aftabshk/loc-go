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

var allResolvers = map[string]OptionResolver{
	"-ignore": IgnoreOptionResolver{},
	"-sort":   SortOptionResolver{},
}

func Resolve(cliArgs []string) domain.Options {
	options := &domain.Options{}
	options.Path = "."

	for i := 0; i < len(cliArgs); {
		if cliArgs[i] == "-ignore" {
			IgnoreOptionResolver{}.resolve(cliArgs[i+1:i+2], options)
			i = i + 2
		} else if cliArgs[i] == "-sort" {
			SortOptionResolver{}.resolve(cliArgs[i+1:i+3], options)
			i = i + 3
		} else if i == (len(cliArgs) - 1) {
			options.Path = cliArgs[i]
			i = i + 1
		}
	}

	return *options
}
