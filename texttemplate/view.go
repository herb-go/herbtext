package texttemplate

//Template text template  interface.
type Template interface {
	//Render render given data and return output string and any error if raised.
	Render(data interface{}) (output string, err error)
}
