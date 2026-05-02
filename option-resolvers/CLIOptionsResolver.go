package option_resolvers

import (
	"loc-go/domain"
	"os"
	"strconv"
	"strings"
	"log"
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
		switch cliArgs[i] {
			case "-ignore": {
				IgnoreOptionResolver{}.resolve(cliArgs[i+1:i+2], options)
				i = i + 2
			}
			case "-sort": {
				SortOptionResolver{}.resolve(cliArgs[i+1:i+3], options)
				i = i + 3
			}
			case "-partial": {
				(*options).Partial = true
				i = i + 1
			}
			case "-partial-read-upto": {
				value, err := strconv.Atoi(cliArgs[i+1])

				if err != nil {
					log.Fatal("Error converting partial read upto option from string to integer. Please provide valid number")
					panic(err)
				}

				(*options).PartialReadUpto = value
				i = i + 2
			}
			default: {
				if i == (len(cliArgs) - 1) {
					path = cliArgs[i]
					i = i + 1
				}
			}
		}
	}

	if options.PartialReadUpto == 0 {
		(*options).PartialReadUpto = domain.DEFAULT_MAX_PARTIAL_LOC_COUNT
	}

	PathOptionResolver{}.resolve([]string{path}, options)

	return *options
}
