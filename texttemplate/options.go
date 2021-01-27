package texttemplate

import "github.com/herb-go/herbtext"

type Options struct {
	Helpers    map[string]herbtext.Helper
	Formatters map[string]herbtext.Formatter
}

func NewOptions() *Options {
	return &Options{
		Helpers:    map[string]herbtext.Helper{},
		Formatters: map[string]herbtext.Formatter{},
	}
}
