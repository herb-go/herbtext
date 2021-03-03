package texttemplate

import (
	"errors"
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	var err error
	if IsParamMissedError(err) || GetParamMissedErrorName(err) != "" {
		t.Fatal()
	}
	err = errors.New("noperr")
	if IsParamMissedError(err) || GetParamMissedErrorName(err) != "" {
		t.Fatal()
	}
	err = NewParamMissedError("test")
	if !IsParamMissedError(err) || GetParamMissedErrorName(err) != "test" {
		t.Fatal()
	}
	msg := err.Error()
	if !strings.Contains(msg, "required") || !strings.Contains(msg, "test") {
		t.Fatal(msg)
	}
}
