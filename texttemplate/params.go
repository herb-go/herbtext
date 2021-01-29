package texttemplate

import (
	"github.com/herb-go/herbtext"
)

//Param param struct
type Param struct {
	//Target dataset key which store parsed data
	Target string
	//Source values key which data to parsed loaded from
	Source string
	//Parser data parsers
	Parser herbtext.Parser
}

//Params param list struct
type Params []*Param

//Load load dataset form given values
func (ps *Params) Load(values herbtext.Set) (Dataset, error) {
	ds := Dataset{}
	for _, v := range *ps {
		data, err := v.Parser(values.Get(v.Source))
		if err != nil {
			return nil, err
		}
		ds[v.Target] = data
	}
	return ds, nil
}
