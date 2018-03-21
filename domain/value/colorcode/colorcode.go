package colorcode

//go:generate stringer -type=ColorCode
type ColorCode uint

const (
	NoColor ColorCode = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
)
