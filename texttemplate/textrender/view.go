package textrender

import (
	"github.com/herb-go/herbtext"
	"github.com/herb-go/herbtext/texttemplate"
)

type View struct {
	Key  string
	View texttemplate.View
}

type Views []*View

func (views *Views) Render(ds Dataset) (*Outputs, error) {
	outputs := make(Outputs, len(*views))
	for k, v := range *views {
		output, err := v.View.Render(ds)
		if err != nil {
			return nil, err
		}
		outputs[k] = &Output{
			Key:    v.Key,
			Output: output,
		}
	}
	return &outputs, nil
}

func Parse(engine texttemplate.Engine, env herbtext.Environment, t *Templates) (*Views, error) {
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
