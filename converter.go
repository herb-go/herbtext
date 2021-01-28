package herbtext

//Converter text converter type.
//Converter should convert give string to new text.
//Converter should panic if any error raised
type Converter func(string) string
