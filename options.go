package herbtext

import "sync/atomic"

type Environment struct {
	Converters map[string]Converter
	Formatters map[string]Formatter
	Parsers    map[string]Parser
}

func (e *Environment) Reset() {
	e.Converters = map[string]Converter{}
	e.Formatters = map[string]Formatter{}
	e.Parsers = map[string]Parser{}
}
func (e *Environment) SetConverter(name string, c Converter) {
	e.Converters[name] = c
}
func (e *Environment) GetConverter(name string) Converter {
	return e.Converters[name]
}
func (e *Environment) SetFormatter(name string, f Formatter) {
	e.Formatters[name] = f
}
func (e *Environment) GetFormatter(name string) Formatter {
	return e.Formatters[name]
}
func (e *Environment) SetParser(name string, p Parser) {
	e.Parsers[name] = p
}
func (e *Environment) GetParser(name string) Parser {
	return e.Parsers[name]
}

func NewEnvironment() *Environment {
	e := &Environment{}
	e.Reset()
	return e
}

var defaultEnvironment atomic.Value

var DefaultEnvironment = func() *Environment {
	return defaultEnvironment.Load().(*Environment)
}

var SetDefaultEnvironment = func(e *Environment) {
	defaultEnvironment.Store(e)
}

func init() {
	SetDefaultEnvironment(NewEnvironment())
}
