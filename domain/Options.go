package domain

type Sort struct {
	Key       string
	Direction string
}

type Options struct {
	Ignore []string
	Sort
}

func (o *Options) appendToIgnore(ignorePaths []string) *Options {
	o.Ignore = append(o.Ignore, ignorePaths...)
	return o
}
