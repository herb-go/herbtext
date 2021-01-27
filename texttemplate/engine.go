package texttemplate

type Engine interface {
	ApplyOptions(*Options) error
	Parse(string, *Options) (View, error)
	Supported() ([]string, error)
}
