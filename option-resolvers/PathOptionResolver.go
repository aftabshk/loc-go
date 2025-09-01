package option_resolvers

import (
	"loc-go/domain"
	"strings"
)

type PathOptionResolver struct{}

func (i PathOptionResolver) resolve(args []string, options *domain.Options) *domain.Options {
	(*options).Path = strings.Trim(args[0], " /")
	for i, ignorePath := range (*options).Ignore {
		(*options).Ignore[i] = strings.Trim((*options).Path + "/" + ignorePath, " /")
	}
	return options
}
