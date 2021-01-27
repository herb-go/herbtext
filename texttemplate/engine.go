package texttemplate

import "github.com/herb-go/herbtext"

type Engine interface {
	ApplyOptions(*herbtext.Environment) error
	Parse(template string, env *herbtext.Environment) (View, error)
	Supported() (directives []string, err error)
}
