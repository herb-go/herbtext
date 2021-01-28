package textrender

import (
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
