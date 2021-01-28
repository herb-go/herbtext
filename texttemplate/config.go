package texttemplate

//Config engine config
type Config struct {
	//EngineType engine type
	Engine string
	//EngineConfig engine config
	EngineConfig func(v interface{}) error `config:", lazyload"`
}

//NewEngine create new engine with config
func (c *Config) NewEngine() (Engine, error) {
	return NewEngine(c.Engine, c.EngineConfig)
}
