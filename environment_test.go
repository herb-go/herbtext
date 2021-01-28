package herbtext

import (
	"testing"
)

func TestDefaultEnvironment(t *testing.T) {
	defer func() {
		SetDefaultEnvironment(NewEnvironment())
	}()
	e := Clone(DefaultEnvironment())
	e.SetParser("nopparser", nopParser)
	if DefaultEnvironment().GetParser("nopparser") != nil {
		t.Fatal()
	}
	SetDefaultEnvironment(e)
	if DefaultEnvironment().GetParser("nopparser") == nil {
		t.Fatal()
	}
}

func nopConverter(string) string {
	return ""
}

func nopParser(string) (interface{}, error) {
	return nil, nil
}

func nopFormatter(string, string) string {
	return ""
}

func TestEnvironment(t *testing.T) {
	env := NewEnvironment()
	f := env.GetFormatter("nopformatter")
	if f != nil {
		t.Fatal(f)
	}
	p := env.GetParser("nopparser")
	if p != nil {
		t.Fatal(p)
	}
	c := env.GetConverter("nopconverter")
	if c != nil {
		t.Fatal(c)
	}
	env.SetFormatter("nopformatter", nopFormatter)
	env.SetFormatter("nopformatter2", nopFormatter)
	env.SetParser("nopparser", nopParser)
	env.SetParser("nopparser2", nopParser)
	env.SetConverter("nopconverter", nopConverter)
	env.SetConverter("nopconverter2", nopConverter)
	f = env.GetFormatter("nopformatter")
	if f == nil {
		t.Fatal(f)
	}
	p = env.GetParser("nopparser")
	if p == nil {
		t.Fatal(p)
	}
	c = env.GetConverter("nopconverter")
	if c == nil {
		t.Fatal(c)
	}
	var count int
	count = 0
	env.RangeConverters(func(key string, v Converter) bool {
		count = count + 1
		return true
	})
	if count != 2 {
		t.Fatal(count)
	}
	count = 0
	env.RangeConverters(func(key string, v Converter) bool {
		count = count + 1
		return false
	})
	if count != 1 {
		t.Fatal(count)
	}

	count = 0
	env.RangeFormatters(func(key string, v Formatter) bool {
		count = count + 1
		return true
	})
	if count != 2 {
		t.Fatal(count)
	}
	count = 0
	env.RangeFormatters(func(key string, v Formatter) bool {
		count = count + 1
		return false
	})
	if count != 1 {
		t.Fatal(count)
	}

	count = 0
	env.RangeParsers(func(key string, v Parser) bool {
		count = count + 1
		return true
	})
	if count != 2 {
		t.Fatal(count)
	}
	count = 0
	env.RangeParsers(func(key string, v Parser) bool {
		count = count + 1
		return false
	})
	if count != 1 {
		t.Fatal(count)
	}
}

func TestCloneAndMerge(t *testing.T) {
	env := NewEnvironment()
	f := env.GetFormatter("nopformatter")
	if f != nil {
		t.Fatal(f)
	}
	env.MergeWith(nil)
	f = env.GetFormatter("nopformatter")
	if f != nil {
		t.Fatal(f)
	}
	env2 := Clone(env)
	f = env.GetFormatter("nopformatter")
	if f != nil {
		t.Fatal(f)
	}
	env2.SetFormatter("nopformatter", nopFormatter)
	env2.SetParser("nopparser", nopParser)
	env2.SetConverter("nopconverter", nopConverter)
	f = env.GetFormatter("nopformatter")
	if f != nil {
		t.Fatal(f)
	}
	f = env2.GetFormatter("nopformatter")
	if f == nil {
		t.Fatal(f)
	}
	p := env.GetParser("nopparser")
	if p != nil {
		t.Fatal(p)
	}
	p = env2.GetParser("nopparser")
	if p == nil {
		t.Fatal(p)
	}
	c := env.GetConverter("nopconverter")
	if c != nil {
		t.Fatal(p)
	}
	c = env2.GetConverter("nopconverter")
	if c == nil {
		t.Fatal(c)
	}
	env.MergeWith(env2)
	f = env.GetFormatter("nopformatter")
	if f == nil {
		t.Fatal(f)
	}
	p = env.GetParser("nopparser")
	if p == nil {
		t.Fatal(p)
	}
	c = env.GetConverter("nopconverter")
	if c == nil {
		t.Fatal(p)
	}
	env2.SetFormatter("nopformatter2", nopFormatter)
	f = env2.GetFormatter("nopformatter2")
	if f == nil {
		t.Fatal(f)
	}
	f = env.GetFormatter("nopformatter2")
	if f != nil {
		t.Fatal(f)
	}
	env3 := Clone(env)
	f = env3.GetFormatter("nopformatter")
	if f == nil {
		t.Fatal(f)
	}
	p = env3.GetParser("nopparser")
	if p == nil {
		t.Fatal(p)
	}
	c = env3.GetConverter("nopconverter")
	if c == nil {
		t.Fatal(p)
	}

}
