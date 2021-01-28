package texttemplate

import "testing"

func RegisterNullFactory() {
	Register("null", func(loader func(v interface{}) error) (Engine, error) {
		return nil, nil
	})
}

func TestFactory(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	UnregisterAll()
	fs := Factories()
	if len(fs) != 0 {
		t.Fatal(fs)
	}
	RegisterNullFactory()
	fs = Factories()
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
	d, err := NewEngine("unknown", nil)
	if d != nil || err == nil {
		t.Fatal(d, err)
	}
}

func TestNull(t *testing.T) {
	defer func() {
		UnregisterAll()
	}()
	RegisterNullFactory()
	c := Config{
		Engine:       "null",
		EngineConfig: nil,
	}
	engine, err := c.NewEngine()
	if err != nil || engine != nil {
		t.Fatal(err)
	}
}
