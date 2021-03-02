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
	//Constant constant value.
	//Data will use constant instead of loading form source if not empty.
	Constant string
	//Required if param is required
	Required bool
	//Parser data parsers
	Parser herbtext.Parser
}

//Params param list struct
type Params []*Param

//Load load dataset form given values.
func (ps *Params) Load(values herbtext.Set) (Dataset, error) {
	ds := Dataset{}
	for _, v := range *ps {
		var value string
		if v.Constant != "" {
			value = v.Constant
		} else {
			value = values.Get(v.Source)
		}
		if v.Required && value == "" {
			return nil, NewParamMissedError(v.Source)
		}
		data, err := v.Parser(value)
		if err != nil {
			return nil, err
		}
		ds[v.Target] = data
	}
	return ds, nil
}
