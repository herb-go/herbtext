package texttemplate

import "github.com/herb-go/herbtext"

//Engine engine interface
type Engine interface {
	//Parse parse given template with given environment to template view.
	Parse(template string, env herbtext.Environment) (View, error)
	//Supported return supported directives which can be used in template string.
	Supported(env herbtext.Environment) (directives []string, err error)
}
