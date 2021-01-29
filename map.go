package herbtext

//Map text key-value map
type Map map[string]string

func (m Map) Set(key, value string) {
	m[key] = value
}
func (m Map) Get(key string) string {
	return m[key]
}

func (m Map) Length() int {
	return len(m)
}
func (m Map) Range(f func(string, string) bool) {
	for k := range m {
		if !f(k, m[k]) {
			return
		}
	}
}

func NewMap() Map {
	return Map{}
}

type Set interface {
	Set(string, string)
	Get(string) string
	Range(func(string, string) bool)
	Length() int
}
