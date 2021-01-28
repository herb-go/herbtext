package texttemplate

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

//Factory engine factory
type Factory func(loader func(v interface{}) error) (Engine, error)

var (
	factorysMu sync.RWMutex
	factories  = make(map[string]Factory)
)

// Register makes a engine creator available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, f Factory) {
	factorysMu.Lock()
	defer factorysMu.Unlock()
	if f == nil {
		panic(errors.New("texttemplate: Register delivery factory is nil"))
	}
	if _, dup := factories[name]; dup {
		panic(errors.New("texttemplate: Register called twice for factory " + name))
	}
	factories[name] = f
}

//UnregisterAll unregister all driver
func UnregisterAll() {
	factorysMu.Lock()
	defer factorysMu.Unlock()
	// For tests.
	factories = make(map[string]Factory)
}

//Factories returns a sorted list of the names of the registered factories.
func Factories() []string {
	factorysMu.RLock()
	defer factorysMu.RUnlock()
	var list []string
	for name := range factories {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

//NewEngine create new driver with given name loader.
//Reutrn driver created and any error if raised.
func NewEngine(name string, loader func(v interface{}) error) (Engine, error) {
	factorysMu.RLock()
	factoryi, ok := factories[name]
	factorysMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("texttemplate: unknown driver %q (forgotten import?)", name)
	}
	return factoryi(loader)
}
