package herbtext

import (
	"sync/atomic"
)

//PlainEnvironment plain environment(writeable) struct
type PlainEnvironment struct {
	Converters map[string]Converter
	Formatters map[string]Formatter
	Parsers    map[string]Parser
}

//Reset reset plain environment to empty value
func (e *PlainEnvironment) Reset() {
	e.Converters = map[string]Converter{}
	e.Formatters = map[string]Formatter{}
	e.Parsers = map[string]Parser{}
}

//MergeWith mergive plain environment with given environment
//Nothing will happened if nil environment given.
func (e *PlainEnvironment) MergeWith(env Environment) {
	if env == nil {
		return
	}
	env.RangeConverters(func(key string, converter Converter) bool {
		e.Converters[key] = converter
		return true
	})
	env.RangeFormatters(func(key string, formatter Formatter) bool {
		e.Formatters[key] = formatter
		return true
	})
	env.RangeParsers(func(key string, parser Parser) bool {
		e.Parsers[key] = parser
		return true
	})
}

//RangeConverters range over environment converters with given func.
//Stop if false returned.
func (e *PlainEnvironment) RangeConverters(f func(key string, value Converter) bool) {
	for k, v := range e.Converters {
		if !f(k, v) {
			return
		}
	}
}

//SetConverter set convernter to enviroment with given name
func (e *PlainEnvironment) SetConverter(name string, c Converter) {
	e.Converters[name] = c
}

//GetConverter get converter from environment by given name.
//Nil should be returned if name not found.
func (e *PlainEnvironment) GetConverter(name string) Converter {
	return e.Converters[name]
}

//RangeFormatters range over environment formatters with given func.
//Stop if false returned.
func (e *PlainEnvironment) RangeFormatters(f func(key string, value Formatter) bool) {
	for k, v := range e.Formatters {
		if !f(k, v) {
			return
		}
	}
}

//SetFormatter set formatter to enviroment with given name
func (e *PlainEnvironment) SetFormatter(name string, f Formatter) {
	e.Formatters[name] = f
}

//GetFormatter get formatter from environment by given name.
//Nil should be returned if name not found.
func (e *PlainEnvironment) GetFormatter(name string) Formatter {
	return e.Formatters[name]
}

//RangeParsers range over environment parsers with given func.
//Stop if false returned.
func (e *PlainEnvironment) RangeParsers(f func(key string, value Parser) bool) {
	for k, v := range e.Parsers {
		if !f(k, v) {
			return
		}
	}
}

//SetParser set parser to enviroment with given name
func (e *PlainEnvironment) SetParser(name string, p Parser) {
	e.Parsers[name] = p
}

//GetParser get parser from environment by given name.
//Nil should be returned if name not found.
func (e *PlainEnvironment) GetParser(name string) Parser {
	return e.Parsers[name]
}

//NewEnvironment create new plain enviroment.
func NewEnvironment() *PlainEnvironment {
	e := &PlainEnvironment{}
	e.Reset()
	return e
}

//Clone clone enviroment to a new plain enviroment
func Clone(env Environment) *PlainEnvironment {
	e := NewEnvironment()
	env.RangeConverters(func(key string, value Converter) bool {
		e.SetConverter(key, value)
		return true
	})
	env.RangeFormatters(func(key string, value Formatter) bool {
		e.SetFormatter(key, value)
		return true
	})
	env.RangeParsers(func(key string, value Parser) bool {
		e.SetParser(key, value)
		return true
	})
	return e
}

//Environment environment(readonly) interface
type Environment interface {
	//GetConverter get converter from environment by given name.
	//Nil should be returned if name not found.
	GetConverter(name string) Converter
	//GetFormatter get formatter from environment by given name.
	//Nil should be returned if name not found.
	GetFormatter(name string) Formatter
	//GetParser get parser from environment by given name.
	//Nil should be returned if name not found.
	GetParser(name string) Parser
	//RangeConverters range over environment converters with given func.
	//Stop if false returned.
	RangeConverters(f func(key string, value Converter) bool)
	//RangeFormatters range over environment formatters with given func.
	//Stop if false returned.
	RangeFormatters(f func(key string, value Formatter) bool)
	//RangeParsers range over environment parsers with given func.
	//Stop if false returned.
	RangeParsers(f func(key string, value Parser) bool)
}

var defaultEnvironment atomic.Value

//DefaultEnvironment return default environment
var DefaultEnvironment = func() Environment {
	return defaultEnvironment.Load().(Environment)
}

//SetDefaultEnvironment set given enviroment as default
var SetDefaultEnvironment = func(e Environment) {
	defaultEnvironment.Store(e)
}

func init() {
	SetDefaultEnvironment(NewEnvironment())
}
