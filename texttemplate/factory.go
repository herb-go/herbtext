package texttemplate

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

var (
	enginesMu sync.RWMutex
	engines   = make(map[string]Engine)
)

// Register makes a engine available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, e Engine) {
	enginesMu.Lock()
	defer enginesMu.Unlock()
	if e == nil {
		panic(errors.New("texttemplate: Register engine is nil"))
	}
	if _, dup := engines[name]; dup {
		panic(errors.New("texttemplate: Register called twice for factory " + name))
	}
	engines[name] = e
}

//UnregisterAll unregister all driver
func UnregisterAll() {
	enginesMu.Lock()
	defer enginesMu.Unlock()
	// For tests.
	engines = make(map[string]Engine)
}

//Factories returns a sorted list of the names of the registered factories.
func Factories() []string {
	enginesMu.RLock()
	defer enginesMu.RUnlock()
	var list []string
	for name := range engines {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

//GetEngine get new driver bt given name .
//Reutrn driver created and any error if raised.
func GetEngine(name string) (Engine, error) {
	enginesMu.RLock()
	factoryi, ok := engines[name]
	enginesMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("texttemplate: unknown engine %q (forgotten import?)", name)
	}
	return factoryi, nil
}
