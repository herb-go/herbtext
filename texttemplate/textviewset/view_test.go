package textviewset_test

import (
	"errors"
	"testing"

	"github.com/herb-go/herbtext"
	_ "github.com/herb-go/herbtext-drivers/engine/handlebars"
	"github.com/herb-go/herbtext/texttemplate"
	"github.com/herb-go/herbtext/texttemplate/textviewset"
)

var nopErr = errors.New("nop")

type testEngine struct{}

//Parse parse given template with given environment to template view.
func (testEngine) Parse(template string, env herbtext.Environment) (texttemplate.View, error) {
	return nil, nopErr
}

//Supported return supported directives which can be used in template string.
func (testEngine) Supported(env herbtext.Environment) (directives []string, err error) {
	return nil, nil
}
func TestView(t *testing.T) {
	templates := herbtext.NewMap()
	templates.Set("template1", "{{testkey1}}")
	templates.Set("template2", "{{testkey2}}")
	data := texttemplate.Dataset{}
	data["testkey1"] = "testvalue1"
	data["testkey2"] = "testvalue2"
	views, err := textviewset.ParseWithEngineName(templates, "handlebars", herbtext.DefaultEnvironment())
	if err != nil {
		panic(err)
	}
	outputs, err := views.Render(data)
	if err != nil {
		panic(err)
	}
	if outputs.Get("template1") != "testvalue1" || outputs.Get("template2") != "testvalue2" {
		t.Fatal(outputs)
	}
	views, err = textviewset.ParseWith(templates, testEngine{}, herbtext.DefaultEnvironment())
	if err != nopErr || views != nil {
		t.Fatal(views, err)
	}
}
