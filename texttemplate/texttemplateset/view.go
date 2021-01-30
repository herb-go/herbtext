package texttemplateset

import (
	"github.com/herb-go/herbtext"
	"github.com/herb-go/herbtext/texttemplate"
)

//ParseWithEngineName parse templates with given engine name and herbtext environment.
//Return views parsed and any error if raised.
func ParseWithEngineName(templates herbtext.Set, enginename string, env herbtext.Environment) (set TemplateSet, err error) {
	eng, err := texttemplate.GetEngine(enginename)
	if err != nil {
		return nil, err
	}
	return ParseWith(templates, eng, env)
}

//ParseWith parse templates with given engine and herbtext environment.
//Return views parsed and any error if raised.
func ParseWith(templates herbtext.Set, engine texttemplate.Engine, env herbtext.Environment) (set TemplateSet, err error) {
	var view texttemplate.Template
	set = make(TemplateSet, templates.Length())
	templates.Range(func(k, v string) bool {
		view, err = engine.Parse(v, env)
		if err != nil {
			return false
		}
		set[k] = view
		return true
	})
	if err != nil {
		return nil, err
	}
	return set, nil
}

//TemplateSet template set type
type TemplateSet map[string]texttemplate.Template

//Render render given dataset to outout set
func (s TemplateSet) Render(ds texttemplate.Dataset) (herbtext.Map, error) {
	outputs := make(herbtext.Map, len(s))
	for k, v := range s {
		output, err := v.Render(ds)
		if err != nil {
			return nil, err
		}
		outputs[k] = output
	}
	return outputs, nil
}
