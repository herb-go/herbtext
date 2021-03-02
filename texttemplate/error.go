package texttemplate

import (
	"fmt"
)

//ParamMissedError param required error
type ParamMissedError struct {
	//Name missed param name
	Name string
}

//Error error message
func (e *ParamMissedError) Error() string {
	return fmt.Sprintf("texttemplate: param [%s] required", e.Name)
}

//NewParamMissedError create new param missed error with given name
func NewParamMissedError(name string) error {
	return &ParamMissedError{
		Name: name,
	}
}

//IsParamMissedError check if given error is ParamMissedError
func IsParamMissedError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ParamMissedError)
	return ok
}

//GetParamMissedErrorName get missed param error name from given error.
//Empty string will be returned if given error is null or not a ParamMissedError.
func GetParamMissedErrorName(err error) string {
	if err == nil {
		return ""
	}
	e, ok := err.(*ParamMissedError)
	if !ok {
		return ""
	}
	return e.Name
}
