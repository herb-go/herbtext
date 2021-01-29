package textviewset

import (
	"github.com/herb-go/herbtext"
	"github.com/herb-go/herbtext/texttemplate"
)

//ParseWithEngineName parse templates with given engine name and herbtext environment.
//Return views parsed and any error if raised.
func ParseWithEngineName(templates herbtext.Set, enginename string, env herbtext.Environment) (views Views, err error) {
	eng, err := texttemplate.GetEngine(enginename)
	if err != nil {
		return nil, err
	}
	return ParseWith(templates, eng, env)
}

//ParseWith parse templates with given engine and herbtext environment.
//Return views parsed and any error if raised.
func ParseWith(templates herbtext.Set, engine texttemplate.Engine, env herbtext.Environment) (views Views, err error) {
	var view texttemplate.View
	views = make(Views, templates.Length())
	templates.Range(func(k, v string) bool {
		view, err = engine.Parse(v, env)
		if err != nil {
			return false
		}
		views[k] = view
		return true
	})
	if err != nil {
		return nil, err
	}
	return views, nil
}

//Views views type which can render dataset to output set.
type Views map[string]texttemplate.View

//Render render given dataset to outout set
func (views Views) Render(ds texttemplate.Dataset) (herbtext.Map, error) {
	outputs := make(herbtext.Map, len(views))
	for k, v := range views {
		output, err := v.Render(ds)
		if err != nil {
			return nil, err
		}
		outputs[k] = output
	}
	return outputs, nil
}
