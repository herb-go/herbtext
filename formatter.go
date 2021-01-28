package herbtext

//Formatter formatter type.
//Formatter should format given data with in given foramt.
//Formatter should panic if any error raised.
type Formatter func(format string, data string) (formatted string)
