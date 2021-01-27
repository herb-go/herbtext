package texttemplate

type View interface {
	Render(data interface{}) (output string, err error)
	Supported() (directives []string, err error)
}
