package texttemplate

import (
	"fmt"

	"github.com/herb-go/herbtext"
)

//ParamConfig param config
type ParamConfig struct {
	//Source source key
	Source string
	//Target target key.
	//Target will be setted to source key if empty
	Target string
	//Parser paraser name in herbtext enviroment
	Parser string
	//Constant constant value.
	//Data will use constant instead of loading form source if not empty.
	Constant string
	//Required if param is required
	Required bool
}

//CreateParam create param with given enviroment.
func (c *ParamConfig) CreateParam(env herbtext.Environment) (*Param, error) {
	if c.Source == "" {
		return nil, fmt.Errorf("texttemplate:empty param source")
	}
	p := env.GetParser(c.Parser)
	if p == nil {
		return nil, fmt.Errorf("texttemplate: parser [%s] not found", c.Parser)
	}
	l := &Param{Target: c.Target, Source: c.Source, Parser: p}
	if l.Target == "" {
		l.Target = l.Source
	}
	l.Required = c.Required
	l.Constant = c.Constant
	return l, nil
}

//ParamDefinition param definition struct
type ParamDefinition struct {
	ParamConfig
	//Description param description
	Description string
	//Example param value exmaple
	Example string
}

//ParamDefinitions paramdefinition list struct
type ParamDefinitions []*ParamDefinition

//CreateParams create params with given herbtext environment
func (d *ParamDefinitions) CreateParams(env herbtext.Environment) (*Params, error) {
	var err error
	p := make(Params, len(*d))
	for k, v := range *d {
		p[k], err = v.CreateParam(env)
		if err != nil {
			return nil, err
		}
	}
	return &p, nil
}
