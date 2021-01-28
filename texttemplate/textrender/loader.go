package textrender

import (
	"github.com/herb-go/herbtext"
)

type Loader struct {
	Target string
	Source string
	Parser herbtext.Parser
}

type LoaderSet map[string]*Loader

func (ls LoaderSet) Add(list ...*Loader) {
	for _, v := range list {
		ls[v.Source] = v
	}
}

func (ls LoaderSet) Load(values *Values) (Dataset, error) {
	ds := Dataset{}
	for _, v := range *values {
		l := ls[v.Key]
		if l != nil {
			data, err := l.Parser(v.Value)
			if err != nil {
				return nil, err
			}
			ds[l.Target] = data
		}
	}
	return ds, nil
}
