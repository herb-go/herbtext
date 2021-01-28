package textrender

type Renderer struct {
	Views     *Views
	LoaderSet LoaderSet
}

func (r *Renderer) RenderValues(values *Values) (*Outputs, error) {
	ds, err := r.LoaderSet.Load(values)
	if err != nil {
		return nil, err
	}
	return r.Views.Render(ds)
}
