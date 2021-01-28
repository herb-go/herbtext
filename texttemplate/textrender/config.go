package textrender

import (
	"fmt"

	"github.com/herb-go/herbtext/texttemplate"

	"github.com/herb-go/herbtext"
)

type LoaderConfig struct {
	Target string
	Source string
	Parser string
}

func (c *LoaderConfig) CreateLoader(env herbtext.Environment) (*Loader, error) {
	p := env.GetParser(c.Parser)
	if p == nil {
		return nil, fmt.Errorf("textrender: parser [%s] not found", c.Parser)
	}
	return &Loader{Target: c.Target, Source: c.Source, Parser: p}, nil
}

type LoaderSetConfig []*LoaderConfig

func (c *LoaderSetConfig) CreateLoaderSet(env herbtext.Environment) (LoaderSet, error) {
	ls := LoaderSet{}
	for _, v := range *c {
		loader, err := v.CreateLoader(env)
		if err != nil {
			return nil, err
		}
		ls.Add(loader)
	}
	return ls, nil
}

type TemplatesConfig struct {
	Engine    string
	Tempaltes Templates
}

func (c *TemplatesConfig) CreateViews(env herbtext.Environment) (*Views, error) {
	e, err := texttemplate.GetEngine(c.Engine)
	if err != nil {
		return nil, err
	}
	return c.Tempaltes.Parse(e, env)
}
