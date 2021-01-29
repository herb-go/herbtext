package herbtext

//Map text key-value map
type Map map[string]string

//Set set given value to set by given key
func (m Map) Set(key, value string) {
	m[key] = value
}

//Get get value from set by given key
func (m Map) Get(key string) string {
	return m[key]
}

//Length return set length
func (m Map) Length() int {
	return len(m)
}

//Range range over set with given function.
//Stop range if function return false,
func (m Map) Range(f func(string, string) bool) {
	for k := range m {
		if !f(k, m[k]) {
			return
		}
	}
}

//NewMap create new map
func NewMap() Map {
	return Map{}
}

//Set text set interface
type Set interface {
	//Set set given value to set by given key
	Set(string, string)
	//Get get value from set by given key
	Get(string) string
	//Range range over set with given function.
	//Stop range if function return false,
	Range(func(string, string) bool)
	//Length return set length
	Length() int
}

//CloneSet clone set
func CloneSet(set Set) Map {
	m := NewMap()
	set.Range(func(key, value string) bool {
		m[key] = value
		return true
	})
	return m
}

//MergeSet merge sources set to target
//Value in later set will overwrite early one with same key
func MergeSet(target Set, sources ...Set) {
	for _, source := range sources {
		source.Range(func(key, value string) bool {
			target.Set(key, value)
			return true
		})
	}
}
