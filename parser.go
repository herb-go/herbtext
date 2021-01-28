package herbtext

//Parser paraser type.
//Parser parse given data to interface{} and return any error if rasied.
type Parser func(data string) (interface{}, error)
