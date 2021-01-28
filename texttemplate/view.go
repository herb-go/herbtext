package texttemplate

//View text template view interface.
type View interface {
	//Render render given data and return output string and any error if raised.
	Render(data interface{}) (output string, err error)
}
