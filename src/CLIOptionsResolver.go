package src

import (
	"os"
	"strings"
)

type Options struct {
	Ignore []string
}

type OptionResolver interface {
	resolve(arg string, options *Options) *Options
}

type IgnoreOptionResolver struct{}

func readLocIgnore() (directoriesOrFilesToIgnore []string) {
	locIgnoreFilePath := os.Getenv("HOME") + "/.locignore"
	locIgnore, _ := os.ReadFile(locIgnoreFilePath)
	directoriesOrFilesToIgnore = strings.Split(string(locIgnore), "\n")
	return
}

func (o *Options) appendLocIgnore() *Options {
	locIgnore := readLocIgnore()
	o.Ignore = append(o.Ignore, locIgnore...)
	return o
}

func (i IgnoreOptionResolver) resolve(arg string, options *Options) *Options {
	fileOrDirNames := strings.Split(arg, ",")
	options.Ignore = fileOrDirNames
	return options
}

var allResolvers = map[string]OptionResolver{
	"ignore": IgnoreOptionResolver{},
}

func Resolve(cliArgs []string) Options {
	options := &Options{}

	for _, arg := range cliArgs {
		option := strings.Split(arg, "=")
		allResolvers[option[0]].resolve(option[1], options)
	}

	options = options.appendLocIgnore()

	return *options
}
