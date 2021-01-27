package texttemplate

type View interface {
	Render(interface{}) (string, error)
}
