package texttemplate

import (
	"testing"

	"github.com/herb-go/herbtext"
)

type nopEngine struct {
}

//Parse parse given template with given environment to template view.
func (nopEngine) Parse(template string, env herbtext.Environment) (View, error) {
	return nil, nil
}

//Supported return supported directives which can be used in template string.
func (nopEngine) Supported(env herbtext.Environment) (directives []string, err error) {
	return nil, nil
}

func RegisterNullFactory() {
	Register("null", nopEngine{})
}

func TestEngines(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	UnregisterAll()
	fs := Engines()
	if len(fs) != 0 {
		t.Fatal(fs)
	}
	RegisterNullFactory()
	fs = Engines()
	if len(fs) != 1 {
		t.Fatal(fs)
	}
}

func TestNilFactory(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
	}()
	Register("test", nil)
}

func TestDupFactory(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
	}()
	RegisterNullFactory()
	RegisterNullFactory()
}

func TestUnknownFactory(t *testing.T) {
	d, err := GetEngine("unknown")
	if d != nil || err == nil {
		t.Fatal(d, err)
	}
}

func TestNull(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	RegisterNullFactory()
	engine, err := GetEngine("null")
	if err != nil || engine == nil {
		t.Fatal(err)
	}
}
