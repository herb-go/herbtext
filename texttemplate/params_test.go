package texttemplate_test

import (
	"testing"

	"github.com/herb-go/herbtext-drivers/engine/handlebars"

	"github.com/herb-go/herbtext"
	_ "github.com/herb-go/herbtext-drivers/commonenvironment"
	"github.com/herb-go/herbtext/texttemplate"
)

func TestParams(t *testing.T) {
	defer texttemplate.UnregisterAll()
	texttemplate.UnregisterAll()
	handlebars.Register()
	defs := &texttemplate.ParamDefinitions{
		{
			ParamConfig: texttemplate.ParamConfig{
				Source: "testsource",
				Target: "testtarget",
				Parser: "",
			},
		},
	}
	ps, err := defs.CreateParams(herbtext.DefaultEnvironment())
	if err != nil {
		panic(err)
	}
	values := herbtext.Map{}
	values.Set("testsource", "testvalue")
	ds, err := ps.Load(values)
	if err != nil {
		panic(err)
	}
	if len(ds) != 1 || ds["testtarget"].(string) != "testvalue" {
		t.Fatal(ds)
	}
	eng, err := texttemplate.GetEngine("handlebars")
	if err != nil {
		panic(err)
	}
	view, err := eng.Parse("output {{{testtarget}}}", herbtext.DefaultEnvironment())
	if err != nil {
		panic(err)
	}
	output, err := view.Render(ds)
	if err != nil {
		panic(err)
	}
	if output != "output testvalue" {
		t.Fatal(output)
	}
}

func TestRequired(t *testing.T) {
	defer texttemplate.UnregisterAll()
	texttemplate.UnregisterAll()
	handlebars.Register()
	defs := &texttemplate.ParamDefinitions{
		{
			ParamConfig: texttemplate.ParamConfig{
				Source:   "test",
				Parser:   "",
				Required: true,
			},
		},
	}
	ps, err := defs.CreateParams(herbtext.DefaultEnvironment())
	if err != nil {
		panic(err)
	}
	values := herbtext.Map{}
	_, err = ps.Load(values)
	if err == nil || texttemplate.GetParamMissedErrorName(err) != "test" {
		t.Fatal()
	}
	values.Set("test", "testvalue")
	_, err = ps.Load(values)
	if err != nil {
		t.Fatal()
	}
}

func TestConstant(t *testing.T) {
	defer texttemplate.UnregisterAll()
	texttemplate.UnregisterAll()
	handlebars.Register()
	defs := &texttemplate.ParamDefinitions{
		{
			ParamConfig: texttemplate.ParamConfig{
				Source:   "test",
				Parser:   "",
				Required: true,
				Constant: "constant",
			},
		},
	}
	ps, err := defs.CreateParams(herbtext.DefaultEnvironment())
	if err != nil {
		panic(err)
	}
	values := herbtext.Map{}
	ds, err := ps.Load(values)
	if err != nil {
		panic(err)
	}
	if ds["test"] != "constant" {
		t.Fatal()
	}
	values.Set("test", "testvalue")
	ds, err = ps.Load(values)
	if err != nil {
		panic(err)
	}
	if ds["test"] != "constant" {
		t.Fatal()
	}
}
