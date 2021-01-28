package textrender

import (
	"github.com/herb-go/herbtext"
	"github.com/herb-go/herbtext/texttemplate"
)

type Template struct {
	Key      string
	Template string
}

type Templates []*Template

func (t *Templates) Parse(engine texttemplate.Engine, env herbtext.Environment) (*Views, error) {
	views := make(Views, len(*t))
	for k, v := range *t {
		view, err := engine.Parse(v.Template, env)
		if err != nil {
			return nil, err
		}
		views[k] = &View{
			Key:  v.Key,
			View: view,
		}
	}
	return &views, nil
}
