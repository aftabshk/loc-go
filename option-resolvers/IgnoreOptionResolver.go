package option_resolvers

import (
	"loc-go/domain"
	"strings"
)

type IgnoreOptionResolver struct{}

func (i IgnoreOptionResolver) resolve(args []string, options *domain.Options) *domain.Options {
	fileOrDirNames := strings.Split(args[0], ",")
	options.Ignore = append(fileOrDirNames, readLocIgnore()...)
	return options
}
