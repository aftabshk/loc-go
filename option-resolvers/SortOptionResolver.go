package option_resolvers

import (
	"loc-go/domain"
)

type SortOptionResolver struct{}

func (i SortOptionResolver) resolve(args []string, options *domain.Options) *domain.Options {
	(*options).Sort = domain.Sort{
		Key:       args[0],
		Direction: args[1],
	}
	return options
}
