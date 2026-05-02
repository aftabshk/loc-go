package domain

type Sort struct {
	Key       string
	Direction string
}

type Options struct {
	Ignore []string
	Sort
	Path string
	Partial bool
	PartialReadUpto int
}

